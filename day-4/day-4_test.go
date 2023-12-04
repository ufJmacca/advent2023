package main

import "testing"

type gamepoints struct {
	input  int
	result int
}
type scratchcards struct {
	input  string
	result int
}

type scratchcardsstart struct {
	input  string
	result int
}

func TestGamePoints(t *testing.T) {
	testData := []gamepoints{
		{4, 8},
		{2, 2},
		{1, 1},
		{0, 0},
	}

	for _, datum := range testData {
		result := GamePoints(datum.input)

		if result != datum.result {
			t.Errorf("GamePoints(%d) FAILED - Expected %d Got %d\n", datum.input, datum.result, result)
		} else {
			t.Logf("GamePoints(%d) PASSED", datum.input)
		}
	}
}

func TestGameResult(t *testing.T) {
	testData := []scratchcards{
		{"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53", 8},
		{"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19", 2},
		{"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1", 2},
		{"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83", 1},
		{"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36", 0},
		{"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11", 0},
	}

	for _, datum := range testData {
		result, _ := GameResult(datum.input)

		if result != datum.result {
			t.Errorf("GameResult(%s) FAILED - Expected %d Got %d\n", datum.input, datum.result, result)
		} else {
			t.Logf("GameResult(%s) PASSED", datum.input)
		}
	}
}

func TestGameResultWithDuplicates(t *testing.T) {
	testData := []scratchcardsstart{
		{`Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`, 30},
	}

	for _, datum := range testData {
		result := GameResultWithDuplicates(datum.input)

		if result != datum.result {
			t.Errorf("GameResultWithDuplicates(%s) FAILED - Expected %d Got %d\n", datum.input, datum.result, result)
		} else {
			t.Logf("GameResultWithDuplicates(%s) PASSED", datum.input)
		}
	}
}
