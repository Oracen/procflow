package graph

// Initial circular vertex with no corresponding edge
func StartingVertex(name, parentFlow string) Vertex {
	return Vertex{
		SiteName: name,
		Data: VertexStyle{
			Colour:      colours.BLUE,
			Shape:       shapes.ELLIPSE,
			ParentFlow:  parentFlow,
			IsFlowStart: true,
		},
	}
}

// Intermediate task node denoting process activity
func TaskVertex(name, parentFlow string) Vertex {
	return Vertex{
		SiteName: name,
		Data: VertexStyle{
			Colour:      colours.NEUTRAL,
			Shape:       shapes.BOX,
			ParentFlow:  parentFlow,
			IsFlowStart: false,
		},
	}
}

// Final circular vertex with no child nodes
func EndingVertex(name, parentFlow string, isError bool, isReturned bool) Vertex {
	col := colours.GREEN
	if isError {
		col = colours.RED
	}
	shape := shapes.ELLIPSE
	if !isReturned {
		shape = shapes.INVHOUSE
	}
	return Vertex{
		SiteName: name,
		Data: VertexStyle{
			Colour:      col,
			Shape:       shape,
			ParentFlow:  parentFlow,
			IsFlowStart: false,
		},
	}
}
