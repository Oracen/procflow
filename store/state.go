package store

import "sync"

var (
	lock           = &sync.Mutex{}
	singletonState = &globalState[Collector]{}
)

type State[T any] interface {
	addObject(*T)
	getState() []T
}

type globalState[T any] struct {
	mu      sync.Mutex
	objects []*T
}

func CreateGlobalCollectorSingleton[T any]() State[Collector] {
	return &globalState[Collector]{}
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
