package main

import (
	"github.com/antholord/poeIndexer/indexer"
	"github.com/antholord/poeIndexer/subscription"
	"log"
	"net/http"
	"os"
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
	port := "1337"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
		log.Println("port " + port)
	}
	http.ListenAndServe(":" + port, nil)
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
