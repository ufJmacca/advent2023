package main

import "testing"

type puzzle_input_1 struct {
	input  string
	result int
}

type arrangement_counter struct {
	records    string
	conditions []int
	result     int
}

func TestPuzzle1(t *testing.T) {
	test_data := []puzzle_input_1{
		{`???.### 1,1,3
.??..??...?##. 1,1,3
?#?#?#?#?#?#?#? 1,3,1,6
????.#...#... 4,1,1
????.######..#####. 1,6,5
?###???????? 3,2,1`, 21},
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

func TestArrangementCounter(t *testing.T) {
	test_data := []arrangement_counter{
		{`???.###`, []int{1, 1, 3}, 1},
		{`.??..??...?##.`, []int{1, 1, 3}, 4},
		{`?#?#?#?#?#?#?#?`, []int{1, 3, 1, 6}, 1},
		{`????.#...#...`, []int{4, 1, 1}, 1},
		{`????.######..#####.`, []int{1, 6, 5}, 4},
		{`?###????????`, []int{3, 2, 1}, 10},
	}

	for _, datum := range test_data {
		result := ArrangementCounter(datum.records, datum.conditions)

		if result != datum.result {
			t.Errorf("Puzzle1(%s, %s) FAILED - Expected %d Got %d\n", datum.records, datum.conditions, datum.result, result)
		} else {
			t.Logf("Puzzle1(%s, %s) PASSED", datum.records, datum.conditions)
		}
	}
}
