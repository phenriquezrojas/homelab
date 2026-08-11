package reporter

import (
	"fmt"
	"homelab/runtime/internal/model"
	"io"
)

// PrintResult formats and prints the ExecutionResult to the given writer.
func PrintResult(w io.Writer, result *model.ExecutionResult) {
	fmt.Fprintf(w, "Execution Result: %s\n\n", result.Outcome)

	var executed, succeeded, failed int

	for _, step := range result.Plan.Steps {
		if step.Result == model.ResultPending {
			continue // Did not reach this step due to Halt-on-Fail
		}

		executed++

		icon := "✓"
		if step.Result == model.ResultFailed {
			icon = "✗"
			failed++
		} else if step.Result == model.ResultSuccess || step.Result == model.ResultSkipped {
			succeeded++
		}

		fmt.Fprintf(w, "Step %d: %s.%s\n", step.StepID, step.ProviderID, step.Operation)
		fmt.Fprintf(w, "  actual_pre:  %s\n", step.ActualPreState)
		
		if step.Result == model.ResultSkipped {
			fmt.Fprintf(w, "  [skipped...]\n")
		} else {
			fmt.Fprintf(w, "  [executing %s...]\n", step.Operation)
		}
		
		fmt.Fprintf(w, "  actual_post: %s\n", step.ActualPostState)
		fmt.Fprintf(w, "  result: %s %s\n", step.Result, icon)
		
		if step.Result == model.ResultFailed && step.Stderr != "" {
			fmt.Fprintf(w, "  error: %s\n", step.Stderr)
		}
		fmt.Fprintln(w)
	}

	fmt.Fprintf(w, "Steps executed: %d/%d\n", executed, len(result.Plan.Steps))
	fmt.Fprintf(w, "Steps succeeded: %d\n", succeeded)
	fmt.Fprintf(w, "Steps failed: %d\n", failed)
}
