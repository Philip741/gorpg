package main

import (
	"log"

	"github.com/Philip741/gorpg/game"
)

func main() {
	// initialize the game
	g, err := game.New()
	if err != nil {
		log.Fatal(err)
	}
}
