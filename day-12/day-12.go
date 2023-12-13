package main

import (
	"strings"
)

func ArrangementCounter(records string, conditions []int) int {
	if records == "" {
		if len(conditions) == 0 {
			return 1
		} else {
			return 0
		}
	}

	if len(conditions) == 0 {
		if strings.Contains(records, `#`) {
			return 0
		} else {
			return 1
		}
	}

	result := 0

	if strings.Contains(".?", string(records[0])) {
		result += ArrangementCounter(records[1:], conditions)
	}

	if strings.Contains("#?", string(records[0])) {
		if conditions[0] <= len(records) && !strings.Contains(records[:conditions[0]], ".") && (conditions[0] == len(records) || string(records[conditions[0]]) != "#") {
			if conditions[0] == len(records) {
				result += ArrangementCounter("", conditions[1:])
			} else {
				result += ArrangementCounter(records[conditions[0]+1:], conditions[1:])
			}

		}
	}

	return result
}

func Puzzle1(input string) int {

	return 0
}
