package flow_test

import (
	"testing"

	"github.com/Oracen/procflow/basic"
)

func proc1func1() (int, error) {
	return 1, nil
}

func proc1func2(int) (string, error) {
	return "string", nil
}

func proc1func3(string) error {
	return nil
}

// Current mode is ugly and manual, but we need to start somewhere
func proc1() {
	// No error branches should fail
	tracker := basic.RegisterTracker()
	nodeStart := tracker.StartFlow("input")

	node1 := tracker.AddNode([]basic.Node{nodeStart}, "intermediate")
	out1, err := proc1func1()

	if err != nil {
		tracker.EndFlow([]basic.Node{node1}, "error")
		return
	}
	node2 := tracker.AddNode([]basic.Node{node1}, "intermediate")
	out2, err2 := proc1func2(out1)
	if err2 != nil {
		tracker.EndFlow([]basic.Node{node1}, "error")
	}

	err3 := proc1func3(out2)
	if err3 != nil {
		tracker.EndFlow([]basic.Node{node2}, "error")
	}
	tracker.EndFlow([]basic.Node{node2}, "finish")

}

func TestPublicApi(t *testing.T) {

}
