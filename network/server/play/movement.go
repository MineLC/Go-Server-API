package play

import "github.com/minelc/go-server-api/network"

type PacketPlayOutRelEntityMove struct {
	EntityID int32
	DeltaX   byte
	DeltaY   byte
	DeltaZ   byte
	OnGround bool
}

func (p *PacketPlayOutRelEntityMove) UUID() int32 {
	return 21
}

func (p *PacketPlayOutRelEntityMove) Push(writer network.Buffer) {
	writer.PushVrI(p.EntityID)
	writer.PushByt(p.DeltaX)
	writer.PushByt(p.DeltaY)
	writer.PushByt(p.DeltaZ)
	writer.PushBit(p.OnGround)
}

type PacketPlayOutEntityHeadRotation struct {
	EntityID int32
	Position byte
}

func (p *PacketPlayOutEntityHeadRotation) UUID() int32 {
	return 25
}

func (p *PacketPlayOutEntityHeadRotation) Push(writer network.Buffer) {
	writer.PushVrI(p.EntityID)
	writer.PushByt(p.Position)
}
