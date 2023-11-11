package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"tictactoe/game"
)

func main() {
	gameInstance := game.NewBoard()

	player1 := game.NewPlayer('X', "Player 1")
	player2 := game.NewPlayer('O', "Player 2")

	currentPlayer := player1
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Obecny gracz:", currentPlayer.Name)
		gameInstance.PrintBoard()

		fmt.Print("Wpisz współrzędne ruchu (np. 1,2): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		coords := strings.Split(input, ",")

		if len(coords) != 2 {
			fmt.Println("Nieprawidłowy format współrzędnych. Spróbuj ponownie.")
			continue
		}

		x, err := strconv.Atoi(coords[0])
		if err != nil {
			fmt.Println("Nieprawidłowa współrzędna X. Spróbuj ponownie.")
			continue
		}

		y, err := strconv.Atoi(coords[1])
		if err != nil {
			fmt.Println("Nieprawidłowa współrzędna Y. Spróbuj ponownie.")
			continue
		}

		err = gameInstance.MakeMove(x, y, currentPlayer)
		if err != nil {
			fmt.Println(err)
			continue
		}

		if gameInstance.CheckWin(currentPlayer.Symbol) {
			gameInstance.PrintBoard()
			fmt.Printf("Gracz %s wygrywa!\n", currentPlayer.Name)
			break
		}

		if gameInstance.IsFull() {
			gameInstance.PrintBoard()
			fmt.Println("Remis!")
			break
		}

		if currentPlayer == player1 {
			currentPlayer = player2
		} else {
			currentPlayer = player1
		}
	}
}
