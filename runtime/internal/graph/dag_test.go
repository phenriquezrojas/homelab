package graph

import (
	"homelab/runtime/internal/model"
	"testing"
)

func TestBuild_ValidDAG(t *testing.T) {
	components := []model.Component{
		{Capability: "container-runtime", ProviderID: "docker-engine", ContractVersion: 1, Dependencies: nil},
		{Capability: "reverse-proxy", ProviderID: "caddy", ContractVersion: 1, Dependencies: []string{"container-runtime"}},
	}

	g, err := Build(components)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(g.Order) != 2 {
		t.Fatalf("expected 2 nodes in order, got %d", len(g.Order))
	}

	// container-runtime must come before reverse-proxy
	indexCR := -1
	indexRP := -1
	for i, cap := range g.Order {
		if cap == "container-runtime" {
			indexCR = i
		}
		if cap == "reverse-proxy" {
			indexRP = i
		}
	}
	if indexCR >= indexRP {
		t.Errorf("container-runtime (index %d) should come before reverse-proxy (index %d)", indexCR, indexRP)
	}
}

func TestBuild_CircularDependency(t *testing.T) {
	components := []model.Component{
		{Capability: "a", ProviderID: "a-provider", ContractVersion: 1, Dependencies: []string{"b"}},
		{Capability: "b", ProviderID: "b-provider", ContractVersion: 1, Dependencies: []string{"a"}},
	}

	_, err := Build(components)
	if err == nil {
		t.Fatal("expected circular dependency error, got nil")
	}
}

func TestBuild_DuplicateCapability(t *testing.T) {
	components := []model.Component{
		{Capability: "x", ProviderID: "x1", ContractVersion: 1},
		{Capability: "x", ProviderID: "x2", ContractVersion: 1},
	}

	_, err := Build(components)
	if err == nil {
		t.Fatal("expected duplicate capability error, got nil")
	}
}

func TestBuild_MissingDependency(t *testing.T) {
	components := []model.Component{
		{Capability: "a", ProviderID: "a-provider", ContractVersion: 1, Dependencies: []string{"nonexistent"}},
	}

	_, err := Build(components)
	if err == nil {
		t.Fatal("expected missing dependency error, got nil")
	}
}

func TestBuild_NoDependencies(t *testing.T) {
	components := []model.Component{
		{Capability: "a", ProviderID: "a-provider", ContractVersion: 1},
		{Capability: "b", ProviderID: "b-provider", ContractVersion: 1},
	}

	g, err := Build(components)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(g.Order) != 2 {
		t.Fatalf("expected 2 nodes, got %d", len(g.Order))
	}
}

func TestBuild_FullRegistry(t *testing.T) {
	// Mirrors the actual registry.yaml
	components := []model.Component{
		{Capability: "container-runtime", ProviderID: "docker-engine", ContractVersion: 1, Dependencies: nil},
		{Capability: "private-network", ProviderID: "tailscale", ContractVersion: 1, Dependencies: nil},
		{Capability: "reverse-proxy", ProviderID: "caddy", ContractVersion: 1, Dependencies: []string{"container-runtime"}},
		{Capability: "internal-dns", ProviderID: "magic-dns", ContractVersion: 1, Dependencies: []string{"private-network"}},
	}

	g, err := Build(components)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(g.Order) != 4 {
		t.Fatalf("expected 4 nodes, got %d", len(g.Order))
	}

	// Verify dependency ordering
	orderIdx := make(map[string]int)
	for i, cap := range g.Order {
		orderIdx[cap] = i
	}

	if orderIdx["container-runtime"] >= orderIdx["reverse-proxy"] {
		t.Error("container-runtime must come before reverse-proxy")
	}
	if orderIdx["private-network"] >= orderIdx["internal-dns"] {
		t.Error("private-network must come before internal-dns")
	}
}
