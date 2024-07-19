package ents

import (
	"github.com/minelc/go-server-api/data/chat"
	"github.com/minelc/go-server-api/data/player"
	"github.com/minelc/go-server-api/network"
)

type Player interface {
	EntityLiving
	Sender

	GetIsOnline() bool
	SetIsOnline(state bool)

	GetSkinParts() byte
	GetLanguage() string
	GetViewDistance() byte
	GetChatMode() byte

	GetGamemode() player.GameMode
	SetGamemode(gamemode player.GameMode)

	GetProfile() *player.Profile

	GetPing() int32
	SetPing(ping_delay int64, server_ping int64)
	GetKeepAlive() int64
	SetKeepAlive(time int64)

	SendMsgPos(pos chat.MessagePosition, messages ...string)
	SendMsgColorPos(pos chat.MessagePosition, messages ...string)
	GetConnection() network.Connection

	Disconnect()
}
