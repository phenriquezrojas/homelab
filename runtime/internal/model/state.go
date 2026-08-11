package model

// State represents the empirical state of a capability.
type State string

const (
	StateAbsent     State = "ABSENT"
	StateInstalled  State = "INSTALLED"
	StateConfigured State = "CONFIGURED"
	StateHealthy    State = "HEALTHY"
	StateFailed     State = "FAILED"
	StateUnknown    State = "UNKNOWN"
)

// StateWeight returns a numerical weight for states to compare hierarchy.
func StateWeight(s State) int {
	switch s {
	case StateAbsent:
		return 0
	case StateInstalled:
		return 1
	case StateConfigured:
		return 2
	case StateHealthy:
		return 3
	default:
		return -1 // FAILED or UNKNOWN are not part of the linear progression
	}
}
