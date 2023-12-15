package main

import (
	"reflect"
	"testing"
)

type puzzle_input_1 struct {
	input  string
	result int
}

type platform struct {
	input  [][]string
	result [][]string
}

func TestPuzzle1(t *testing.T) {
	test_data := []puzzle_input_1{
		{`O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....`, 136},
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

func TestNorthTilt(t *testing.T) {
	test_data := []platform{
		{[][]string{
			{`O`, `.`, `.`, `.`, `.`, `#`, `.`, `.`, `.`, `.`},
			{`O`, `.`, `O`, `O`, `#`, `.`, `.`, `.`, `.`, `#`},
			{`.`, `.`, `.`, `.`, `.`, `#`, `#`, `.`, `.`, `.`},
			{`O`, `O`, `.`, `#`, `O`, `.`, `.`, `.`, `.`, `O`},
			{`.`, `O`, `.`, `.`, `.`, `.`, `.`, `O`, `#`, `.`},
			{`O`, `.`, `#`, `.`, `.`, `O`, `.`, `#`, `.`, `#`},
			{`.`, `.`, `O`, `.`, `.`, `#`, `O`, `.`, `.`, `O`},
			{`.`, `.`, `.`, `.`, `.`, `.`, `.`, `O`, `.`, `.`},
			{`#`, `.`, `.`, `.`, `.`, `#`, `#`, `#`, `.`, `.`},
			{`#`, `O`, `O`, `.`, `.`, `#`, `.`, `.`, `.`, `.`},
		}, [][]string{
			{`O`, `O`, `O`, `O`, `.`, `#`, `.`, `O`, `.`, `.`},
			{`O`, `O`, `.`, `.`, `#`, `.`, `.`, `.`, `.`, `#`},
			{`O`, `O`, `.`, `.`, `O`, `#`, `#`, `.`, `.`, `O`},
			{`O`, `.`, `.`, `#`, `.`, `O`, `O`, `.`, `.`, `.`},
			{`.`, `.`, `.`, `.`, `.`, `.`, `.`, `.`, `#`, `.`},
			{`.`, `.`, `#`, `.`, `.`, `.`, `.`, `#`, `.`, `#`},
			{`.`, `.`, `O`, `.`, `.`, `#`, `.`, `O`, `.`, `O`},
			{`.`, `.`, `O`, `.`, `.`, `.`, `.`, `.`, `.`, `.`},
			{`#`, `.`, `.`, `.`, `.`, `#`, `#`, `#`, `.`, `.`},
			{`#`, `.`, `.`, `.`, `.`, `#`, `.`, `.`, `.`, `.`},
		},
		},
	}

	for _, datum := range test_data {
		result := NorthTilt(datum.input)

		correct := true

		size := len(result)
		if size > 0 {
			for i := 0; i < size; i++ {
				if !reflect.DeepEqual(result[i], datum.result[i]) {
					correct = false
				}
			}
		} else {
			t.Errorf("NorthTilt(%s) FAILED - Empty result", datum.input)
		}

		if !correct {
			t.Errorf("NorthTilt(%s) FAILED - Expect %s got %s", datum.input, datum.result, result)
		} else if size > 0 {
			t.Logf("NorthTilt(%s) PASSED", datum.input)
		}
	}
}

func TestPuzzle2(t *testing.T) {
	test_data := []puzzle_input_1{
		{`O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....`, 64},
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
