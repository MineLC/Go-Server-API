package ents

import "github.com/minelc/go-server-api/data"

type EntityLiving interface {
	Entity

	GetHeadPos() *data.FloatPosition

	GetHealth() float64
	SetHealth(health float64)
}
