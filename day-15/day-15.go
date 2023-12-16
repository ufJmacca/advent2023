package main

import "strings"

func Hash(input string) int {
	current_value := 0

	for _, char := range input {
		current_value += int(byte(char))
		current_value = (current_value * 17) % 256
	}
	return current_value
}

func Puzzle1(input string) int {
	lines := strings.Split(input, ",")

	var sum int

	for _, line := range lines {
		if len(line) > 0 {
			sum += Hash(line)
		}
	}

	return sum
}
