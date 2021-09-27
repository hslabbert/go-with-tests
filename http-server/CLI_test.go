package poker

import (
	"strings"
	"testing"
)

func TestCLI(t *testing.T) {
	in := strings.NewReader("Chris wins\n")

	playerstore := &StubPlayerStore{}
	cli := &CLI{playerstore, in}
	cli.PlayPoker()

	assertPlayerWin(t, playerstore, "Chris")
}
