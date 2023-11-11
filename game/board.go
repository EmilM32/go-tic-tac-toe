package game

import "fmt"

type Board struct {
	Cells [3][3]rune
}

func NewBoard() *Board {
	return &Board{
		Cells: [3][3]rune{{' ', ' ', ' '}, {' ', ' ', ' '}, {' ', ' ', ' '}},
	}
}

func (b *Board) CheckWin(symbol rune) bool {
	// Sprawdzanie wierszy
	for i := 0; i < 3; i++ {
		if b.Cells[i][0] == symbol && b.Cells[i][1] == symbol && b.Cells[i][2] == symbol {
			return true
		}
	}

	// Sprawdzanie kolumn
	for i := 0; i < 3; i++ {
		if b.Cells[0][i] == symbol && b.Cells[1][i] == symbol && b.Cells[2][i] == symbol {
			return true
		}
	}

	// Sprawdzanie przekątnych
	if b.Cells[0][0] == symbol && b.Cells[1][1] == symbol && b.Cells[2][2] == symbol {
		return true
	}
	if b.Cells[0][2] == symbol && b.Cells[1][1] == symbol && b.Cells[2][0] == symbol {
		return true
	}

	return false
}

func (b *Board) IsFull() bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if b.Cells[i][j] == ' ' {
				return false
			}
		}
	}
	return true
}

func (b *Board) PrintBoard() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%c", b.Cells[i][j])
			if j < 2 {
				fmt.Printf("|")
			}
		}
		if i < 2 {
			fmt.Printf("\n-+-+-\n")
		} else {
			fmt.Println()
		}
	}
}
