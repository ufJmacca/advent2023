package main

import "strings"

func Puzzle1(input string) int {
	lines := strings.Split(input, "\n")

	var grid [][]string

	for _, line := range lines {
		if len(line) > 0 {
			elements := strings.Split(line, "")
			grid = append(grid, elements)
		}
	}

	return 0
}
