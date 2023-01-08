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
	EndFlow(inputs []Node[T], data T)
	CloseTrace() bool
}

type BasicTracker[S comparable] struct {
	traceClosed bool
	Collector   *collections.BasicCollectorAdapter[S]
	wg          *sync.WaitGroup
}

func RegisterBasicTracker[S comparable](collector *collections.BasicCollector[S]) BasicTracker[S] {
	wg := sync.WaitGroup{}
	adapter := collections.CreateNewCollectorAdapter(collector)
	return BasicTracker[S]{traceClosed: false, Collector: &adapter, wg: &wg}
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

func (b *BasicTracker[S]) EndFlow(inputs []Node[S], data S) {
	for range inputs {
		b.Collector.AddRelationship(data)
	}
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
type graphConstructorInner[S, T comparable] struct {
	VertexName string
	Vertex     topo.Vertex[S]
	EdgeName   string
	Edge       topo.Edge[T]
}

type GraphCollector[S, T comparable] struct {
	Graph  topo.Graph[S, T]
	Errors map[string]string
	Wg     *sync.WaitGroup
}

func (g *GraphCollector[S, T]) appendError(site, errorMsg string) {
	g.Errors[site] = errorMsg
}

func CreateNewGraphCollection[S, T comparable]() GraphCollector[S, T] {
	errors := map[string]string{}
	wg := sync.WaitGroup{}
	return GraphCollector[S, T]{Graph: topo.CreateNewGraph[S, T](), Errors: errors, Wg: &wg}
}

func (g *GraphCollector[S, T]) Add(item graphConstructorInner[S, T]) error {
	err := g.Graph.AddNewVertex(item.VertexName, item.Vertex)
	if err != nil {
		return err
	}
	if item.EdgeName == "" {
		return nil
	}

	err = g.Graph.AddNewEdge(item.EdgeName, item.Edge)
	if err != nil {
		return err
	}
	return nil
}

func (g *GraphCollector[S, T]) Union(other GraphCollector[S, T]) (merged GraphCollector[S, T], err error) {
	collection, err := topo.MergeGraphs(g.Graph, other.Graph)
	if err != nil {
		return
	}
	return GraphCollector[S, T]{Graph: collection}, nil
}

func (g *GraphCollector[S, T]) AddTask() {
	g.Wg.Add(1)
}

func (g *GraphCollector[S, T]) FinishTask() {
	g.Wg.Done()
}
func (g *GraphCollector[S, T]) WaitForFinish() {
	g.Wg.Wait()
}

type GraphCollectorAdapter[S comparable, T comparable] struct {
	Object *GraphCollector[S, T]
	wg     *sync.WaitGroup
	mu     *sync.Mutex
}

func CreateNewGraphCollector[S comparable, T comparable](object *GraphCollector[S, T]) GraphCollectorAdapter[S, T] {
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	return GraphCollectorAdapter[S, T]{Object: object, wg: &wg, mu: &mu}
}

func (c *GraphCollectorAdapter[S, T]) AddRelationship(obj graphConstructorInner[S, T]) (err error) {
	c.wg.Add(1)
	defer c.wg.Done()
	c.mu.Lock()
	defer c.mu.Unlock()
	deref := *(c.Object)

	err = deref.Add(obj)
	if err == nil {
		c.Object = &deref
	}
	return
}

func (c *GraphCollectorAdapter[S, T]) UnionRelationships(obj GraphCollector[S, T]) (merged GraphCollector[S, T], err error) {
	c.wg.Wait()
	c.mu.Lock()
	defer c.mu.Unlock()
	deref := *(c.Object)
	return deref.Union(obj)
}

type GraphTracker[S comparable, T comparable] struct {
	traceClosed    bool
	Collector      *GraphCollectorAdapter[S, T]
	NameParentNode string
	wg             *sync.WaitGroup
}

func RegisterGraphTracker[S comparable, T comparable](
	collector *GraphCollector[S, T],
	parentNode string,
) GraphTracker[S, T] {
	wg := sync.WaitGroup{}
	adapter := CreateNewGraphCollector(collector)
	collector.AddTask()
	return GraphTracker[S, T]{
		traceClosed:    false,
		Collector:      &adapter,
		NameParentNode: parentNode,
		wg:             &wg,
	}
}

func (g *GraphTracker[S, T]) handleAddRelationship(inner graphConstructorInner[S, T]) {
	if len(g.Collector.Object.Errors) > 0 {
		return
	}
	err := g.Collector.AddRelationship(inner)
	if err != nil {
		g.Collector.Object.appendError(inner.VertexName, fmt.Sprint(err))
	}
}

func (g *GraphTracker[S, T]) StartFlow(data GraphConstructor[S, T]) Node[GraphConstructor[S, T]] {
	var emptyEdge topo.Edge[T]

	inner := graphConstructorInner[S, T]{
		VertexName: data.Name,
		Vertex:     data.Vertex,
		EdgeName:   "",
		Edge:       emptyEdge,
	}
	g.handleAddRelationship(inner)
	return Node[GraphConstructor[S, T]]{data}
}

func constructGraphInner[S, T comparable](new, old GraphConstructor[S, T]) graphConstructorInner[S, T] {
	return graphConstructorInner[S, T]{
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

func (g *GraphTracker[S, T]) EndFlow(inputs []Node[GraphConstructor[S, T]], data GraphConstructor[S, T]) {
	for _, item := range inputs {
		inner := constructGraphInner(data, item.Data)
		g.handleAddRelationship(inner)
	}
}

func (g *GraphTracker[S, T]) CloseTrace() bool {
	g.wg.Wait()
	g.Collector.Object.FinishTask()
	g.traceClosed = true
	return g.traceClosed
}
