package custom

import (
	"io/ioutil"
	"github.com/antholord/poeIndexer/itemData"
	"encoding/json"
	"log"
)

type Category struct {
	TopCategory string
	SubCategory string
}
type CustomParser struct {
	BasesMap map[string]Category
	ItemsMap map[string]bool
}
func NewCustomParser() *CustomParser {
	return &CustomParser{
		BasesMap : ParseItemTypes(),
		ItemsMap : ParseItems(),
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

