package play

import (
	"github.com/minelc/go-server-api/api/data"
	block "github.com/minelc/go-server-api/api/data/block"
	"github.com/minelc/go-server-api/api/network"
)

type PacketPlayInBlockDig struct {
	Action    block.BlockAction
	Position  data.PositionI
	Direction block.BlockDirection
}

func (p *PacketPlayInBlockDig) UUID() int32 {
	return 7
}

func (p *PacketPlayInBlockDig) Pull(reader network.Buffer) {
	p.Action = block.BlockAction(reader.PullVrI())
	p.Position = reader.PullPos()
	p.Direction = block.BlockDirection(reader.PullVrI())
}
