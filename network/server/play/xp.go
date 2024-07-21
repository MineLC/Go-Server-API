package play

import "github.com/minelc/go-server-api/network"

type PacketPlayOutExperience struct {
	Xp    int32
	Level int32
}

func (p *PacketPlayOutExperience) UUID() int32 {
	return 0x1F
}

func (p *PacketPlayOutExperience) Push(writer network.Buffer) {
	nextLevel := GetXP(p.Level + 1)
	if p.Xp == 0 {
		writer.PushF32(0.0)
	} else {
		writer.PushF32(float32(p.Xp) / float32(nextLevel))
	}
	writer.PushVrI(p.Level)
	writer.PushVrI(p.Xp)
}

func GetXP(level int32) int32 {
	if level <= 16 {
		return level*level + 6*level
	} else if level <= 31 {
		return int32(2.5*float32(level*level)) + int32(40.5*float32(level)) + 360
	} else {
		return int32(4.5*float32(level*level)) + int32(162.5*float32(level)) + 2220
	}
}

func GetLevel(xp int) int32 {
	i := 0

	for {
		if xp < 0 {
			break
		}
		if i < 16 {
			xp -= (2 * i) + 7
		} else if i < 31 {
			xp -= (5 * i) - 38
		} else {
			xp -= (9 * i) - 158
		}
		i++
	}
	return int32(i - 1)
}
