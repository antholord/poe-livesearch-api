package api

import (
	"fmt"
)

type ItemResult struct {
	Item              Item   `json:"Item"`
	AccountName       string `json:"accountName"`
	LastCharacterName string `json:"lastCharacterName"`
	StashId           string `json:"id"`
	StashLabel        string `json:"stash"`
	StashType         string `json:"stashType"`
	Error             string `json:"error"`
}

func (i *ItemResult) ToString() string {
	return fmt.Sprintf("Ancient Reliquary Key: account = %v, league = %v, note = %v, tab = %v", i.AccountName, i.Item.League, i.Item.Note, i.StashLabel)
}
