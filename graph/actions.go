package graph

func Start(tracker *Tracker, name, description string) Node {
	node := Constructor{
		Name:     name,
		Vertex:   StartingVertex(description),
		EdgeData: StandardEdge(),
	}
	return tracker.StartFlow(node)
}

func Task(tracker *Tracker, inputs []Node, name, description string) Node {
	node := Constructor{
		Name:     name,
		Vertex:   TaskVertex(description),
		EdgeData: StandardEdge(),
	}
	return tracker.AddNode(inputs, node)
}

func End(tracker *Tracker, inputs []Node, name, description string, isError bool) {
	edge := StandardEdge()
	if isError {
		edge = ErrorEdge()
	}
	node := Constructor{
		Name:     name,
		Vertex:   EndingVertex(description, isError),
		EdgeData: edge,
	}
	tracker.EndFlow(inputs, node)
}
