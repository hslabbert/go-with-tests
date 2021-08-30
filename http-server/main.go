package main

import (
	"log"
	"net/http"
)

func main() {
	store, _ := NewSqlitePlayerStore("./store.db")
	server := NewPlayerServer(store)
	log.Fatal(http.ListenAndServe(":5000", server))
}
