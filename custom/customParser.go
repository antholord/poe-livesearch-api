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
}
func NewCustomParser() *CustomParser {

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
	return &CustomParser{
		BasesMap : m,
	}
}

