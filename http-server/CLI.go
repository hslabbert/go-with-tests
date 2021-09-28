package poker

import (
	"bufio"
	"io"
	"strings"
	"time"
)

// A CLI wraps a playerstore and supports reading an io.Reader
// to record user input.
type CLI struct {
	playerStore PlayerStore
	in          *bufio.Scanner
}

type BlindAlerter interface {
	ScheduleAlertAt(duration time.Duration, amount int)
}

// NewCLI constructs a *CLI from the provided PlayerStore and
// user input.
func NewCLI(store PlayerStore, in io.Reader, alerter BlindAlerter) *CLI {
	return &CLI{
		playerStore: store,
		in:          bufio.NewScanner(in),
	}
}

// PlayPoker records a win on the playerStore in the provided
// *CLI
func (cli *CLI) PlayPoker() {
	userInput := cli.readLine()

	cli.playerStore.RecordWin(extractWinner(userInput))
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}

func (cli *CLI) readLine() string {
	cli.in.Scan()
	return cli.in.Text()
}
