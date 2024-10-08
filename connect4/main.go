package main

import (
	"fmt"
	"log"
	"strconv"
)

const (
	rows     = 6
	cols     = 7
	PrevLine = "\033[F"
)

var firstPaint = true

type Column []string

func NewColumn(rows int) Column {
	col := Column{}
	col = make([]string, rows)

	return col
}

func (c Column) Len() int {
	for i := 0; i < len(c); i++ {
		if c[i] != "" {
			return i + 1
		}
	}

	return 0
}

func (c Column) Push(val string) {
	for i := 0; i < len(c); i++ {
		if c[i] != "" {
			continue
		}

		c[i] = val
		break
	}
}

type Board []Column

func (b Board) Point(col, row int) string {
	return b[col][row]
}

func NewBoard(cols, rows int) Board {
	board := Board{}
	board = make([]Column, cols)

	for i := 0; i < cols; i++ {
		board[i] = NewColumn(rows)
	}

	return board
}

func main() {
	board := NewBoard(cols, rows)
	playerTurn := NextTurn("")
	input := ""

	fmt.Println("Welcome to  Connect 4")

	for {
		PrintBoard(board, playerTurn)

		if CheckWin(board, NextTurn(playerTurn)) {
			prevPlayer := PlayerName(NextTurn(playerTurn))
			fmt.Println("\n🎖️🎖️🎖️Player", prevPlayer, "wins the game🎖️🎖️🎖️")
			break
		}

		_, err := fmt.Scanln(&input)
		if err != nil {
			if err.Error() == "unexpected newline" {
				fmt.Println("No move detected", PrevLine)
				continue
			}

			log.Panic("An error has occurred", err)
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
		if col.Len() < rows {
			col.Push(playerTurn)
			// board.Columns[inputCol] = col
		} else {
			fmt.Println("this column is already full", cols)
		}

		playerTurn = NextTurn(playerTurn)
	}
}

func PlayerName(player string) string {
	if player == "D" {
		return "DIANA 👸"
	}

	return "EDGAR"
}

func NextTurn(playerTurn string) string {
	if playerTurn == "D" {
		return "E"
	}

	return "D"
}

func PrintBoard(board Board, playerTurn string) {
	rewind := ""
	if firstPaint {
		firstPaint = false
	} else {
		for i := 0; i <= rows+1; i++ {
			rewind += PrevLine
		}
	}
	fmt.Printf("%vConnect 4: %v turn\n", rewind, PlayerName(playerTurn))

	for i := rows - 1; i >= 0; i-- {
		for j := 0; j < cols; j++ {
			col := board[j]
			if col[i] == "" {
				fmt.Printf("- ")
			} else {
				fmt.Printf("%v ", col[i])
			}
		}

		fmt.Printf("\n")
	}
}

func CheckWin(board Board, player string) bool {
	// Check vertically
	for col := 0; col < cols; col++ {
		for row := 0; row <= rows-4; row++ {
			if board.Point(col, row) == player &&
				board.Point(col, row+1) == player &&
				board.Point(col, row+2) == player &&
				board.Point(col, row+3) == player {
				return true
			}
		}
	}

	// Check horizontally
	for col := 0; col <= cols-4; col++ {
		for row := 0; row < rows; row++ {
			if board.Point(col, row) == player &&
				board.Point(col+1, row) == player &&
				board.Point(col+2, row) == player &&
				board.Point(col+3, row) == player {
				return true
			}
		}
	}

	// Check upslope diagonal
	for col := 0; col <= cols-4; col++ {
		for row := 0; row <= rows-4; row++ {
			if board.Point(col, row) == player &&
				board.Point(col+1, row+1) == player &&
				board.Point(col+2, row+2) == player &&
				board.Point(col+3, row+3) == player {
				return true
			}
		}
	}

	// Check downslope diagonal
	for col := 0; col <= cols-4; col++ {
		for row := 3; row < rows; row++ {
			if board.Point(col, row) == player &&
				board.Point(col+1, row-1) == player &&
				board.Point(col+2, row-2) == player &&
				board.Point(col+3, row-3) == player {
				return true
			}
		}
	}

	return false
}
