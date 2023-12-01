package main

import "testing"

type part1 struct {
	str    string
	result int
}

func TestCalibration(t *testing.T) {
	testData := []part1{
		{"1abc2", 12},
		{"pqr3stu8vwx", 38},
		{"a1b2c3d4e5f", 15},
		{"treb7uchet", 77},
	}

	for _, datum := range testData {
		result := Calibration(datum.str)

		if result != datum.result {
			t.Errorf("Calibation(%s) FAILED - Expected %d Got %d\n", datum.str, datum.result, result)
		} else {
			t.Logf("Calibration(%s) PASSED", datum.str)
		}
	}
}

func TestCalibrationStrings(t *testing.T) {
	testData := []part1{
		{"two1nine", 29},
		{"eightwothree", 83},
		{"abcone2threexyz", 13},
		{"xtwone3four", 24},
		{"4nineeightseven2", 42},
		{"zoneight234", 14},
		{"7pqrstsixteen", 76},
	}

	for _, datum := range testData {
		result := CalibrationString(datum.str)

		if result != datum.result {
			t.Errorf("Calibation(%s) FAILED - Expected %d Got %d\n", datum.str, datum.result, result)
		} else {
			t.Logf("Calibration(%s) PASSED", datum.str)
		}
	}
}
