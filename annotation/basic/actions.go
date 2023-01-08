package basic

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

// Initialises string-based basic tracking object with required params
func RegisterTracker(ctx context.Context) (t Tracker) {
	// Set up simple values
	parentName, ok := ctx.Value(constants.ContextParentFlowKey).(string)
	collection := collections.CreateNewBasicCollector[string]()

	if StateManager.UseGlobalState() {
		// If shared state enabled, use the singleton to bring the object in
		export := registerGlobal(&singletonPtr, &collection, fileio.CreateFileEncapsulation("flow.txt"))
		StateManager.AddExporter("basic", export)
	}
	if parentName == "" || !ok {
		// If no parent, initialise to default parent flow name
		parentName = constants.ContextParentDefault
	}
	return tracker.RegisterBasicTracker(&collection, parentName)
}

// Action annotation for marking start of a process flow
func Start(ctx context.Context, tracker *Tracker, name, description string) (ctxNew context.Context, node Node) {
	if StateManager.TrackState() {
		taskName := stringhandle.PackNames(name, taskLabel.START)
		node = tracker.StartFlow(stringhandle.PackNames(tracker.NameParentNode, taskName))
	}
	ctxNew = context.WithValue(ctx, constants.ContextParentFlowKey, name)
	return
}

// Action annotation for marking intermediate tasks of a process
func Task(ctx context.Context, tracker *Tracker, inputs []Node, name, description string) (ctxNew context.Context, node Node) {
	if StateManager.TrackState() {
		taskName := stringhandle.PackNames(name, taskLabel.TASK)
		node = tracker.AddNode(inputs, stringhandle.PackNames(tracker.NameParentNode, taskName))
	}
	ctxNew = context.WithValue(ctx, constants.ContextParentFlowKey, name)
	return
}

// Action annotation for marking the endpoints of a process flow
func End(tracker *Tracker, inputs []Node, name, description string, isError bool) {
	if StateManager.TrackState() {
		nodeType := taskLabel.END
		if isError {
			nodeType = taskLabel.ERROR
		}
		taskName := stringhandle.PackNames(name, nodeType)
		tracker.EndFlow(inputs, stringhandle.PackNames(tracker.NameParentNode, taskName))
	}
}
