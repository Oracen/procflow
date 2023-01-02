package graph

type EdgeData struct{}

type Edge struct {
	vertexFrom string
	vertexTo   string
	data       EdgeData
}

type EdgeCollection map[string]Edge
