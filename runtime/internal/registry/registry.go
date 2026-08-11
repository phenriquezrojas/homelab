package registry

import (
	"fmt"
	"homelab/runtime/internal/model"
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

type RegistryFile struct {
	SchemaVersion int               `yaml:"schema_version"`
	Capabilities  []model.Component `yaml:"capabilities"`
}

// Load reads and parses the registry file from the given path.
func Load(path string) (*RegistryFile, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open registry file: %w", err)
	}
	defer f.Close()

	b, err := io.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("failed to read registry file: %w", err)
	}

	var reg RegistryFile
	if err := yaml.Unmarshal(b, &reg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal registry yaml: %w", err)
	}

	if reg.SchemaVersion != 1 {
		return nil, fmt.Errorf("unsupported registry schema version: %d", reg.SchemaVersion)
	}

	return &reg, nil
}
