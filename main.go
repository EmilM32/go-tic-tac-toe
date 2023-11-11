package main

import (
	"bufio"
	"fmt"
	"os"
	"tictactoe/game"
)

func main() {
	fmt.Println("Welcome to Tic Tac Toe!")

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("Enter Player1 name: ")
	scanner.Scan()
	player1Name := scanner.Text()

	fmt.Printf("Enter Player2 name: ")
	scanner.Scan()
	player2Name := scanner.Text()

	player1 := game.NewPlayer('X', player1Name)
	player2 := game.NewPlayer('O', player2Name)

	gameInstance := game.NewGame(player1, player2)
	gameInstance.PlayGame()
}
