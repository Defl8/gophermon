package cards

//type Card interface {
//}

type EnergyType int

const (
	NA EnergyType = iota
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

func (e EnergyType) String() string {
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
		return "dragon"
	default:
		return "NA"
	}
}

type WeaknessInfo struct {
	EnergyType EnergyType
	Multiplier int
}

type ResistanceInfo struct {
	EnergyType     EnergyType
	DamageResisted int
}

type Stage int

const (
	Basic Stage = iota
	Stage1
	Stage2
)

type Attack struct {
	Name   string
	Desc   string
	Damage int
	Cost   []EnergyType
}

func (a Attack) GetName() string {
	return a.Name
}

func (a Attack) GetDesc() string {
	return a.Desc
}

type Ability struct {
	Name string
	Desc string
}

func (a Ability) GetName() string {
	return a.Name
}

func (a Ability) GetDesc() string {
	return a.Desc
}

type Action interface {
	GetName() string
	GetDesc() string
}

type PokemonCard struct {
	Name        string
	Hp          int
	Energy      EnergyType
	Move        []Action
	Weakness    WeaknessInfo
	Resistance  ResistanceInfo
	RetreatCost int
}

type TrainerType int

const (
	Item TrainerType = iota
	Tool
	Supporter
	Stadium
)

type TrainerCard struct {
	Name    string
	Type    TrainerType
	Effect  Action
	AceSpec bool
}
