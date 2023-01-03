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
	nodeStart := tracker.StartFlow()

	node1 := tracker.RegisterNode("node1", []basic.Node{nodeStart})
	out1, err := proc1func1()

	if err != nil {
		tracker.TerminalNode("error1", []basic.Node{node1})
		return
	}
	node2 := tracker.RegisterNode("node2", []basic.Node{node1})
	out2, err2 := proc1func2(out1)
	if err2 != nil {
		tracker.TerminalNode("error2", []basic.Node{node1})
	}

	err3 := proc1func3(out2)
	if err3 != nil {
		tracker.TerminalNode("error3", []basic.Node{node2})
	}
	tracker.TerminalNode("end3", []basic.Node{node2})

}

func TestPublicApi(t *testing.T) {

}
