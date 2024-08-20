package character

type Character struct {
	Name string
	HP   int
	MP   int
}

func NewCharacter(name string, hp int, mp int) (*Character, error) {
	return &Character{
		Name: name,
		HP:   hp,
		MP:   mp,
	}, nil
}
