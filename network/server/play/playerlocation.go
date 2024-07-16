package play

import (
	"github.com/minelc/go-server-api/data"
	"github.com/minelc/go-server-api/network"
)

type PacketPlayOutPosition struct {
	Position data.PositionF
	Head     data.HeadPosition
}

func (p *PacketPlayOutPosition) UUID() int32 {
	return 8
}

func (p *PacketPlayOutPosition) Push(writer network.Buffer) {
	writer.PushF64(p.Position.X)
	writer.PushF64(p.Position.Y)
	writer.PushF64(p.Position.Z)

	writer.PushF32(p.Head.AxisX)
	writer.PushF32(p.Head.AxisY)

	// No relativity
	writer.PushByt(0)
}
