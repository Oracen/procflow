package store

import "sync"

type State[T any] interface {
	addObject(*T)
	getState() []T
}

type GlobalState[T any] struct {
	objects []*T

	mu sync.Mutex
	wg sync.WaitGroup
}

func createEmptyGlobalState[T any]() (state GlobalState[T]) {
	return GlobalState[T]{objects: []*T{}}
}

func CreateGlobalSingleton[T any](singletonPtr *GlobalState[T], safetyLock *sync.Mutex) (state *GlobalState[T]) {
	safetyLock.Lock()
	defer safetyLock.Unlock()
	if singletonPtr == nil {
		state := createEmptyGlobalState[T]()
		singletonPtr = &state
	}
	return singletonPtr
}

func (g *GlobalState[T]) AddObject(obj *T) {
	g.wg.Add(1)
	defer g.wg.Done()
	g.mu.Lock()
	defer g.mu.Unlock()
	g.objects = append(g.objects, obj)
}

func (g *GlobalState[T]) GetState() (state []T) {
	g.wg.Wait()
	g.mu.Lock()
	defer g.mu.Unlock()
	state = []T{}
	for _, item := range g.objects {
		state = append(state, *item)
	}
	return
}
