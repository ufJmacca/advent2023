package main

import "testing"

type puzzle_input_1 struct {
	input  string
	result int
}

type decode struct {
	hexa      string
	direction string
	steps     int
}

func TestPuzzle1(t *testing.T) {
	test_data := []puzzle_input_1{
		{`R 6 (#70c710)
D 5 (#0dc571)
L 2 (#5713f0)
D 2 (#d2c081)
R 2 (#59c680)
D 2 (#411b91)
L 5 (#8ceee2)
U 2 (#caa173)
L 1 (#1b58a2)
U 2 (#caa171)
R 2 (#7807d2)
U 3 (#a77fa3)
L 2 (#015232)
U 2 (#7a21e3)`, 62},
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

func TestInstructionDecode(t *testing.T) {
	test_data := []decode{
		{`#70c710`, `R`, 461937},
		{`#0dc571`, `D`, 56407},
		{`#5713f0`, `R`, 356671},
		{`#d2c081`, `D`, 863240},
		{`#59c680`, `R`, 367720},
		{`#411b91`, `D`, 266681},
		{`#8ceee2`, `L`, 577262},
		{`#caa173`, `U`, 829975},
		{`#1b58a2`, `L`, 112010},
		{`#caa171`, `D`, 829975},
		{`#7807d2`, `L`, 491645},
		{`#a77fa3`, `U`, 686074},
		{`#015232`, `L`, 5411},
		{`#7a21e3`, `U`, 500254},
	}

	for _, datum := range test_data {
		direction, steps := InstructionDecode(datum.hexa)

		if direction != datum.direction {
			t.Errorf("InstructionDecode(%s) FAILED - Expected %s Got %s\n", datum.hexa, datum.direction, direction)
		} else if steps != datum.steps {
			t.Errorf("InstructionDecode(%s) FAILED - Expected %d Got %d\n", datum.hexa, datum.steps, steps)
		} else {
			t.Logf("InstructionDecode(%s) PASSED", datum.hexa)
		}
	}
}

func TestPuzzle2(t *testing.T) {
	test_data := []puzzle_input_1{
		{`R 6 (#70c710)
D 5 (#0dc571)
L 2 (#5713f0)
D 2 (#d2c081)
R 2 (#59c680)
D 2 (#411b91)
L 5 (#8ceee2)
U 2 (#caa173)
L 1 (#1b58a2)
U 2 (#caa171)
R 2 (#7807d2)
U 3 (#a77fa3)
L 2 (#015232)
U 2 (#7a21e3)`, 952408144115},
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
