package plugin

import (
	"github.com/minelc/go-server-api/cmd"
	"github.com/minelc/go-server-api/ents"
)

const (
	Join int32 = iota
	Quit
	Chat
	BlockBreak
)

type PluginManager interface {
	GetPlugin(name string) Plugin
	GetCommandManager() cmd.CommandManager
	GetEntityCreator() ents.EntityCreator

	CallEvent(event interface{}, eventType int32)
	AddListener(listener func(event interface{}), eventType int32, plugin Plugin)
	RemoveListener(eventType int32, plugin Plugin)
}
