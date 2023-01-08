package flow_test

import (
	"context"
	"testing"

	"github.com/Oracen/procflow/annotation/graph"
)

func proc1func1(ctx context.Context) (int, error) {
	// Fake some functions in here, just add the nodes
	tracker := graph.RegisterTracker(ctx)
	defer tracker.CloseTrace()
	ctx, nodeStart := graph.Start(ctx, &tracker, "inner1Input", "First inner input node")

	_, node1 := graph.Task(ctx, &tracker, []graph.Node{nodeStart}, "inner1Intermediate", "Task in first process")

	graph.End(&tracker, []graph.Node{node1}, "inner1Finish", "Finish first process", false)
	return 1, nil
}

func proc1func2(ctx context.Context, input int) (string, error) {
	// Fake more nodes
	tracker := graph.RegisterTracker(ctx)
	defer tracker.CloseTrace()
	ctx, nodeStart := graph.Start(ctx, &tracker, "inner2Input", "Our input node")

	ctx, node1 := graph.Task(ctx, &tracker, []graph.Node{nodeStart}, "inner2Intermediate", "Second process, first task")

	_, node2 := graph.Task(ctx, &tracker, []graph.Node{node1}, "inner2Intermediate2", "Second process, second task")

	graph.End(&tracker, []graph.Node{node2}, "inner2Finish", "Finish second process", false)
	return "string", nil
}

func proc1func3(ctx context.Context, inputInt int, inputStr string) error {
	// More dummy input
	tracker := graph.RegisterTracker(ctx)
	defer tracker.CloseTrace()
	ctx, nodeStart := graph.Start(ctx, &tracker, "inner3Input", "Third inner input node")

	_, node1 := graph.Task(ctx, &tracker, []graph.Node{nodeStart}, "inner3Intermediate", "Task in third process")

	graph.End(&tracker, []graph.Node{node1}, "inner3Finish", "Finish third process", false)
	return nil
}

// Current mode is ugly and manual, but we need to start somewhere
func proc1(ctx context.Context, willFailOn int) {
	// No error branches should fail
	tracker := graph.RegisterTracker(ctx)
	defer tracker.CloseTrace()

	ctx, nodeStart := graph.Start(ctx, &tracker, "input", "Our input node")

	ctx, node1 := graph.Task(ctx, &tracker, []graph.Node{nodeStart}, "intermediate", "Top-level task")
	out1, err := proc1func1(ctx)

	if err != nil || willFailOn == 0 {
		graph.End(&tracker, []graph.Node{node1}, "error", "first error node", true)
		return
	}

	ctx, node2 := graph.Task(ctx, &tracker, []graph.Node{node1}, "intermediate2", "Top-level task with int input")
	out2, err2 := proc1func2(ctx, out1)
	if err2 != nil || willFailOn == 1 {
		graph.End(&tracker, []graph.Node{node2}, "error2", "second error node", true)
		return
	}

	ctx, node3 := graph.Task(ctx, &tracker, []graph.Node{node1, node2}, "intermediate3", "Top-level task with mixed input")
	err3 := proc1func3(ctx, out1, out2)
	if err3 != nil || willFailOn == 2 {
		graph.End(&tracker, []graph.Node{node3}, "error3", "third error node", true)
		return
	}
	graph.End(&tracker, []graph.Node{node3}, "finish", "Final top level node - success", false)

}

func TestDemoFlow(t *testing.T) {
	t.Run(
		"test node name conversion functions",
		func(t *testing.T) {
			ctx := context.Background()
			for idx := 0; idx < 3; idx++ {
				proc1(ctx, idx)
			}

		},
	)
}
