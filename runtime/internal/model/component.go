package model

// Component represents a single capability in the registry.
type Component struct {
	Capability      string   `yaml:"capability"`
	ProviderID      string   `yaml:"provider_id"`
	ContractVersion int      `yaml:"contract_version"`
	Dependencies    []string `yaml:"dependencies"`
}
