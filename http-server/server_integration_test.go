package poker

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

	fileDatabase, cleanFileDatabase := createTempFile(t, `[]`)
	defer cleanFileDatabase()

	fileStore, err := NewFileSystemPlayerStore(fileDatabase)

	assertNoError(t, err)

	player := "Pepper"

	cases := []struct {
		storeType string
		server    *PlayerServer
	}{
		{"InMemoryPlayerStore", mustMakePlayerServer(t, inMemoryStore)},
		{"SqlitePlayerStore", mustMakePlayerServer(t, sqliteStore)},
		{"FilePlayerStore", mustMakePlayerServer(t, fileStore)},
	}

	for _, c := range cases {
		c.server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
		c.server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
		c.server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

		t.Run(fmt.Sprintf("get %v score", c.storeType), func(t *testing.T) {
			response := httptest.NewRecorder()
			c.server.ServeHTTP(response, newGetScoreRequest(player))
			AssertStatus(t, response.Code, http.StatusOK)

			AssertResponseBody(t, response.Body.String(), "3")
		})

		t.Run(fmt.Sprintf("get %v league", c.storeType), func(t *testing.T) {
			response := httptest.NewRecorder()
			c.server.ServeHTTP(response, newLeagueRequest())
			AssertStatus(t, response.Code, http.StatusOK)

			got := getLeagueFromResponse(t, response.Body)
			want := League{
				{"Pepper", 3},
			}
			AssertLeague(t, got, want)
		})
	}
}
