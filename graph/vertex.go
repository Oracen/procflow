package graph

type VertexData struct {
	taskName string
}

type Vertex struct {
	siteName string
	data     VertexData
}

type VertexCollection map[string]Vertex
