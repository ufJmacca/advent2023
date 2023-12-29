package main

import (
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

type queue_item struct {
	cell            [2]int
	steps_remaining int
}

type queue struct {
	items []queue_item
}

var directions = [][]int{
	{1, 0},  // Down
	{-1, 0}, // Up
	{0, 1},  // Right
	{0, -1}, // Left
}

func (q *queue) Enqueue(item queue_item) {
	q.items = append(q.items, item)
}

func (q *queue) Dequeue() queue_item {
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

func modulo(a, b int) int {
	result := a % b
	if result < 0 && b > 0 {
		result += b
	}
	return result
}

func Function(grid [][]string, start [2]int, steps int) int {
	rows := len(grid)
	cols := len(grid[0])

	answer := make(map[[2]int]int)
	seen := make(map[[2]int]int)

	p_queue := queue{}
	p_queue.Enqueue(queue_item{
		cell:            start,
		steps_remaining: steps,
	})

	for len(p_queue.items) > 0 {
		element := p_queue.Dequeue()

		if element.steps_remaining%2 == 0 {
			answer[[2]int(element.cell)] = 1
		}
		if element.steps_remaining == 0 {
			continue
		}

		for _, direction := range directions {
			next_x, next_y := element.cell[0]+direction[0], element.cell[1]+direction[1]
			_, already_seen := seen[[2]int{next_x, next_y}]

			if grid[modulo(next_x, rows)][modulo(next_y, cols)] != "#" && !already_seen {
				seen[[2]int{next_x, next_y}] = 1
				p_queue.Enqueue(queue_item{
					cell:            [2]int{next_x, next_y},
					steps_remaining: element.steps_remaining - 1,
				})
			}
		}

	}

	return len(answer)
}

func Puzzle1(input string, steps int) int {
	var grid [][]string
	var start [2]int
	lines := strings.Split(input, "\n")
	for row, line := range lines {
		if len(line) > 0 {
			elements := strings.Split(line, "")
			grid = append(grid, elements)
			for col, value := range elements {
				if value == "S" {
					start = [2]int{row, col}
				}
			}
		}
	}

	return Function(grid, start, steps)
}

func FunctionQuadratic(grid [][]string, start [2]int, steps int, x int, original_x int) int {
	rows := len(grid)
	cols := len(grid[0])

	answer := make(map[[2]int]int)
	seen := make(map[[2]int]int)

	p_queue := queue{}
	p_queue.Enqueue(queue_item{
		cell:            start,
		steps_remaining: original_x + (2 * rows * x),
	})

	for len(p_queue.items) > 0 {
		element := p_queue.Dequeue()

		if element.steps_remaining%2 == 0 {
			answer[[2]int(element.cell)] = 1
		}
		if element.steps_remaining == 0 {
			continue
		}

		for _, direction := range directions {
			next_x, next_y := element.cell[0]+direction[0], element.cell[1]+direction[1]
			_, already_seen := seen[[2]int{next_x, next_y}]

			if grid[modulo(next_x, rows)][modulo(next_y, cols)] != "#" && !already_seen {
				seen[[2]int{next_x, next_y}] = 1
				p_queue.Enqueue(queue_item{
					cell:            [2]int{next_x, next_y},
					steps_remaining: element.steps_remaining - 1,
				})
			}
		}

	}

	return len(answer)
}

func Puzzle2(input string, steps int) int {
	var grid [][]string
	var start [2]int
	lines := strings.Split(input, "\n")
	for row, line := range lines {
		if len(line) > 0 {
			elements := strings.Split(line, "")
			grid = append(grid, elements)
			for col, value := range elements {
				if value == "S" {
					start = [2]int{row, col}
				}
			}
		}
	}

	size := len(grid)
	original_x := steps % (2 * size)

	var a []int
	var first_differences []int
	var second_differences []int

	x := 0

	for {
		a = append(a, FunctionQuadratic(grid, start, steps, x, original_x))
		x += 1

		if len(a) >= 4 {
			first_differences = []int{
				a[1] - a[0],
				a[2] - a[1],
				a[3] - a[2],
			}
			second_differences = []int{
				first_differences[1] - first_differences[0],
				first_differences[2] - first_differences[1],
			}
			if second_differences[0] == second_differences[1] {
				break
			} else {
				a = a[1:]
			}
		}

	}

	alpha, beta, gamma := a[0], a[1], a[2]

	c := alpha
	a_eq := (gamma - 2*beta + c) / 2
	b := beta - c - a_eq

	x_result := steps/(2*size) - (x - 4)

	return a_eq*int(math.Pow(float64(x_result), float64(2))) + b*x_result + c

}

func main() {
	c := colly.NewCollector()

	// Sets cookie from environment variable
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("cookie", os.Getenv("COOKIE"))
	})

	c.OnResponse(func(r *colly.Response) {
		inputs := string(r.Body)

		puzzle_1 := Puzzle1(inputs, 64)
		fmt.Println(puzzle_1)

		puzzle_2 := Puzzle2(inputs, 26501365)
		fmt.Println(puzzle_2)
	})

	c.Visit("https://adventofcode.com/2023/day/21/input")
}
