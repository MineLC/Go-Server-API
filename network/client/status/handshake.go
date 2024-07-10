package status

import (
	"github.com/minelc/go-server-api/network"
)

type PacketHandshakingInSetProtocol struct {
	Version int32

	host string
	port uint16

	State network.PacketState
}

func (p *PacketHandshakingInSetProtocol) Pull(reader network.Buffer) {
	p.Version = reader.PullVrI()

	p.host = reader.PullTxt()
	p.port = reader.PullU16()

	p.State = network.StateValueOf(int(reader.PullVrI()))
}

func (p *PacketHandshakingInSetProtocol) UUID() int32 {
	return 0
}
