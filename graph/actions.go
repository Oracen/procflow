package graph

import (
	"context"
	"fmt"
	"strings"

	"github.com/Oracen/procflow/core/constants"
	"github.com/Oracen/procflow/core/tracker"
)

func packNames(parentNodeName, currentNodeName string) string {
	return fmt.Sprintf("%s%s%s", parentNodeName, constants.StandardDelimiter, currentNodeName)
}

func unpackNames(parentNodeName, packedName string) string {
	prefix := fmt.Sprintf("%s%s", parentNodeName, constants.StandardDelimiter)
	return strings.TrimPrefix(packedName, prefix)
}

func RegisterTracker(ctx context.Context) Tracker {
	parentName, ok := ctx.Value(constants.ContextParentFlowKey).(string)
	if parentName == "" || !ok {
		parentName = constants.ContextParentDefault
	}
	collectable := tracker.CreateNewGraphCollectable[VertexStyle, EdgeStyle]()
	collector := tracker.CreateNewGraphCollector(&collectable)
	return tracker.RegisterGraphTracker(&collector, parentName)
}

func Start(ctx context.Context, tracker *Tracker, name, description string) (ctxNew context.Context, node Node) {
	params := Constructor{
		Name:     packNames(tracker.NameParentNode, name),
		Vertex:   StartingVertex(description, tracker.NameParentNode),
		EdgeData: StandardEdge(),
	}
	ctxNew = context.WithValue(ctx, constants.ContextParentFlowKey, name)
	node = tracker.StartFlow(params)
	return
}

func Task(ctx context.Context, tracker *Tracker, inputs []Node, name, description string) (ctxNew context.Context, node Node) {
	params := Constructor{
		Name:     packNames(tracker.NameParentNode, name),
		Vertex:   TaskVertex(description, tracker.NameParentNode),
		EdgeData: StandardEdge(),
	}
	ctxNew = context.WithValue(ctx, constants.ContextParentFlowKey, name)
	node = tracker.AddNode(inputs, params)
	return
}

func End(tracker *Tracker, inputs []Node, name, description string, isError bool) {
	edge := StandardEdge()
	if isError {
		edge = ErrorEdge()
	}
	params := Constructor{
		Name:     packNames(tracker.NameParentNode, name),
		Vertex:   EndingVertex(description, tracker.NameParentNode, isError),
		EdgeData: edge,
	}
	tracker.EndFlow(inputs, params)
}
