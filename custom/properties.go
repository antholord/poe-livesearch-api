package custom

import (
	"strings"
	"strconv"
)

type Mod struct {
	ModStr string `json:"modStr"`
	Value1 float64 `json:"value1"`
	Value2 float64 `json:"value2"`
}

type CProperties struct {
	Armour  int `json:"armour"`
	Es      int `json:"es"`
	Evasion int `json:"evasion"`
	Block   float64 `json:"block"`
	Crit    float64 `json:"crit"`
	Quality float64 `json:"quality"`
	WeaponRange float64 `json:"weaponRange"`
	Type string `json:"type"`
	Category string `json:"category"`
	SubCategory string `json:"subCategory"`
	MapTier int `json:"mapTier"`

	APS   float64 `json:"aps"`
	Phys  float64 `json:"phys"`
	Ele   float64 `json:"ele"`
	Chaos float64 `json:"chaos"`
	Cdps  float64 `json:"cdps"`
	Pdps  float64 `json:"pdps"`
	Edps  float64 `json:"edps"`
	Dps   float64 `json:"dps"`

	Mods []Mod `json:"mods"`
}

func ParseDmgRange(s string) float64 {
	arr := strings.Split(s, "-")
	f1, _:= strconv.ParseFloat(arr[0], 32)
	f2, _:= strconv.ParseFloat(arr[1], 32)
	return (f1 + f2) / 2
}

func CalculateFinalValues(cp *CProperties){
	if (cp.APS ==0) {return}
	if (cp.Chaos != 0){
		cp.Cdps = cp.Chaos * cp.APS
	}
	if (cp.Phys != 0){
		cp.Pdps = cp.Phys * cp.APS
	}
	if (cp.Ele != 0){
		cp.Edps = cp.Ele * cp.APS
	}
	cp.Dps = cp.Cdps + cp.Pdps + cp.Edps
}
