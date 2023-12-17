package main

import "testing"

type puzzle_input struct {
	input  string
	result int
}

func TestPuzzle1(t *testing.T) {
	test_data := []puzzle_input{
		{`.|...\....
|.-.\.....
.....|-...
........|.
..........
.........\
..../.\\..
.-.-/..|..
.|....-|.\
..//.|....`, 46},
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

func TestPuzzle2(t *testing.T) {
	test_data := []puzzle_input{
		{`.|...\....
|.-.\.....
.....|-...
........|.
..........
.........\
..../.\\..
.-.-/..|..
.|....-|.\
..//.|....`, 46},
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
