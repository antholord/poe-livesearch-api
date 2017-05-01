package subscription

type ItemSearch struct {

	Name       string
	MultiName bool
	League     string

	Category string
	SubCategory string
	Type       string

	MinSockets int
	MaxSockets int
	MinLinks   int
	MaxLinks   int
	MinIlvl    int
	MaxIlvl    int

	Armour  float64
	Es      float64
	Evasion float64
	Block   float64
	Crit    float64
	Quality float64
	WeaponRange float64
	MapTier float64

	APS   float64
	Phys  float64
	Ele   float64
	Chaos float64
	Cdps  float64
	Pdps  float64
	Edps  float64
	Dps   float64
}
