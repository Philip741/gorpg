package main

import (
	"log"

	//"github.com/Philip741/gorpg/game"
	"github.com/Philip741/gorpg/ui"
)

func main() {
	// initialize the game
	//g, err := game.New()
	//if err != nil {
	//log.Fatal(err)
	//}

	// initialize the ui
	gameUI, err := ui.New()
	if err != nil {
		log.Fatalf("Failed to initialize ui: %v", err)
	}
	gameUI.UpdateGraphics("Welcome to the TextRPG!")
	gameUI.UpdateCharacterStats("Character Stats will appear here")
	gameUI.AppendGameText("Your adventure begins...")
	//gameUI.UpdateActions("[Move] [Attack] [Inventory]")
	// run the ui
	//gameUI.Run()
	go func() {
		if err := gameUI.Run(); err != nil {
			log.Fatalf("Error running UI: %v", err)
		}
	}()

	// load default image
	if err := gameUI.UpdateGraphics("../assets/images/labyrinth.jpeg"); err != nil {
		log.Printf("Error setting image: %v", err)
	}
	// Main game loop
	for {
		select {
		case input := <-gameUI.GetInputChannel():
			switch input {
			case "move":
				gameUI.AppendGameText("You move forward.")
			case "attack":
				gameUI.AppendGameText("You attack.")
			case "quit":
				gameUI.AppendGameText("Thanks for playing!")
				defer gameUI.Stop()
				return
			}
		}
	}

}
