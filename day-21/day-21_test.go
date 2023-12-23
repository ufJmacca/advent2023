package main

import "testing"

type puzzle_input_1 struct {
	input  string
	steps  int
	result int
}

func TestPuzzle1(t *testing.T) {
	test_data := []puzzle_input_1{
		{`...........
.....###.#.
.###.##..#.
..#.#...#..
....#.#....
.##..S####.
.##..#...#.
.......##..
.##.#.####.
.##..##.##.
...........
`, 6, 16},
	}

	for _, datum := range test_data {
		result := Puzzle1(datum.input, datum.steps)

		if result != datum.result {
			t.Errorf("Puzzle1(%s. %d) FAILED - Expected %d Got %d\n", datum.input, datum.steps, datum.result, result)
		} else {
			t.Logf("Puzzle1(%s, %d) PASSED", datum.input, datum.steps)
		}
	}
}
