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

type Column struct {
	Rows [rows]string
}

func (c *Column) Length() int {
	for i := 0; i < len(c.Rows); i++ {
		if c.Rows[i] != "" {
			return i + 1
		}
	}

	return 0
}

func (c *Column) Push(val string) {
	for i := 0; i < len(c.Rows); i++ {
		if c.Rows[i] != "" {
			continue
		}

		c.Rows[i] = val
		break
	}
}

type Board struct {
	Columns [cols]Column
}

func (b *Board) Point(col, row int) string {
	return b.Columns[col].Rows[row]
}

func main() {
	board := Board{}
	for i := 0; i < cols; i++ {
		board.Columns[i] = Column{}
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
		col := board.Columns[inputCol]
		if col.Length() < rows {
			col.Push(playerTurn)
			board.Columns[inputCol] = col
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

func PrintBoard(board Board, playerTurn string) {
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
			col := board.Columns[j]
			if col.Rows[i] == "" {
				fmt.Printf("- ")
			} else {
				fmt.Printf("%v ", col.Rows[i])
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
