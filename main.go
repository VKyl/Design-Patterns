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

func CommandEditorExample() {
	e := editor.NewEditor[Text]()

	resetButton := controls.NewButton("Reset", command.NewResetCommand[Text](e))
	saveButton := controls.NewButton("Save", command.NewSaveCommand[Text](e))

	e.SetValue("Hello")
	saveButton.Click()
	fmt.Printf("Is Edited: %v, Current value: %s\n", e.IsEdited(), e.Value())
	e.SetValue("Hello, World!")
	fmt.Printf("Is Edited: %v, Current value: %s\n", e.IsEdited(), e.Value())
	resetButton.Click()
	fmt.Printf("Is Edited: %v, Current value: %s\n", e.IsEdited(), e.Value())
}

func VersionControlExample() {
	vc := editor.NewVersionControlledEditor[Text]()
	resetButton := controls.NewButton("Reset", command.NewResetCommand[Text](vc.VersionControl()))
	saveButton := controls.NewButton("Save", command.NewSaveCommand[Text](vc.VersionControl()))
	undoButton := controls.NewButton("Undo", command.NewUndoCommand[Text](vc.VersionControl()))

	vc.SetValue("Hello")
	saveButton.Click()
	fmt.Printf("Is Edited: %v, Current value: %s\n", vc.IsEdited(), vc.Value())

	vc.SetValue("Hello, World!")
	saveButton.Click()
	fmt.Printf("Is Edited: %v, Current value: %s\n", vc.IsEdited(), vc.Value())

	resetButton.Click()
	fmt.Printf("Is Edited: %v, Current value: %s\n", vc.IsEdited(), vc.Value())

	vc.SetValue("Hello Again")
	fmt.Printf("Is Edited: %v, Current value: %s\n", vc.IsEdited(), vc.Value())

	vc.SetValue("Hello again 2")
	fmt.Printf("Is Edited: %v, Current value: %s\n", vc.IsEdited(), vc.Value())

	undoButton.Click()
	fmt.Printf("Is Edited: %v, Current value: %s\n", vc.IsEdited(), vc.Value())

	undoButton.Click()
	fmt.Printf("Is Edited: %v, Current value: %s\n", vc.IsEdited(), vc.Value())
}

func main() {
	VersionControlExample()
}
