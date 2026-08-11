package observer

import (
	"bytes"
	"context"
	"fmt"
	"homelab/runtime/internal/model"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

const operationTimeout = 30 * time.Second

// Observer is responsible for determining the empirical state of a capability.
type Observer struct {
	componentsDir string
}

func New(componentsDir string) *Observer {
	return &Observer{
		componentsDir: componentsDir,
	}
}

// Observe executes the validate script of a component and returns its state.
func (o *Observer) Observe(comp model.Component) (model.State, error) {
	scriptPath := filepath.Join(o.componentsDir, comp.ProviderID, string(model.OpValidate)+".sh")
	
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, "bash", scriptPath)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	
	if ctx.Err() == context.DeadlineExceeded {
		return model.StateUnknown, &model.ContractFailure{
			Capability: comp.Capability,
			ProviderID: comp.ProviderID,
			Operation:  model.OpValidate,
			Err:        model.ErrTimeout,
			Message:    "validate operation timed out",
		}
	}

	if err != nil {
		return model.StateUnknown, &model.ContractFailure{
			Capability: comp.Capability,
			ProviderID: comp.ProviderID,
			Operation:  model.OpValidate,
			Err:        model.ErrProcessCrash,
			Message:    fmt.Sprintf("exit code %v, stderr: %s", err, strings.TrimSpace(stderr.String())),
		}
	}

	outStr := strings.TrimSpace(stdout.String())
	state := model.State(outStr)

	// Validate output strictly
	switch state {
	case model.StateAbsent, model.StateInstalled, model.StateConfigured, model.StateHealthy, model.StateFailed:
		return state, nil
	default:
		return model.StateUnknown, &model.ContractFailure{
			Capability: comp.Capability,
			ProviderID: comp.ProviderID,
			Operation:  model.OpValidate,
			Err:        model.ErrInvalidOutput,
			Message:    fmt.Sprintf("unrecognized state string: '%s'", outStr),
		}
	}
}
