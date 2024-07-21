package ents

import (
	"github.com/minelc/go-server-api/data/entity/living"
)

type HorseData struct {
	IsTame    bool
	HasSaddle bool
	HasChest  bool
	IsBred    bool
	IsEating  bool
	IsRealing bool
	MouthOpen bool

	Type  living.HorseType
	Color living.HorseColor
	Style living.HorseStyle
	Armor living.HorseArmor
}

type Horse interface {
	EntityLiving

	GetAge() byte
	SetAge(byte)

	GetHorseData() *HorseData
}
