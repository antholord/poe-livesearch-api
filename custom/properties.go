package custom

import (
	"strings"
	"strconv"
)

type CProperties struct {
	Armour  float64 `json:"armour"`
	Es      float64 `json:"es"`
	Evasion float64 `json:"evasion"`
	Block   float64 `json:"block"`
	Crit    float64 `json:"crit"`
	Quality float64 `json:"quality"`
	WeaponRange float64 `json:"weaponRange"`
	Type string `json:"type"`
	Category string `json:"category"`
	SubCategory string `json:"subCategory"`
	MapTier float64 `json:"mapTier"`

	APS   float64 `json:"aps"`
	Phys  float64 `json:"phys"`
	Ele   float64 `json:"ele"`
	Chaos float64 `json:"chaos"`
	Cdps  float64 `json:"cdps"`
	Pdps  float64 `json:"pdps"`
	Edps  float64 `json:"edps"`
	Dps   float64 `json:"dps"`
}

func ParseDmgRange(s interface{}) float64 {
	arr := strings.Split(s.(string), "-")
	f1, _:= strconv.ParseFloat(arr[0], 10)
	f2, _:= strconv.ParseFloat(arr[1], 10)
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
