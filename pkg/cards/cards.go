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
	Special
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
	case Special:
		return "special"
	default:
		return "NA"
	}
}

type CardRarity int

const (
	Common CardRarity = iota
	Uncommon
	Rare
	DoubleRare
	UltraRare
	IllustrationRare
	SpecialIllustrationRare
	HyperRare
)

func (r CardRarity) String() string {
	switch r {
	case Common:
		return "common"
	case Uncommon:
		return "uncommon"
	case Rare:
		return "rare"
	case DoubleRare:
		return "double_rare"
	case UltraRare:
		return "ultra_rare"
	case IllustrationRare:
		return "illustration_rare"
	case SpecialIllustrationRare:
		return "special_illustration_rare"
	case HyperRare:
		return "hyper_rare"
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

type CardSet int

const (
	SVI CardSet = iota
	PROMO
	PAL
	SVE
	OBF
	MEW
	PAR
	TEF
	TWM
	SFA
	SCR
	SSP
	PRE
	JTG
	DRI
	BLK
	WHT
)

func (s CardSet) String() string {
	switch s {
	case SVI:
		return "SVI"
	case PROMO:
		return "PROMO"
	case PAL:
		return "PAL"
	case SVE:
		return "SVE"
	case OBF:
		return "OBF"
	case MEW:
		return "MEW"
	case PAR:
		return "PAR"
	case TEF:
		return "TEF"
	case TWM:
		return "TWM"
	case SFA:
		return "SFA"
	case SCR:
		return "SCR"
	case SSP:
		return "SSP"
	case PRE:
		return "PRE"
	case JTG:
		return "JTG"
	case DRI:
		return "DRI"
	case BLK:
		return "BLK"
	case WHT:
		return "WHT"
	default:
		return "NA"
	}
}

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
	Name             string
	Hp               int
	Energy           EnergyType
	Move             []Action
	Weakness         WeaknessInfo
	Resistance       ResistanceInfo
	RetreatCost      int
	Set              CardSet
	CollectionNumber int
	Rarity           CardRarity
}

type TrainerType int

const (
	Item TrainerType = iota
	Tool
	Supporter
	Stadium
)

func (t TrainerType) String() string {
	switch t {
	case Item:
		return "item"
	case Tool:
		return "tool"
	case Supporter:
		return "supporter"
	case Stadium:
		return "stadium"
	default:
		return "NA"
	}
}

type TrainerCard struct {
	Name      string
	Type      TrainerType
	Effect    Action
	IsAceSpec bool
	Set       CardSet
	Rarity    CardRarity
}

type EnergyCard struct {
	Name   string
	Type   EnergyType
	Effect Action
	Set    CardSet
	Rarity CardRarity
}
