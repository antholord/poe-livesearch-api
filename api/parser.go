package api

import (
	"github.com/antholord/poeIndexer/custom"
)

func ParseProperties(op []ItemProperty) custom.CProperties {
	//var lastRequestTime time.Time
	p := custom.CProperties{}
	//lastRequestTime = time.Now()
	for _, prop := range op {

		//log.Println(prop.Name)
		switch prop.Name {
		case "Armour" :
			//log.Println(prop.Values[0])
			//log.Println(reflect.TypeOf(prop.Values[0][0]))
			//p.Armour = string.Atoi(prop.Values[0][0].(string))
			//s:= reflect.ValueOf(prop.Values[0])
			//p.Armour = s.Index(0).Int()
			//t := make([]int, 2)
			//log.Println(reflect.TypeOf(prop.Values[0]).Kind())

		}
	}

	//timeToQuery := time.Now().Sub(lastRequestTime)
	//log.Println("Unmarshall took : ", timeToQuery)
	return p
}
