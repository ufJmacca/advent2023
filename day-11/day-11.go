package main

import (
	"container/list"
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

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

func GridExpansion(input [][]string, empty_rows int) [][]string {
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
				break
			}
		}
		if duplicate {
			slice := []string(input[i])
			double_slice := make([][]string, 1)
			double_slice[0] = slice
			for k := 0; k < empty_rows; k++ {
				copied_input = append(copied_input[:i+row_duplicates_required+k], append([][]string(double_slice), copied_input[i+row_duplicates_required+k:]...)...)
			}
			row_duplicates_required += empty_rows
		}
	}

	col_duplicates_required := 0
	for j := 0; j < cols; j++ {
		duplicate := true
		for i := 0; i < rows; i++ {
			if input[i][j] != "." {
				duplicate = false
				break
			}
		}
		if duplicate {
			new_rows := len(copied_input)
			for i := 0; i < new_rows; i++ {
				// fmt.Println("col iteration start")
				// fmt.Println(len(copied_input[i]))
				for k := 0; k < empty_rows; k++ {
					copied_input[i] = append(copied_input[i], "")
					copy(copied_input[i][j+1+col_duplicates_required+k:], copied_input[i][j+col_duplicates_required+k:])
					copied_input[i][j+col_duplicates_required+k] = "."
				}
			}
			col_duplicates_required += empty_rows
		}
		// fmt.Println("col duplication ended")
	}
	// fmt.Println(len(copied_input))
	// fmt.Println(len(copied_input[0]))
	// for _, row := range copied_input {
	// 	fmt.Println(row)
	// }
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

func Puzzle1(input string, empty_rows int) int {
	lines := strings.Split(input, "\n")

	var grid [][]string

	for _, line := range lines {
		if len(line) > 0 {
			elements := strings.Split(line, "")
			grid = append(grid, elements)
		}
	}

	expanded_grid := GridExpansion(grid, empty_rows)

	rows, cols := len(expanded_grid), len(expanded_grid[0])
	galaxy_distance := make(map[[2]int]int)
	var galaxies []coords

	for x := 0; x < rows; x++ {
		for y := 0; y < cols; y++ {
			if expanded_grid[x][y] == "#" {
				galaxies = append(galaxies, coords{x, y})
			}
		}
	}

	for i, galaxy_1 := range galaxies {
		for j, galaxy_2 := range galaxies {
			if galaxy_1 == galaxy_2 {
				continue
			} else {
				galaxy_distance[[2]int{i, j}] = ShortestPath(expanded_grid, galaxy_1, galaxy_2)
			}
		}
	}

	sum := 0
	for _, distance := range galaxy_distance {
		sum += distance
	}

	return sum / 2
}

func ManhattanDistance(coord1 coords, coord2 coords) int {
	return int(math.Abs(float64(coord2.x-coord1.x)) + math.Abs(float64(coord2.y-coord1.y)))
}

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

func EmptyRows(grid [][]string) []int {
	var empty_rows []int
	rows := len(grid)
	cols := len(grid[0])
	for i := 0; i < rows; i++ {
		duplicate := true
		for j := 0; j < cols; j++ {
			if grid[i][j] != "." {
				duplicate = false
				break
			}
		}
		if duplicate {
			empty_rows = append(empty_rows, i)
		}
	}
	return empty_rows
}

func IsInSlice(variable int, slice []int) bool {
	for _, value := range slice {
		if value == variable {
			return true
		}
	}
	return false
}

func Puzzle2(input string, multiple int) int {
	lines := strings.Split(input, "\n")

	var grid [][]string

	for _, line := range lines {
		if len(line) > 0 {
			elements := strings.Split(line, "")
			grid = append(grid, elements)
		}
	}

	empty_rows := EmptyRows(grid)
	empty_cols := EmptyRows(Transpose(grid))

	rows, cols := len(grid), len(grid[0])
	var galaxies []coords

	for x := 0; x < rows; x++ {
		for y := 0; y < cols; y++ {
			if grid[x][y] == "#" {
				galaxies = append(galaxies, coords{x, y})
			}
		}
	}

	sum := 0
	for i, galaxy_1 := range galaxies {
		for _, galaxy_2 := range galaxies[:i] {
			for r := min(galaxy_1.x, galaxy_2.x); r < max(galaxy_1.x, galaxy_2.x); r++ {
				if IsInSlice(r, empty_rows) {
					sum += multiple
				} else {
					sum += 1
				}
			}
			for c := min(galaxy_1.y, galaxy_2.y); c < max(galaxy_2.y, galaxy_1.y); c++ {
				if IsInSlice(c, empty_cols) {
					sum += multiple
				} else {
					sum += 1
				}
			}
		}
	}

	return sum
}

func main() {
	c := colly.NewCollector()

	// Sets cookie from environment variable
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("cookie", os.Getenv("COOKIE"))
	})

	c.OnResponse(func(r *colly.Response) {
		inputs := string(r.Body)

		// puzzle_1 := Puzzle1(inputs, 1)
		// fmt.Println(puzzle_1)

		puzzle_2 := Puzzle2(inputs, 1000000)
		fmt.Println(puzzle_2)
	})

	c.Visit("https://adventofcode.com/2023/day/11/input")
}
