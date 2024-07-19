package play

import (
	"github.com/minelc/go-server-api/network"
)

type ChatMode byte

const (
	FULL ChatMode = iota
	SYSTEM
	HIDDEN
)

type PacketPlayInSettings struct {
	Language     string
	ViewDistance byte
	ChatMode     ChatMode
	ChatColors   bool
	SkinBitMask  uint8
}

func (p *PacketPlayInSettings) UUID() int32 {
	return 21
}

func (p *PacketPlayInSettings) Pull(reader network.Buffer) {
	p.Language = reader.PullTxt()
	p.ViewDistance = reader.PullByt()
	p.ChatMode = ChatMode(reader.PullByt())
	p.ChatColors = reader.PullBit()
	p.SkinBitMask = reader.PullByt()
}
