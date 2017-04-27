package indexer

import (

	"github.com/antholord/poeIndexer/api"
	"github.com/antholord/poeIndexer/subscription"
	"github.com/mailru/easyjson"
	"log"
	"strconv"
)

func processStash(stash *api.Stash, m *subscription.Manager) {
	m.MapLock.Lock()
	for _, item := range stash.Items {
		for itemSearch, clients := range m.SubMap {
			if (matchesCriterias(&itemSearch, &item)){
				go broadcast(clients, api.ItemResult{item, stash.AccountName, stash.LastCharacterName, stash.Id, stash.Label, stash.Type, "",})
			}
		}
	}
	m.MapLock.Unlock()
}

func matchesCriterias(s *subscription.ItemSearch, item *api.Item) bool{
	if (s.League == "" || s.League != item.League){
		return false
	}else if(s.Type != "" && s.Type != item.Type){
		return false
	}else if(s.Name != "" && s.Name != item.Name) {
		return false
	}else if(s.MinSockets != 0 && item.NbSockets < s.MinSockets){
		return false
	}else if(s.MaxSockets != 0 && item.NbSockets > s.MaxSockets){
		return false
	}else if(s.MinLinks != 0 && item.BiggestLink < s.MinLinks){
		return false
	}else if(s.MaxLinks != 0 && item.BiggestLink > s.MaxLinks){
		return false
	}else if(s.MaxIlvl != 0 && item.ItemLevel > s.MaxIlvl){
		return false
	}


	return true
}

func broadcast(clients map[*subscription.Client]bool, s api.ItemResult){
	log.Println("Broadcasting to " + strconv.Itoa(len(clients)) + " clients item : " + s.Item.Name)
	for client,_:= range clients {
		json, _ := easyjson.Marshal(s)
		client.Send <- json
	}
	return
}