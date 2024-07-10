package login

import (
	"github.com/minelc/go-server-api/network"
)

type PacketIEncryptionResponse struct {
	Secret []byte
	Verify []byte
}

func (p *PacketIEncryptionResponse) UUID() int32 {
	return 0x01
}

func (p *PacketIEncryptionResponse) Pull(reader network.Buffer) {
	p.Secret = reader.PullUAS()
	p.Verify = reader.PullUAS()
}
