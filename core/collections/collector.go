package collections

import (
	"sync"
)

type Collection[S, T any] interface {
	Add(S) error
	Union(T) (T, error)
	AddTask()
	FinishTask()
	WaitForFinish()
}

type BasicCollectorAdapter[S comparable] struct {
	Object *BasicCollector[S]
	wg     *sync.WaitGroup
	mu     *sync.Mutex
}

func CreateNewCollectorAdapter[S comparable](object *BasicCollector[S]) BasicCollectorAdapter[S] {
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

type BasicCollector[T comparable] struct {
	Collection []T
}

func (m *BasicCollector[T]) Add(item T) error {
	array := append(m.Collection, item)
	m.Collection = array
	return nil
}

func (m *BasicCollector[T]) Union(other BasicCollector[T]) (BasicCollector[T], error) {
	collection := append(m.Collection, other.Collection...)
	return BasicCollector[T]{collection}, nil
}

func (b *BasicCollector[T]) AddTask() {

}

func (b *BasicCollector[T]) FinishTask() {

}
func (b *BasicCollector[T]) WaitForFinish() {

}
