package play

import (
	"github.com/minelc/go-server-api/api/data/player"
	"github.com/minelc/go-server-api/api/network"
)

type PacketPlayOutHeldItemChange struct {
	Slot player.HotBarSlot // 0-8
}

func (p *PacketPlayOutHeldItemChange) UUID() int32 {
	return 9
}

func (p *PacketPlayOutHeldItemChange) Push(writer network.Buffer) {
	writer.PushByt(byte(p.Slot))
}
