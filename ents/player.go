package ents

import (
	"github.com/minelc/go-server-api/data/chat"
	"github.com/minelc/go-server-api/data/player"
	"github.com/minelc/go-server-api/network"
)

type ClientSettings struct {
	SkinParts    byte
	Language     string
	ViewDistance byte
	ChatMode     byte
}

type Player interface {
	EntityLiving
	Sender

	GetClientSettings() *ClientSettings
	GetAbsorption() byte
	GetFood() float32

	GetGamemode() player.GameMode
	SetGamemode(gamemode player.GameMode)

	GetXP() int32
	GetLevel() int32
	SetXP(int32)
	SetLevel(int32)

	GetProfile() *player.Profile

	GetPing() int32
	SetPing(ping_delay int64, server_ping int64)
	GetKeepAlive() int64
	SetKeepAlive(time int64)

	SendMsgPos(pos chat.MessagePosition, messages ...string)
	SendMsgColorPos(pos chat.MessagePosition, messages ...string)
	GetConnection() network.Connection

	IsSneaking() bool
	IsSprinting() bool
	SetSprinting(bool)
	SetSneaking(bool)

	Disconnect()
}
