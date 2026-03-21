package editor

import "github.com/VKyl/Design-Patterns/util"

type Editor[T util.Comperable[T]] struct {
	previousValue T
	currentValue  T
}

func NewEditor[T util.Comperable[T]]() *Editor[T] {
	return &Editor[T]{}
}

func (e *Editor[T]) Value() T {
	return e.currentValue
}

func (e *Editor[T]) IsEdited() bool {
	return e.currentValue.Equals(e.previousValue) == false
}

func (e *Editor[T]) SetValue(value T) {
	e.previousValue = e.currentValue
	e.currentValue = value
}

func (e *Editor[T]) Reset() {
	e.currentValue = e.previousValue
}

func (e *Editor[T]) Save() {
	e.previousValue = e.currentValue
}
