package cmd

import "github.com/minelc/go-server-api/ents"

type CommandManager interface {
	ReplaceManager(newManager CommandManager)

	Add(cmd Command, values ...string)
	AddStruct(cmd StructCommand, values ...string)
	Delete(values ...string)
	Get(prefix string) *StructCommand
}

type Command interface {
	Execute(sender ents.Sender, args []string)
}

type StructCommand struct {
	Execute func(sender ents.Sender, args []string)
	Tab     func(sender ents.Sender, args []string) []string
}

type TabCommand interface {
	Command
	Tab(sender ents.Sender, args []string) []string
}
