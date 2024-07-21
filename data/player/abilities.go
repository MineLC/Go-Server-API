package player

import (
	"github.com/minelc/go-server-api/network"
	"github.com/minelc/go-server-api/util"
)

type PlayerAbilities struct {
	Invulnerable bool
	Flying       bool
	AllowFlight  bool
	InstantBuild bool
}

func (p *PlayerAbilities) Push(writer network.Buffer) {
	flags := byte(0)

	util.Set(&flags, 0x01, p.Invulnerable)
	util.Set(&flags, 0x02, p.Flying)
	util.Set(&flags, 0x04, p.AllowFlight)
	util.Set(&flags, 0x08, p.InstantBuild)

	writer.PushByt(flags)
}

func (p *PlayerAbilities) Pull(reader network.Buffer) {
	flags := reader.PullByt()

	p.Invulnerable = util.Has(flags, 0x01)
	p.Flying = util.Has(flags, 0x02)
	p.AllowFlight = util.Has(flags, 0x04)
	p.InstantBuild = util.Has(flags, 0x08)
}
