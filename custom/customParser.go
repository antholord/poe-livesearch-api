package custom

import (
	"io/ioutil"
	"github.com/antholord/poeIndexer/itemData"
	"encoding/json"
	"log"
	"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
)

type Category struct {
	TopCategory string
	SubCategory string
}
type CustomParser struct {
	BasesMap map[string]Category
	ItemsMap map[string]bool
	ModsMap map[string]bool
	session *mgo.Session
}
func NewCustomParser() *CustomParser {
	session, err := mgo.Dial("mongodb://test:test@ds123371.mlab.com:23371/heroku_lnc7sl64")
	if err != nil {
		log.Panic(err)
	}
	defer session.Close()

	return &CustomParser{
		BasesMap : ParseItemTypes(),
		ItemsMap : ParseItems(),
		ModsMap : make(map[string]bool),
		session : session,
	}
}

func ParseItemTypes() map[string]Category{
	file, e := ioutil.ReadFile("../itemData/item-types.json")
	if e != nil{
		log.Println(e)
	}
	var data itemData.ItemData
	_ = json.Unmarshal(file, &data)
	m := make(map[string]Category)
	for _, topCat := range data.ItemTypes{
		for _, subCat := range topCat.Value{
			for _, base := range subCat.Base{
				m[base] = Category{TopCategory:topCat.Category, SubCategory:subCat.SubCategory}
			}
		}
	}
	return m
}

func ParseItems() map[string]bool {
	file, e := ioutil.ReadFile("../itemData/items.json")
	if e != nil{
		log.Println(e)
	}
	var data itemData.ItemData
	_ = json.Unmarshal(file, &data)
	m := make(map[string]bool)

	for _, item := range data.Items {
		m[item] = true
	}

	return m
}

