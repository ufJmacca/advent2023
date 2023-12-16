package main

import (
	"fmt"
	"os"
	"reflect"
	"sort"
	"strings"

	"github.com/gocolly/colly"
	orderedmap "github.com/wk8/go-ordered-map"
)

func GridEquality(grid_1 [][]string, grid_2 [][]string) bool {
	size := min(len(grid_1), len(grid_2))

	for i := 0; i < size; i++ {
		if !reflect.DeepEqual(grid_1[i], grid_2[i]) {
			return false
		}
	}

	return true
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

func StringifyGrid(key [][]string) string {
	return fmt.Sprintf("%s", key)
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

func SouthTilt(input [][]string) [][]string {
	transposed_input := Transpose(input)
	rows := len(transposed_input)

	var result [][]string
	var result_row []string

	for i := 0; i < rows; i++ {
		reversed_row := transposed_input[i]
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

	transposed_result := Transpose(result)

	return transposed_result
}

func WestTilt(input [][]string) [][]string {
	transposed_input := input
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

	return unreversed_result
}

func EastTilt(input [][]string) [][]string {
	transposed_input := input
	rows := len(transposed_input)

	var result [][]string
	var result_row []string

	for i := 0; i < rows; i++ {
		reversed_row := transposed_input[i]
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

	return result
}

func TiltCycle(input [][]string) [][]string {
	input = NorthTilt(input)
	input = WestTilt(input)
	input = SouthTilt(input)
	input = EastTilt(input)

	return input
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

func Puzzle2(input string, cycles int) int {
	lines := strings.Split(input, "\n")

	var grid [][]string

	for _, line := range lines {
		if len(line) > 0 {
			elements := strings.Split(line, "")
			grid = append(grid, elements)
		}
	}

	tilted_grid := grid
	grid_row := len(tilted_grid)
	ordered_map := orderedmap.New()
	var loop_start string

	fmt.Println("starting tilt cycle")

	for i := 0; i < cycles; i++ {
		tilted_grid = TiltCycle(tilted_grid)
		key := StringifyGrid(tilted_grid)
		fmt.Println(i)
		if _, ok := ordered_map.Get(key); ok {
			fmt.Println("next rotation")
			loop_start = key
			fmt.Println(loop_start)
			break
		}
		ordered_map.Set(key, true)
	}

	fmt.Println(ordered_map)

	loop_start_i := 0
	for pair := ordered_map.Oldest(); pair != nil; pair = pair.Next() {
		loop_start_i++
		if pair.Key == loop_start {
			break
		}
	}
	fmt.Println(loop_start_i)

	loop_end_i := 0
	for pair := ordered_map.Oldest(); pair != nil; pair = pair.Next() {
		loop_end_i++
	}
	fmt.Println(loop_end_i)

	cycles_left := cycles - (loop_start_i)
	fmt.Println(cycles_left)
	modulo := cycles_left % ((loop_end_i - loop_start_i) + 1)
	fmt.Println(modulo)

	i := 0
	var final_grid [][]string
	for pair := ordered_map.Oldest(); pair != nil; pair = pair.Next() {
		i++
		if i == modulo+loop_start_i {
			final := pair.Key.(string)
			fmt.Println("set titled_grid")
			// tilted_grid = final
			fmt.Println(final)
			rows := strings.Split(strings.Trim(final[2:len(final)-2], " "), "] [")

			for _, row := range rows {
				// Split each row into elements
				elements := strings.Split(row, " ")
				final_grid = append(final_grid, elements)
			}
		}
	}

	fmt.Println()
	fmt.Println(final_grid)

	total_load := 0

	for i, row := range final_grid {
		count := 0
		for _, str := range row {
			if str == `O` {
				count++
			}
		}

		total_load += count * (grid_row - i)
		fmt.Println(final_grid[i])
	}
	fmt.Println()

	return total_load
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

		// puzzle_2 := Puzzle2(inputs, 1000000000)
		// fmt.Println(puzzle_2)
	})

	c.Visit("https://adventofcode.com/2023/day/14/input")
}
