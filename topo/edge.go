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

func MergeEdges(edges1, edges2 EdgeCollection) (merged EdgeCollection, err error) {

	merged = EdgeCollection{}
	for key, value := range edges1 {
		// Create copy for safety
		merged[key] = value
	}

	for key, value := range edges2 {
		err = addGraphItem(merged, key, value, errVertexMergeDuplicate)
		if err != nil {
			// TODO: Better error communication
			return EdgeCollection{}, err
		}
	}
	return
}
