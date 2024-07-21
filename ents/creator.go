package ents

import "github.com/minelc/go-server-api/data/entity"

type EntityCreator interface {
	NewEntity(entity.Entity) Entity
}
