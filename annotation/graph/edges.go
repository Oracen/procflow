package graph

// Standard black edge for conventional program links
func StandardEdge() EdgeStyle {
	return EdgeStyle{Colour: colours.BLACK}
}

// Creates red edge to highlight error paths
func ErrorEdge() EdgeStyle {
	return EdgeStyle{Colour: colours.RED}
}
