package graph

import (
	"context"
	"fmt"
	"strings"

	"github.com/Oracen/procflow/core/constants"
	"github.com/Oracen/procflow/core/store"
	"github.com/Oracen/procflow/core/tracker"
)

var (
	singletonPtr Singleton
	StateManager = &store.StateManager
)

func packNames(parentNodeName, currentNodeName string) string {
	return fmt.Sprintf("%s%s%s", parentNodeName, constants.StandardDelimiter, currentNodeName)
}

func unpackNames(parentNodeName, packedName string) string {
	prefix := fmt.Sprintf("%s%s", parentNodeName, constants.StandardDelimiter)
	return strings.TrimPrefix(packedName, prefix)
}

func mockStateManagement() {
	stateManager := store.CreateNewStateManager()
	StateManager = &stateManager
}

func registerGlobal(singletonPointer *Singleton, collectable *Collectable) *exporter {
	storage := store.CreateGlobalSingleton(singletonPointer, store.StateManager.GetLock())
	storage.AddObject(collectable)
	export := exporter{storage, CreateFile}
	return &export
}

func RegisterTracker(ctx context.Context) (t Tracker) {
	// Set up simple values
	parentName, ok := ctx.Value(constants.ContextParentFlowKey).(string)
	collectable := tracker.CreateNewGraphCollectable[VertexStyle, EdgeStyle]()
	collector := tracker.CreateNewGraphCollector(&collectable)

	if StateManager.UseGlobalState() {
		// If shared state enabled, use the singleton to bring the object in
		export := registerGlobal(&singletonPtr, &collectable)
		StateManager.AddExporter("graph", export)
	}
	if parentName == "" || !ok {
		// If no parent, initialise to default parent flow name
		parentName = constants.ContextParentDefault
	}
	return tracker.RegisterGraphTracker(&collector, parentName)
}

func Start(ctx context.Context, tracker *Tracker, name, description string) (ctxNew context.Context, node Node) {
	if StateManager.TrackState() {
		params := Constructor{
			Name:     packNames(tracker.NameParentNode, name),
			Vertex:   StartingVertex(description, tracker.NameParentNode),
			EdgeData: StandardEdge(),
		}
		node = tracker.StartFlow(params)
	}
	ctxNew = context.WithValue(ctx, constants.ContextParentFlowKey, name)
	return
}

func Task(ctx context.Context, tracker *Tracker, inputs []Node, name, description string) (ctxNew context.Context, node Node) {
	if StateManager.TrackState() {
		params := Constructor{
			Name:     packNames(tracker.NameParentNode, name),
			Vertex:   TaskVertex(description, tracker.NameParentNode),
			EdgeData: StandardEdge(),
		}
		node = tracker.AddNode(inputs, params)
	}
	ctxNew = context.WithValue(ctx, constants.ContextParentFlowKey, name)
	return
}

func End(tracker *Tracker, inputs []Node, name, description string, isError bool) {
	if StateManager.TrackState() {
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
}
