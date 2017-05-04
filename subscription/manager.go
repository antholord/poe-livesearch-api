package subscription

import (
	"log"
	"net/http"
	"strconv"
	"sync"
	"github.com/antholord/poeIndexer/custom"
)

type Manager struct {
	SubMap map[ItemSearch]map[*Client]bool
	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client


	Quit chan bool

	MapLock sync.Mutex

	CustomParser *custom.CustomParser
}

func NewManager(cp *custom.CustomParser) *Manager {
	return &Manager{
		SubMap:     make(map[ItemSearch]map[*Client]bool),
		register:   make(chan *Client),
		Quit:       make(chan bool),
		unregister: make(chan *Client),
		CustomParser: cp,
	}
}

func (m *Manager) Run() {
	for {
		select {
		case client := <-m.register:
			log.Println("Trying to register client")
			m.MapLock.Lock()
			log.Println(client.ItemSearch)
			if _, ok := m.SubMap[client.ItemSearch]; !ok {
				log.Println("New search, creating client map")
				m.SubMap[client.ItemSearch] = make(map[*Client]bool)
			} else {
				log.Println("Adding client to existing search!")
			}
			m.SubMap[client.ItemSearch][client] = true
			m.MapLock.Unlock()
		case client := <-m.unregister:
			log.Println("Deleting client")
			m.MapLock.Lock()
			delete(m.SubMap[client.ItemSearch], client)
			log.Println(len(m.SubMap[client.ItemSearch]))
			if !(len(m.SubMap[client.ItemSearch]) > 0) {
				delete(m.SubMap, client.ItemSearch)
			}
			log.Println(len(m.SubMap))
			m.MapLock.Unlock()
		}
	}
}

func (manager *Manager) ServeWs(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	minSockets, err := strconv.ParseInt(r.FormValue("minSockets"), 10, 32)
	maxSockets, err := strconv.ParseInt(r.FormValue("maxSockets"), 10, 32)
	minLinks, err := strconv.ParseInt(r.FormValue("minLinks"), 10, 32)
	maxLinks, err := strconv.ParseInt(r.FormValue("maxLinks"), 10, 32)
	minIlvl, err := strconv.ParseInt(r.FormValue("minIlvl"), 10, 32)
	maxIlvl, err := strconv.ParseInt(r.FormValue("maxIlvl"), 10, 32)


	search := ItemSearch{
		CustomParser: manager.CustomParser,
		Type:       r.FormValue("type"),
		Category: r.FormValue("category"),
		SubCategory : r.FormValue("subCategory"),
		League:     r.FormValue("league"),
		MinSockets: int(minSockets),
		MaxSockets: int(maxSockets),
		MinLinks:   int(minLinks),
		MaxLinks:   int(maxLinks),
		MinIlvl:    int(minIlvl),
		MaxIlvl:    int(maxIlvl),
	}
	search.GenerateName(r.FormValue("name"), manager.CustomParser)

	//if search valid
	if search.League != "" && (search.Type != "" || search.NameObj.Name != "" || search.Category != "" || search.SubCategory != "") {
		client := &Client{manager: manager, conn: conn, Send: make(chan []byte, 1024), ItemSearch: search}
		manager.register <- client
		go client.writePump()
	} else {
		log.Println("Client sent invalid search ", search)
	}
}
