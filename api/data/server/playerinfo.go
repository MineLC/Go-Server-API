package server

import (
	"github.com/minelc/go-server-api/api/ents"
	"github.com/minelc/go-server-api/api/network"
)

type PlayerInfoAction int32

const (
	AddPlayer PlayerInfoAction = iota
	UpdateGameMode
	UpdateLatency
	UpdateDisplayName
	RemovePlayer
)

type PlayerInfo interface {
	network.BufferPush
}

type PlayerInfoAddPlayer struct {
	Player ents.Player
}

func (p *PlayerInfoAddPlayer) Push(writer network.Buffer) {
	prof := p.Player.GetProfile()
	writer.PushUID(prof.UUID)
	writer.PushTxt(prof.Name)

	writer.PushVrI(int32(len(prof.Properties)))

	for _, prop := range prof.Properties {
		writer.PushTxt(prop.Name)
		writer.PushTxt(prop.Value)

		if prop.Signature == nil {
			writer.PushBit(false)
		} else {
			writer.PushBit(true)
			writer.PushTxt(*prop.Signature)
		}
	}

	writer.PushVrI(int32(p.Player.GetGamemode()))

	writer.PushVrI(0) // update this to the player's actual ping

	writer.PushBit(false) // update this to be whether the player has a custom display name or not, write that name as json if they do
}

type PlayerInfoUpdateLatency struct {
}
