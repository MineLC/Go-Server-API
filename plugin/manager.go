package plugin

import "github.com/minelc/go-server-api/cmd"

const (
	Join int32 = iota
	Quit
	Chat
	BlockBreak
)

type PluginManager interface {
	GetPlugin(name string) Plugin
	GetCommandManager() cmd.CommandManager

	CallEvent(event interface{}, eventType int32)
	AddListener(listener func(event interface{}), eventType int32, plugin Plugin)
	RemoveListener(eventType int32, plugin Plugin)
}
