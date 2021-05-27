package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubPlayerStore struct {
	scores map[string]int
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func TestGETPlayers(t *testing.T) {

	store := StubPlayerStore{
		map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
	}

	server := &PlayerServer{&store}

	cases := []struct {
		Player string
		Score  string
	}{
		{Player: "Pepper", Score: "20"},
		{Player: "Floyd", Score: "10"},
	}

	for _, test := range cases {
		t.Run(fmt.Sprintf("returns %s's score", test.Player), func(t *testing.T) {
			request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", test.Player), nil)
			response := httptest.NewRecorder()

			server.ServeHTTP(response, request)

			got := response.Body.String()
			want := test.Score

			if got != want {
				t.Errorf("got %q, want %q", got, want)
			}
		})
	}
}
