package ents

import "github.com/minelc/go-server-api/data"

type ArmorData struct {
	Small      bool
	Gravity    bool
	Arms       bool
	RemoveBase bool
	Marker     bool
}
type ArmorPos struct {
	BodyPos *data.FloatPosition

	LeftArm  *data.FloatPosition
	RightArm *data.FloatPosition

	LeftLeg  *data.FloatPosition
	RightLeg *data.FloatPosition
}

type ArmorStand interface {
	EntityLiving

	GetArmorData() ArmorData
	SetArmorData(ArmorData)

	GetArmorPos() *ArmorPos
}
