package main

import (
	"fmt"
	"net/http"
	"strings"
)

// A PlayerServer implements the needed methods to satisfy the
// http.Handler interface (ServeHTTP), and holds a PlayerStore.
type PlayerServer struct {
	store PlayerStore
}

// A PlayerStore implements the needed methods to update and retrieve
// a player's score.
type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string) error
}

// ServerHTTP handles HTTP requests on a PlayerStore,
// satisfying the http.Handler interface.
func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	router := http.NewServeMux()

	router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	router.Handle("/players/", http.HandlerFunc(p.playersHandler))

	router.ServeHTTP(w, r)
}

func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (p *PlayerServer) playersHandler(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	switch r.Method {
	case http.MethodPost:
		p.processWin(w, player)
	case http.MethodGet:
		p.showScore(w, player)
	}
}

func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {
	score := p.store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}

func (p *PlayerServer) processWin(w http.ResponseWriter, player string) {
	_ = p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}
