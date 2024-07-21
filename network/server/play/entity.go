package play

import (
	"github.com/minelc/go-server-api/ents"
	"github.com/minelc/go-server-api/network"
)

type PacketPlayOutEntityMetadata struct {
	Entity ents.Entity
}

func (p *PacketPlayOutEntityMetadata) UUID() int32 {
	return 0x1C
}

func (p *PacketPlayOutEntityMetadata) Push(writer network.Buffer) {
	writer.PushVrI(int32(p.Entity.EntityUUID()))
	p.Entity.PushMetadata(writer)
	writer.PushByt(127)
}

type PacketPlayOutSpawnEntityLiving struct {
	Entity ents.EntityLiving
}

func (p *PacketPlayOutSpawnEntityLiving) UUID() int32 {
	return 15
}

func (p *PacketPlayOutSpawnEntityLiving) Push(writer network.Buffer) {
	writer.PushVrI(int32(p.Entity.EntityUUID()))
	writer.PushByt(byte(p.Entity.GetType()))

	writer.PushI32(int32(p.Entity.GetPosition().X * 32.0))
	writer.PushI32(int32(p.Entity.GetPosition().Y * 32.0))
	writer.PushI32(int32(p.Entity.GetPosition().Z * 32.0))

	writer.PushByt(byte(p.Entity.GetHeadPos().Yaw * 256.0 / 360.0))
	writer.PushByt(byte(p.Entity.GetHeadPos().Pitch * 256.0 / 360.0))
	writer.PushByt(byte(p.Entity.GetHeadPos().Yaw * 256.0 / 360.0))

	writer.PushI16(0)
	writer.PushI16(0)
	writer.PushI16(0)

	p.Entity.PushMetadata(writer)
	writer.PushByt(127)
}

type PacketPlayOutSpawnPlayer struct {
	Player ents.Player
}

func (p *PacketPlayOutSpawnPlayer) UUID() int32 {
	return 12
}

func (p *PacketPlayOutSpawnPlayer) Push(writer network.Buffer) {
	writer.PushVrI(int32(p.Player.EntityUUID()))
	writer.PushUID(p.Player.GetProfile().UUID)

	writer.PushI32(int32(p.Player.GetPosition().X * 32.0))
	writer.PushI32(int32(p.Player.GetPosition().Y * 32.0))
	writer.PushI32(int32(p.Player.GetPosition().Z * 32.0))

	writer.PushByt(byte(p.Player.GetHeadPos().Yaw * 256.0 / 360.0))
	writer.PushByt(byte(p.Player.GetHeadPos().Pitch * 256.0 / 360.0))

	writer.PushI16(0) // TODO: ItemStack

	p.Player.PushMetadata(writer)
	writer.PushByt(127)
}
