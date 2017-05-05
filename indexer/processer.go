package indexer

import (
	"github.com/antholord/poeIndexer/api"
	"github.com/antholord/poeIndexer/subscription"
	"github.com/mailru/easyjson"
	"log"
	"strconv"
	"strings"
)

func processStash(stash *api.Stash, m *subscription.Manager) {
	m.MapLock.Lock()
	for _, item := range stash.Items {
		for itemSearch, clients := range m.SubMap {
			if matchesCriterias(&itemSearch, &item) {
				//log.Println(item.CProperties)
				go broadcast(clients, api.ItemResult{item, stash.AccountName, stash.LastCharacterName, stash.Id, stash.Label, stash.Type, ""})
			}
		}
	}
	m.MapLock.Unlock()
}

func matchesCriterias(s *subscription.ItemSearch, item *api.Item) bool {

	if s.League == "" || s.League != item.League {
		return false
	//} else if s.Type != "" && s.Type != item.Type {
	} else if s.Type != "" && !strings.Contains(item.Type, s.Type) {
		return false
	//If search has a name
	} else if s.NameObj.Name != "" {
		//If its a single name, and not a concanated one
		if (!s.NameObj.IsMultiName){
			//Simple equality comparison if the name entered is in items list
			//Contains check if the name entered is partial or not in the items list
			if ((s.NameObj.IsFullName && s.NameObj.Name != item.FName) || (!s.NameObj.IsFullName && !strings.Contains(item.FNameUpper, s.NameObj.Name))){
				return false
			}
		}else {
			//There are multiple concanated names
			var found bool = false
			for _, i := range s.NameObj.MultiName {
				if ((i.IsFullName && i.Name == item.FName) || (!i.IsFullName && strings.Contains(item.FNameUpper, i.Name))){
					found = true
				}
			}
			if (!found){ return false }
		}
	}
	if (s.Category != "" && s.Category != item.CProperties.Category) {
		return false
	} else if (s.SubCategory != "" && s.SubCategory != item.CProperties.SubCategory) {
		return false
	} else if s.MinSockets != 0 && item.NbSockets < s.MinSockets {
		return false
	} else if s.MaxSockets != 0 && item.NbSockets > s.MaxSockets {
		return false
	} else if s.MinLinks != 0 && item.BiggestLink < s.MinLinks {
		return false
	} else if s.MaxLinks != 0 && item.BiggestLink > s.MaxLinks {
		return false
	} else if s.MinIlvl != 0 && item.ItemLevel < s.MinIlvl {
		return false
	} else if s.MaxIlvl != 0 && item.ItemLevel > s.MaxIlvl {
		return false
	}
	return true
}

func broadcast(clients map[*subscription.Client]bool, s api.ItemResult) {
	log.Println("Broadcasting to " + strconv.Itoa(len(clients)) + " clients item : " + s.Item.Name + " --- " + s.Item.Type )
	for client, _ := range clients {
		json, _ := easyjson.Marshal(s)
		client.Send <- json
	}
	return
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
