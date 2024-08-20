package game

import (
	"fmt"

	"github.com/Philip741/gorpg/character"
	"github.com/Philip741/gorpg/world"
)

type Game struct {
	Player    *character.Character
	World     *world.World
	GameState string
}

func New() (*Game, error) {
	fmt.Println("Welcome to the TextRPG!")
	player, err := character.NewCharacter("Player", 100, 50)
	if err != nil {
		return nil, err
	}
	world, err := world.NewWorld()
	if err != nil {
		return nil, err
	}
	return &Game{
		Player:    player,
		World:     world,
		GameState: "playing",
	}, nil
}
