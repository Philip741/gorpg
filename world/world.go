package world

import (
	"fmt"
)

type World struct {
	mainroom string
}

func NewWorld() (*World, error) {
	fmt.Println("World created")
	return &World{}, nil
}
