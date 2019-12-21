package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/nkex606/chatroom-server/ws"
)

var port = os.Getenv("port")

func main() {
	s := ws.NewServer()
	go s.Run()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		ws.Hanlder(s, w, r)
	})

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		healthz(w, r)
	})
	log.Printf("Start listening on port %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func healthz(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ok")
}
