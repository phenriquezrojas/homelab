package discovery

import (
	"fmt"
	"homelab/runtime/internal/graph"
	"homelab/runtime/internal/registry"
)

// Run executes the discovery pipeline: parse registry -> validate -> build DAG
func Run(registryPath string) (*graph.ComponentGraph, error) {
	reg, err := registry.Load(registryPath)
	if err != nil {
		return nil, fmt.Errorf("discovery failed during registry load: %w", err)
	}

	g, err := graph.Build(reg.Capabilities)
	if err != nil {
		return nil, fmt.Errorf("discovery failed during graph build: %w", err)
	}

	return g, nil
}
