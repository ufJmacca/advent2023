package main

import (
	"reflect"
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

func ReverseArray(grid [][]string) [][]string {
	size := len(grid)

	for i := 0; i < size/2; i++ {
		grid[i], grid[size-1-i] = grid[size-1-i], grid[i]
	}

	return grid
}

func IsMirror(grid_1 [][]string, grid_2 [][]string) bool {
	size := min(len(grid_1), len(grid_2))

	for i := 0; i < size; i++ {
		if !reflect.DeepEqual(grid_1[i], grid_2[i]) {
			return false
		}
	}

	return true
}

func MirrorDetection(grid [][]string) int {
	size := len(grid)

	for i := 1; i < size; i++ {
		copied := make([][]string, len(grid[:i]))
		copy(copied, grid[:i])
		grid_1 := ReverseArray(copied)
		grid_2 := grid[i:]
		if IsMirror(grid_1, grid_2) {
			return (i + 1) * 100
		}
	}

	grid = Transpose(grid)
	size = len(grid)

	for i := 1; i < size; i++ {
		copied := make([][]string, len(grid[:i]))
		copy(copied, grid[:i])
		grid_1 := ReverseArray(copied)
		grid_2 := grid[i:]
		if IsMirror(grid_1, grid_2) {
			return i
		}
	}

	return 0
}

func Puzzle1(input string) int {

	return 0
}