package itemData

type ItemData struct {
	ItemTypes []struct {
		Category string `json:"key"`
		Value []struct {
			SubCategory string `json:"key"`
			Base []string `json:"value"`
		} `json:"value"`
	} `json:"itemTypes"`
	Items []string `json:"items"`
}

