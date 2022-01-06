package main

import "fmt"

var sudoku [9][9]int

func sudokuSolver() bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if sudoku[i][j] == 0 {
				for k := 1; k <= 9; k++ {
					if !numberVerifier(k, i, j) {
						sudoku[i][j] = k

						//printer()
						//fmt.Println()

						if sudokuSolver() {
							return true
						} else {
							sudoku[i][j] = 0
						}
					}
				}

				return false
			}
		}
	}

	return true
}

func numberVerifier(num int, row int, column int) bool {
	for i := 0; i < 9; i++ {
		if num == sudoku[row][i] {
			return true
		}
	}

	for i := 0; i < 9; i++ {
		if num == sudoku[i][column] {
			return true
		}
	}

	y := row - (row % 3)
	x := column - (column % 3)

	for i := y; i < y+3; i++ {
		for j := x; j < x+3; j++ {
			if num == sudoku[i][j] {
				return true
			}
		}
	}

	return false
}

func printer() {
	fmt.Println("┌───────┬───────┬───────┐")

	for i := 0; i < 9; i++ {
		if i%3 == 0 && i != 0 {
			fmt.Println("├───────┼───────┼───────┤")
		}
		for j := 0; j < 9; j++ {
			if j == 0 || j%3 == 0 {
				fmt.Print("| ")
			}
			fmt.Printf("%d ", sudoku[i][j])
		}
		fmt.Println("|")
	}
	fmt.Println("└───────┴───────┴───────┘")
}

func main() {
	sudoku = [9][9]int{{8, 0, 7, 0, 9, 0, 2, 0, 4},
		{0, 0, 0, 0, 0, 4, 0, 0, 0},
		{5, 0, 4, 2, 0, 8, 3, 0, 9},
		{0, 7, 3, 0, 0, 0, 4, 0, 0},
		{2, 0, 0, 0, 0, 0, 0, 0, 7},
		{0, 0, 6, 0, 0, 0, 1, 3, 0},
		{4, 0, 2, 9, 0, 7, 8, 0, 6},
		{0, 0, 0, 4, 0, 0, 0, 0, 0},
		{7, 0, 8, 0, 1, 0, 9, 0, 3}}

	printer()
	sudokuSolver()
	fmt.Println()
	printer()
}
