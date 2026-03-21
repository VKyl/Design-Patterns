package command

import (
	"github.com/VKyl/Design-Patterns/util"
)

type ResetEditor[T util.Comperable[T]] interface {
	Reset()
}

type ResetCommand[T util.Comperable[T]] struct {
	editor ResetEditor[T]
}

func NewResetCommand[T util.Comperable[T]](editor ResetEditor[T]) *ResetCommand[T] {
	return &ResetCommand[T]{editor: editor}
}

func (c *ResetCommand[T]) Execute() {
	c.editor.Reset()
}

type SaveEditor[T util.Comperable[T]] interface {
	Save()
}

type SaveCommand[T util.Comperable[T]] struct {
	editor SaveEditor[T]
}

func NewSaveCommand[T util.Comperable[T]](editor SaveEditor[T]) *SaveCommand[T] {
	return &SaveCommand[T]{editor: editor}
}

func (c *SaveCommand[T]) Execute() {
	c.editor.Save()
}

type UndoEditor[T util.Comperable[T]] interface {
	Undo()
}

type UndoCommand[T util.Comperable[T]] struct {
	editor UndoEditor[T]
}

func NewUndoCommand[T util.Comperable[T]](editor UndoEditor[T]) *UndoCommand[T] {
	return &UndoCommand[T]{editor: editor}
}

func (c *UndoCommand[T]) Execute() {
	c.editor.Undo()
}
