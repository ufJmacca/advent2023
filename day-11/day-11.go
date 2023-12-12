package main

import "fmt"

type coords struct {
	x int
	y int
}

func GridExpansion(input [][]string) [][]string {
	rows := len(input)
	cols := len(input[0])

	fmt.Println(rows)

	copied_input := make([][]string, rows)
	for i := range input {
		copied_input[i] = make([]string, cols)
		copy(copied_input[i], input[i])
	}

	row_duplicates_required := 0
	for i := 0; i < rows; i++ {
		duplicate := true
		for j := 0; j < cols; j++ {
			if input[i][j] != "." {
				duplicate = false
				fmt.Printf("# in row - %d\n", i)
			}
		}
		if duplicate {
			fmt.Printf("adding a row - %s\n", input[i])
			slice := []string(input[i])
			double_slice := make([][]string, 1)
			double_slice[0] = slice
			copied_input = append(copied_input[:i+row_duplicates_required], append([][]string(double_slice), copied_input[i+row_duplicates_required:]...)...)
			row_duplicates_required += 1
		}
	}

	col_duplicates_required := 0
	for j := 0; j < cols; j++ {
		duplicate := true
		for i := 0; i < rows; i++ {
			if input[i][j] != "." {
				duplicate = false
				fmt.Printf("# in col - %d\n", j)
			}
		}
		if duplicate {
			fmt.Printf("adding a col - %d\n", j)
			new_rows := len(copied_input)
			for i := 0; i < new_rows; i++ {
				copied_input[i] = append(copied_input[i], "")
				copy(copied_input[i][j+1+col_duplicates_required:], copied_input[i][j+col_duplicates_required:])
				copied_input[i][j+col_duplicates_required] = "."
			}
			col_duplicates_required += 1
		}
	}

	return copied_input
}

func ShortestPath(grid [][]string, start coords, end coords) int {

	return 0
}

func Puzzle1(input string) int {

	return 0
}
