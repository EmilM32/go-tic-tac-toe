package game

import (
	"testing"
)

func TestCheckWin(t *testing.T) {
	game := NewGame(NewPlayer('X', "Player 1"), NewPlayer('O', "Player 2"))
	game.Board.Cells = [3][3]rune{
		{'X', 'X', 'X'},
		{' ', ' ', ' '},
		{' ', ' ', ' '},
	}

	if !game.Board.CheckWin('X') {
		t.Errorf("Expected CheckWin('X') to be true, got false")
	}
}

func TestIsFull(t *testing.T) {
	game := NewGame(NewPlayer('X', "Player 1"), NewPlayer('O', "Player 2"))
	game.Board.Cells = [3][3]rune{
		{'X', 'O', 'X'},
		{'X', 'X', 'O'},
		{'O', 'X', 'O'},
	}

	if !game.Board.IsFull() {
		t.Errorf("Expected IsFull() to be true, got false")
	}
}
