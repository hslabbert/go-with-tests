package poker_test

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"time"

	poker "github.com/hslabbert/go-with-tests/http-server"
)

var dummySpyAlerter = &SpyBlindAlerter{}
var dummyPlayerStore = &poker.StubPlayerStore{}
var dummyStdOut = &bytes.Buffer{}

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
	//	cases := []string{
	//		"Cleo",
	//		"Chris",
	//	}
	//
	//	for _, c := range cases {
	//		t.Run(fmt.Sprintf("record %s win from user input", c), func(t *testing.T) {
	//			in := strings.NewReader(fmt.Sprintf("%s wins\n", c))
	//			playerStore := &poker.StubPlayerStore{}
	//
	//			cli := poker.NewCLI(playerStore, in, dummySpyAlerter)
	//			cli.PlayPoker()
	//
	//			poker.AssertPlayerWin(t, playerStore, c)
	//		})
	//	}
	//
	//	t.Run("it schedules present of blind values", func(t *testing.T) {
	//		in := strings.NewReader("Chris wins \n")
	//		playerStore := &poker.StubPlayerStore{}
	//		blindAlerter := &SpyBlindAlerter{}
	//
	//		cli := poker.NewCLI(playerStore, in, blindAlerter)
	//		cli.PlayPoker()
	//
	//		cases := []scheduledAlert{
	//			{0 * time.Second, 100},
	//			{10 * time.Minute, 200},
	//			{20 * time.Minute, 300},
	//			{30 * time.Minute, 400},
	//			{40 * time.Minute, 500},
	//			{50 * time.Minute, 600},
	//			{60 * time.Minute, 800},
	//			{70 * time.Minute, 1000},
	//			{80 * time.Minute, 2000},
	//			{90 * time.Minute, 4000},
	//			{100 * time.Minute, 8000},
	//		}
	//
	//		for i, want := range cases {
	//			t.Run(fmt.Sprint(want), func(t *testing.T) {
	//
	//				if len(blindAlerter.alerts) <= i {
	//					t.Fatalf("alert %d was not scheduled %v", i, blindAlerter.alerts)
	//				}
	//
	//				got := blindAlerter.alerts[i]
	//				assertScheduledAlert(t, got, want)
	//			})
	//		}
	//	})

	t.Run("it prompts the user to enter the number of players", func(t *testing.T) {
		in := strings.NewReader("7\n")
		game := poker.NewGame(dummySpyAlerter, dummyPlayerStore)

		cli := poker.NewCLI(in, dummyStdOut, game)
		cli.PlayPoker()

		got := dummyStdOut.String()
		want := poker.PlayerPrompt

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

		cases := []scheduledAlert{
			{0 * time.Second, 100},
			{12 * time.Minute, 200},
			{24 * time.Minute, 300},
			{36 * time.Minute, 400},
		}

		for i, want := range cases {
			t.Run(fmt.Sprint(want), func(t *testing.T) {
				if len(dummySpyAlerter.alerts) <= i {
					t.Fatalf("alert %d was not scheduled %v", i, dummySpyAlerter.alerts)
				}

				got := dummySpyAlerter.alerts[i]
				assertScheduledAlert(t, got, want)
			})
		}
	})
}

func assertScheduledAlert(t testing.TB, got, want scheduledAlert) {
	t.Helper()
	if got != want {
		t.Errorf("got wrong scheduledAlert, got %v want %v", got, want)
	}
}
