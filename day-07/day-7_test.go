package main

import "testing"

type puzzle_input_1 struct {
	input  string
	result int
}

type hand_type struct {
	input  string
	result int
}

type hand_ranks struct {
	input  []string
	result []string
}

func TestHandType(t *testing.T) {
	test_data := []hand_type{
		{"AAAAA", 6},
		{"AA8AA", 5},
		{"23332", 4},
		{"TTT98", 3},
		{"23432", 2},
		{"A23A4", 1},
		{"23456", 0},
	}

	for _, datum := range test_data {
		result := HandType(datum.input)

		if result != datum.result {
			t.Errorf("HandType(%s) FAILED - Expected %d Got %d\n", datum.input, datum.result, result)
		} else {
			t.Logf("HandType(%s) PASSED", datum.input)
		}
	}
}

func TestPuzzle1(t *testing.T) {
	test_data := []puzzle_input_1{
		{`32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`, 6440},
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
		{`32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`, 5905},
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
