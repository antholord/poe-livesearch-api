package custom

import (
	"strings"
	"strconv"
)

type CProperties struct {
	Armour  float64
	Es      float64
	Evasion float64
	Block   float64
	Crit    float64
	Quality float64
	WeaponRange float64
	Type string
	Category string
	SubCategory string
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
