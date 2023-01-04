package collection

import (
	"sync"
)

type Collectable[S, T any] interface {
	Add(S) error
	Union(T) (T, error)
}

type Collector[S, T any] struct {
	object *Collectable[S, T]
	wg     *sync.WaitGroup
	mu     *sync.Mutex
}

func CreateNewCollector[S, T any](object Collectable[S, T]) Collector[S, T] {
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	return Collector[S, T]{object: &object, wg: &wg, mu: &mu}
}

func (c *Collector[S, T]) AddRelationship(obj S) (err error) {
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

func (c *Collector[S, T]) UnionRelationships(obj T) (merged T, err error) {
	c.wg.Wait()
	c.mu.Lock()
	defer c.mu.Unlock()
	deref := *(c.object)
	return deref.Union(obj)
}

// Used as a placeholder for development
type BasicCollectable[T comparable] struct {
	Collection []T
}

func (m *BasicCollectable[T]) Add(item T) error {
	m.Collection = append(m.Collection, item)
	return nil
}

func (m *BasicCollectable[T]) Union(other BasicCollectable[T]) (BasicCollectable[T], error) {
	collection := append(m.Collection, other.Collection...)
	return BasicCollectable[T]{collection}, nil
}
