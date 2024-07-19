package ents

import (
	"github.com/minelc/go-server-api/data"
	"github.com/minelc/go-server-api/network"
)

type Entity interface {
	EntityUUID() int64
	GetPosition() *data.PositionF

	PushMetadata(network.Buffer)
}
