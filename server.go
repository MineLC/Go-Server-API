package main

import (
	"github.com/madflojo/tasks"
	"github.com/minelc/go-server-api/ents"
	"github.com/minelc/go-server-api/game/world"
	"github.com/minelc/go-server-api/network"
	"github.com/minelc/go-server-api/plugin"
)

type Server interface {
	GetPlayer(network.Connection) ents.Player
	GetPlayers() *map[network.Connection]ents.Player

	AddPlayer(conn network.Connection, player ents.Player)
	Disconnect(network.Connection)

	Broadcast(messages ...string)

	GetConsole() ents.Console
	GetMspt() Mspt
	GetScheduler() *tasks.Scheduler

	GetWorldManager() world.WorldManager
	GetPluginManager() plugin.PluginManager
	GetPacketManager() network.PacketManager

	Stop()
}

var server Server

func GetServer() Server {
	return server
}

func SetServer(newServer Server) {
	if server == nil {
		server = newServer
	}
}
