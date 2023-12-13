package main

import "testing"

type puzzle_input_1 struct {
	input  string
	result int
}

type mirrow_detection struct {
	input  [][]string
	result int
}

func TestPuzzle1(t *testing.T) {
	test_data := []puzzle_input_1{
		{`#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.

#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#`, 405},
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

func TestMirrorDetection(t *testing.T) {
	test_data := []mirrow_detection{
		{[][]string{
			{`#`, `.`, `#`, `#`, `.`, `.`, `#`, `#`, `.`},
			{`.`, `.`, `#`, `.`, `#`, `#`, `.`, `#`, `.`},
			{`#`, `#`, `.`, `.`, `.`, `.`, `.`, `.`, `#`},
			{`#`, `#`, `.`, `.`, `.`, `.`, `.`, `.`, `#`},
			{`.`, `.`, `#`, `.`, `#`, `#`, `.`, `#`, `.`},
			{`.`, `.`, `#`, `#`, `.`, `.`, `#`, `#`, `.`},
			{`#`, `.`, `#`, `.`, `#`, `#`, `.`, `#`, `.`},
		}, 5},
		{[][]string{
			{`#`, `.`, `.`, `.`, `.`, `#`, `.`, `.`, `#`},
			{`.`, `.`, `#`, `#`, `.`, `.`, `#`, `#`, `#`},
			{`#`, `#`, `#`, `#`, `#`, `.`, `#`, `#`, `.`},
			{`#`, `#`, `#`, `#`, `#`, `.`, `#`, `#`, `.`},
			{`.`, `.`, `#`, `#`, `.`, `.`, `#`, `#`, `#`},
			{`#`, `.`, `.`, `.`, `.`, `#`, `.`, `.`, `#`},
		}, 400},
	}

	for _, datum := range test_data {
		result := MirrorDetection(datum.input)

		if result != datum.result {
			t.Errorf("MirrorDetection(%s) FAILED - Expected %d Got %d\n", datum.input, datum.result, result)
		} else {
			t.Logf("MirrorDetection(%s) PASSED", datum.input)
		}
	}
}
