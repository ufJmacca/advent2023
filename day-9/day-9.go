package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

func NextInSequence(input []int) int {
	if len(input) == 0 {
		return 0
	}
	var next_sequence []int

	last_int := input[len(input)-1]

	for index, value := range input {
		if index == 0 {
			continue
		}

		next_sequence = append(next_sequence, value-input[index-1])
	}

	subSum := NextInSequence(next_sequence)

	return last_int + subSum
}

func Puzzle1(input string) int {
	lines := strings.Split(input, "\n")
	re := regexp.MustCompile(`(\-?\d+)`)

	sum := 0

	for _, line := range lines {
		if len(line) > 0 {
			var sequence []int
			for _, number := range re.FindAllStringSubmatch(line, -1) {
				number_int, _ := strconv.Atoi(number[0])
				sequence = append(sequence, number_int)
			}
			sum += NextInSequence(sequence)
		}
	}

	return sum
}

func ReverseIntSlice(slice []int) []int {
	length := len(slice)
	for i := 0; i < length/2; i++ {
		slice[i], slice[length-i-1] = slice[length-i-1], slice[i]
	}
	return slice
}

func Puzzle2(input string) int {
	lines := strings.Split(input, "\n")
	re := regexp.MustCompile(`(\-?\d+)`)

	sum := 0

	for _, line := range lines {
		if len(line) > 0 {
			var sequence []int
			for _, number := range re.FindAllStringSubmatch(line, -1) {
				number_int, _ := strconv.Atoi(number[0])
				sequence = append(sequence, number_int)
			}
			rev_sequence := ReverseIntSlice(sequence)
			sum += NextInSequence(rev_sequence)
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

		puzzle_2 := Puzzle2(inputs)
		fmt.Println(puzzle_2)
	})

	c.Visit("https://adventofcode.com/2023/day/9/input")
}
