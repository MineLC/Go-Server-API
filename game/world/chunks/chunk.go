package world

import block "github.com/minelc/go-server-api/data/block"

type ChunkSection struct {
	Blocks       [4096]uint16
	BlocksPlaced int16
}

func (c *ChunkSection) GetBlocksPlaced() int16 {
	return c.BlocksPlaced
}

type Chunk interface {
	GetX() int32
	GetZ() int32

	SetBlock(x int32, y int32, z int32, blockType block.Block, data uint16)
	SetAll(blockType block.Block, start int, end, data uint16)
	GetSections() [16]*ChunkSection
}
