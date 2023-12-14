package main

import (
	"reflect"
	"testing"
)

type puzzle_input_1 struct {
	input  string
	result int
}

type sequence struct {
	input  []int
	result int
}

type reverse_slice struct {
	input  []int
	result []int
}

func TestPuzzle1(t *testing.T) {
	test_data := []puzzle_input_1{
		{`0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`, 114},
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

func TestNextInSequence(t *testing.T) {
	test_data := []sequence{
		{[]int{0, 3, 6, 9, 12, 15}, 18},
		{[]int{1, 3, 6, 10, 15, 21}, 28},
		{[]int{10, 13, 16, 21, 30, 45}, 68},
	}

	for _, datum := range test_data {
		result := NextInSequence(datum.input)

		if result != datum.result {
			t.Errorf("NextInSequence(%v) FAILED - Expected %d Got %d\n", datum.input, datum.result, result)
		} else {
			t.Logf("NextInSequence(%v) PASSED", datum.input)
		}
	}
}

func TestPuzzle2(t *testing.T) {
	test_data := []puzzle_input_1{
		{`0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`, 2},
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

func TestReverseIntSlice(t *testing.T) {
	test_data := []reverse_slice{
		{[]int{0, 3, 6, 9, 12, 15}, []int{15, 12, 9, 6, 3, 0}},
		{[]int{1, 3, 6, 10, 15, 21}, []int{21, 15, 10, 6, 3, 1}},
		{[]int{10, 13, 16, 21, 30, 45}, []int{45, 30, 21, 16, 13, 10}},
	}

	for _, datum := range test_data {
		result := ReverseIntSlice(datum.input)

		if !reflect.DeepEqual(datum.input, datum.result) {
			t.Errorf("NextInSequence(%v) FAILED - Expected %d Got %d\n", datum.input, datum.result, result)
		} else {
			t.Logf("NextInSequence(%v) PASSED", datum.input)
		}
	}
}

// func TestFirstInSequence(t *testing.T) {
// 	test_data := []sequence{
// 		{[]int{0, 3, 6, 9, 12, 15}, -3},
// 		{[]int{1, 3, 6, 10, 15, 21}, 0},
// 		{[]int{10, 13, 16, 21, 30, 45}, 5},
// 	}

// 	for _, datum := range test_data {
// 		result := NextInSequence(datum.input)

// 		if result != datum.result {
// 			t.Errorf("NextInSequence(%v) FAILED - Expected %d Got %d\n", datum.input, datum.result, result)
// 		} else {
// 			t.Logf("NextInSequence(%v) PASSED", datum.input)
// 		}
// 	}
// }
