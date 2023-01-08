package graph

func StartingVertex(name, parentFlow string) Vertex {
	return Vertex{
		SiteName: name,
		Data: VertexStyle{
			Colour:      colours.BLUE,
			Shape:       shapes.ELLIPSE,
			ParentFlow:  parentFlow,
			IsFlowStart: true,
			IsFlowEnd:   false,
		},
	}
}

func TaskVertex(name, parentFlow string) Vertex {
	return Vertex{
		SiteName: name,
		Data: VertexStyle{
			Colour:      colours.NEUTRAL,
			Shape:       shapes.BOX,
			ParentFlow:  parentFlow,
			IsFlowStart: false,
			IsFlowEnd:   false,
		},
	}
}

func EndingVertex(name, parentFlow string, isError bool) Vertex {
	col := colours.GREEN
	if isError {
		col = colours.RED
	}
	return Vertex{
		SiteName: name,
		Data: VertexStyle{
			Colour:      col,
			Shape:       shapes.ELLIPSE,
			ParentFlow:  parentFlow,
			IsFlowStart: false,
			IsFlowEnd:   true,
		},
	}
}
