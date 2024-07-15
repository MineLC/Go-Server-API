package world

import chunks "github.com/minelc/go-server-api/game/world/chunks"

type World interface {
	Compress(chunkX int32, chunkZ int32) uint64

	GetChunk(chunkX int32, chunkZ int32) *chunks.Chunk
	SetChunk(chunkX int32, chunkZ int32, chunk *chunks.Chunk)

	UnloadChunk(chunkX int32, chunkZ int32)
	GetAllChunks() []*chunks.Chunk
}
