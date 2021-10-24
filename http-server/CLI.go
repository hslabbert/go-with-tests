package poker

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// PlayerPrompt is simply a string fired at the start of a
// Game to provide the number of players in the Game.
const PlayerPrompt = "Please enter the number of players: "

// A Game implements Start and Finish methods to start a game
// with the specified number of players, and finish a game by recording
// the winner.
type Game interface {
	Start(numberOfPlayers int)
	Finish(winner string)
}

// A CLI wraps a playerstore and supports reading an io.Reader
// to record user input.
type CLI struct {
	in   *bufio.Scanner
	out  io.Writer
	game Game
}

// NewCLI constructs a *CLI from the provided PlayerStore and
// user input.
func NewCLI(in io.Reader, out io.Writer, game Game) *CLI {
	return &CLI{
		in:   bufio.NewScanner(in),
		out:  out,
		game: game,
	}
}

// PlayPoker records a win on the playerStore in the provided
// *CLI
func (cli *CLI) PlayPoker() {
	fmt.Fprint(cli.out, PlayerPrompt)

	numberOfPlayers, err := strconv.Atoi(cli.readLine())

	if err != nil {
		return
	}

	cli.game.Start(numberOfPlayers)
	winnerInput := cli.readLine()
	winner := extractWinner(winnerInput)

	cli.game.Finish(winner)
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}

func (cli *CLI) readLine() string {
	cli.in.Scan()
	return cli.in.Text()
}
