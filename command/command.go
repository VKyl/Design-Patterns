package command

import "github.com/VKyl/Design-Patterns/util"

type Editor[T util.Comperable[T]] interface {
	Reset()
	Save()
}

type ResetCommand[T util.Comperable[T]] struct {
	editor Editor[T]
}

func NewResetCommand[T util.Comperable[T]](editor Editor[T]) *ResetCommand[T] {
	return &ResetCommand[T]{editor: editor}
}

func (c *ResetCommand[T]) Execute() {
	c.editor.Reset()
}

type SaveCommand[T util.Comperable[T]] struct {
	editor Editor[T]
}

func NewSaveCommand[T util.Comperable[T]](editor Editor[T]) *SaveCommand[T] {
	return &SaveCommand[T]{editor: editor}
}

func (c *SaveCommand[T]) Execute() {
	c.editor.Save()
}
