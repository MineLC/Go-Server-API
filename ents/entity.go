package ents

import "github.com/minelc/go-server-api/data"

type Entity interface {
	EntityUUID() int64
	GetPosition() *data.PositionF
}
