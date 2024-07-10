package play

import (
	"github.com/minelc/go-server-api/network"
)

type PacketPlayInKeepAlive struct {
	KeepAliveID int32
}

func (p *PacketPlayInKeepAlive) UUID() int32 {
	return 0
}

func (p *PacketPlayInKeepAlive) Pull(reader network.Buffer) {
	p.KeepAliveID = reader.PullI32()
}
