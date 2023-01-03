package collection

import (
	"sync"
)

type Collectable[S, T any] interface {
	AddRelationship(S) error
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
	return
}
