package procflow_test

import (
	"context"
	"errors"
	"flag"
	"os"
	"testing"

	"github.com/Oracen/procflow"
	"github.com/Oracen/procflow/annotation/graph"
	"github.com/stretchr/testify/assert"
)

func returnInt(context.Context) int {
	return 1
}

func int2string(context.Context, int) string {
	return "a string"
}

func string2string(context.Context, string) string {
	return "another string"
}

func mixEmUp(context.Context, int, string) {
	// In the tradition of the Black Pearl, takes what it can and gives nothing back
}

func proc1func1(ctx context.Context) (int, error) {
	// Fake some logic in here
	tracker := graph.RegisterTracker(ctx)
	defer tracker.CloseTrace()
	ctx, nodeStart := graph.Start(ctx, &tracker, "inner1Input", "First inner input node")

	ctx, node1 := graph.Task(ctx, &tracker, []graph.Node{nodeStart}, "inner1Intermediate", "Task in first process")
	aNumber := returnInt(ctx)

	graph.End(&tracker, []graph.Node{node1}, "inner1Finish", "Finish first process", false)
	return aNumber, nil
}

func proc1func2(ctx context.Context, input int) (string, error) {
	// Fake more nodes
	tracker := graph.RegisterTracker(ctx)
	defer tracker.CloseTrace()
	ctx, nodeStart := graph.Start(ctx, &tracker, "inner2Input", "Our input node")

	ctx, node1 := graph.Task(ctx, &tracker, []graph.Node{nodeStart}, "inner2Intermediate", "Second process, first task")
	out1 := int2string(ctx, input)

	ctx, node2 := graph.Task(ctx, &tracker, []graph.Node{node1}, "inner2Intermediate2", "Second process, second task")
	out2 := string2string(ctx, out1)

	graph.End(&tracker, []graph.Node{node2}, "inner2Finish", "Finish second process", false)
	return out2, nil
}

func proc1func3(ctx context.Context, inputInt int, inputStr string) error {
	// More dummy input
	tracker := graph.RegisterTracker(ctx)
	defer tracker.CloseTrace()
	ctx, nodeStart := graph.Start(ctx, &tracker, "inner3Input", "Third inner input node")

	ctx, node1 := graph.Task(ctx, &tracker, []graph.Node{nodeStart}, "inner3Intermediate", "Task in third process")
	mixEmUp(ctx, inputInt, inputStr)

	graph.End(&tracker, []graph.Node{node1}, "inner3Finish", "Finish third process", false)
	return nil
}

// Current mode is ugly and manual, but we need to start somewhere
func proc1(ctx context.Context, willFailOn int) error {
	// No error branches should fail without our input
	tracker := graph.RegisterTracker(ctx)
	defer tracker.CloseTrace()

	ctx, nodeStart := graph.Start(ctx, &tracker, "input", "Our input node")

	ctx, node1 := graph.Task(ctx, &tracker, []graph.Node{nodeStart}, "intermediate", "Top-level task")
	out1, err := proc1func1(ctx)

	if err != nil || willFailOn == 0 {
		graph.End(&tracker, []graph.Node{node1}, "error", "first error node", true)
		return errors.New("error1")
	}

	ctx, node2 := graph.Task(ctx, &tracker, []graph.Node{node1}, "intermediate2", "Top-level task with int input")
	out2, err2 := proc1func2(ctx, out1)
	if err2 != nil || willFailOn == 1 {
		graph.End(&tracker, []graph.Node{node2}, "error2", "second error node", true)
		return errors.New("error2")
	}

	ctx, node3 := graph.Task(ctx, &tracker, []graph.Node{node1, node2}, "intermediate3", "Top-level task with mixed input")
	err3 := proc1func3(ctx, out1, out2)
	if err3 != nil || willFailOn == 2 {
		graph.End(&tracker, []graph.Node{node3}, "error3", "third error node", true)
		return errors.New("error3")
	}
	graph.End(&tracker, []graph.Node{node3}, "finish", "Final top level node - success!", false)
	return nil
}

func TestDemoFlow(t *testing.T) {
	t.Run(
		"test node name conversion functions",
		func(t *testing.T) {
			ctx := context.Background()
			for idx := 0; idx < 4; idx++ {
				err := proc1(ctx, idx)
				if idx == 3 {
					assert.Nil(t, err)
					continue
				}
				assert.NotNil(t, err)
			}
		},
	)
}

func TestMain(m *testing.M) {
	var recordFlow bool
	flag.BoolVar(&recordFlow, "recordflow", false, "use unit tests to measure program process flow")
	flag.Parse()

	procflow.StartFlowRecord(recordFlow)
	exitVal := m.Run()
	if exitVal == 0 {
		procflow.StopFlowRecord(recordFlow, ".")
	}

	os.Exit(exitVal)
}
