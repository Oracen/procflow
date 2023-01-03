package tracker

import "github.com/Oracen/procflow/core/collection"

type Node[T comparable] struct {
	nodeSite string
	data     T
}

type Tracker[S any, T comparable] struct {
	collector collection.Collector[T, S]
}

func RegisterTracker[S, T comparable]() Tracker[S, T] {
	return Tracker[S, T]{}
}

func (t *Tracker[S, T]) StartFlow(name string, data T) Node[T] {
	return Node[T]{name, data}
}

func (t *Tracker[S, T]) AddNode(name string, inputs []Node[T], data T) Node[T] {
	return Node[T]{name, data}
}

func (t *Tracker[S, T]) EndFlow(name string, inputs []Node[T], data T) {
}
