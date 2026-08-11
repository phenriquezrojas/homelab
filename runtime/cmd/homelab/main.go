package main

import (
	"fmt"
	"homelab/runtime/internal/discovery"
	"homelab/runtime/internal/executor"
	"homelab/runtime/internal/model"
	"homelab/runtime/internal/observer"
	"homelab/runtime/internal/planner"
	"homelab/runtime/internal/reporter"
	"homelab/runtime/internal/resolver"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		printHelp()
		os.Exit(1)
	}

	command := os.Args[1]
	registryPath := "runtime/registry.yaml" // Default for testing
	if p := os.Getenv("HOMELAB_REGISTRY"); p != "" {
		registryPath = p
	}
	componentsDir := "components" // Default for testing
	if p := os.Getenv("HOMELAB_COMPONENTS"); p != "" {
		componentsDir = p
	}

	switch command {
	case "converge":
		runConverge(registryPath, componentsDir)
	case "plan":
		runPlan(registryPath, componentsDir)
	case "validate":
		runValidate(registryPath, componentsDir)
	default:
		fmt.Printf("Unknown command: %s\n", command)
		printHelp()
		os.Exit(1)
	}
}

func printHelp() {
	fmt.Println("Usage: homelab <command>")
	fmt.Println("Commands:")
	fmt.Println("  converge   Converges the system to the desired state")
	fmt.Println("  plan       Shows the execution plan without executing it")
	fmt.Println("  validate   Shows the observed state of all capabilities")
}

func runValidate(registryPath, componentsDir string) {
	g, err := discovery.Run(registryPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Discovery failed: %v\n", err)
		os.Exit(1)
	}

	obs := observer.New(componentsDir)
	fmt.Println("Current State:")
	for _, capName := range g.Order {
		comp := g.Components[capName]
		state, err := obs.Observe(comp)
		if err != nil {
			fmt.Printf("  %s (%s): %s\n", comp.Capability, comp.ProviderID, err)
		} else {
			fmt.Printf("  %s (%s): %s\n", comp.Capability, comp.ProviderID, state)
		}
	}
}

func runPlan(registryPath, componentsDir string) {
	g, err := discovery.Run(registryPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Discovery failed: %v\n", err)
		os.Exit(1)
	}

	obs := observer.New(componentsDir)
	observedStates := make(map[string]model.State)

	for _, capName := range g.Order {
		comp := g.Components[capName]
		state, err := obs.Observe(comp)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Observation failed for %s: %v\n", comp.Capability, err)
			os.Exit(1)
		}
		observedStates[comp.Capability] = state
	}

	targetStates := make(map[string]model.State)
	// Default to HEALTHY for all capabilities

	transitions := planner.Plan(g, observedStates, targetStates)
	plan := resolver.Resolve(transitions)

	fmt.Println("Execution Plan:")
	if len(plan.Steps) == 0 {
		fmt.Println("  No operations needed. System is in the desired state.")
		return
	}

	for _, step := range plan.Steps {
		fmt.Printf("  Step %d: %s.%s (%s -> %s)\n", step.StepID, step.ProviderID, step.Operation, step.ExpectedPreState, step.ExpectedPostState)
	}
}

func runConverge(registryPath, componentsDir string) {
	g, err := discovery.Run(registryPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Discovery failed: %v\n", err)
		os.Exit(1)
	}

	obs := observer.New(componentsDir)
	observedStates := make(map[string]model.State)

	for _, capName := range g.Order {
		comp := g.Components[capName]
		state, err := obs.Observe(comp)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Observation failed for %s: %v\n", comp.Capability, err)
			os.Exit(1)
		}
		observedStates[comp.Capability] = state
	}

	targetStates := make(map[string]model.State)
	// Default to HEALTHY

	transitions := planner.Plan(g, observedStates, targetStates)
	plan := resolver.Resolve(transitions)

	if len(plan.Steps) == 0 {
		fmt.Println("No operations needed. System is in the desired state.")
		return
	}

	exec := executor.New(componentsDir, obs)
	result := exec.Execute(plan, g)

	reporter.PrintResult(os.Stdout, result)
	
	if result.Outcome == model.OutcomeFailed {
		os.Exit(1)
	}
}
