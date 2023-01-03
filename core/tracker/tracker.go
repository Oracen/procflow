package tracker

type Node[T comparable] struct {
	nodeSite string
	data     T
}

type Tracker[T comparable] struct {
}

func RegisterTracker[T comparable]() Tracker[T] {
	return Tracker[T]{}
}

func (t *Tracker[T]) StartFlow(name string, data T) Node[T] {
	return Node[T]{name, data}
}

func (t *Tracker[T]) AddNode(name string, inputs []Node[T], data T) Node[T] {
	return Node[T]{name, data}
}

func (t *Tracker[T]) EndFlow(name string, inputs []Node[T], data T) {
}
