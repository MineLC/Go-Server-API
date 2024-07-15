package world

type WorldManager interface {
	GetWorld(name string) *World

	UnloadWorld(world *World)
	LoadWorld(filePath string)
}
