package model

import "errors"

var (
	// ErrInvalidOutput is returned when a component's output does not conform to the contract.
	ErrInvalidOutput = errors.New("INVALID_OUTPUT")
	
	// ErrProcessCrash is returned when a component exits with an error code during a non-validate operation.
	ErrProcessCrash = errors.New("PROCESS_CRASH")
	
	// ErrTimeout is returned when a component operation takes too long to complete.
	ErrTimeout = errors.New("TIMEOUT")
)

// ContractFailure represents a failure that aborts the execution cycle.
type ContractFailure struct {
	Capability string
	ProviderID string
	Operation  Operation
	Err        error
	Message    string
}

func (c *ContractFailure) Error() string {
	return c.Capability + " (" + c.ProviderID + ") failed during " + string(c.Operation) + ": " + c.Err.Error() + " - " + c.Message
}
