package living

type HorseType byte
type HorseColor byte
type HorseStyle byte
type HorseArmor byte

const (
	Horse HorseType = iota
	Donkey
	Mule
	Zombie
	Skeleton
)

const (
	White HorseColor = iota
	Creamy
	Chestnut
	Brown
	Black
	Gray
	DarkDown
)

const (
	None HorseStyle = iota
	WhiteStyle
	Whitefield
	WhiteDots
	BlackDots
)

const (
	NoArmor HorseArmor = iota
	Iron
	Gold
	Diamond
)
