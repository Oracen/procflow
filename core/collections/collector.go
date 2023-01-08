package collections

import (
	"sync"

	"github.com/Oracen/procflow/core/topo"
)

type Collection[S, T any] interface {
	Add(S) error
	Union(T) (T, error)
	AddTask()
	FinishTask()
	WaitForFinish()
}

type BasicCollector[T comparable] struct {
	Array  []T
	Errors map[string]string
	Wg     *sync.WaitGroup
}

func createNewBasicCollector[T comparable](array []T) BasicCollector[T] {
	errors := map[string]string{}
	wg := sync.WaitGroup{}
	return BasicCollector[T]{Array: array, Errors: errors, Wg: &wg}
}

func CreateNewBasicCollector[T comparable]() BasicCollector[T] {
	return createNewBasicCollector([]T{})
}

func (m *BasicCollector[T]) Add(item T) error {
	array := append(m.Array, item)
	m.Array = array
	return nil
}

func (m *BasicCollector[T]) Union(other BasicCollector[T]) (BasicCollector[T], error) {
	collection := append(m.Array, other.Array...)

	return createNewBasicCollector(collection), nil
}

func (b *BasicCollector[T]) AddTask() {

}

func (b *BasicCollector[T]) FinishTask() {

}
func (b *BasicCollector[T]) WaitForFinish() {

}

type BasicCollectorAdapter[S comparable] struct {
	Object *BasicCollector[S]
	wg     *sync.WaitGroup
	mu     *sync.Mutex
}

func CreateNewBasicAdapter[S comparable](object *BasicCollector[S]) BasicCollectorAdapter[S] {
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	return BasicCollectorAdapter[S]{Object: object, wg: &wg, mu: &mu}
}

func (b *BasicCollectorAdapter[S]) AddRelationship(obj S) (err error) {
	b.wg.Add(1)
	defer b.wg.Done()
	b.mu.Lock()
	defer b.mu.Unlock()

	err = b.Object.Add(obj)
	return
}

func (b *BasicCollectorAdapter[S]) UnionRelationships(obj BasicCollector[S]) (merged BasicCollector[S], err error) {
	b.wg.Wait()
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.Object.Union(obj)
}

type GraphConstructorInner[S, T comparable] struct {
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

func (g *GraphCollector[S, T]) AppendError(site, errorMsg string) {
	g.Errors[site] = errorMsg
}

func createNewGraphCollector[S, T comparable](g topo.Graph[S, T]) GraphCollector[S, T] {
	errors := map[string]string{}
	wg := sync.WaitGroup{}
	return GraphCollector[S, T]{Graph: g, Errors: errors, Wg: &wg}
}

func CreateNewGraphCollector[S, T comparable]() GraphCollector[S, T] {
	return createNewGraphCollector(topo.CreateNewGraph[S, T]())
}

func (g *GraphCollector[S, T]) Add(item GraphConstructorInner[S, T]) error {
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
	return createNewGraphCollector(collection), nil
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

func CreateNewGraphAdapter[S comparable, T comparable](object *GraphCollector[S, T]) GraphCollectorAdapter[S, T] {
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	return GraphCollectorAdapter[S, T]{Object: object, wg: &wg, mu: &mu}
}

func (c *GraphCollectorAdapter[S, T]) AddRelationship(obj GraphConstructorInner[S, T]) (err error) {
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
