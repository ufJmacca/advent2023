package main

import "testing"

type puzzle_input_1 struct {
	input  string
	result int
}

func TestPuzzle1(t *testing.T) {
	test_data := []puzzle_input_1{
		{`broadcaster -> a, b, c
%a -> b
%b -> c
%c -> inv
&inv -> a
`, 32000000}, {`broadcaster -> a
%a -> inv, con
&inv -> b
%b -> con
&con -> output
`, 11687500},
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
