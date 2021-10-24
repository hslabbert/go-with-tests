package poker_test

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"time"

	poker "github.com/hslabbert/go-with-tests/http-server"
)

var dummyBlindAlerter = &poker.SpyBlindAlerter{}
var dummyPlayerStore = &poker.StubPlayerStore{}
var dummyStdIn = &bytes.Buffer{}
var dummyStdOut = &bytes.Buffer{}

type GameSpy struct {
	StartedWith  int
	FinishedWith string
	StartCalled  bool
}

func (g *GameSpy) Start(numberOfPlayers int) {
	g.StartedWith = numberOfPlayers
}

func (g *GameSpy) Finish(winner string) {
	g.FinishedWith = winner
}

type scheduledAlert struct {
	at     time.Duration
	amount int
}

func (s scheduledAlert) String() string {
	return fmt.Sprintf("%d chips at %v", s.amount, s.at)
}

type SpyBlindAlerter struct {
	alerts []scheduledAlert
}

func (s *SpyBlindAlerter) ScheduleAlertAt(at time.Duration, amount int) {
	s.alerts = append(s.alerts, scheduledAlert{at, amount})
}

func TestCLI(t *testing.T) {

	t.Run("it prompts the user to enter the number of players and starts the game", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("7\n")
		game := &GameSpy{}

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		gotPrompt := stdout.String()
		wantPrompt := poker.PlayerPrompt

		if gotPrompt != wantPrompt {
			t.Errorf("got %q, want %q", gotPrompt, wantPrompt)
		}

		if game.StartedWith != 7 {
			t.Errorf("wanted Start called with 7 but got %d", game.StartedWith)
		}

	})

	t.Run("finish game with 'Chris' as winner", func(t *testing.T) {
		in := strings.NewReader("1\nChris wins\n")
		game := &GameSpy{}
		cli := poker.NewCLI(in, dummyStdOut, game)

		cli.PlayPoker()

		if game.FinishedWith != "Chris" {
			t.Errorf("expected finish called with 'Chris' but got %q", game.FinishedWith)
		}
	})

	t.Run("record 'Cleo' win from user input", func(t *testing.T) {
		in := strings.NewReader("1\nCleo wins\n")
		game := &GameSpy{}
		cli := poker.NewCLI(in, dummyStdOut, game)

		cli.PlayPoker()

		if game.FinishedWith != "Cleo" {
			t.Errorf("expected finish called with 'Cleo' but got %q", game.FinishedWith)
		}
	})

	t.Run("it prints an eror when a non numeric value is entered and does not start the game", func(T *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("Pies\n")
		game := &GameSpy{}

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		wantPrompt := poker.PlayerPrompt + poker.BadPlayerInputErrMsg

		assertMessagesSentToUser(t, stdout, wantPrompt)

		if game.StartCalled {
			t.Errorf("game should not have started")
		}
	})
}

func assertScheduledAlert(t testing.TB, got, want poker.ScheduledAlert) {
	t.Helper()
	if got != want {
		t.Errorf("got wrong scheduledAlert, got %v want %v", got, want)
	}
}

func assertMessagesSentToUser(t testing.TB, stdout *bytes.Buffer, messages ...string) {
	t.Helper()
	want := strings.Join(messages, "")
	got := stdout.String()
	if got != want {
		t.Errorf("got %q sent to stdout but expect %+v", got, messages)
	}
}
