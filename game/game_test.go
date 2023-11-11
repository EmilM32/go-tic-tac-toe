package game

import (
	"testing"
)

func TestMakeMove(t *testing.T) {
	board := NewBoard()
	player := NewPlayer('X', "Player 1")
	err := board.MakeMove(0, 0, player)

	if err != nil {
		t.Errorf("MakeMove(0, 0, player) failed, expected no error, got %v", err)
	}

	if board.Cells[0][0] != player.Symbol {
		t.Errorf("Expected board.Cells[0][0] to be %v, got %v", player.Symbol, board.Cells[0][0])
	}
}

func TestSwitchPlayer(t *testing.T) {
	player1 := NewPlayer('X', "Player 1")
	player2 := NewPlayer('O', "Player 2")
	game := NewGame(player1, player2)

	if game.CurrentPlayer != player1 {
		t.Errorf("Expected current player to be %v, but got %v", player1, game.CurrentPlayer)
	}

	game.switchPlayer()
	if game.CurrentPlayer != player2 {
		t.Errorf("Expected current player to be %v, but got %v", player2, game.CurrentPlayer)
	}

	game.switchPlayer()
	if game.CurrentPlayer != player1 {
		t.Errorf("Expected current player to be %v, but got %v", player1, game.CurrentPlayer)
	}
}
