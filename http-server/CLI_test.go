package poker_test

import (
	"fmt"
	"strings"
	"testing"

	poker "github.com/hslabbert/go-with-tests/http-server"
)

func TestCLI(t *testing.T) {
	cases := []string{
		"Cleo",
		"Chris",
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("record %s win from user input", c), func(t *testing.T) {
			in := strings.NewReader(fmt.Sprintf("%s wins\n", c))
			playerStore := &poker.StubPlayerStore{}

			cli := poker.NewCLI(playerStore, in)
			cli.PlayPoker()

			poker.AssertPlayerWin(t, playerStore, c)
		})
	}
}
