package events

import "github.com/minelc/go-server-api/ents"

type PlayerJoinEvent struct {
	Player ents.Player
}
type PlayerQuitEvent struct {
	Player ents.Player
}
