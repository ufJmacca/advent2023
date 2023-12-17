package main

import "testing"

type puzzle_input_1 struct {
	input  string
	result int
	start  beam
}

type puzzle_input_2 struct {
	input  string
	result int
}

func TestPuzzle1(t *testing.T) {
	test_data := []puzzle_input_1{
		{`.|...\....
|.-.\.....
.....|-...
........|.
..........
.........\
..../.\\..
.-.-/..|..
.|....-|.\
..//.|....`, 46, beam{direction: "east", current_cell: []int{0, 0}}},
	}

	for _, datum := range test_data {
		result := Puzzle1(datum.input, datum.start)

		if result != datum.result {
			t.Errorf("Puzzle1(%s) FAILED - Expected %d Got %d\n", datum.input, datum.result, result)
		} else {
			t.Logf("Puzzle1(%s) PASSED", datum.input)
		}
	}
}

func TestPuzzle2(t *testing.T) {
	test_data := []puzzle_input_2{
		{`.|...\....
|.-.\.....
.....|-...
........|.
..........
.........\
..../.\\..
.-.-/..|..
.|....-|.\
..//.|....`, 54},
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
