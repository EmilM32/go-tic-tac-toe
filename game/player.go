package game

type Player struct {
	Symbol rune
	Name   string
}

func NewPlayer(symbol rune, name string) *Player {
	return &Player{
		Symbol: symbol,
		Name:   name,
	}
}
