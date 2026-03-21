package editor

import (
	"fmt"

	versioncontrol "github.com/VKyl/Design-Patterns/editor/versionControl"
	"github.com/VKyl/Design-Patterns/util"
)

type VersionControll[T util.Comperable[T]] interface {
	AddChange(v Snapshot[T])
	Undo()
	Save()
	Reset()
	IsEdited() bool
}

type VersionControlledEditor[T util.Comperable[T]] struct {
	editor *MementoEditor[T]
	vc     VersionControll[T]
}

func NewVersionControlledEditor[T util.Comperable[T]]() *VersionControlledEditor[T] {
	editor := NewMementoEditor[T]()
	return &VersionControlledEditor[T]{
		editor: editor,
		vc:     versioncontrol.NewVersionControl(editor),
	}
}

func (vc *VersionControlledEditor[T]) Value() T {
	return vc.editor.Value()
}

func (vc *VersionControlledEditor[T]) SetValue(value T) {
	fmt.Printf("Setting value to %v\n", value)
	vc.saveChange()
	vc.editor.SetValue(value)
}

func (vc *VersionControlledEditor[T]) SetFont(font string) {
	fmt.Printf("Setting font to %v\n", font)
	vc.saveChange()
	vc.editor.SetFont(font)
}

func (vc *VersionControlledEditor[T]) SetFontSize(size int) {
	fmt.Printf("Setting font size to %v\n", size)
	vc.saveChange()
	vc.editor.SetFontSize(size)
}

func (vc *VersionControlledEditor[T]) IsEdited() bool {
	return vc.vc.IsEdited()
}

func (vc *VersionControlledEditor[T]) Reset() {
	vc.vc.Reset()
}

func (vc *VersionControlledEditor[T]) Save() {
	vc.vc.Save()
}

func (vc *VersionControlledEditor[T]) Undo() {
	vc.vc.Undo()
}

func (vc *VersionControlledEditor[T]) saveChange() {
	vc.vc.AddChange(vc.editor.MakeSnapshot())
}
