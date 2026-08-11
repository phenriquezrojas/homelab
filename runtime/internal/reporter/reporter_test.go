package reporter

import (
	"bytes"
	"homelab/runtime/internal/model"
	"strings"
	"testing"
)

func TestPrintResult_Converged(t *testing.T) {
	result := &model.ExecutionResult{
		Plan: &model.ExecutionPlan{
			Steps: []model.ExecutionStep{
				{
					StepID:            1,
					Capability:        "container-runtime",
					ProviderID:        "docker-engine",
					Operation:         model.OpInstall,
					ExpectedPreState:  model.StateAbsent,
					ExpectedPostState: model.StateInstalled,
					ActualPreState:    model.StateAbsent,
					ActualPostState:   model.StateInstalled,
					Result:            model.ResultSuccess,
				},
			},
		},
		Outcome: model.OutcomeConverged,
	}

	var buf bytes.Buffer
	PrintResult(&buf, result)
	output := buf.String()

	if !strings.Contains(output, "CONVERGED") {
		t.Error("output should contain CONVERGED")
	}
	if !strings.Contains(output, "docker-engine") {
		t.Error("output should contain provider ID")
	}
	if !strings.Contains(output, "Steps executed: 1/1") {
		t.Errorf("unexpected output: %s", output)
	}
}

func TestPrintResult_Failed(t *testing.T) {
	result := &model.ExecutionResult{
		Plan: &model.ExecutionPlan{
			Steps: []model.ExecutionStep{
				{
					StepID:     1,
					Capability: "container-runtime",
					ProviderID: "docker-engine",
					Operation:  model.OpInstall,
					Result:     model.ResultFailed,
					Stderr:     "permission denied",
				},
				{
					StepID:     2,
					Capability: "container-runtime",
					ProviderID: "docker-engine",
					Operation:  model.OpConfigure,
					Result:     model.ResultPending,
				},
			},
		},
		Outcome: model.OutcomeFailed,
	}

	var buf bytes.Buffer
	PrintResult(&buf, result)
	output := buf.String()

	if !strings.Contains(output, "FAILED") {
		t.Error("output should contain FAILED")
	}
	if !strings.Contains(output, "permission denied") {
		t.Error("output should contain error message")
	}
	// Step 2 is PENDING so it should not be counted as executed
	if !strings.Contains(output, "Steps executed: 1/2") {
		t.Errorf("unexpected output: %s", output)
	}
}

func TestPrintResult_EmptyPlan(t *testing.T) {
	result := &model.ExecutionResult{
		Plan:    &model.ExecutionPlan{},
		Outcome: model.OutcomeConverged,
	}

	var buf bytes.Buffer
	PrintResult(&buf, result)
	output := buf.String()

	if !strings.Contains(output, "Steps executed: 0/0") {
		t.Errorf("unexpected output for empty plan: %s", output)
	}
}
