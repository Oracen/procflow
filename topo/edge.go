package topo

type EdgeData struct {
	InvocationName string
}

type Edge struct {
	VertexFrom string
	VertexTo   string
	Data       EdgeData
}

type EdgeCollection map[string]Edge
