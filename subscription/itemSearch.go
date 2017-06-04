package subscription

import (
	"strings"
	"github.com/antholord/poeIndexer/custom"
	"github.com/antholord/poeIndexer/api"
)

type NameObj struct {
	Name string
	IsFullName bool
	IsMultiName bool
	MultiName 	[6]MultiName
}

type MultiName struct {
	Name string
	IsFullName bool
}

type ItemSearch struct {

	League     string
	NameObj	NameObj

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

	Mods *[]api.MinMaxStr
	CustomParser *custom.CustomParser
}

func (s *ItemSearch) GenerateName (str string, cp *custom.CustomParser) {
	s.NameObj.IsMultiName = strings.Contains(str, "|")

	s.NameObj.Name = str
	s.NameObj.IsFullName = s.CheckIfNameIsFull(str)
	if (!s.NameObj.IsFullName){
		s.NameObj.Name = strings.ToUpper(s.NameObj.Name)
	}
	if (s.NameObj.IsMultiName) {
		for i, name := range strings.Split(str, "|"){
			s.NameObj.MultiName[i].Name = name
			s.NameObj.MultiName[i].IsFullName = s.CheckIfNameIsFull(name)
			if (!s.NameObj.MultiName[i].IsFullName){
				s.NameObj.MultiName[i].Name = strings.ToUpper(s.NameObj.MultiName[i].Name)
			}
		}
	}
}

func (s *ItemSearch) CheckIfNameIsFull (str string) bool{
	if (s.CustomParser.ItemsMap[str]){
		return true
	}
	return false
}
