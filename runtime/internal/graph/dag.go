package graph

import (
	"fmt"
	"homelab/runtime/internal/model"
)

// ComponentGraph represents the directed acyclic graph of capabilities.
type ComponentGraph struct {
	// Components mapped by capability name
	Components map[string]model.Component
	
	// Ordered list of capabilities (topological sort)
	Order []string
}

// Build creates a ComponentGraph from a list of components, detecting cycles.
func Build(components []model.Component) (*ComponentGraph, error) {
	compMap := make(map[string]model.Component)
	adj := make(map[string][]string)

	for _, c := range components {
		if _, exists := compMap[c.Capability]; exists {
			return nil, fmt.Errorf("duplicate capability found in registry: %s", c.Capability)
		}
		compMap[c.Capability] = c
		adj[c.Capability] = c.Dependencies
	}

	// Verify all dependencies exist in the registry
	for _, c := range components {
		for _, dep := range c.Dependencies {
			if _, exists := compMap[dep]; !exists {
				return nil, fmt.Errorf("capability %s depends on %s which is not in the registry", c.Capability, dep)
			}
		}
	}

	order, err := topologicalSort(adj)
	if err != nil {
		return nil, err
	}

	return &ComponentGraph{
		Components: compMap,
		Order:      order,
	}, nil
}

func topologicalSort(adj map[string][]string) ([]string, error) {
	var order []string
	visited := make(map[string]bool)
	tempMark := make(map[string]bool)

	var visit func(node string) error
	visit = func(node string) error {
		if tempMark[node] {
			return fmt.Errorf("circular dependency detected involving %s", node)
		}
		if visited[node] {
			return nil
		}
		
		tempMark[node] = true
		for _, dep := range adj[node] {
			if err := visit(dep); err != nil {
				return err
			}
		}
		
		tempMark[node] = false
		visited[node] = true
		order = append(order, node)
		return nil
	}

	for node := range adj {
		if !visited[node] {
			if err := visit(node); err != nil {
				return nil, err
			}
		}
	}
	
	return order, nil
}
