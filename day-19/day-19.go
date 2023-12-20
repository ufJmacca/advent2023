package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/PaesslerAG/gval"
	"github.com/gocolly/colly"
)

func PartEvaluation(part map[string]int, workflow []string, workflows map[string][]string) int {

	if len(workflow) == 0 {
		workflow = workflows["in"]
	}

	if workflow[0] == "A" {
		sum := 0
		for _, value := range part {
			sum += value
		}
		return sum
	}

	if workflow[0] == "R" {
		return 0
	}

	for _, flow := range workflow {
		criteria := strings.Split(flow, ":")
		if len(criteria) > 1 {

			value, _ := gval.Evaluate(criteria[0], part)

			if value == true {
				if strings.Contains("AR", criteria[1]) {
					return PartEvaluation(part, []string{criteria[1]}, workflows)
				}
				return PartEvaluation(part, workflows[criteria[1]], workflows)
			} else {
				return PartEvaluation(part, workflow[1:], workflows)
			}
		} else {
			return PartEvaluation(part, workflows[criteria[0]], workflows)
		}
	}

	return 0
}

func Puzzle1(input string) int {
	blocks := strings.Split(input, "\n\n")

	flows_regex := regexp.MustCompile(`(\w+){(.*?)}`)
	parts_regex := regexp.MustCompile(`(\w+)=(\d+)`)

	flows := make(map[string][]string)
	part := make(map[string]int)

	for _, line := range strings.Split(blocks[0], "\n") {
		if len(line) > 0 {
			matches := flows_regex.FindAllStringSubmatch(line, -1)
			flows[matches[0][1]] = append(flows[matches[0][1]], strings.Split(matches[0][2], ",")...)
		}
	}

	total := 0

	for _, line := range strings.Split(blocks[1], "\n") {
		if len(line) > 0 {
			matches := parts_regex.FindAllStringSubmatch(line, -1)
			for i := range matches {
				part_value, _ := strconv.Atoi(string(matches[i][2]))
				part[string(matches[i][1])] = part_value
			}
			total += PartEvaluation(part, []string{}, flows)
		}
	}

	return total
}

func Puzzle2(input string) int {

	return 0
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

		// puzzle_2 := Puzzle2(inputs)
		// fmt.Println(puzzle_2)
	})

	c.Visit("https://adventofcode.com/2023/day/19/input")
}
