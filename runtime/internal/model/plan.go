package model

// Operation defines the operations supported by the contract.
type Operation string

const (
	OpInstall   Operation = "install"
	OpConfigure Operation = "configure"
	OpValidate  Operation = "validate"
	OpRepair    Operation = "repair"
)

// DesiredTransition represents a required transition of state for a capability.
type DesiredTransition struct {
	Capability string
	ProviderID string
	FromState  State
	ToState    State
}

// ExecutionStep represents a single step in an ExecutionPlan.
type ExecutionStep struct {
	StepID            int
	Capability        string
	ProviderID        string
	Operation         Operation
	ExpectedPreState  State
	ExpectedPostState State
	ActualPreState    State
	ActualPostState   State
	Result            StepResult
	Stderr            string
}

type StepResult string

const (
	ResultPending StepResult = "PENDING"
	ResultSuccess StepResult = "SUCCESS"
	ResultFailed  StepResult = "FAILED"
	ResultSkipped StepResult = "SKIPPED"
)

// ExecutionPlan holds the ordered steps to reach the Desired State.
type ExecutionPlan struct {
	Steps []ExecutionStep
}

// ExecutionResult represents the outcome of an executed plan.
type ExecutionResult struct {
	Plan    *ExecutionPlan
	Outcome Outcome
}

type Outcome string

const (
	OutcomeConverged Outcome = "CONVERGED"
	OutcomePartial   Outcome = "PARTIAL"
	OutcomeFailed    Outcome = "FAILED"
)
