package main

import (
	"fmt"
	"log"
	"strconv"
)

const (
	rows = 6
	cols = 7
)

func main() {
	board := [][]string{}
	for i := 0; i < cols; i++ {
		col := make([]string, 0, cols)
		board = append(board, col)
	}

	playerTurn := "A"
	input := ""

	fmt.Println("Hello Connect 4")

	for {
		PrintBoard(board, playerTurn)

		if CheckWin(board, NextTurn(playerTurn)) {
			fmt.Println("Player ", NextTurn(playerTurn), "wins the game!!!")
			break
		}

		_, err := fmt.Scanln(&input)
		if err != nil {
			log.Panic("An error occurred")
		}

		if input == "exit" {
			fmt.Println("Game Over!")
			return
		}

		inputCol, err := strconv.Atoi(input)
		if err != nil || inputCol < 1 || inputCol > 7 {
			fmt.Println("incorrect input must be between 1 and", cols)
			continue
		}

		inputCol--
		col := board[inputCol]

		if len(col) < rows {
			col := append(col, playerTurn)
			board[inputCol] = col
		} else {
			fmt.Println("this column is already full", cols)
		}

		playerTurn = NextTurn(playerTurn)
	}
}

func NextTurn(playerTurn string) string {
	if playerTurn == "A" {
		return "B"
	}

	return "A"
}

const PrevLine = "\033[F"

var firstPaint = true

func PrintBoard(board [][]string, playerTurn string) {
	rewind := ""
	if firstPaint {
		firstPaint = false
	} else {
		for i := 0; i <= rows+1; i++ {
			rewind += PrevLine
		}
	}
	fmt.Printf("%vConnect 4: player %v turn\n", rewind, playerTurn)

	for i := rows - 1; i >= 0; i-- {
		for j := 0; j < cols; j++ {
			l := len(board[j])
			if l > 0 && i < l {
				fmt.Printf("%v ", board[j][i])
			} else {
				fmt.Printf("- ")
			}
		}
		fmt.Printf("\n")
	}
}

func checkColRows(col []string) bool {
	rows := len(col)
	return rows < 4
}

func CheckWin(board [][]string, player string) bool {
	// Check vertically
	for col := 0; col < cols; col++ {
		colRows := len(board[col])
		if colRows < 4 {
			continue
		}

		for row := 0; row <= colRows-4; row++ {
			if board[col][row] == player &&
				board[col][row+1] == player &&
				board[col][row+2] == player &&
				board[col][row+3] == player {
				return true
			}
		}
	}

	// Check horizontally
	for col := 0; col <= cols-4; col++ {
		col1Rows := len(board[col])
		col2Rows := len(board[col+1])
		col3Rows := len(board[col+2])
		col4Rows := len(board[col+3])

		if col1Rows <= 0 || col2Rows <= 0 || col3Rows <= 0 || col4Rows <= 0 {
			continue
		}

		for row := 0; row < rows; row++ {
			if row >= col1Rows || row >= col2Rows || row >= col3Rows || row >= col4Rows {
				break
			}

			if board[col][row] == player &&
				board[col+1][row] == player &&
				board[col+2][row] == player &&
				board[col+3][row] == player {
				return true
			}
		}
	}

	for col := 0; col <= cols-4; col++ {
		col1Rows := len(board[col])
		col2Rows := len(board[col+1])
		col3Rows := len(board[col+2])
		col4Rows := len(board[col+3])

		if col1Rows <= 0 || col2Rows <= 0 || col3Rows <= 0 || col4Rows <= 0 {
			continue
		}

		for row := 0; row <= rows-4; row++ {
			if row >= col1Rows || row+1 >= col2Rows || row+2 >= col3Rows || row+3 >= col4Rows {
				break
			}

			if board[col][row] == player &&
				board[col+1][row+1] == player &&
				board[col+2][row+2] == player &&
				board[col+3][row+3] == player {
				return true
			}
		}
	}

	for col := 0; col <= cols-4; col++ {
		col1Rows := len(board[col])
		col2Rows := len(board[col+1])
		col3Rows := len(board[col+2])
		col4Rows := len(board[col+3])

		if col1Rows <= 0 || col2Rows <= 0 || col3Rows <= 0 || col4Rows <= 0 {
			continue
		}

		for row := 3; row < rows; row++ {
			if row >= col1Rows || row-1 >= col2Rows || row-2 >= col3Rows || row-3 >= col4Rows {
				break
			}

			if board[col][row] == player &&
				board[col+1][row-1] == player &&
				board[col+2][row-2] == player &&
				board[col+3][row-3] == player {
				return true
			}
		}
	}

	// Check diagonal (bottom-left to top-right)
	// for row := 3; row < rows; row++ {
	// 	for col := 0; col <= cols-4; col++ {
	// 		if board[row][col] == player &&
	// 			board[row-1][col+1] == player &&
	// 			board[row-2][col+2] == player &&
	// 			board[row-3][col+3] == player {
	// 			return true
	// 		}
	// 	}
	// }

	return false
}
