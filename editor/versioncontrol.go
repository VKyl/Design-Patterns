package editor

import (
	"github.com/VKyl/Design-Patterns/memento"
	"github.com/VKyl/Design-Patterns/memento/history"
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

type VersionControl[T util.Comperable[T]] struct {
	editor         *MementoEditor[T]
	versionHistory *history.History[Snapshot[T]]
	changesHistory *history.History[Snapshot[T]]
}

func NewVersionControl[T util.Comperable[T]](editor *MementoEditor[T]) *VersionControl[T] {
	return &VersionControl[T]{
		editor:         editor,
		versionHistory: history.NewHistory[Snapshot[T]](),
		changesHistory: history.NewHistory[Snapshot[T]](),
	}
}

func (vc *VersionControl[T]) Value() T {
	return vc.editor.Value()
}

func (vc *VersionControl[T]) IsEdited() bool {
	return !vc.changesHistory.IsEmpty()
}

func (vc *VersionControl[T]) SetValue(value T) {
	memento := memento.NewMemento(vc.editor, vc.editor.MakeSnapshot())
	vc.changesHistory.AddSnapshot(memento)
	vc.editor.SetValue(value)
}

func (vc *VersionControl[T]) Reset() {
	state := vc.versionHistory.Peek()
	if state != nil {
		state.Restore()
	}
	vc.changesHistory.Clear()
}

func (vc *VersionControl[T]) Save() {
	state := vc.changesHistory.Peek()
	if state != nil {
		memento := memento.NewMemento(vc.editor, vc.editor.MakeSnapshot())
		vc.versionHistory.AddSnapshot(memento)
	}
	vc.changesHistory.Clear()
}

func (vc *VersionControl[T]) Undo() {
	state := vc.changesHistory.Pop()
	if state != nil {
		state.Restore()
	}
}
