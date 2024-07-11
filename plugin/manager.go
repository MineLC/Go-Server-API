package plugin

import "github.com/minelc/go-server-api/cmd"

type Event int32

const (
	Join Event = iota
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
	RemoveListener(eventType Event, plugin Plugin)
}
