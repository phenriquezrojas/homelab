package resolver

import "homelab/runtime/internal/model"

// Resolve translates a list of DesiredTransitions into an ordered ExecutionPlan.
func Resolve(transitions []model.DesiredTransition) *model.ExecutionPlan {
	plan := &model.ExecutionPlan{}
	stepID := 1

	for _, t := range transitions {
		// Only supporting transitions to HEALTHY right now as per the Matrix in Design Baseline.
		if t.ToState != model.StateHealthy {
			continue // In a full implementation, we might support demotion/destruction
		}

		switch t.FromState {
		case model.StateAbsent:
			// ABSENT -> HEALTHY: install, configure
			plan.Steps = append(plan.Steps, model.ExecutionStep{
				StepID:            stepID,
				Capability:        t.Capability,
				ProviderID:        t.ProviderID,
				Operation:         model.OpInstall,
				ExpectedPreState:  model.StateAbsent,
				ExpectedPostState: model.StateInstalled,
				Result:            model.ResultPending,
			})
			stepID++
			plan.Steps = append(plan.Steps, model.ExecutionStep{
				StepID:            stepID,
				Capability:        t.Capability,
				ProviderID:        t.ProviderID,
				Operation:         model.OpConfigure,
				ExpectedPreState:  model.StateInstalled,
				ExpectedPostState: model.StateHealthy,
				Result:            model.ResultPending,
			})
			stepID++

		case model.StateInstalled:
			// INSTALLED -> HEALTHY: configure
			plan.Steps = append(plan.Steps, model.ExecutionStep{
				StepID:            stepID,
				Capability:        t.Capability,
				ProviderID:        t.ProviderID,
				Operation:         model.OpConfigure,
				ExpectedPreState:  model.StateInstalled,
				ExpectedPostState: model.StateHealthy,
				Result:            model.ResultPending,
			})
			stepID++

		case model.StateConfigured:
			// CONFIGURED -> HEALTHY: configure
			plan.Steps = append(plan.Steps, model.ExecutionStep{
				StepID:            stepID,
				Capability:        t.Capability,
				ProviderID:        t.ProviderID,
				Operation:         model.OpConfigure,
				ExpectedPreState:  model.StateConfigured,
				ExpectedPostState: model.StateHealthy,
				Result:            model.ResultPending,
			})
			stepID++

		case model.StateFailed:
			// FAILED -> HEALTHY: repair, configure
			plan.Steps = append(plan.Steps, model.ExecutionStep{
				StepID:            stepID,
				Capability:        t.Capability,
				ProviderID:        t.ProviderID,
				Operation:         model.OpRepair,
				ExpectedPreState:  model.StateFailed,
				ExpectedPostState: model.StateConfigured,
				Result:            model.ResultPending,
			})
			stepID++
			plan.Steps = append(plan.Steps, model.ExecutionStep{
				StepID:            stepID,
				Capability:        t.Capability,
				ProviderID:        t.ProviderID,
				Operation:         model.OpConfigure,
				ExpectedPreState:  model.StateConfigured,
				ExpectedPostState: model.StateHealthy,
				Result:            model.ResultPending,
			})
			stepID++
		}
	}

	return plan
}
