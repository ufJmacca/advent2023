package main

import "testing"

type puzzle_input_1 struct {
	input  string
	result int
}

type ways_to_win struct {
	time     int
	distance int
	result   int
}

func TestWaysToWin(t *testing.T) {
	test_data := []ways_to_win{
		{7, 9, 4},
		{15, 40, 8},
		{30, 200, 9},
	}

	for _, datum := range test_data {
		result := WaysToWin(datum.time, datum.distance)

		if result != datum.result {
			t.Errorf("WaysToWin(%d, %d) FAILED - Expected %d Got %d\n", datum.time, datum.distance, datum.result, result)
		} else {
			t.Logf("WaysToWin(%d, %d) PASSED", datum.time, datum.distance)
		}
	}
}

func TestPuzzle1(t *testing.T) {
	test_data := []puzzle_input_1{
		{`Time:      7  15   30
Distance:  9  40  200`, 288},
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
	test_data := []puzzle_input_1{
		{`Time:      7  15   30
Distance:  9  40  200`, 71503},
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
