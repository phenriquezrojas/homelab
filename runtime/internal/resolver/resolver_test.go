package resolver

import (
	"homelab/runtime/internal/model"
	"testing"
)

func TestResolve_AbsentToHealthy(t *testing.T) {
	transitions := []model.DesiredTransition{
		{
			Capability: "container-runtime",
			ProviderID: "docker-engine",
			FromState:  model.StateAbsent,
			ToState:    model.StateHealthy,
		},
	}

	plan := Resolve(transitions)

	if len(plan.Steps) != 2 {
		t.Fatalf("expected 2 steps (install + configure), got %d", len(plan.Steps))
	}

	if plan.Steps[0].Operation != model.OpInstall {
		t.Errorf("step 1 should be install, got %s", plan.Steps[0].Operation)
	}
	if plan.Steps[1].Operation != model.OpConfigure {
		t.Errorf("step 2 should be configure, got %s", plan.Steps[1].Operation)
	}
}

func TestResolve_InstalledToHealthy(t *testing.T) {
	transitions := []model.DesiredTransition{
		{
			Capability: "private-network",
			ProviderID: "tailscale",
			FromState:  model.StateInstalled,
			ToState:    model.StateHealthy,
		},
	}

	plan := Resolve(transitions)

	if len(plan.Steps) != 1 {
		t.Fatalf("expected 1 step (configure), got %d", len(plan.Steps))
	}

	if plan.Steps[0].Operation != model.OpConfigure {
		t.Errorf("step 1 should be configure, got %s", plan.Steps[0].Operation)
	}
}

func TestResolve_FailedToHealthy(t *testing.T) {
	transitions := []model.DesiredTransition{
		{
			Capability: "reverse-proxy",
			ProviderID: "caddy",
			FromState:  model.StateFailed,
			ToState:    model.StateHealthy,
		},
	}

	plan := Resolve(transitions)

	if len(plan.Steps) != 2 {
		t.Fatalf("expected 2 steps (repair + configure), got %d", len(plan.Steps))
	}

	if plan.Steps[0].Operation != model.OpRepair {
		t.Errorf("step 1 should be repair, got %s", plan.Steps[0].Operation)
	}
	if plan.Steps[1].Operation != model.OpConfigure {
		t.Errorf("step 2 should be configure, got %s", plan.Steps[1].Operation)
	}
}

func TestResolve_EmptyTransitions(t *testing.T) {
	plan := Resolve(nil)

	if len(plan.Steps) != 0 {
		t.Fatalf("expected 0 steps, got %d", len(plan.Steps))
	}
}

func TestResolve_MultipleTransitions(t *testing.T) {
	transitions := []model.DesiredTransition{
		{Capability: "a", ProviderID: "a-prov", FromState: model.StateAbsent, ToState: model.StateHealthy},
		{Capability: "b", ProviderID: "b-prov", FromState: model.StateInstalled, ToState: model.StateHealthy},
		{Capability: "c", ProviderID: "c-prov", FromState: model.StateFailed, ToState: model.StateHealthy},
	}

	plan := Resolve(transitions)

	// a: install + configure = 2
	// b: configure = 1
	// c: repair + configure = 2
	// Total: 5
	if len(plan.Steps) != 5 {
		t.Fatalf("expected 5 steps, got %d", len(plan.Steps))
	}

	// Verify step IDs are sequential
	for i, step := range plan.Steps {
		if step.StepID != i+1 {
			t.Errorf("step %d has ID %d, expected %d", i, step.StepID, i+1)
		}
	}
}
