package graph

func StartingVertex(name string) Vertex {
	return Vertex{
		SiteName: name,
		Data:     VertexStyle{Colour: colours.BLUE, Shape: shapes.ELLIPSE},
	}
}

func TaskVertex(name string) Vertex {
	return Vertex{
		SiteName: name,
		Data:     VertexStyle{Colour: colours.NEUTRAL, Shape: shapes.BOX},
	}
}

func EndingVertex(name string, isError bool) Vertex {
	col := colours.GREEN
	if isError {
		col = colours.RED
	}
	return Vertex{
		SiteName: name,
		Data:     VertexStyle{Colour: col, Shape: shapes.ELLIPSE},
	}
}
