package network

type PacketInput int32

const (
	KeepAlive PacketInput = iota
	ChatMessage
	UseEntity
	Flying
	Position
	Look
	PositionLook
	BlockDig
	BlockPlace
	HeldItemSlot
	ArmAnimation
	EntityAction
	SteerVehicle
	CloseWindow
	WindowClick
	Transaction
	SetCreativeSlot
	EnchantItem
	UpdateSign
	Abilities
	TabComplete
	Settings
	ClientCommand
	CustomPayload
	Spectate
)

type PacketManager interface {
	SetHandler(packetId PacketInput, handler func(Connection, PacketI))
	RemoveHandler(packetId PacketInput)
	GetPlayFunc(id int32) func(c Connection, packet PacketI)

	GetCompression() int
}

type PacketI interface {
	Pull(in Buffer)
	UUID() int32
}

type PacketO interface {
	UUID() int32
	Push(out Buffer)
}
