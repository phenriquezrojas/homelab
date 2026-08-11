package planner

import (
	"homelab/runtime/internal/graph"
	"homelab/runtime/internal/model"
)

// Plan calculates the desired transitions required to move from the observed state to the target state.
// Target states are assumed to be HEALTHY for all components in the graph unless specified otherwise.
func Plan(g *graph.ComponentGraph, observedStates map[string]model.State, targetStates map[string]model.State) []model.DesiredTransition {
	var transitions []model.DesiredTransition

	// Iterate in topological order so dependencies are processed first.
	for _, capName := range g.Order {
		comp := g.Components[capName]
		observed := observedStates[capName]
		
		target, ok := targetStates[capName]
		if !ok {
			target = model.StateHealthy
		}

		if observed != target {
			transitions = append(transitions, model.DesiredTransition{
				Capability: comp.Capability,
				ProviderID: comp.ProviderID,
				FromState:  observed,
				ToState:    target,
			})
		}
	}

	return transitions
}
