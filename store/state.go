package store

import "sync"

type State[T any] interface {
	addObject(*T)
	getState() []T
}

type globalState[T any] struct {
	objects []*T

	mu sync.Mutex
	wg sync.WaitGroup
}

func createEmptyGlobalState[T any]() (state globalState[T]) {
	return globalState[T]{objects: []*T{}}
}

func createGlobalSingleton[T any](singletonPtr *globalState[T], safetyLock *sync.Mutex) (state *globalState[T]) {
	safetyLock.Lock()
	defer safetyLock.Unlock()
	if singletonPtr == nil {
		state := createEmptyGlobalState[T]()
		singletonPtr = &state
	}
	return singletonPtr
}

func (g *globalState[T]) addObject(obj *T) {
	g.wg.Add(1)
	defer g.wg.Done()
	g.mu.Lock()
	defer g.mu.Unlock()
	g.objects = append(g.objects, obj)
}

func (g *globalState[T]) getState() (state []T) {
	g.wg.Wait()
	g.mu.Lock()
	defer g.mu.Unlock()
	state = []T{}
	for _, item := range g.objects {
		state = append(state, *item)
	}
	return
}
