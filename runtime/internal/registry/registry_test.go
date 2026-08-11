package registry

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoad_ValidRegistry(t *testing.T) {
	content := `schema_version: 1

capabilities:
  - capability: container-runtime
    provider_id: docker-engine
    contract_version: 1
    dependencies: []

  - capability: reverse-proxy
    provider_id: caddy
    contract_version: 1
    dependencies: [container-runtime]
`
	tmpDir := t.TempDir()
	path := filepath.Join(tmpDir, "registry.yaml")
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatalf("failed to write test registry: %v", err)
	}

	reg, err := Load(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if reg.SchemaVersion != 1 {
		t.Errorf("expected schema version 1, got %d", reg.SchemaVersion)
	}

	if len(reg.Capabilities) != 2 {
		t.Fatalf("expected 2 capabilities, got %d", len(reg.Capabilities))
	}

	if reg.Capabilities[0].Capability != "container-runtime" {
		t.Errorf("expected first capability to be container-runtime, got %s", reg.Capabilities[0].Capability)
	}

	if reg.Capabilities[1].ProviderID != "caddy" {
		t.Errorf("expected second provider to be caddy, got %s", reg.Capabilities[1].ProviderID)
	}
}

func TestLoad_InvalidSchemaVersion(t *testing.T) {
	content := `schema_version: 99
capabilities: []
`
	tmpDir := t.TempDir()
	path := filepath.Join(tmpDir, "registry.yaml")
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatalf("failed to write test registry: %v", err)
	}

	_, err := Load(path)
	if err == nil {
		t.Fatal("expected error for unsupported schema version")
	}
}

func TestLoad_FileNotFound(t *testing.T) {
	_, err := Load("/nonexistent/path/registry.yaml")
	if err == nil {
		t.Fatal("expected error for missing file")
	}
}

func TestLoad_InvalidYAML(t *testing.T) {
	tmpDir := t.TempDir()
	path := filepath.Join(tmpDir, "registry.yaml")
	if err := os.WriteFile(path, []byte("{{invalid yaml"), 0644); err != nil {
		t.Fatalf("failed to write test file: %v", err)
	}

	_, err := Load(path)
	if err == nil {
		t.Fatal("expected error for invalid YAML")
	}
}
