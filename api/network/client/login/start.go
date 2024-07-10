package login

import (
	"github.com/minelc/go-server-api/api/network"
)

type PacketILoginStart struct {
	PlayerName string
}

func (p *PacketILoginStart) UUID() int32 {
	return 0x00
}

func (p *PacketILoginStart) Pull(reader network.Buffer) {
	p.PlayerName = reader.PullTxt()
}
