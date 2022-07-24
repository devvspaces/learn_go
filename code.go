package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func removeIndexFromArr(arr []Position, index int) []Position {
	return append(arr[:index], arr[index+1:]...)
}

func getPosition(positions *[]Position) Position {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	index := r1.Intn(len(*positions))
	new := (*positions)[index]
	*positions = removeIndexFromArr(*positions, index)
	return new
}

func getStartPlayer() string {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	if r1.Intn(2) == 1 {
		return "X"
	}
	return "O"
}

func playBoard(board [][]string, position Position, player string) {
	board[position.x][position.y] = player
}

func switchPlay(player string) string {
	if player == "X" {
		return "O"
	}
	return "X"
}

func checkRows(board [][]string) string {
	for row := 0; row < 3; row++ {
		side := ""
		for col := 0; col < 3; col++ {
			if board[row][col] == "_" || (side != "" && side != board[row][col]) {
				break
			}
			side = board[row][col]
			if col == 2 {
				return side
			}
		}

	}
	return ""
}

func checkColumns(board [][]string) string {
	for row := 0; row < 3; row++ {
		side := ""
		for col := 0; col < 3; col++ {
			if board[col][row] == "_" || (side != "" && side != board[col][row]) {
				break
			}
			side = board[col][row]
			if col == 2 {
				return side
			}
		}

	}
	return ""
}

func checkWinner(board [][]string) string {
	if side := checkRows(board); side != "" {
		return side
	} else if side := checkColumns(board); side != "" {
		return side
	}
	return ""
}

type Position struct {
	x int
	y int
}

func main() {
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	positions := make([]Position, 9)

	for i := 0; i < 9; i++ {
		positions[i].x = i % 3
		positions[i].y = i / 3
	}

	// The players take turns.
	nextPlayer := getStartPlayer()
	winner := "Draw"
	for i := 0; i < 9; i++ {
		playBoard(board, getPosition(&positions), nextPlayer)
		nextPlayer = switchPlay(nextPlayer)
		if win := checkWinner(board); win != "" {
			winner = win
			break
		}
	}

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	fmt.Println("\nThe winner is: ", winner)
}
