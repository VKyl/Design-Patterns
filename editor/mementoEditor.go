package editor

import (
	"github.com/VKyl/Design-Patterns/util"
)

type Snapshot[T util.Comperable[T]] struct {
	value    T
	font     string
	fontSize int
}

func (s Snapshot[T]) IsEqual(other Snapshot[T]) bool {
	return s.value.Equals(other.value) && s.font == other.font && s.fontSize == other.fontSize
}

type MementoEditor[T util.Comperable[T]] struct {
	value    T
	font     string
	fontSize int
}

func NewMementoEditor[T util.Comperable[T]]() *MementoEditor[T] {
	return &MementoEditor[T]{}
}

func (e *MementoEditor[T]) Value() T {
	return e.value
}

func (e *MementoEditor[T]) SetValue(value T) {
	e.value = value
}

func (e *MementoEditor[T]) SetFont(font string) {
	e.font = font
}

func (e *MementoEditor[T]) SetFontSize(size int) {
	e.fontSize = size
}

func (e *MementoEditor[T]) MakeSnapshot() Snapshot[T] {
	return Snapshot[T]{
		value:    e.value,
		font:     e.font,
		fontSize: e.fontSize,
	}
}

func (e *MementoEditor[T]) RestoreState(state Snapshot[T]) {
	e.value = state.value
	e.font = state.font
	e.fontSize = state.fontSize
}
