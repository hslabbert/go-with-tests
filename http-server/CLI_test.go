package poker_test

import (
	"fmt"
	"strings"
	"testing"
	"time"

	poker "github.com/hslabbert/go-with-tests/http-server"
)

var dummySpyAlerter = &SpyBlindAlerter{}

type SpyBlindAlerter struct {
	alerts []struct {
		scheduledAt time.Duration
		amount      int
	}
}

func (s *SpyBlindAlerter) ScheduleAlertAt(duration time.Duration, amount int) {
	s.alerts = append(s.alerts, struct {
		scheduledAt time.Duration
		amount      int
	}{duration, amount})
}

func TestCLI(t *testing.T) {
	cases := []string{
		"Cleo",
		"Chris",
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("record %s win from user input", c), func(t *testing.T) {
			in := strings.NewReader(fmt.Sprintf("%s wins\n", c))
			playerStore := &poker.StubPlayerStore{}

			cli := poker.NewCLI(playerStore, in, dummySpyAlerter)
			cli.PlayPoker()

			poker.AssertPlayerWin(t, playerStore, c)
		})
	}

	t.Run("it schedules present of blind values", func(t *testing.T) {
		in := strings.NewReader("Chris wins \n")
		playerStore := &poker.StubPlayerStore{}
		blindAlerter := &SpyBlindAlerter{}

		cli := poker.NewCLI(playerStore, in, blindAlerter)
		cli.PlayPoker()

		if len(blindAlerter.alerts) != 1 {
			t.Fatalf("expected a blind alert to be scheduled")
		}
	})
}
