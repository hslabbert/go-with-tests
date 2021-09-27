package poker

import "io"

type CLI struct {
	playerstore PlayerStore
	in          io.Reader
}

// PlayPoker records a win on the playerStore in the provided
// *CLI
func (cli *CLI) PlayPoker() {
	//var player []byte
	//cli.in.Read(player)
	_ = cli.playerstore.RecordWin("Chris")

}
