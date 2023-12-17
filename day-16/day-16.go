package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

type grid_element struct {
	value     string
	energised bool
	visits    int
}

type beam struct {
	direction    string
	current_cell []int
}

var next_cell = map[string][]int{
	"east":  []int{0, 1},
	"west":  []int{0, -1},
	"north": []int{-1, 0},
	"south": []int{1, 0},
}

func NextStep(input_beam beam, grid [][]grid_element, rows int, cols int) []beam {
	var beams []beam
	grid[input_beam.current_cell[0]][input_beam.current_cell[1]].energised = true
	grid[input_beam.current_cell[0]][input_beam.current_cell[1]].visits++
	if grid[input_beam.current_cell[0]][input_beam.current_cell[1]].visits < 100 {
		next_x, next_y := input_beam.current_cell[0]+next_cell[input_beam.direction][0], input_beam.current_cell[1]+next_cell[input_beam.direction][1]
		// fmt.Printf("Current step (%d, %d) - direction %s\n", input_beam.current_cell[0], input_beam.current_cell[1], input_beam.direction)

		if next_x >= 0 && next_x < rows && next_y >= 0 && next_y < cols {
			switch input_beam.direction {
			case "east":
				next_step := grid[next_x][next_y]
				if next_step.value == "|" {
					// NORTH / SOUTH
					// fmt.Println(next_step.value)
					return []beam{
						{"north", []int{next_x, next_y}},
						{"south", []int{next_x, next_y}},
					}
				} else if next_step.value == "\\" {
					// SOUTH
					// fmt.Println(next_step.value)
					return []beam{
						{"south", []int{next_x, next_y}},
					}
				} else if next_step.value == "/" {
					// NORTH
					// fmt.Println(next_step.value)
					return []beam{
						{"north", []int{next_x, next_y}},
					}
				} else {
					// fmt.Println(next_step.value)
					return []beam{
						{"east", []int{next_x, next_y}},
					}
				}

			case "west":
				next_step := grid[next_x][next_y]
				if next_step.value == "|" {
					// NORTH / SOUTH
					// fmt.Println(next_step.value)
					return []beam{
						{"north", []int{next_x, next_y}},
						{"south", []int{next_x, next_y}},
					}
				} else if next_step.value == "\\" {
					// NORTH
					// fmt.Println(next_step.value)
					return []beam{
						{"north", []int{next_x, next_y}},
					}
				} else if next_step.value == "/" {
					// south
					// fmt.Println(next_step.value)
					return []beam{
						{"south", []int{next_x, next_y}},
					}
				} else {
					// fmt.Println(next_step.value)
					return []beam{
						{"west", []int{next_x, next_y}},
					}
				}
			case "north":
				next_step := grid[next_x][next_y]
				if next_step.value == "-" {
					// west/east
					// fmt.Println(next_step.value)
					return []beam{
						{"west", []int{next_x, next_y}},
						{"east", []int{next_x, next_y}},
					}
				} else if next_step.value == "\\" {
					// west
					// fmt.Println(next_step.value)
					return []beam{
						{"west", []int{next_x, next_y}},
					}
				} else if next_step.value == "/" {
					// east
					// fmt.Println(next_step.value)
					return []beam{
						{"east", []int{next_x, next_y}},
					}
				} else {
					// fmt.Println(next_step.value)
					return []beam{
						{"north", []int{next_x, next_y}},
					}
				}
			case "south":
				next_step := grid[next_x][next_y]
				if next_step.value == "-" {
					// west/east
					// fmt.Println(next_step.value)
					return []beam{
						{"west", []int{next_x, next_y}},
						{"east", []int{next_x, next_y}},
					}
				} else if next_step.value == "\\" {
					// west
					// fmt.Println(next_step.value)
					return []beam{
						{"east", []int{next_x, next_y}},
					}
				} else if next_step.value == "/" {
					// east
					// fmt.Println(next_step.value)
					return []beam{
						{"west", []int{next_x, next_y}},
					}
				} else {
					// fmt.Println(next_step.value)
					return []beam{
						{"south", []int{next_x, next_y}},
					}
				}
			}
		}

		return beams
	} else {
		return beams
	}
}

func Puzzle1(input string) int {
	lines := strings.Split(input, "\n")

	var grid [][]grid_element

	for _, line := range lines {
		if len(line) > 0 {
			elements := strings.Split(line, "")

			var row []grid_element

			for _, element := range elements {
				row = append(row, grid_element{value: element, energised: false, visits: 0})
			}
			grid = append(grid, row)
		}
	}

	var beams []beam
	rows := len(grid)
	cols := len(grid[0])

	if strings.Contains("\\|", grid[0][0].value) {
		beams = append(beams, beam{direction: "south", current_cell: []int{0, 0}})
	} else {
		beams = append(beams, beam{direction: "east", current_cell: []int{0, 0}})
	}

	for len(beams) > 0 {
		next_step := NextStep(beams[0], grid, rows, cols)
		beams = append(beams[1:], next_step...)
	}

	count := 0
	for _, row := range grid {
		for _, col := range row {
			if col.energised {
				count++
				// fmt.Printf("#")
			} else {
				// fmt.Printf(".")
			}
		}
		// fmt.Printf("\n")
	}

	return count
}

func main() {
	c := colly.NewCollector()

	// Sets cookie from environment variable
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("cookie", os.Getenv("COOKIE"))
	})

	c.OnResponse(func(r *colly.Response) {
		inputs := string(r.Body)

		puzzle_1 := Puzzle1(inputs)
		fmt.Println(puzzle_1)

		// puzzle_2 := Puzzle2(inputs)
		// fmt.Println(puzzle_2)
	})

	c.Visit("https://adventofcode.com/2023/day/16/input")
}
