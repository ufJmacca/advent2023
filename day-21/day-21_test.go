package main

import "testing"

type puzzle_input_1 struct {
	input  string
	steps  int
	result int
}

type function_input struct {
	grid   [][]string
	start  [2]int
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

func TestFunction(t *testing.T) {
	test_data := []function_input{
		{[][]string{
			{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
			{".", ".", ".", ".", ".", "#", "#", "#", ".", "#", "."},
			{".", "#", "#", "#", ".", "#", "#", ".", ".", "#", "."},
			{".", ".", "#", ".", "#", ".", ".", ".", "#", ".", "."},
			{".", ".", ".", ".", "#", ".", "#", ".", ".", ".", "."},
			{".", "#", "#", ".", ".", "S", "#", "#", "#", "#", "."},
			{".", "#", "#", ".", ".", "#", ".", ".", ".", "#", "."},
			{".", ".", ".", ".", ".", ".", ".", "#", "#", ".", "."},
			{".", "#", "#", ".", "#", ".", "#", "#", "#", "#", "."},
			{".", "#", "#", ".", ".", "#", "#", ".", "#", "#", "."},
			{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		}, [2]int{5, 5}, 6, 16},
		{[][]string{
			{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
			{".", ".", ".", ".", ".", "#", "#", "#", ".", "#", "."},
			{".", "#", "#", "#", ".", "#", "#", ".", ".", "#", "."},
			{".", ".", "#", ".", "#", ".", ".", ".", "#", ".", "."},
			{".", ".", ".", ".", "#", ".", "#", ".", ".", ".", "."},
			{".", "#", "#", ".", ".", "S", "#", "#", "#", "#", "."},
			{".", "#", "#", ".", ".", "#", ".", ".", ".", "#", "."},
			{".", ".", ".", ".", ".", ".", ".", "#", "#", ".", "."},
			{".", "#", "#", ".", "#", ".", "#", "#", "#", "#", "."},
			{".", "#", "#", ".", ".", "#", "#", ".", "#", "#", "."},
			{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		}, [2]int{5, 5}, 2, 4},
		{[][]string{
			{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
			{".", ".", ".", ".", ".", "#", "#", "#", ".", "#", "."},
			{".", "#", "#", "#", ".", "#", "#", ".", ".", "#", "."},
			{".", ".", "#", ".", "#", ".", ".", ".", "#", ".", "."},
			{".", ".", ".", ".", "#", ".", "#", ".", ".", ".", "."},
			{".", "#", "#", ".", ".", "S", "#", "#", "#", "#", "."},
			{".", "#", "#", ".", ".", "#", ".", ".", ".", "#", "."},
			{".", ".", ".", ".", ".", ".", ".", "#", "#", ".", "."},
			{".", "#", "#", ".", "#", ".", "#", "#", "#", "#", "."},
			{".", "#", "#", ".", ".", "#", "#", ".", "#", "#", "."},
			{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		}, [2]int{5, 5}, 50, 1594},
		{[][]string{
			{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
			{".", ".", ".", ".", ".", "#", "#", "#", ".", "#", "."},
			{".", "#", "#", "#", ".", "#", "#", ".", ".", "#", "."},
			{".", ".", "#", ".", "#", ".", ".", ".", "#", ".", "."},
			{".", ".", ".", ".", "#", ".", "#", ".", ".", ".", "."},
			{".", "#", "#", ".", ".", "S", "#", "#", "#", "#", "."},
			{".", "#", "#", ".", ".", "#", ".", ".", ".", "#", "."},
			{".", ".", ".", ".", ".", ".", ".", "#", "#", ".", "."},
			{".", "#", "#", ".", "#", ".", "#", "#", "#", "#", "."},
			{".", "#", "#", ".", ".", "#", "#", ".", "#", "#", "."},
			{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		}, [2]int{5, 5}, 100, 6536},
		{[][]string{
			{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
			{".", ".", ".", ".", ".", "#", "#", "#", ".", "#", "."},
			{".", "#", "#", "#", ".", "#", "#", ".", ".", "#", "."},
			{".", ".", "#", ".", "#", ".", ".", ".", "#", ".", "."},
			{".", ".", ".", ".", "#", ".", "#", ".", ".", ".", "."},
			{".", "#", "#", ".", ".", "S", "#", "#", "#", "#", "."},
			{".", "#", "#", ".", ".", "#", ".", ".", ".", "#", "."},
			{".", ".", ".", ".", ".", ".", ".", "#", "#", ".", "."},
			{".", "#", "#", ".", "#", ".", "#", "#", "#", "#", "."},
			{".", "#", "#", ".", ".", "#", "#", ".", "#", "#", "."},
			{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		}, [2]int{5, 5}, 500, 167004},
		{[][]string{
			{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
			{".", ".", ".", ".", ".", "#", "#", "#", ".", "#", "."},
			{".", "#", "#", "#", ".", "#", "#", ".", ".", "#", "."},
			{".", ".", "#", ".", "#", ".", ".", ".", "#", ".", "."},
			{".", ".", ".", ".", "#", ".", "#", ".", ".", ".", "."},
			{".", "#", "#", ".", ".", "S", "#", "#", "#", "#", "."},
			{".", "#", "#", ".", ".", "#", ".", ".", ".", "#", "."},
			{".", ".", ".", ".", ".", ".", ".", "#", "#", ".", "."},
			{".", "#", "#", ".", "#", ".", "#", "#", "#", "#", "."},
			{".", "#", "#", ".", ".", "#", "#", ".", "#", "#", "."},
			{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		}, [2]int{5, 5}, 1000, 668697},
		{[][]string{
			{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
			{".", ".", ".", ".", ".", "#", "#", "#", ".", "#", "."},
			{".", "#", "#", "#", ".", "#", "#", ".", ".", "#", "."},
			{".", ".", "#", ".", "#", ".", ".", ".", "#", ".", "."},
			{".", ".", ".", ".", "#", ".", "#", ".", ".", ".", "."},
			{".", "#", "#", ".", ".", "S", "#", "#", "#", "#", "."},
			{".", "#", "#", ".", ".", "#", ".", ".", ".", "#", "."},
			{".", ".", ".", ".", ".", ".", ".", "#", "#", ".", "."},
			{".", "#", "#", ".", "#", ".", "#", "#", "#", "#", "."},
			{".", "#", "#", ".", ".", "#", "#", ".", "#", "#", "."},
			{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		}, [2]int{5, 5}, 5000, 16733044},
	}

	for _, datum := range test_data {
		result := Function(datum.grid, datum.start, datum.steps)

		if result != datum.result {
			t.Errorf("Puzzle1(%s. %d, %d) FAILED - Expected %d Got %d\n", datum.grid, datum.start, datum.steps, datum.result, result)
		} else {
			t.Logf("Puzzle1(%s, %d, %d) PASSED", datum.grid, datum.start, datum.steps)
		}
	}
}
