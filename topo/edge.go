package topo

type EdgeData struct {
	invocationName string
}

type Edge struct {
	vertexFrom string
	vertexTo   string
	data       EdgeData
}

type EdgeCollection map[string]Edge
