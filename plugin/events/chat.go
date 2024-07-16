package events

import "github.com/minelc/go-server-api/ents"

type PlayerChat struct {
	Cancel  bool
	Player  ents.Player
	Message string
}
