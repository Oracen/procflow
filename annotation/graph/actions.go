package graph

import (
	"context"
	"io"

	"github.com/Oracen/procflow/core/collections"
	"github.com/Oracen/procflow/core/constants"
	"github.com/Oracen/procflow/core/fileio"
	"github.com/Oracen/procflow/core/store"
	"github.com/Oracen/procflow/core/stringhandle"
	"github.com/Oracen/procflow/core/tracker"
)

var (
	singletonPtr Singleton
	// Local pointer to global state, can be reset for testing
	StateManager = &store.StateManager
)

func mockStateManagement() {
	stateManager := store.CreateNewStateManager()
	StateManager = &stateManager
}

func registerGlobal(singletonPointer *Singleton, collection *Collection, writer func(string) io.Writer) *exporter {
	storage := store.CreateGlobalSingleton(singletonPointer, store.StateManager.GetLock())
	storage.AddObject(collection)
	export := exporter{storage, writer}
	return &export
}

// Initialises graphviz-based graph tracking object with required params
func RegisterTracker(ctx context.Context) (t Tracker) {
	// Set up simple values
	parentName := stringhandle.GetParentFlow(ctx)
	collection := collections.CreateNewGraphCollector[VertexStyle, EdgeStyle]()

	if StateManager.UseGlobalState() {
		// If shared state enabled, use the singleton to bring the object in
		export := registerGlobal(&singletonPtr, &collection, fileio.CreateFileEncapsulation("graph.gv"))
		StateManager.AddExporter("graph", export)
	}
	return tracker.RegisterGraphTracker(&collection, parentName)
}

// Action annotation for marking start of a process flow
func Start(tracker *Tracker, name, description string) (ctxNew context.Context, node Node) {
	if StateManager.TrackState() {
		params := Constructor{
			Name:     stringhandle.PackNames(tracker.NameParentNode, name),
			Vertex:   StartingVertex(description, tracker.NameParentNode),
			EdgeData: StandardEdge(),
		}
		node = tracker.StartFlow(params)
	}

	return
}

// Action annotation for marking intermediate tasks of a process
func Task(ctx context.Context, tracker *Tracker, inputs []Node, name, description string) (ctxNew context.Context, node Node) {
	if StateManager.TrackState() {
		params := Constructor{
			Name:     stringhandle.PackNames(tracker.NameParentNode, name),
			Vertex:   TaskVertex(description, tracker.NameParentNode),
			EdgeData: StandardEdge(),
		}
		node = tracker.AddNode(inputs, params)
	}
	ctxNew = context.WithValue(ctx, constants.ContextParentFlowKey, name)
	return
}

// Action annotation for marking the endpoints of a process flow
func End(tracker *Tracker, inputs []Node, name, description string, isReturned, isError bool) (node Node) {
	if StateManager.TrackState() {
		edge := StandardEdge()
		if isError {
			edge = ErrorEdge()
		}
		params := Constructor{
			Name:     stringhandle.PackNames(tracker.NameParentNode, name),
			Vertex:   EndingVertex(description, tracker.NameParentNode, isError, isReturned),
			EdgeData: edge,
		}
		node = tracker.EndFlow(inputs, params)
	}
	return
}
