package executor

import (
	"bytes"
	"context"
	"fmt"
	"homelab/runtime/internal/graph"
	"homelab/runtime/internal/model"
	"homelab/runtime/internal/observer"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

const operationTimeout = 2 * time.Minute // Operations like docker pull might take time

type Executor struct {
	componentsDir string
	obs           *observer.Observer
}

func New(componentsDir string, obs *observer.Observer) *Executor {
	return &Executor{
		componentsDir: componentsDir,
		obs:           obs,
	}
}

// Execute runs the plan, enforcing Halt-on-Fail and Unexpected Success.
func (e *Executor) Execute(plan *model.ExecutionPlan, g *graph.ComponentGraph) *model.ExecutionResult {
	result := &model.ExecutionResult{
		Plan:    plan,
		Outcome: model.OutcomeConverged,
	}

	for i := range plan.Steps {
		step := &plan.Steps[i]
		comp := g.Components[step.Capability]

		// 1. Pre-Observation
		preState, err := e.obs.Observe(comp)
		if err != nil {
			step.Result = model.ResultFailed
			step.Stderr = fmt.Sprintf("pre-observation failed: %v", err)
			result.Outcome = model.OutcomeFailed
			break // Halt-on-Fail
		}
		step.ActualPreState = preState

		// Check for Unexpected Success before mutating
		if model.StateWeight(preState) >= model.StateWeight(step.ExpectedPostState) {
			// Already achieved or surpassed the expected post state
			step.Result = model.ResultSkipped
			step.ActualPostState = preState
			continue
		}

		// 2. Mutation
		err = e.runOperation(comp, step.Operation)
		if err != nil {
			step.Result = model.ResultFailed
			step.Stderr = fmt.Sprintf("operation failed: %v", err)
			result.Outcome = model.OutcomeFailed
			break // Halt-on-Fail
		}

		// 3. Post-Verification
		postState, err := e.obs.Observe(comp)
		if err != nil {
			step.Result = model.ResultFailed
			step.Stderr = fmt.Sprintf("post-observation failed: %v", err)
			result.Outcome = model.OutcomeFailed
			break // Halt-on-Fail
		}
		step.ActualPostState = postState

		// 4. Contrast
		if model.StateWeight(postState) >= model.StateWeight(step.ExpectedPostState) {
			step.Result = model.ResultSuccess
		} else {
			step.Result = model.ResultFailed
			step.Stderr = fmt.Sprintf("post state %s did not meet expected %s", postState, step.ExpectedPostState)
			result.Outcome = model.OutcomeFailed
			break // Halt-on-Fail
		}
	}

	return result
}

func (e *Executor) runOperation(comp model.Component, op model.Operation) error {
	scriptPath := filepath.Join(e.componentsDir, comp.ProviderID, string(op)+".sh")
	
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, "bash", scriptPath)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	// stdout is ignored for non-validate operations (used for debug internally if needed)

	err := cmd.Run()
	if ctx.Err() == context.DeadlineExceeded {
		return model.ErrTimeout
	}
	if err != nil {
		return fmt.Errorf("%w: %s", model.ErrProcessCrash, strings.TrimSpace(stderr.String()))
	}
	return nil
}
