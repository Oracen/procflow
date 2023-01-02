package graph

type VertexData struct {
	taskName string
}

type Vertex struct {
	siteName string
	data     VertexData
}

type VertexCollection map[string]Vertex

func MergeVertices(vertices1, vertices2 VertexCollection) (merged VertexCollection) {
	merged = VertexCollection{}
	for key, value := range vertices1 {
		merged[key] = value
	}
	for key, value := range vertices2 {
		merged[key] = value
	}
	return
}
