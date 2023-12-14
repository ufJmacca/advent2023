package main

import "testing"

type testSymbolInput struct {
	input  string
	result bool
}

type testInput1 struct {
	input  string
	result int
}

func TestIsSymbol(t *testing.T) {
	testData := []testSymbolInput{
		{"6", false},
		{"7", false},
		{"4", false},
		{".", false},
		{"1", false},
		{"*", true},
		{"#", true},
		{"+", true},
		{"$", true},
	}

	for _, datum := range testData {
		result := isSymbol(rune(datum.input[0]))

		if result != datum.result {
			t.Errorf("EngineSchematic(%s) FAILED - Expected %t Got %t\n", datum.input, datum.result, result)
		} else {
			t.Logf("EngineSchematic(%s) PASSED", datum.input)
		}
	}
}

func TestEngineSchematic(t *testing.T) {
	testData := []testInput1{
		{`467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`, 4361},
	}

	for _, datum := range testData {
		result := EngineSchematic(datum.input)

		if result != datum.result {
			t.Errorf("EngineSchematic(%s) FAILED - Expected %d Got %d\n", datum.input, datum.result, result)
		} else {
			t.Logf("EngineSchematic(%s) PASSED", datum.input)
		}
	}
}

func TestGearSchematic(t *testing.T) {
	testData := []testInput1{
		{`467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`, 467835},
	}

	for _, datum := range testData {
		result := GearSchematic(datum.input)

		if result != datum.result {
			t.Errorf("GearSchematic(%s) FAILED - Expected %d Got %d\n", datum.input, datum.result, result)
		} else {
			t.Logf("GearSchematic(%s) PASSED", datum.input)
		}
	}
}
