package entity

type Entity byte

const (
	Item          Entity = 1
	Xporb         Entity = 2
	Thrownegg     Entity = 7
	Leashknot     Entity = 8
	Painting      Entity = 9
	Arrow         Entity = 10
	Snowball      Entity = 11
	Fireball      Entity = 12
	Smallfireball Entity = 13
	Enderpearl    Entity = 14
	Eyeofender    Entity = 15
	Thrownpotion  Entity = 16
	Expbottle     Entity = 17
	Itemframe     Entity = 18
	Witherskull   Entity = 19
	Primedtnt     Entity = 20
	Fallingsand   Entity = 21
	Fireworks     Entity = 22
	Armorstand    Entity = 30
	Boat          Entity = 41

	Mob         Entity = 48
	Monster     Entity = 49
	Creeper     Entity = 50
	Skeleton    Entity = 51
	Spider      Entity = 52
	Giant       Entity = 53
	Zombie      Entity = 54
	Slime       Entity = 55
	Ghast       Entity = 56
	Pigzombie   Entity = 57
	Enderman    Entity = 58
	Cavespider  Entity = 59
	Silverfish  Entity = 60
	Blaze       Entity = 61
	Lavaslime   Entity = 62
	Enderdragon Entity = 63
	Witherboss  Entity = 64
	Bat         Entity = 65
	Witch       Entity = 66
	Endermite   Entity = 67
	Guardian    Entity = 68
	Pig         Entity = 90
	Sheep       Entity = 91
	Cow         Entity = 92
	Chicken     Entity = 93
	Squid       Entity = 94
	Wolf        Entity = 95
	MushroomCow Entity = 96
	Snowman     Entity = 97
	Ozelot      Entity = 98
	Golem       Entity = 99
	Horse       Entity = 100
	Rabbit      Entity = 101
	Villager    Entity = 102
)

func GetEntity(name string) Entity {
	switch name {
	case "item":
		return 1
	case "xporb":
		return 2
	case "thrownegg":
		return 7
	case "leashknot":
		return 8
	case "painting":
		return 9
	case "arrow":
		return 10
	case "snowball":
		return 11
	case "fireball":
		return 12
	case "smallfireball":
		return 13
	case "thrownenderpearl":
		return 14
	case "eyeofendersignal":
		return 15
	case "thrownpotion":
		return 16
	case "thrownexpbottle":
		return 17
	case "itemframe":
		return 18
	case "witherskull":
		return 19
	case "primedtnt":
		return 20
	case "fallingsand":
		return 21
	case "fireworksrocket":
		return 22
	case "armorstand":
		return 30
	case "boat":
		return 41
	case "mob":
		return 48
	case "monster":
		return 49
	case "creeper":
		return 50
	case "skeleton":
		return 51
	case "spider":
		return 52
	case "giant":
		return 53
	case "zombie":
		return 54
	case "slime":
		return 55
	case "ghast":
		return 56
	case "pigzombie":
		return 57
	case "enderman":
		return 58
	case "cavespider":
		return 59
	case "silverfish":
		return 60
	case "blaze":
		return 61
	case "lavaslime":
		return 62
	case "enderdragon":
		return 63
	case "witherboss":
		return 64
	case "bat":
		return 65
	case "witch":
		return 66
	case "endermite":
		return 67
	case "guardian":
		return 68
	case "pig":
		return 90
	case "sheep":
		return 91
	case "cow":
		return 92
	case "chicken":
		return 93
	case "squid":
		return 94
	case "wolf":
		return 95
	case "mushroomcow":
		return 96
	case "snowman":
		return 97
	case "ozelot":
		return 98
	case "villagergolem":
		return 99
	case "entityhorse":
		return 100
	case "rabbit":
		return 101
	case "villager":
		return 102
	}
	return Mob
}
