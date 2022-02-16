package poker

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/websocket"
)

func TestGETPlayers(t *testing.T) {

	store := StubPlayerStore{
		scores: map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
	}

	server := mustMakePlayerServer(t, &store)

	cases := []struct {
		Player string
		Score  string
	}{
		{Player: "Pepper", Score: "20"},
		{Player: "Floyd", Score: "10"},
	}

	for _, test := range cases {
		t.Run(fmt.Sprintf("returns %s's score", test.Player), func(t *testing.T) {
			request := newGetScoreRequest(test.Player)
			response := httptest.NewRecorder()

			server.ServeHTTP(response, request)

			AssertStatus(t, response.Code, http.StatusOK)
			AssertResponseBody(t, response.Body.String(), test.Score)
		})
	}

	t.Run("returns 404 on missing players", func(t *testing.T) {
		request := newGetScoreRequest("Apollo")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		AssertStatus(t, response.Code, http.StatusNotFound)
	})
}

func TestStoreWins(t *testing.T) {
	store := StubPlayerStore{
		scores: map[string]int{},
	}

	server := mustMakePlayerServer(t, &store)

	t.Run("it records wins on POST", func(t *testing.T) {
		player := "Pepper"

		request := newPostWinRequest(player)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		AssertStatus(t, response.Code, http.StatusAccepted)

		AssertPlayerWin(t, &store, player)
	})
}

func TestLeague(t *testing.T) {

	t.Run("it returns league table as JSON", func(t *testing.T) {
		wantedLeague := League{
			{"Cleo", 32},
			{"Chris", 20},
			{"Tiest", 14},
		}

		store := StubPlayerStore{nil, nil, wantedLeague}
		server := mustMakePlayerServer(t, &store)

		request := newLeagueRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := getLeagueFromResponse(t, response.Body)
		AssertStatus(t, response.Code, http.StatusOK)
		AssertLeague(t, got, wantedLeague)
		AssertContentType(t, response, jsonContentType)
	})
}

func TestGame(t *testing.T) {
	t.Run("GET /game returns 200", func(t *testing.T) {
		server := mustMakePlayerServer(t, &StubPlayerStore{})

		request := newGameRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		AssertStatus(t, response.Code, http.StatusOK)

	})

	t.Run("when we get a message over a websocket, it is a winner of a game", func(t *testing.T) {
		store := &StubPlayerStore{}
		winner := "Ruth"

		server := httptest.NewServer(mustMakePlayerServer(t, store))
		defer server.Close()

		wsURL := "ws" + strings.TrimPrefix(server.URL, "http") + "/ws"

		ws := mustDialWS(t, wsURL)
		defer ws.Close()

		writeWSMessage(t, ws, winner)

		// bad; very bad
		time.Sleep(10 * time.Millisecond)
		AssertPlayerWin(t, store, winner)
	})
}

func newPostWinRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func newGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func newGameRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/game", nil)
	return req
}

func getLeagueFromResponse(t testing.TB, body io.Reader) (league League) {
	t.Helper()
	league, err := NewLeague(body)

	if err != nil {
		t.Fatalf("Unable to parse response from server %q into slice of Player, '%v'", body, err)
	}

	return league
}

func newLeagueRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return req
}

func mustMakePlayerServer(t *testing.T, store PlayerStore) *PlayerServer {
	server, err := NewPlayerServer(store)
	if err != nil {
		t.Fatal("could not create player server ", err)
	}
	return server
}

func mustDialWS(t *testing.T, url string) *websocket.Conn {
	ws, _, err := websocket.DefaultDialer.Dial(url, nil)

	if err != nil {
		t.Fatalf("could not open a ws connection on %s %v", url, err)
	}

	return ws
}

func writeWSMessage(t testing.TB, conn *websocket.Conn, message string) {
	t.Helper()
	if err := conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
		t.Fatalf("could not send message over ws connection %v", err)
	}

}
