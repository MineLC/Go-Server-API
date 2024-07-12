package events

import "github.com/minelc/go-server-api/ents"

type PlayerJoinEvent struct {
	Player ents.Player
}

func (p *PlayerJoinEvent) GetPlayer() ents.Player {
	return p.Player
}

type PlayerQuitEvent struct {
	Player ents.Player
}

func (p *PlayerQuitEvent) GetPlayer() ents.Player {
	return p.Player
}
