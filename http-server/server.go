package main

import (
	"fmt"
	"net/http"
	"strings"
)

// A PlayerServer implements the needed methods to satisfy the
// http.Handler interface, and holds a PlayerStore.
type PlayerServer struct {
	store PlayerStore
}

// A PlayerStore implements the needed methods to update and retrieve
// a player's score.
type PlayerStore interface {
	GetPlayerScore(name string) int
}

// ServerHTTP handles HTTP requests on a PlayerStore,
// satisfying the http.Handler interface.
func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	fmt.Fprint(w, p.store.GetPlayerScore(player))
}

// GetPlayerScore takes a Player's name and retrieves their score.
func GetPlayerScore(name string) string {
	if name == "Pepper" {
		return "20"
	}

	if name == "Floyd" {
		return "10"
	}

	return ""
}
