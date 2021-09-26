package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	inMemoryStore := NewInMemoryPlayerStore()

	sqliteStore, _ := NewSqlitePlayerStore("test.db")
	_ = sqliteStore.DeletePlayerScores()
	defer sqliteStore.DeletePlayerScores()

	fileDatabase, cleanFileDatabase := createTempFile(t, "")
	defer cleanFileDatabase()

	fileStore := &FileSystemPlayerStore{fileDatabase}

	player := "Pepper"

	cases := []struct {
		storeType string
		server    *PlayerServer
	}{
		{"InMemoryPlayerStore", NewPlayerServer(inMemoryStore)},
		{"SqlitePlayerStore", NewPlayerServer(sqliteStore)},
		{"FilePlayerStore", NewPlayerServer(fileStore)},
	}

	for _, c := range cases {
		c.server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
		c.server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
		c.server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

		t.Run(fmt.Sprintf("get %v score", c.storeType), func(t *testing.T) {
			response := httptest.NewRecorder()
			c.server.ServeHTTP(response, newGetScoreRequest(player))
			assertStatus(t, response.Code, http.StatusOK)

			assertResponseBody(t, response.Body.String(), "3")
		})

		t.Run(fmt.Sprintf("get %v league", c.storeType), func(t *testing.T) {
			response := httptest.NewRecorder()
			c.server.ServeHTTP(response, newLeagueRequest())
			assertStatus(t, response.Code, http.StatusOK)

			got := getLeagueFromResponse(t, response.Body)
			want := League{
				{"Pepper", 3},
			}
			assertLeague(t, got, want)
		})
	}
}
