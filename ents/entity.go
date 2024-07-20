package ents

import (
	"github.com/minelc/go-server-api/data"
	"github.com/minelc/go-server-api/network"
)

type EntityMeta struct {
	OnFire    bool
	Crouched  bool
	Sprinting bool
	Eating    bool
	Invisible bool
}

type Entity interface {
	EntityUUID() int32
	GetPosition() *data.PositionF

	SetData(EntityMeta)
	GetData() EntityMeta

	PushMetadata(network.Buffer)
}
