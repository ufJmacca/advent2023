package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

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
			fmt.Println(strings.ReplaceAll(line, "\n", ""))
			sum += Hash(strings.ReplaceAll(line, "\n", ""))
		}
	}

	return sum
}

func main() {
	c := colly.NewCollector()

	// Sets cookie from environment variable
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("cookie", os.Getenv("COOKIE"))
	})

	c.OnResponse(func(r *colly.Response) {
		inputs := string(r.Body)

		puzzle_1 := Puzzle1(inputs)
		fmt.Println(puzzle_1)

		// puzzle_2 := Puzzle2(inputs, 1000000000)
		// fmt.Println(puzzle_2)
	})

	c.Visit("https://adventofcode.com/2023/day/15/input")
}
