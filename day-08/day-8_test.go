package main

import "testing"

type puzzle_input_1 struct {
	input  string
	result int
}

type all_finished_input struct {
	input  []map[string]path_nodes
	result bool
}

type lcm_input struct {
	input  []int
	result int
}

func TestPathMap(t *testing.T) {
	test_data := []puzzle_input_1{
		{`RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)`, 2},
		{`LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)`, 6},
	}

	for _, datum := range test_data {
		result := PathMap(datum.input)

		if result != datum.result {
			t.Errorf("PathMap(%s) FAILED - Expected %d Got %d\n", datum.input, datum.result, result)
		} else {
			t.Logf("PathMap(%s) PASSED", datum.input)
		}
	}
}

func TestGhostPathMap(t *testing.T) {
	test_data := []puzzle_input_1{
		{`LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)`, 6},
	}

	for _, datum := range test_data {
		result := GhostPathMap(datum.input)

		if result != datum.result {
			t.Errorf("GhostPathMap(%s) FAILED - Expected %d Got %d\n", datum.input, datum.result, result)
		} else {
			t.Logf("GhostPathMap(%s) PASSED", datum.input)
		}
	}
}

func TestAllFinished(t *testing.T) {
	test_data := []all_finished_input{
		{
			input: []map[string]path_nodes{
				{
					"11B": path_nodes{R: "XXX", L: "11Z"},
					"11Z": path_nodes{R: "XXX", L: "11Z"},
				},
			},
			result: false,
		},
		{
			input: []map[string]path_nodes{
				{
					"12Z": path_nodes{R: "XXX", L: "11Z"},
					"11Z": path_nodes{R: "XXX", L: "11Z"},
				},
			},
			result: true,
		},
	}

	for _, datum := range test_data {
		result := AllFinished(datum.input)

		if result != datum.result {
			t.Errorf("AllFinished(%s) FAILED - Expected %t Got %t\n", datum.input, datum.result, result)
		} else {
			t.Logf("AllFinished(%s) PASSED", datum.input)
		}
	}
}

func TestLCM(t *testing.T) {
	test_data := []lcm_input{
		{
			input:  []int{10, 15},
			result: 30,
		},
		{
			input:  []int{10, 15, 20},
			result: 60,
		},
		{
			input:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			result: 2520,
		},
	}

	for _, datum := range test_data {
		result := LCM(datum.input)

		if result != datum.result {
			t.Errorf("LCM(%v) FAILED - Expected %d Got %d\n", datum.input, datum.result, result)
		} else {
			t.Logf("LCM(%v) PASSED", datum.input)
		}
	}
}
