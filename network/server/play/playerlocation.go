package play

import (
	"github.com/minelc/go-server-api/data"
	"github.com/minelc/go-server-api/network"
)

type PacketPlayOutPosition struct {
	Position data.PositionF
	Float    data.FloatPosition
}

func (p *PacketPlayOutPosition) UUID() int32 {
	return 8
}

func (p *PacketPlayOutPosition) Push(writer network.Buffer) {
	writer.PushF64(p.Position.X)
	writer.PushF64(p.Position.Y)
	writer.PushF64(p.Position.Z)

	writer.PushF32(p.Float.Yaw)
	writer.PushF32(p.Float.Pitch)

	// No relativity
	writer.PushByt(0)
}
