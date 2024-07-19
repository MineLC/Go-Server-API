package events

import "github.com/minelc/go-server-api/ents"

type PlayerDamageEvent struct {
	Cancel  bool
	Victim  ents.Player
	Damager *ents.Entity
}
