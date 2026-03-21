package editorui

import (
	"fmt"

	"github.com/VKyl/Design-Patterns/editorUI/controls"
	"github.com/VKyl/Design-Patterns/util"
)

type Command interface {
	Execute()
}

type Editor[T util.Comperable[T]] interface {
	IsEdited() bool
	Value() T
	SetValue(value T)
}

type EditorUI[T util.Comperable[T]] struct {
	buttons map[string]*controls.Button
	editor  Editor[T]
}

func NewEditorUI[T util.Comperable[T]](editor Editor[T]) *EditorUI[T] {
	return &EditorUI[T]{editor: editor}
}

func (ui *EditorUI[T]) DisplayValue() {
	fmt.Printf("IsEdited: %t; CurrentValue: %v\n",
		ui.editor.IsEdited(),
		ui.editor.Value(),
	)
}

func (ui *EditorUI[T]) SetValue(value T) {
	fmt.Printf("Setting value to %v\n", value)
	ui.editor.SetValue(value)
}

func (ui *EditorUI[T]) ClickButton(buttonName string) {
	if button, exists := ui.buttons[buttonName]; exists {
		button.Click()
	}
}

func (ui *EditorUI[T]) AddButton(name string, button *controls.Button) {
	ui.buttons[name] = button
}
