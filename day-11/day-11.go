package main

import "container/list"

type coords struct {
	x int
	y int
}

var directions = []coords{
	{1, 0},  // Down
	{-1, 0}, // Up
	{0, 1},  // Right
	{0, -1}, // Left
}

func GridExpansion(input [][]string) [][]string {
	rows := len(input)
	cols := len(input[0])

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
			}
		}
		if duplicate {
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
			}
		}
		if duplicate {
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
	rows, cols := len(grid), len(grid[0])
	visited := make(map[coords]bool)
	queue := list.New()

	queue.PushBack(start)
	visited[start] = true
	distance := 0

	for queue.Len() > 0 {
		size := queue.Len()
		distance += 1

		for i := 0; i < size; i++ {
			current := queue.Remove(queue.Front()).(coords)

			if current == end {
				return distance - 1
			}

			for _, direction := range directions {
				next_x, next_y := current.x+direction.x, current.y+direction.y

				if next_x >= 0 && next_x < rows && next_y >= 0 && next_y < cols {
					next := coords{next_x, next_y}
					if !visited[next] {
						queue.PushBack(next)
						visited[next] = true
					}
				}
			}
		}
	}

	return 0
}

func Puzzle1(input string) int {

	return 0
}
