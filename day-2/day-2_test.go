package main

import (
	"reflect"
	"testing"
)

type part1 struct {
	str    string
	result string
}

type part2 struct {
	str    string
	result map[string]int
}

type part2b struct {
	input  map[string]int
	result int
}

func TestGamePossible(t *testing.T) {
	testData := []part1{
		{"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", "possible"},
		{"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue", "possible"},
		{"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red", "impossible"},
		{"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red", "impossible"},
		{"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green", "possible"},
	}

	for _, datum := range testData {
		_, result := GamePossible(datum.str)

		if result != datum.result {
			t.Errorf("GamePossible(%s) FAILED - Expected %s Got %s\n", datum.str, datum.result, result)
		} else {
			t.Logf("GamePossibel(%s) PASSED", datum.str)
		}
	}
}

func TestMinimumPossibleCubes(t *testing.T) {
	testData := []part2{
		{"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", map[string]int{"red": 4, "green": 2, "blue": 6}},
		{"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue", map[string]int{"red": 1, "green": 3, "blue": 4}},
		{"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red", map[string]int{"red": 20, "green": 13, "blue": 6}},
		{"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red", map[string]int{"red": 14, "green": 3, "blue": 15}},
		{"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green", map[string]int{"red": 6, "green": 3, "blue": 2}},
	}

	for _, datum := range testData {
		result := MinimumPossibleCubes(datum.str)

		if !reflect.DeepEqual(result, datum.result) {
			t.Errorf("GamePossible(%s) FAILED - Expected %#v Got %#v\n", datum.str, datum.result, result)
		} else {
			t.Logf("GamePossibel(%s) PASSED", datum.str)
		}
	}
}

func TestBagPower(t *testing.T) {
	testData := []part2b{
		{map[string]int{"red": 4, "green": 2, "blue": 6}, 48},
		{map[string]int{"red": 1, "green": 3, "blue": 4}, 12},
		{map[string]int{"red": 20, "green": 13, "blue": 6}, 1560},
		{map[string]int{"red": 14, "green": 3, "blue": 15}, 630},
		{map[string]int{"red": 6, "green": 3, "blue": 2}, 36},
	}

	for _, datum := range testData {
		result := BagPower(datum.input)

		if !reflect.DeepEqual(result, datum.result) {
			t.Errorf("GamePossible(%#v) FAILED - Expected %d Got %d\n", datum.input, datum.result, result)
		} else {
			t.Logf("GamePossibel(%#v) PASSED", datum.input)
		}
	}
}
