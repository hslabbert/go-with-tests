package main

import (
	"fmt"
	"log"
	"os"

	poker "github.com/hslabbert/go-with-tests/http-server"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}
	defer close()

	game := poker.NewTexasHoldem(poker.BlindAlerterFunc(poker.StdOutAlerter), store)
	cli := poker.NewCLI(os.Stdin, os.Stdout, game)
	fmt.Println("Let's play poker")
	fmt.Println("type {Name} wins to record a win")
	cli.PlayPoker()
}
