package cards

//type Card interface {
//}

type Energy int

const (
	NA Energy = iota
	Grass
	Water
	Fire
	Lightning
	Fighting
	Psychic
	Darkness
	Metal
	Fairy
	Dragon
	Colorless
)

func (e Energy) String() string {
	switch e {
	case Colorless:
		return "colorless"
	case Water:
		return "water"
	case Fire:
		return "fire"
	case Grass:
		return "grass"
	case Lightning:
		return "lightning"
	case Fighting:
		return "fighting"
	case Psychic:
		return "psychic"
	case Darkness:
		return "darkness"
	case Metal:
		return "metal"
	case Fairy:
		return "fairy"
	case Dragon:
		return "Dragon"
	default:
		return "NA"
	}
}

type PokemonCard struct {
	Name string
	Hp   int
	EnergyType Energy
}
