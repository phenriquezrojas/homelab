package model

import "testing"

func TestStateWeight(t *testing.T) {
	tests := []struct {
		state    State
		expected int
	}{
		{StateAbsent, 0},
		{StateInstalled, 1},
		{StateConfigured, 2},
		{StateHealthy, 3},
		{StateFailed, -1},
		{StateUnknown, -1},
	}

	for _, tt := range tests {
		got := StateWeight(tt.state)
		if got != tt.expected {
			t.Errorf("StateWeight(%s) = %d, want %d", tt.state, got, tt.expected)
		}
	}
}

func TestStateWeight_Ordering(t *testing.T) {
	// ABSENT < INSTALLED < CONFIGURED < HEALTHY
	if StateWeight(StateAbsent) >= StateWeight(StateInstalled) {
		t.Error("ABSENT should be less than INSTALLED")
	}
	if StateWeight(StateInstalled) >= StateWeight(StateConfigured) {
		t.Error("INSTALLED should be less than CONFIGURED")
	}
	if StateWeight(StateConfigured) >= StateWeight(StateHealthy) {
		t.Error("CONFIGURED should be less than HEALTHY")
	}
}

func TestContractFailure_Error(t *testing.T) {
	cf := &ContractFailure{
		Capability: "container-runtime",
		ProviderID: "docker-engine",
		Operation:  OpValidate,
		Err:        ErrInvalidOutput,
		Message:    "got empty output",
	}

	errMsg := cf.Error()
	if errMsg == "" {
		t.Error("expected non-empty error message")
	}

	// Should contain all relevant info
	if !contains(errMsg, "container-runtime") {
		t.Error("error should contain capability name")
	}
	if !contains(errMsg, "docker-engine") {
		t.Error("error should contain provider ID")
	}
	if !contains(errMsg, "validate") {
		t.Error("error should contain operation")
	}
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && searchSubstring(s, substr)
}

func searchSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
