package main

import (
	"github.com/antholord/poeIndexer/indexer"
	"github.com/antholord/poeIndexer/subscription"
	"log"
	"net/http"
)

var manager *subscription.Manager

func init() {

}
func main() {

	manager = subscription.NewManager()
	go manager.Run()
	_ = indexer.Run(manager)
	/*

	go func() {
		for {
			select {
			case quit := <-manager.Quit:
				if quit {
					return
				}
			}
		}
	}()
	*/
	http.HandleFunc("/", index)
	http.HandleFunc("/ws/", serveWS)
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("../web/livesearch/src/client/public"))))
	http.ListenAndServe(":1337", nil)
}

func index(w http.ResponseWriter, req *http.Request) {

	http.ServeFile(w, req, "../web/livesearch/src/client/index.html")
}

func serveWS(w http.ResponseWriter, req *http.Request) {
	manager.ServeWs(w, req)
}

func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}
