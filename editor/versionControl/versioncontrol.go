package versioncontrol

import (
	"github.com/VKyl/Design-Patterns/memento"
	"github.com/VKyl/Design-Patterns/memento/history"
	"github.com/VKyl/Design-Patterns/util"
)

type Editor[T util.Comperable[T], Snapshot any] interface {
	Value() T
	SetValue(value T)
	MakeSnapshot() Snapshot
	RestoreState(state Snapshot)
}

type VersionControl[T util.Comperable[T], Snapshot any] struct {
	editor         Editor[T, Snapshot]
	versionHistory *history.History[Snapshot]
	changesHistory *history.History[Snapshot]
}

func NewVersionControl[T util.Comperable[T], Snapshot any](editor Editor[T, Snapshot]) *VersionControl[T, Snapshot] {
	return &VersionControl[T, Snapshot]{
		editor:         editor,
		versionHistory: history.NewHistory[Snapshot](),
		changesHistory: history.NewHistory[Snapshot](),
	}
}

func (vc *VersionControl[T, Snapshot]) AddChange(v Snapshot) {
	vc.changesHistory.AddSnapshot(memento.NewMemento(vc.editor, v))
}

func (vc *VersionControl[T, Snapshot]) IsEdited() bool {
	return !vc.changesHistory.IsEmpty()
}

func (vc *VersionControl[T, Snapshot]) Reset() {
	state := vc.versionHistory.Peek()
	if state != nil {
		state.Restore()
	}
	vc.changesHistory.Clear()
}

func (vc *VersionControl[T, Snapshot]) Save() {
	if !vc.changesHistory.IsEmpty() {
		memento := memento.NewMemento(vc.editor, vc.editor.MakeSnapshot())
		vc.versionHistory.AddSnapshot(memento)
	}
	vc.changesHistory.Clear()
}

func (vc *VersionControl[T, Snapshot]) Undo() {
	if vc.changesHistory.IsEmpty() {
		vc.undoVersion()
	} else {
		vc.undoChange()
	}
}

func (vc *VersionControl[T, Snapshot]) undoChange() {
	state := vc.changesHistory.Pop()
	if state != nil {
		state.Restore()
	}
}

func (vc *VersionControl[T, Snapshot]) undoVersion() {
	vc.versionHistory.Pop()
	state := vc.versionHistory.Peek()
	if state != nil {
		state.Restore()
	}
}
