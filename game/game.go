package game

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	Board         *Board
	CurrentPlayer *Player
	Player1       *Player
	Player2       *Player
}

func NewGame(player1, player2 *Player) *Game {
	return &Game{
		Board:         NewBoard(),
		CurrentPlayer: player1,
		Player1:       player1,
		Player2:       player2,
	}
}

func (g *Game) switchPlayer(player1, player2 *Player) {
	if g.CurrentPlayer == player1 {
		g.CurrentPlayer = player2
	} else {
		g.CurrentPlayer = player1
	}
}

func (b *Board) MakeMove(x int, y int, player *Player) error {
	if b.Cells[x][y] != ' ' {
		return fmt.Errorf("Field is already occupied")
	}

	b.Cells[x][y] = player.Symbol
	return nil
}

func (g *Game) PlayGame() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		g.Board.PrintBoard()
		fmt.Printf("%s (enter 'x, y'): ", g.CurrentPlayer.Name)
		scanner.Scan()
		move := scanner.Text()

		coords := strings.Split(move, ",")
		if len(coords) != 2 {
			fmt.Println("Incorrect movement")
			continue
		}

		x, errX := strconv.Atoi(coords[0])
		y, errY := strconv.Atoi(coords[1])
		if errX != nil || errY != nil || x < 0 || x > 2 || y < 0 || y > 2 {
			fmt.Println("Incorrect movement")
			continue
		}

		err := g.Board.MakeMove(x, y, g.CurrentPlayer)
		if err != nil {
			fmt.Println(err)
			continue
		}

		if g.Board.CheckWin(g.CurrentPlayer.Symbol) {
			g.Board.PrintBoard()
			fmt.Printf("%s won!\n", g.CurrentPlayer.Name)
			break
		}

		if g.Board.IsFull() {
			g.Board.PrintBoard()
			fmt.Println("Draw!")
			break
		}

		g.switchPlayer(g.Player1, g.Player2)
	}
}
