package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

type map_struct struct {
	id           string
	focal_length int
}

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
			sum += Hash(strings.ReplaceAll(line, "\n", ""))
		}
	}

	return sum
}

func Puzzle2(input string) int {
	lines := strings.Split(input, ",")

	var sum int
	hashmap := make(map[int][]map_struct)

	pattern := `(\w+)(\-|\=)(\d?)`
	regex := regexp.MustCompile(pattern)

	for _, line := range lines {
		if len(line) > 0 {
			matches := regex.FindAllStringSubmatch(strings.ReplaceAll(line, "\n", ""), -1)

			box := Hash(matches[0][1])
			id := matches[0][1]
			action := matches[0][2]
			if action == "-" {
				for i, queue := range hashmap[box] {
					if queue.id == id {
						queue_before := hashmap[box][:i]
						queue_after := hashmap[box][i+1:]
						hashmap[box] = append(queue_before, queue_after...)
					}
				}
			} else {
				var exists bool
				focal_length, _ := strconv.Atoi(matches[0][3])
				for i, queue := range hashmap[box] {
					if queue.id == id {
						exists = true
						queue.focal_length = focal_length

						hashmap[box][i].focal_length = focal_length
					}
				}
				if !exists {
					hashmap[box] = append(hashmap[box], []map_struct{{id, focal_length}}...)
				}
			}
		}
	}

	for i, val := range hashmap {
		for j, slice := range val {
			sum += ((i + 1) * (j + 1) * slice.focal_length)
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

	c.Visit("https://adventofcode.com/2023/day/15/input")
}
