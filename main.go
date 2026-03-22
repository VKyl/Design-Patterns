package main

import (
	"fmt"

	"github.com/VKyl/Design-Patterns/command"
	"github.com/VKyl/Design-Patterns/controls"
	"github.com/VKyl/Design-Patterns/editor"
)

type Text string

func (t Text) Equals(other Text) bool {
	return t == other
}

type EditedValue[T any] interface {
	Value() T
	IsEdited() bool
}

func DisplayValue[T any](v EditedValue[T]) {
	fmt.Printf("Is Edited: %v, Current value: %v\n", v.IsEdited(), v.Value())
}

func CommandEditorExample() {
	e := editor.NewEditor[Text]()

	resetButton := controls.NewButton("Reset", command.NewResetCommand[Text](e))
	saveButton := controls.NewButton("Save", command.NewSaveCommand[Text](e))

	e.SetValue("Hello")
	saveButton.Click()
	DisplayValue(e)
	e.SetValue("Hello, World!")
	DisplayValue(e)
	resetButton.Click()
	DisplayValue(e)
}

func VersionControlExample() {
	vc := editor.NewVersionControlledEditor[Text]()
	resetButton := controls.NewButton("Reset", command.NewResetCommand[Text](vc.VersionControl()))
	saveButton := controls.NewButton("Save", command.NewSaveCommand[Text](vc.VersionControl()))
	undoButton := controls.NewButton("Undo", command.NewUndoCommand[Text](vc.VersionControl()))

	vc.SetValue("Hello")
	saveButton.Click()
	DisplayValue(vc)

	vc.SetValue("Hello, World!")
	saveButton.Click()
	DisplayValue(vc)

	resetButton.Click()
	DisplayValue(vc)

	vc.SetValue("Hello Again")
	DisplayValue(vc)

	vc.SetValue("Hello again 2")
	DisplayValue(vc)

	undoButton.Click()
	DisplayValue(vc)

	undoButton.Click()
	DisplayValue(vc)
}

func main() {
	VersionControlExample()
}
