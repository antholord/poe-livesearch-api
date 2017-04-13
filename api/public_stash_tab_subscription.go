//Credit to https://github.com/ccbrown/poe-go/
package api

import (
	"net/http"
	"net/url"
	"time"
	"github.com/mailru/easyjson"
)

type PublicStashTabSubscriptionResult struct {
	PublicStashTabs *PublicStashTabs
	Error           error
}

type PublicStashTabSubscription struct {
	Channel      chan PublicStashTabSubscriptionResult
	closeChannel chan bool
	host         string
}

// Opens a subscription that begins with the given change id. To subscribe from the beginning, pass
// an empty string.
func OpenPublicStashTabSubscription(firstChangeId string) *PublicStashTabSubscription {
	return OpenPublicStashTabSubscriptionForHost("www.pathofexile.com", firstChangeId)
}

// Opens a subscription for an alternative host. Can be used for beta or foreign servers.
func OpenPublicStashTabSubscriptionForHost(host, firstChangeId string) *PublicStashTabSubscription {
	ret := &PublicStashTabSubscription{
		Channel:      make(chan PublicStashTabSubscriptionResult),
		closeChannel: make(chan bool),
		host:         host,
	}
	go ret.run(firstChangeId)
	return ret
}

func (s *PublicStashTabSubscription) Close() {
	s.closeChannel <- true
}

func (s *PublicStashTabSubscription) run(firstChangeId string) {
	defer close(s.Channel)

	nextChangeId := firstChangeId

	const requestInterval = time.Second * 1
	var lastRequestTime time.Time

	for {
		waitTime := requestInterval - time.Now().Sub(lastRequestTime)
		if waitTime > 0 {
			time.Sleep(waitTime)
		}

		select {
		case <-s.closeChannel:
			return
		default:
			response, err := http.Get("https://" + s.host + "/api/public-stash-tabs?id=" + url.QueryEscape(nextChangeId))
			if err != nil {
				s.Channel <- PublicStashTabSubscriptionResult{
					Error: err,
				}
				continue
			}
			lastRequestTime = time.Now()
			tabs := new(PublicStashTabs)
			err = easyjson.UnmarshalFromReader(response.Body, tabs)
			//decoder := json.NewDecoder(response.Body)
			//err = decoder.Decode(tabs)
			//timeToQuery := time.Now().Sub(lastRequestTime)
			//log.Println("Unmarshall took : ", timeToQuery)
			if err != nil {

				s.Channel <- PublicStashTabSubscriptionResult{
					Error: err,
				}
				//panic(err)
				continue
			}
			nextChangeId = tabs.NextChangeId

			if len(tabs.Stashes) > 0 {
				s.Channel <- PublicStashTabSubscriptionResult{
					PublicStashTabs: tabs,
				}
			}
		}
	}
}