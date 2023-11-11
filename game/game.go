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
	CurrentPlayer rune
}

func NewGame() *Game {
	return &Game{
		Board:         NewBoard(),
		CurrentPlayer: 'X',
	}
}

func (b *Board) MakeMove(x int, y int, player *Player) error {
	if b.Cells[x][y] != ' ' {
		return fmt.Errorf("pole jest już zajęte")
	}

	b.Cells[x][y] = player.Symbol
	return nil
}

func (g *Game) PlayGame(player1, player2 *Player) {
	scanner := bufio.NewScanner(os.Stdin)

	var currentPlayer *Player = player1

	for {
		g.Board.PrintBoard()
		fmt.Printf("Ruch gracza %s (wprowadź 'x y'): ", currentPlayer.Name)
		scanner.Scan()
		move := scanner.Text()

		coords := strings.Split(move, ",")
		if len(coords) != 2 {
			fmt.Println("Nieprawidłowy ruch")
			continue
		}

		x, errX := strconv.Atoi(coords[0])
		y, errY := strconv.Atoi(coords[1])
		if errX != nil || errY != nil || x < 0 || x > 2 || y < 0 || y > 2 {
			fmt.Println("Nieprawidłowy ruch")
			continue
		}

		err := g.Board.MakeMove(x, y, currentPlayer)
		if err != nil {
			fmt.Println(err)
			continue
		}

		// Sprawdzenie warunków zwycięstwa
		if g.Board.CheckWin(currentPlayer.Symbol) {
			g.Board.PrintBoard()
			fmt.Printf("Wygrana gracza %s!\n", currentPlayer.Name)
			break
		}

		// Sprawdzenie remisu
		if g.Board.IsFull() {
			g.Board.PrintBoard()
			fmt.Println("Remis!")
			break
		}

		// Zmiana gracza
		if currentPlayer == player1 {
			currentPlayer = player2
		} else {
			currentPlayer = player1
		}
	}
}
