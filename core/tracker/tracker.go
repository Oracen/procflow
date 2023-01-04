package tracker

import (
	"sync"

	"github.com/Oracen/procflow/core/collection"
	"github.com/Oracen/procflow/core/topo"
)

type Node[T comparable] struct {
	nodeSite string
	data     T
	previous T
}

type Tracker[T comparable] interface {
	StartFlow(name string, data T) Node[T]
	AddNode(name string, inputs []Node[T], data T) Node[T]
	EndFlow(name string, inputs []Node[T], data T)
	CloseTrace() bool
}

type BasicTracker[S comparable, T any] struct {
	traceClosed bool
	collector   *collection.Collector[S, T]
	wg          *sync.WaitGroup
}

func RegisterBasicTracker[S comparable, T any](collector *collection.Collector[S, T]) BasicTracker[S, T] {
	wg := sync.WaitGroup{}
	return BasicTracker[S, T]{traceClosed: false, collector: collector, wg: &wg}
}

func (b *BasicTracker[S, T]) StartFlow(name string, data S) Node[S] {
	var empty S
	b.collector.AddRelationship(data)
	return Node[S]{name, data, empty}
}

func (b *BasicTracker[S, T]) AddNode(name string, inputs []Node[S], data S) Node[S] {
	var empty S
	b.collector.AddRelationship(data)
	return Node[S]{name, data, empty}
}

func (b *BasicTracker[S, T]) EndFlow(name string, inputs []Node[S], data S) {
	b.collector.AddRelationship(data)
}

func (b *BasicTracker[S, T]) CloseTrace() bool {
	b.wg.Wait()
	b.traceClosed = true
	return b.traceClosed
}

type GraphConstructor[S, T comparable] struct {
	Name   string
	Vertex topo.Vertex[S]
	Edge   topo.Edge[T]
}

type GraphCollectable[S, T comparable] struct {
	Graph topo.Graph[S, T]
}

func (g *GraphCollectable[S, T]) Add(item GraphConstructor[S, T]) error {
	err := g.Graph.AddNewVertex(item.Name, item.Vertex)
	if err != nil {
		return err
	}
	err = g.Graph.AddNewEdge(item.Name, item.Edge)
	if err != nil {
		return err
	}
	return nil
}

func (g *GraphCollectable[S, T]) Union(other GraphCollectable[S, T]) (merged GraphCollectable[S, T], err error) {
	collection, err := topo.MergeGraphs(g.Graph, other.Graph)
	if err != nil {
		return
	}
	return GraphCollectable[S, T]{Graph: collection}, nil
}

type GraphCollector[S comparable, T comparable] struct {
	object *GraphCollectable[S, T]
	wg     *sync.WaitGroup
	mu     *sync.Mutex
}

func CreateNewGraphCollector[S comparable, T comparable](object *GraphCollectable[S, T]) GraphCollector[S, T] {
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	return GraphCollector[S, T]{object: object, wg: &wg, mu: &mu}
}

func (c *GraphCollector[S, T]) AddRelationship(obj GraphConstructor[S, T]) (err error) {
	c.wg.Add(1)
	defer c.wg.Done()
	c.mu.Lock()
	defer c.mu.Unlock()
	deref := *(c.object)
	err = deref.Add(obj)
	if err == nil {
		c.object = &deref
	}
	return
}

func (c *GraphCollector[S, T]) UnionRelationships(obj GraphCollectable[S, T]) (merged GraphCollectable[S, T], err error) {
	c.wg.Wait()
	c.mu.Lock()
	defer c.mu.Unlock()
	deref := *(c.object)
	return deref.Union(obj)
}

type GraphTracker[S comparable, T comparable] struct {
	traceClosed bool
	collector   *GraphCollector[S, T]
	wg          *sync.WaitGroup
}

func RegisterGraphTracker[S comparable, T comparable](collector *GraphCollector[S, T]) GraphTracker[S, T] {
	wg := sync.WaitGroup{}
	return GraphTracker[S, T]{traceClosed: false, collector: collector, wg: &wg}
}

func (g *GraphTracker[S, T]) StartFlow(name string, data GraphConstructor[S, T]) Node[GraphConstructor[S, T]] {
	var empty GraphConstructor[S, T]
	g.collector.AddRelationship(data)
	return Node[GraphConstructor[S, T]]{name, data, empty}
}

func (g *GraphTracker[S, T]) AddNode(name string, inputs []Node[GraphConstructor[S, T]], data GraphConstructor[S, T]) Node[GraphConstructor[S, T]] {
	var empty GraphConstructor[S, T]
	return Node[GraphConstructor[S, T]]{name, data, empty}
}

func (g *GraphTracker[S, T]) EndFlow(name string, inputs []Node[GraphConstructor[S, T]], data GraphConstructor[S, T]) {
}

func (g *GraphTracker[S, T]) CloseTrace() bool {
	g.wg.Wait()
	g.traceClosed = true
	return g.traceClosed
}
