package basic

import (
	"context"

	"github.com/Oracen/procflow/annotation/basic"
	"github.com/Oracen/procflow/core/stringhandle"
)

type Message[T any] struct {
	Nodes   []basic.Node
	Payload T
}

type TrackerCtx struct {
	Ctx   context.Context
	Track *basic.Tracker
}

func CreateNewTrackerCtx(
	ctx context.Context,
) TrackerCtx {
	tracker := basic.RegisterTracker(ctx)
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
	node := basic.Start(d.Track, name, stringhandle.StartFlowName(parentName))
	ctx, node := basic.Task(d.Ctx, d.Track, []basic.Node{node}, name, description)
	ret, err := function(ctx, data)
	nodeList := []basic.Node{node}
	if err != nil {
		node = basic.End(d.Track, nodeList, name, err.Error()[:50], true)
		nodeList = []basic.Node{node}
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
	node := basic.Start(d.Track, name, stringhandle.StartFlowName(parentName))
	ctx, node := basic.Task(d.Ctx, d.Track, []basic.Node{node}, name, description)
	ret, err := function(ctx)
	nodeList := []basic.Node{node}
	if err != nil {
		node = basic.End(d.Track, nodeList, name, err.Error()[:50], true)
		nodeList = []basic.Node{node}
	}
	msg = Message[T]{
		Nodes:   nodeList,
		Payload: ret,
	}
	return
}

func Task[S, T any](
	d *TrackerCtx,
	name, description string,
	function func(context.Context, S) (T, error),
	message Message[S],
) (msg Message[T], err error) {
	ctx, node := basic.Task(d.Ctx, d.Track, message.Nodes, name, description)
	ret, err := function(ctx, message.Payload)
	nodeList := []basic.Node{node}
	if err != nil {
		node = basic.End(d.Track, nodeList, name, err.Error()[:50], true)
		nodeList = []basic.Node{node}
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
	ctx, node := basic.Task(d.Ctx, d.Track, message.Nodes, name, description)
	ret, err := function(ctx, message.Payload)
	nodeList := []basic.Node{node}
	if err != nil {
		node = basic.End(d.Track, nodeList, name, err.Error()[:50], true)
	} else {
		node = basic.End(d.Track, nodeList, name, stringhandle.EndFlowName(parentName), false)
	}
	nodeList = []basic.Node{node}
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
	ctx, node := basic.Task(d.Ctx, d.Track, message.Nodes, name, description)
	err = function(ctx, message.Payload)
	nodeList := []basic.Node{node}
	if err != nil {
		node = basic.End(d.Track, nodeList, name, err.Error()[:50], true)
	} else {
		node = basic.End(d.Track, nodeList, name, stringhandle.EndFlowName(parentName), false)
	}
	nodeList = []basic.Node{node}
	msg = Message[error]{
		Nodes:   nodeList,
		Payload: err,
	}

	return
}

func RepackMessage[T any](payload T, nodesVar ...[]basic.Node) Message[T] {
	nodes := []basic.Node{}
	for _, item := range nodesVar {
		nodes = append(nodes, item...)
	}
	return Message[T]{Payload: payload, Nodes: nodes}
}

func UnpackMessage[T any](message Message[T]) (T, []basic.Node) {
	return message.Payload, message.Nodes
}
