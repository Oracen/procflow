package store

import "github.com/Oracen/bpmn-flow/topo"

type Collector struct {
	state State[Collector]
}

func CreateNewCollector(state State[Collector]) Collector {
	return Collector{state: state}
}

func (c *Collector) TrackNewGraph() (graph topo.Graph) {
	graph = topo.CreateNewGraph()
	return
}
