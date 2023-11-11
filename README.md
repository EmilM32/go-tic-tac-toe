# Tic-Tac-Toe Terminal Game

## Description
This Tic-Tac-Toe game is a simple, terminal-based version of the classic game. It's implemented in Go and allows two players to play the game on the command line. The game features a basic text interface and turn-based gameplay.

## Usage
To play the game, run the following command in the terminal:

```bash
go run main.go
```

## Testing
To run the tests, run the following command in the terminal:

```bash
go test ./...
```

## How to Play

- The game starts with an empty 3x3 grid.
- Player 1 uses the 'X' symbol, and Player 2 uses the 'O' symbol.
- Players take turns to place their symbols on the grid.
- The first player to align three of their symbols horizontally, vertically, or diagonally wins the game.
- If the grid is filled and no player has aligned three symbols, the game ends in a draw.