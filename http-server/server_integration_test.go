package poker_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	poker "github.com/hslabbert/go-with-tests/http-server"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	inMemoryStore := poker.NewInMemoryPlayerStore()

	sqliteStore, _ := poker.NewSqlitePlayerStore("test.db")
	_ = sqliteStore.DeletePlayerScores()
	defer sqliteStore.DeletePlayerScores()

	fileDatabase, cleanFileDatabase := createTempFile(t, `[]`)
	defer cleanFileDatabase()

	fileStore, err := poker.NewFileSystemPlayerStore(fileDatabase)

	assertNoError(t, err)

	player := "Pepper"

	cases := []struct {
		storeType string
		server    *poker.PlayerServer
	}{
		{"InMemoryPlayerStore", mustMakePlayerServer(t, inMemoryStore, dummyGame)},
		{"SqlitePlayerStore", mustMakePlayerServer(t, sqliteStore, dummyGame)},
		{"FilePlayerStore", mustMakePlayerServer(t, fileStore, dummyGame)},
	}

	for _, c := range cases {
		c.server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
		c.server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
		c.server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

		t.Run(fmt.Sprintf("get %v score", c.storeType), func(t *testing.T) {
			response := httptest.NewRecorder()
			c.server.ServeHTTP(response, newGetScoreRequest(player))
			poker.AssertStatus(t, response.Code, http.StatusOK)

			poker.AssertResponseBody(t, response.Body.String(), "3")
		})

		t.Run(fmt.Sprintf("get %v league", c.storeType), func(t *testing.T) {
			response := httptest.NewRecorder()
			c.server.ServeHTTP(response, newLeagueRequest())
			poker.AssertStatus(t, response.Code, http.StatusOK)

			got := getLeagueFromResponse(t, response.Body)
			want := poker.League{
				{"Pepper", 3},
			}
			poker.AssertLeague(t, got, want)
		})
	}
}
