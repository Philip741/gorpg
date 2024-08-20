package main

import (
	"go-learning/textrpg/game"
	"log"
)

func main() {
	// initialize the game
	g, err := game.New()
	if err != nil {
		log.Fatal(err)
	}
}
