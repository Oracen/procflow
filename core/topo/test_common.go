package topo

var (
	defaultVertexName     = "A name"
	altVertexName         = "a random name"
	defaultEdgeName       = "An Edge name"
	altEdgeName           = "An Edge nom de gurre"
	nonexistentVertexName = "Cest ne pas une vertex"
)

type utGraph = Graph[vertexData, edgeData]
type utVertex = Vertex[vertexData]
type utVertCol = VertexCollection[vertexData]
type utEdgeCol = EdgeCollection[edgeData]
type utEdge = Edge[edgeData]

type vertexData struct {
	TaskName string
}

type edgeData struct {
	InvocationName string
}

type edgeBuild struct {
	name string
	edge utEdge
}

type buildSlices struct {
	start, stop int
}

func createDefaultVertex(name string) utVertex {
	return utVertex{SiteName: name}
}

func initGraph() utGraph {
	graph := CreateNewGraph[vertexData, edgeData]()
	graph.AddNewVertex(defaultVertexName, createDefaultVertex(defaultVertexName))
	return graph
}

func createBasicVertices() utGraph {
	graph := initGraph()
	vertex := createDefaultVertex(altVertexName)
	graph.AddNewVertex(altVertexName, vertex)
	return graph
}

func createEdgePair(name1, name2 string) []edgeBuild {
	return []edgeBuild{
		{defaultEdgeName, utEdge{name1, name2, edgeData{}}},
		{altEdgeName, utEdge{name2, name1, edgeData{}}},
	}
}
