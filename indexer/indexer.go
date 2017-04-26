package indexer

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
	"fmt"
	"log"
	"github.com/antholord/poeIndexer/api"
	"github.com/antholord/poeIndexer/subscription"
	"time"
)

func init(){

}

type ninjaJSON struct {
	Id int `json:"id"`
	NextChangeId string `json:"nextChangeId"`
}

var id string
const ninjaAPIURL string = "http://api.poe.ninja/api/Data/GetStats"

func Run(m *subscription.Manager) bool{
	//Setup
	id = getNextChangeID()
	log.Printf("\nStarting indexing at : %v",id)
	APIsubscription := api.OpenPublicStashTabSubscription(id)
	var lastRequestTime time.Time
	//Loop over results
	go func() {
		for result := range APIsubscription.Channel {

			lastRequestTime = time.Now()
			if result.Error != nil {
				log.Printf("error: %v", result.Error.Error())
				continue
			}
			var count int = 0;
			for _, stash := range result.PublicStashTabs.Stashes {
				count+=len(stash.Items)
				processStash(&stash, m)
			}

			timeToQuery := time.Now().Sub(lastRequestTime)
			log.Println("Processing took : ", timeToQuery)
			//log.Printf("Processing %v items",count)
		}
	}()
	return true

}

func getNextChangeID() string{
	//Get ninja API info
	res, err := http.Get(ninjaAPIURL)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	resp , _:= ioutil.ReadAll(res.Body)

	defer res.Body.Close()
	//Unmarshal response into s
	var s ninjaJSON
	err = json.Unmarshal([]byte(resp), &s)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	//Assign next change id to package variable

	if s.NextChangeId != "" {
		fmt.Print(s.NextChangeId)
	}else {
		panic("No change ID found")
	}
	return s.NextChangeId
}
