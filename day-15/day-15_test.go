package main

import "testing"

type puzzle_input_1 struct {
	input  string
	result int
}

type hash_input struct {
	input  string
	result int
}

func TestPuzzle1(t *testing.T) {
	test_data := []puzzle_input_1{
		{`rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7`, 1320},
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

func TestHash(t *testing.T) {
	test_data := []puzzle_input_1{
		{`rn=1`, 30},
		{`cm-`, 253},
		{`qp=3`, 97},
		{`cm=2`, 47},
		{`qp-`, 14},
		{`pc=4`, 180},
		{`ot=9`, 9},
		{`ab=5`, 197},
		{`pc-`, 48},
		{`pc=6`, 214},
		{`ot=7`, 231},
	}

	for _, datum := range test_data {
		result := Hash(datum.input)

		if result != datum.result {
			t.Errorf("Hash(%s) FAILED - Expected %d Got %d\n", datum.input, datum.result, result)
		} else {
			t.Logf("Hash(%s) PASSED", datum.input)
		}
	}
}
