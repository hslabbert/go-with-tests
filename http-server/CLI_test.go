package poker

import (
	"fmt"
	"strings"
	"testing"
)

func TestCLI(t *testing.T) {
	cases := []string{
		"Cleo",
		"Chris",
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("record %s win from user input", c), func(t *testing.T) {
			in := strings.NewReader(fmt.Sprintf("%s wins\n", c))
			playerStore := &StubPlayerStore{}

			cli := &CLI{playerStore, in}
			cli.PlayPoker()

			assertPlayerWin(t, playerStore, c)
		})
	}
}
