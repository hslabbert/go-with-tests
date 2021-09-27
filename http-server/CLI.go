package poker

import (
	"bufio"
	"io"
	"strings"
)

// A CLI wraps a playerstore and supports reading an io.Reader
// to record user input.
type CLI struct {
	playerstore PlayerStore
	in          io.Reader
}

// PlayPoker records a win on the playerStore in the provided
// *CLI
func (cli *CLI) PlayPoker() {

	reader := bufio.NewScanner(cli.in)
	reader.Scan()

	player := extractWinner(reader.Text())
	cli.playerstore.RecordWin(player)
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}
