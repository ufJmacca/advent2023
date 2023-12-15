package main

import (
	"sort"
	"strings"
)

func Transpose(grid [][]string) [][]string {
	rows := len(grid)
	cols := len(grid[0])

	transposed := make([][]string, cols)
	for i := range transposed {
		transposed[i] = make([]string, rows)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			transposed[j][i] = grid[i][j]
		}
	}

	return transposed
}

func ReverseArray(grid []string) []string {
	size := len(grid)

	for i := 0; i < size/2; i++ {
		grid[i], grid[size-1-i] = grid[size-1-i], grid[i]
	}

	return grid
}

func SplitArray(array []string, value string) [][]string {
	var result [][]string
	var part []string

	for _, str := range array {
		// part = append(part, str)
		if value == str {
			result = append(result, part)
			part = nil
		} else {
			part = append(part, str)
		}
	}

	if len(part) > 0 {
		result = append(result, part)
	}

	return result
}

func NorthTilt(input [][]string) [][]string {
	transposed_input := Transpose(input)
	rows := len(transposed_input)

	var result [][]string
	var result_row []string

	for i := 0; i < rows; i++ {
		reversed_row := ReverseArray(transposed_input[i])
		split_rows := SplitArray(reversed_row, "#")
		for j, row := range split_rows {
			sort.Strings(row)
			if j == len(split_rows)-1 {
				result_row = append(result_row, row...)
			} else {
				result_row = append(result_row, append(row, []string{"#"}...)...)
			}
		}
		for len(result_row) < len(transposed_input[i]) {
			result_row = append(result_row, "#")
		}
		result = append(result, result_row)
		result_row = nil
	}

	var unreversed_result [][]string

	for _, row := range result {
		unreversed_result = append(unreversed_result, ReverseArray(row))
	}

	transposed_result := Transpose(unreversed_result)

	return transposed_result
}

func Puzzle1(input string) int {
	lines := strings.Split(input, "\n")

	var grid [][]string

	for _, line := range lines {
		if len(line) > 0 {
			elements := strings.Split(line, "")
			grid = append(grid, elements)
		}
	}

	tilted_grid := NorthTilt(grid)
	grid_row := len(tilted_grid)
	total_load := 0

	for i, row := range tilted_grid {
		count := 0
		for _, str := range row {
			if str == `O` {
				count++
			}
		}

		total_load += count * (grid_row - i)
	}

	return total_load
}
