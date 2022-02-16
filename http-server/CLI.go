package poker

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// PlayerPrompt is simply a string fired at the start of a
// Game to provide the number of players in the Game.
const PlayerPrompt = "Please enter the number of players: "

// BadPlayerInputErrMsg is returned when the wrong type of value is
// supplied for starting a game, in providing the player count.
const BadPlayerInputErrMsg = "Bad value received for number of players, please try again with a number"

// BadWinnerInputErrMsg is returned when the winner's name is
// inputted incorrectly during a game.
const BadWinnerInputErrMsg = "Bad value for winner, please try again with the string '<winner> wins'\n"

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
		fmt.Fprint(cli.out, BadPlayerInputErrMsg)
		return
	}

	cli.game.Start(numberOfPlayers, cli.out)

	var winner string
	for winner == "" {
		winnerInput := cli.readLine()
		winner, err = extractWinner(winnerInput)
		if err != nil {
			fmt.Fprint(cli.out, err)
		}
	}

	cli.game.Finish(winner)
}

func extractWinner(userInput string) (string, error) {
	err := validateWinnerInput(userInput)
	if err != nil {
		return "", err
	}
	return strings.Replace(userInput, " wins", "", 1), nil
}

func (cli *CLI) readLine() string {
	cli.in.Scan()
	return cli.in.Text()
}

func validateWinnerInput(userInput string) error {
	s := strings.Split(userInput, " ")
	if len(s) != 2 && userInput[len(userInput)-4:] != "wins" {
		err := errors.New(BadWinnerInputErrMsg)
		return err
	}
	return nil
}
