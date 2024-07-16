package world

type WorldManager interface {
	Compress(chunkX int32, chunkZ int32) uint64

	GetWorld(name string) World

	UnloadWorld(name string)
	LoadWorld(filePath string)
}
