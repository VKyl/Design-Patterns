package memento

type Originator[T any] interface {
	RestoreState(state T)
}

type Memento[T any, O Originator[T]] struct {
	originator O
	state      T
}

func NewMemento[T any, O Originator[T]](originator O, state T) *Memento[T, O] {
	return &Memento[T, O]{originator: originator, state: state}
}

func (m *Memento[T, O]) GetState() T {
	return m.state
}

func (m *Memento[T, O]) Restore() {
	m.originator.RestoreState(m.state)
}
