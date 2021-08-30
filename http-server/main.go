package main

import (
	"log"
	"net/http"
)

func main() {
	store, _ := NewSqlitePlayerStore("./store.db")
	server := &PlayerServer{store}
	log.Fatal(http.ListenAndServe(":5000", server))
}
