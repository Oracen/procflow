package tracker

import (
	"fmt"
	"sync"

	"github.com/Oracen/procflow/core/collections"
	"github.com/Oracen/procflow/core/topo"
)

type Node[T comparable] struct {
	Data T
	// previous T
}

type Tracker[T comparable] interface {
	StartFlow(data T) Node[T]
	AddNode(inputs []Node[T], data T) Node[T]
	EndFlow(inputs []Node[T], data T) Node[T]
	CloseTrace() bool
}

type BasicTracker[S comparable] struct {
	traceClosed    bool
	Collector      *collections.BasicCollectorAdapter[S]
	NameParentNode string
	wg             *sync.WaitGroup
}

func RegisterBasicTracker[S comparable](collector *collections.BasicCollector[S], parentName string) BasicTracker[S] {
	wg := sync.WaitGroup{}
	adapter := collections.CreateNewBasicAdapter(collector)
	return BasicTracker[S]{traceClosed: false, Collector: &adapter, NameParentNode: parentName, wg: &wg}
}

func (b *BasicTracker[S]) StartFlow(data S) Node[S] {
	b.Collector.AddRelationship(data)
	return Node[S]{data}
}

func (b *BasicTracker[S]) AddNode(inputs []Node[S], data S) Node[S] {
	for range inputs {
		b.Collector.AddRelationship(data)
	}

	return Node[S]{data}
}

func (b *BasicTracker[S]) EndFlow(inputs []Node[S], data S) Node[S] {
	for range inputs {
		b.Collector.AddRelationship(data)
	}
	return Node[S]{data}
}

func (b *BasicTracker[S]) CloseTrace() bool {
	b.wg.Wait()
	b.traceClosed = true
	return b.traceClosed
}

type GraphConstructor[S, T comparable] struct {
	Name     string
	Vertex   topo.Vertex[S]
	EdgeData T
}

type GraphTracker[S comparable, T comparable] struct {
	traceClosed    bool
	Collector      *collections.GraphCollectorAdapter[S, T]
	NameParentNode string
	wg             *sync.WaitGroup
}

func RegisterGraphTracker[S comparable, T comparable](
	collector *collections.GraphCollector[S, T],
	parentNode string,
) GraphTracker[S, T] {
	wg := sync.WaitGroup{}
	adapter := collections.CreateNewGraphAdapter(collector)
	collector.AddTask()
	return GraphTracker[S, T]{
		traceClosed:    false,
		Collector:      &adapter,
		NameParentNode: parentNode,
		wg:             &wg,
	}
}

func (g *GraphTracker[S, T]) handleAddRelationship(inner collections.GraphConstructorInner[S, T]) {
	if len(g.Collector.Object.Errors) > 0 {
		return
	}
	err := g.Collector.AddRelationship(inner)
	if err != nil {
		g.Collector.Object.AppendError(inner.VertexName, fmt.Sprint(err))
	}
}

func (g *GraphTracker[S, T]) StartFlow(data GraphConstructor[S, T]) Node[GraphConstructor[S, T]] {
	var emptyEdge topo.Edge[T]

	inner := collections.GraphConstructorInner[S, T]{
		VertexName: data.Name,
		Vertex:     data.Vertex,
		EdgeName:   "",
		Edge:       emptyEdge,
	}
	g.handleAddRelationship(inner)
	return Node[GraphConstructor[S, T]]{data}
}

func constructGraphInner[S, T comparable](new, old GraphConstructor[S, T]) collections.GraphConstructorInner[S, T] {
	return collections.GraphConstructorInner[S, T]{
		VertexName: new.Name,
		Vertex:     new.Vertex,
		EdgeName:   fmt.Sprintf("%s:%s", old.Name, new.Name),
		Edge: topo.Edge[T]{
			VertexFrom: old.Name,
			VertexTo:   new.Name,
			Data:       new.EdgeData,
		},
	}
}

func (g *GraphTracker[S, T]) AddNode(inputs []Node[GraphConstructor[S, T]], data GraphConstructor[S, T]) Node[GraphConstructor[S, T]] {
	for _, item := range inputs {
		inner := constructGraphInner(data, item.Data)
		g.handleAddRelationship(inner)
	}
	return Node[GraphConstructor[S, T]]{data}
}

func (g *GraphTracker[S, T]) EndFlow(inputs []Node[GraphConstructor[S, T]], data GraphConstructor[S, T]) Node[GraphConstructor[S, T]] {
	for _, item := range inputs {
		inner := constructGraphInner(data, item.Data)
		g.handleAddRelationship(inner)
	}
	return Node[GraphConstructor[S, T]]{data}
}

func (g *GraphTracker[S, T]) CloseTrace() bool {
	g.wg.Wait()
	g.Collector.Object.FinishTask()
	g.traceClosed = true
	return g.traceClosed
}
