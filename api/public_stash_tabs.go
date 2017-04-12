//Credit to https://github.com/ccbrown/poe-go/
package api
//easyjson:json
type PublicStashTabs struct {
	NextChangeId string  `json:"next_change_id"`
	Stashes      []Stash `json:"stashes"`
}