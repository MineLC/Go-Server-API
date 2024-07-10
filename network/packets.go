package network

type PacketI interface {
	Pull(in Buffer)
	UUID() int32
}

type PacketO interface {
	UUID() int32
	Push(out Buffer)
}
