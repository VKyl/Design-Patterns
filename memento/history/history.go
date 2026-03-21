package history

type Memento[T any] interface {
	GetState() T
	Restore()
}

type History[T any] struct {
	mementos []Memento[T]
}

func NewHistory[T any]() *History[T] {
	return &History[T]{
		mementos: make([]Memento[T], 0),
	}
}

func (h *History[T]) AddSnapshot(m Memento[T]) {
	h.mementos = append(h.mementos, m)
}

func (h *History[T]) Peek() Memento[T] {
	if len(h.mementos) == 0 {
		return nil
	}
	return h.mementos[len(h.mementos)-1]
}

func (h *History[T]) Pop() Memento[T] {
	if len(h.mementos) == 0 {
		return nil
	}
	m := h.mementos[len(h.mementos)-1]
	h.mementos = h.mementos[:len(h.mementos)-1]
	return m
}

func (h *History[T]) Clear() {
	h.mementos = nil
}

func (h *History[T]) IsEmpty() bool {
	return len(h.mementos) == 0
}