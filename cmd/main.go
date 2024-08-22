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
	gameUI.UpdateActions("[Move] [Attack] [Inventory]")
	// run the ui
	gameUI.Run()

}
