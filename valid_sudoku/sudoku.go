package validsudoku

import (
	"fmt"
)

type (
	Empty struct{}
	Set   map[string]Empty
)

func ValidSudoku(board [9][9]string) bool {
	rows := make(map[int]Set, 9)
	cols := make(map[int]Set, 9)
	squares := make(map[string]Set, 9)
	// Initialize each map in the array
	for i := 0; i < 9; i++ {
		rows[i] = Set{}
		cols[i] = Set{}
	}

	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			squareKey := fmt.Sprintf("%d,%d", r, c)
			squares[squareKey] = Set{}
		}
	}

	for r := 0; r < len(board); r++ {
		for c := 0; c < len(board); c++ {
			cellval := board[r][c]
			if cellval == "." {
				continue
			}
			if _, existsInRow := rows[r][cellval]; existsInRow {
				fmt.Println("Exist in row", cellval)
				return false
			}
			if _, existsInCol := cols[c][cellval]; existsInCol {
				fmt.Println("Exist in col", cellval)
				return false
			}
			squareKey := fmt.Sprintf("%d,%d", r, c)
			fmt.Println("squareKey", squareKey)
			if _, ok := squares[squareKey][cellval]; ok {
				return false
			}

			fmt.Println("CHECK 1")
			rows[r][cellval] = Empty{}
			fmt.Println("CHECK 2")
			cols[c][cellval] = Empty{}
			fmt.Println("CHECK 3")
			squares[squareKey][cellval] = Empty{}

		}
	}

	return true
}
