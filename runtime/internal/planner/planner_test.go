package planner

import (
	"homelab/runtime/internal/graph"
	"homelab/runtime/internal/model"
	"testing"
)

func buildTestGraph(t *testing.T) *graph.ComponentGraph {
	t.Helper()
	components := []model.Component{
		{Capability: "container-runtime", ProviderID: "docker-engine", ContractVersion: 1, Dependencies: nil},
		{Capability: "private-network", ProviderID: "tailscale", ContractVersion: 1, Dependencies: nil},
		{Capability: "reverse-proxy", ProviderID: "caddy", ContractVersion: 1, Dependencies: []string{"container-runtime"}},
		{Capability: "internal-dns", ProviderID: "magic-dns", ContractVersion: 1, Dependencies: []string{"private-network"}},
	}
	g, err := graph.Build(components)
	if err != nil {
		t.Fatalf("failed to build test graph: %v", err)
	}
	return g
}

func TestPlan_AllAbsent(t *testing.T) {
	g := buildTestGraph(t)
	observed := map[string]model.State{
		"container-runtime": model.StateAbsent,
		"private-network":   model.StateAbsent,
		"reverse-proxy":     model.StateAbsent,
		"internal-dns":      model.StateAbsent,
	}
	target := map[string]model.State{} // default HEALTHY

	transitions := Plan(g, observed, target)

	if len(transitions) != 4 {
		t.Fatalf("expected 4 transitions, got %d", len(transitions))
	}

	for _, tr := range transitions {
		if tr.FromState != model.StateAbsent {
			t.Errorf("expected FromState ABSENT for %s, got %s", tr.Capability, tr.FromState)
		}
		if tr.ToState != model.StateHealthy {
			t.Errorf("expected ToState HEALTHY for %s, got %s", tr.Capability, tr.ToState)
		}
	}
}

func TestPlan_AllHealthy(t *testing.T) {
	g := buildTestGraph(t)
	observed := map[string]model.State{
		"container-runtime": model.StateHealthy,
		"private-network":   model.StateHealthy,
		"reverse-proxy":     model.StateHealthy,
		"internal-dns":      model.StateHealthy,
	}
	target := map[string]model.State{}

	transitions := Plan(g, observed, target)

	if len(transitions) != 0 {
		t.Fatalf("expected 0 transitions (system converged), got %d", len(transitions))
	}
}

func TestPlan_PartialConvergence(t *testing.T) {
	g := buildTestGraph(t)
	observed := map[string]model.State{
		"container-runtime": model.StateHealthy,
		"private-network":   model.StateHealthy,
		"reverse-proxy":     model.StateAbsent,
		"internal-dns":      model.StateHealthy,
	}
	target := map[string]model.State{}

	transitions := Plan(g, observed, target)

	if len(transitions) != 1 {
		t.Fatalf("expected 1 transition, got %d", len(transitions))
	}
	if transitions[0].Capability != "reverse-proxy" {
		t.Errorf("expected transition for reverse-proxy, got %s", transitions[0].Capability)
	}
}

func TestPlan_FailedComponent(t *testing.T) {
	g := buildTestGraph(t)
	observed := map[string]model.State{
		"container-runtime": model.StateHealthy,
		"private-network":   model.StateFailed,
		"reverse-proxy":     model.StateHealthy,
		"internal-dns":      model.StateHealthy,
	}
	target := map[string]model.State{}

	transitions := Plan(g, observed, target)

	if len(transitions) != 1 {
		t.Fatalf("expected 1 transition, got %d", len(transitions))
	}
	if transitions[0].FromState != model.StateFailed {
		t.Errorf("expected FromState FAILED, got %s", transitions[0].FromState)
	}
}
