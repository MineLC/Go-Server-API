package play

import (
	"github.com/minelc/go-server-api/api/data"
	block "github.com/minelc/go-server-api/api/data/block"
	"github.com/minelc/go-server-api/api/network"
)

type PacketPlayOutBlockChange struct {
	Pos     data.PositionI
	BlockID block.Block
	Data    int32
}

func (p *PacketPlayOutBlockChange) UUID() int32 {
	return 35
}

func (p *PacketPlayOutBlockChange) Push(writer network.Buffer) {
	writer.PushPos(p.Pos)
	writer.PushVrI(int32(p.BlockID) + p.Data<<12)
}
