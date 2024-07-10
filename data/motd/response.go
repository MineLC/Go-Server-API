package motd

import (
	"encoding/json"

	"github.com/minelc/go-server-api/data/chat"
)

type Response struct {
	Version     Version `json:"version"`
	Players     Players `json:"players"`
	Description Message `json:"description"`
	Favicon     string  `json:"favicon"`
}

type Version struct {
	Name     string `json:"name"`
	Protocol int    `json:"protocol"`
}

type Players struct {
	Max    int            `json:"max"`
	Online int            `json:"online"`
	Sample []SamplePlayer `json:"sample"`
}

type SamplePlayer struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

type Message struct {
	Text string `json:"text"`
}

type Motd struct {
	VersionName string
	Line        string
	Favicon     string
	Protocol    int
	MaxPlayers  int
	Online      int
	Sample      []SamplePlayer
}

func CreateResponse(motd Motd) string {
	res := Response{
		Version: Version{
			Name:     "GoLang Server",
			Protocol: 47,
		},
		Players: Players{
			Max:    motd.MaxPlayers,
			Online: motd.Online,
			Sample: motd.Sample,
		},
		Description: Message{
			Text: chat.Translate(motd.Line),
		},
		Favicon: motd.Favicon,
	}
	text, err := json.Marshal(res)
	if err != nil {
		return ""
	}
	return string(text)
}
