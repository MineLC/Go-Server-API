package events

import (
	"github.com/minelc/go-server-api/data"
	block "github.com/minelc/go-server-api/data/block"
	"github.com/minelc/go-server-api/ents"
)

type BlockBreakEvent struct {
	Cancel bool
	Player ents.Player

	Block     block.Block
	BlockData int16

	Action    block.BlockAction
	Position  data.PositionI
	Direction block.BlockDirection
}
