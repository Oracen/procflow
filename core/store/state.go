package store

import (
	"sync"
)

var StateManager = CreateNewStateManager()

type Exporter interface {
	ExportRun(string)
}

type stateManager struct {
	useGlobalState bool
	trackState     bool
	lock           sync.Mutex
	exporters      map[string]Exporter
}

func (s *stateManager) EnableUseGlobalState() {
	s.useGlobalState = true
	s.trackState = true
}

func (s *stateManager) EnableTrackState() {
	s.trackState = true
}

func (s *stateManager) UseGlobalState() bool {
	return s.useGlobalState
}

func (s *stateManager) TrackState() bool {
	return s.trackState
}

func (s *stateManager) GetLock() *sync.Mutex {
	return &s.lock
}

func (s *stateManager) AddExporter(name string, exporter Exporter) {
	s.lock.Lock()
	defer s.lock.Unlock()
	_, ok := s.exporters[name]
	if !ok {
		s.exporters[name] = exporter
	}
}

func (s *stateManager) RunExport(filepath string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	for _, value := range s.exporters {
		value.ExportRun(filepath)
	}
}

func CreateNewStateManager() stateManager {
	return stateManager{false, false, sync.Mutex{}, map[string]Exporter{}}
}

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
