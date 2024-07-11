package plugin

import "github.com/minelc/go-server-api/cmd"

const (
	Join int32 = iota
	Quit
	Chat
	Command
	Tab
)

type PluginManager interface {
	GetPlugin(name string) Plugin
	GetCommandManager() *cmd.CommandManager

	RegisterCustomEvent(eventType int32)
	AddListener(func(event interface{}, eventType int32, plugin Plugin))
	RemoveListener(eventType int32, plugin Plugin)
}
