package basic

import (
	"context"

	"github.com/Oracen/procflow/annotation/graph"
	"github.com/Oracen/procflow/core/stringhandle"
)

type Message[T any] struct {
	Nodes   []graph.Node
	Payload T
}

type TrackerCtx struct {
	Ctx   context.Context
	Track *graph.Tracker
}

func CreateNewTrackerCtx(
	ctx context.Context,
) TrackerCtx {
	tracker := graph.RegisterTracker(ctx)
	return TrackerCtx{Ctx: ctx, Track: &tracker}
}

func (d *TrackerCtx) CloseTrace() {
	d.Track.CloseTrace()
}

func Start[S, T any](
	d *TrackerCtx,
	name, description string,
	function func(context.Context, S) (T, error),
	data S,
) (msg Message[T], err error) {
	parentName := stringhandle.GetParentFlow(d.Ctx)
	node := graph.Start(d.Track, name, stringhandle.StartFlowName(parentName))
	ctx, node := graph.Task(d.Ctx, d.Track, []graph.Node{node}, name, description)
	ret, err := function(ctx, data)
	nodeList := []graph.Node{node}
	if err != nil {
		node = graph.End(d.Track, nodeList, name, stringhandle.ErrorFlowName(parentName), false, true)
		nodeList = []graph.Node{node}
	}
	msg = Message[T]{
		Nodes:   nodeList,
		Payload: ret,
	}
	return
}

// Because nullary or niladic was a bit too alien
func StartEmpty[S, T any](
	d *TrackerCtx,
	name, description string,
	function func(context.Context) (T, error),
	data S,
) (msg Message[T], err error) {
	parentName := stringhandle.GetParentFlow(d.Ctx)
	node := graph.Start(d.Track, name, stringhandle.StartFlowName(parentName))
	ctx, node := graph.Task(d.Ctx, d.Track, []graph.Node{node}, name, description)
	ret, err := function(ctx)
	nodeList := []graph.Node{node}
	if err != nil {
		node = graph.End(d.Track, nodeList, name, stringhandle.ErrorFlowName(parentName), false, true)
		nodeList = []graph.Node{node}
	}
	msg = Message[T]{
		Nodes:   nodeList,
		Payload: ret,
	}
	return
}

func TaskVoid[S, T any](
	d *TrackerCtx,
	name, description string,
	function func(context.Context) error,
	data S,
) (msg Message[error], err error) {
	parentName := stringhandle.GetParentFlow(d.Ctx)
	node := graph.Start(d.Track, name, stringhandle.StartFlowName(parentName))
	ctx, node := graph.Task(d.Ctx, d.Track, []graph.Node{node}, name, description)
	err = function(ctx)
	nodeList := []graph.Node{node}
	if err != nil {
		node = graph.End(d.Track, nodeList, name, stringhandle.ErrorFlowName(parentName), false, true)
	} else {
		node = graph.End(d.Track, nodeList, name, stringhandle.EndFlowName(parentName), true, false)
	}
	nodeList = []graph.Node{node}
	msg = Message[error]{
		Nodes:   nodeList,
		Payload: err,
	}
	return
}

func Task[S, T any](
	d *TrackerCtx,
	name, description string,
	function func(context.Context, S) (T, error),
	message Message[S],
) (msg Message[T], err error) {
	parentName := stringhandle.GetParentFlow(d.Ctx)
	ctx, node := graph.Task(d.Ctx, d.Track, message.Nodes, name, description)
	ret, err := function(ctx, message.Payload)
	nodeList := []graph.Node{node}
	if err != nil {
		node = graph.End(d.Track, nodeList, name, stringhandle.ErrorFlowName(parentName), false, true)
		nodeList = []graph.Node{node}
	}
	msg = Message[T]{
		Nodes:   nodeList,
		Payload: ret,
	}
	return
}

func End[S, T any](
	d *TrackerCtx,
	name, description string,
	function func(context.Context, S) (T, error),
	message Message[S],
) (msg Message[T], err error) {
	parentName := stringhandle.GetParentFlow(d.Ctx)
	ctx, node := graph.Task(d.Ctx, d.Track, message.Nodes, name, description)
	ret, err := function(ctx, message.Payload)
	nodeList := []graph.Node{node}
	if err != nil {
		node = graph.End(d.Track, nodeList, name, stringhandle.ErrorFlowName(parentName), true, true)
	} else {
		node = graph.End(d.Track, nodeList, name, stringhandle.EndFlowName(parentName), true, false)
	}
	nodeList = []graph.Node{node}
	msg = Message[T]{
		Nodes:   nodeList,
		Payload: ret,
	}

	return
}

func EndEmpty[S, T any](
	d *TrackerCtx,
	name, description string,
	function func(context.Context, S) error,
	message Message[S],
) (msg Message[error], err error) {
	parentName := stringhandle.GetParentFlow(d.Ctx)
	ctx, node := graph.Task(d.Ctx, d.Track, message.Nodes, name, description)
	err = function(ctx, message.Payload)
	nodeList := []graph.Node{node}
	if err != nil {
		node = graph.End(d.Track, nodeList, name, stringhandle.ErrorFlowName(parentName), true, true)
	} else {
		node = graph.End(d.Track, nodeList, name, stringhandle.EndFlowName(parentName), true, false)
	}
	nodeList = []graph.Node{node}
	msg = Message[error]{
		Nodes:   nodeList,
		Payload: err,
	}

	return
}

func RepackMessage[T any](payload T, nodesVar ...[]graph.Node) Message[T] {
	nodes := []graph.Node{}
	for _, item := range nodesVar {
		nodes = append(nodes, item...)
	}
	return Message[T]{Payload: payload, Nodes: nodes}
}

func UnpackMessage[T any](message Message[T]) (T, []graph.Node) {
	return message.Payload, message.Nodes
}
