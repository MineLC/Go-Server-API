package network

import (
	"net"
)

type Connection interface {
	Address() net.Addr

	GetState() PacketState
	SetState(state PacketState)

	Encrypt(data []byte) (output []byte)
	Decrypt(data []byte) (output []byte)

	CertifyName() string

	CertifyData() []byte

	CertifyValues(name string)
	CertifyUpdate(secret []byte)

	Pull(data []byte) (len int, err error)
	Push(data []byte) (len int, err error)

	Stop() (err error)

	SendPacket(packet PacketO)
}
