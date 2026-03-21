package controls

import "fmt"

type Command interface {
	Execute()
}

type Button struct {
	name    string
	command Command
}

func NewButton(name string, command Command) *Button {
	return &Button{
		name:    name,
		command: command,
	}
}

func (b *Button) Click() {
	if b.command == nil {
		return
	}
	fmt.Printf("Button %s clicked, executing command\n", b.name)
	b.command.Execute()
}
