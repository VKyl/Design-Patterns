package main

import (
	"fmt"

	"github.com/VKyl/Design-Patterns/command"
	"github.com/VKyl/Design-Patterns/editor"
	versioncontrol "github.com/VKyl/Design-Patterns/editor/versionControl"
	editorui "github.com/VKyl/Design-Patterns/editorUI"
	"github.com/VKyl/Design-Patterns/editorUI/controls"
)

type Text string

func (t Text) Equals(other Text) bool {
	return t == other
}

func initEditorUICommand(editor *editor.Editor[Text]) *editorui.EditorUI[Text] {
	editorUI := editorui.NewEditorUI(editor)
	editorUI.AddButton("Reset", controls.NewButton("Reset", command.NewResetCommand[Text](editor)))
	editorUI.AddButton("Save", controls.NewButton("Save", command.NewSaveCommand[Text](editor)))
	fmt.Println("Editor UI created with Reset and Save buttons")
	return editorUI
}

func CommandEditorExample() {
	editor := editor.NewEditor[Text]()
	editorUI := initEditorUICommand(editor)

	editorUI.SetValue("Hello")
	editorUI.ClickButton("Save")
	editorUI.DisplayValue()
	editorUI.SetValue("Hello, World!")
	editorUI.DisplayValue()
	editorUI.ClickButton("Reset")
	editorUI.DisplayValue()
}

func initEditorUIMemento(
	editor *versioncontrol.VersionControl[Text, editor.Snapshot[Text]],
) *editorui.EditorUI[Text] {
	editorUI := editorui.NewEditorUI(editor)
	editorUI.AddButton("Reset", controls.NewButton("Reset", command.NewResetCommand[Text](editor)))
	editorUI.AddButton("Save", controls.NewButton("Save", command.NewSaveCommand[Text](editor)))
	editorUI.AddButton("Undo", controls.NewButton("Undo", command.NewUndoCommand[Text](editor)))
	fmt.Println("Editor UI created with Reset, Save, and Undo buttons")
	return editorUI
}

func VersionControlExample() {
	e := editor.NewMementoEditor[Text]()
	vc := versioncontrol.NewVersionControl(e)
	editorUI := initEditorUIMemento(vc)

	editorUI.SetValue("Hello")
	editorUI.ClickButton("Save")
	editorUI.DisplayValue()

	editorUI.SetValue("Hello, World!")
	editorUI.DisplayValue()

	editorUI.ClickButton("Reset")
	editorUI.DisplayValue()

	editorUI.SetValue("Hello Again")
	editorUI.DisplayValue()

	editorUI.SetValue("Hello again 2")
	editorUI.DisplayValue()

	editorUI.ClickButton("Undo")
	editorUI.DisplayValue()

	editorUI.ClickButton("Undo")
	editorUI.DisplayValue()
}

func main() {
	VersionControlExample()
}
