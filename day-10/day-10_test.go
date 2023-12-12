package main

import (
	"testing"
)

type puzzle_input_1 struct {
	input  string
	result int
}

type find_start struct {
	input  [][]string
	result coords
}

type find_connected struct {
	grid     [][]string
	location coords
	result   []coords
}

func TestPuzzle1(t *testing.T) {
	test_data := []puzzle_input_1{
		{`.....
.S-7.
.|.|.
.L-J.
.....`, 4},
		{`..F7.
.FJ|.
SJ.L7
|F--J
LJ...`, 8},
	}

	for _, datum := range test_data {
		result := Puzzle1(datum.input)

		if result != datum.result {
			t.Errorf("Puzzle1(%s) FAILED - Expected %d Got %d\n", datum.input, datum.result, result)
		} else {
			t.Logf("Puzzle1(%s) PASSED", datum.input)
		}
	}
}

func TestFindStart(t *testing.T) {
	test_data := []find_start{
		{[][]string{
			{".", ".", ".", ".", "."},
			{".", "S", "-", "7", "."},
			{".", "|", ".", "|", "."},
			{".", "L", "-", "J", "."},
			{".", ".", ".", ".", "."},
		}, coords{x: 1, y: 1}},
		{[][]string{
			{".", ".", "F", "7", "."},
			{".", "F", "J", "|", "."},
			{"S", "J", ".", "L", "7"},
			{"|", "F", "-", "-", "J"},
			{"L", "J", ".", ".", "."},
		}, coords{x: 2, y: 0}},
	}

	for _, datum := range test_data {
		result := FindStart(datum.input)

		if result != datum.result {
			t.Errorf("FindStart(%s) FAILED - Expected %d Got %d\n", datum.input, datum.result, result)
		} else {
			t.Logf("FindStart(%s) PASSED", datum.input)
		}
	}
}

func TestFindConnectedTiles(t *testing.T) {
	test_data := []find_connected{
		{[][]string{
			{".", ".", ".", ".", "."},
			{".", "S", "-", "7", "."},
			{".", "|", ".", "|", "."},
			{".", "L", "-", "J", "."},
			{".", ".", ".", ".", "."},
		}, coords{x: 1, y: 1}, []coords{{x: 2, y: 1}, {x: 1, y: 2}}},
		{[][]string{
			{".", ".", "F", "7", "."},
			{".", "F", "J", "|", "."},
			{"S", "J", ".", "L", "7"},
			{"|", "F", "-", "-", "J"},
			{"L", "J", ".", ".", "."},
		}, coords{x: 2, y: 0}, []coords{{x: 2, y: 1}, {x: 3, y: 0}}},
	}

	for _, datum := range test_data {
		result := FindConnectedTiles(datum.grid, datum.location)

		if !equalStructSlicesIgnoreOrder(result, datum.result) {
			t.Errorf("FindConnectedTiles(%v, %v) FAILED - Expected %d Got %d\n", datum.grid, datum.location, datum.result, result)
		} else {
			t.Logf("FindConnectedTiles(%v, %v) PASSED", datum.grid, datum.location)
		}
	}
}

func TestPuzzle2(t *testing.T) {
	test_data := []puzzle_input_1{
		{`...........
.S-------7.
.|F-----7|.
.||.....||.
.||.....||.
.|L-7.F-J|.
.|..|.|..|.
.L--J.L--J.
...........`, 4},
		{`..........
.S------7.
.|F----7|.
.||OOOO||.
.||OOOO||.
.|L-7F-J|.
.|II||II|.
.L--JL--J.
..........`, 4},
		{`.F----7F7F7F7F-7....
.|F--7||||||||FJ....
.||.FJ||||||||L7....
FJL7L7LJLJ||LJ.L-7..
L--J.L7...LJS7F-7L7.
....F-J..F7FJ|L7L7L7
....L7.F7||L7|.L7L7|
.....|FJLJ|FJ|F7|.LJ
....FJL-7.||.||||...
....L---J.LJ.LJLJ...`, 8},
		{`FF7FSF7F7F7F7F7F---7
L|LJ||||||||||||F--J
FL-7LJLJ||||||LJL-77
F--JF--7||LJLJ7F7FJ-
L---JF-JLJ.||-FJLJJ7
|F|F-JF---7F7-L7L|7|
|FFJF7L7F-JF7|JL---7
7-L-JL7||F7|L7F-7F7|
L.L7LFJ|||||FJL7||LJ
L7JLJL-JLJLJL--JLJ.L`, 10},
	}

	for _, datum := range test_data {
		result := Puzzle2(datum.input)

		if result != datum.result {
			t.Errorf("Puzzle2(%s) FAILED - Expected %d Got %d\n", datum.input, datum.result, result)
		} else {
			t.Logf("Puzzle2(%s) PASSED", datum.input)
		}
	}
}
