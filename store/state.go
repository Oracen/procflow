package store

import "sync"

type State[T any] interface {
	addObject(*T)
	getState() []T
}

type globalState[T any] struct {
	mu      sync.Mutex
	objects []*T
}

func createGlobalSingleton[T any](singletonPtr *globalState[T], mu *sync.Mutex) (state *globalState[T]) {
	mu.Lock()
	defer mu.Unlock()
	if singletonPtr == nil {
		singletonPtr = &globalState[T]{}
	}
	return singletonPtr
}

func (g *globalState[T]) addObject(obj *T) {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.objects = append(g.objects, obj)
}

func (g *globalState[T]) getState() (state []T) {
	state = []T{}
	for _, item := range g.objects {
		state = append(state, *item)
	}
	return
}
