package main

import (
	"fmt"

	"github.com/VKyl/Design-Patterns/command"
	"github.com/VKyl/Design-Patterns/editorUI/controls"
	"github.com/VKyl/Design-Patterns/editor"
	editorui "github.com/VKyl/Design-Patterns/editorUI"
)

type Text string

func (t Text) Equals(other Text) bool {
	return t == other
}

func initEditorUI(editor *editor.Editor[Text]) *editorui.EditorUI[Text] {
	editorUI := editorui.NewEditorUI(editor)
	editorUI.AddButton("Reset", controls.NewButton("Reset", command.NewResetCommand[Text](editor)))
	editorUI.AddButton("Save", controls.NewButton("Save", command.NewSaveCommand[Text](editor)))
	fmt.Println("Editor UI created with Reset and Save buttons")
	return editorUI
}

func CommandEditorExample() {
	editor := editor.NewEditor[Text]()
	editorUI := initEditorUI(editor)

	editorUI.SetValue("Hello")
	editorUI.ClickButton("Save")
	editorUI.DisplayValue()
	editorUI.SetValue("Hello, World!")
	editorUI.DisplayValue()
	editorUI.ClickButton("Reset")
	editorUI.DisplayValue()
}

func main() {
	CommandEditorExample()
}
