package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

func WaysToWin(time int, distance int) int {
	possibility_of_record := 0

	for i := 0; i < time; i++ {
		speed := i
		distance_travelable := speed * (time - i)
		if distance_travelable > distance {
			possibility_of_record += 1
		}
	}
	return possibility_of_record
}

func Puzzle1(input string) int {
	lines := strings.Split(input, "\n")

	re := regexp.MustCompile(`\d+`)

	times := re.FindAllString(lines[0], -1)
	distances := re.FindAllString(lines[1], -1)

	margin_of_error := 1

	for i := range times {
		time, _ := strconv.Atoi(times[i])
		distance, _ := strconv.Atoi(distances[i])
		margin_of_error *= WaysToWin(time, distance)
	}

	return margin_of_error
}

func Puzzle2(input string) int {
	lines := strings.Split(input, "\n")

	re := regexp.MustCompile(`\d+`)

	times := strings.Join(re.FindAllString(lines[0], -1), "")
	distances := strings.Join(re.FindAllString(lines[1], -1), "")

	time, _ := strconv.Atoi(times)
	distance, _ := strconv.Atoi(distances)
	margin_of_error := WaysToWin(time, distance)

	return margin_of_error
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

	c.Visit("https://adventofcode.com/2023/day/6/input")
}
