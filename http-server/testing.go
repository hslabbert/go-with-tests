package poker

import (
	"fmt"
	"io"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

// A StubPlayerStore is a stub implementation of the PlayerStore
// interface for testing
type StubPlayerStore struct {
	Scores   map[string]int
	WinCalls []string
	League   League
}

// GetPlayerScore returns the score of the named player from the
// provided *StubPlayerStore.
func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.Scores[name]
	return score
}

// RecordWin increments the score of the named player in the
// provided *StubPlayerStore.
func (s *StubPlayerStore) RecordWin(name string) error {
	s.WinCalls = append(s.WinCalls, name)
	return nil
}

// GetLeague returns the League from the provided *StubPlayerStore.
func (s *StubPlayerStore) GetLeague() League {
	return s.League
}

func AssertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}

func AssertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q want %q", got, want)
	}
}

func AssertContentType(t testing.TB, response *httptest.ResponseRecorder, want string) {
	t.Helper()
	if response.Result().Header.Get("content-type") != want {
		t.Errorf("repsonse did not have content-type of %v, got %v", want, response.Result().Header)
	}
}

func AssertLeague(t testing.TB, got, want League) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func AssertPlayerWin(t testing.TB, store *StubPlayerStore, winner string) {
	t.Helper()

	if len(store.WinCalls) != 1 {
		t.Fatalf("got %d calls to RecordWin want %d", len(store.WinCalls), 1)
	}

	if store.WinCalls[0] != winner {
		t.Errorf("did not store correct winner; got %q want %q", store.WinCalls[0], winner)
	}
}

// ScheduledAlert holds information about when an alert is scheduled.
type ScheduledAlert struct {
	At     time.Duration
	Amount int
}

func (s ScheduledAlert) String() string {
	return fmt.Sprintf("%d chips at %v", s.Amount, s.At)
}

// SpyBlindAlerter allows you to spy on ScheduleAlertAt calls.
type SpyBlindAlerter struct {
	Alerts []ScheduledAlert
}

// ScheduleAlertAt records alerts that have been scheduled.
func (s *SpyBlindAlerter) ScheduleAlertAt(at time.Duration, amount int, to io.Writer) {
	s.Alerts = append(s.Alerts, ScheduledAlert{at, amount})
}
