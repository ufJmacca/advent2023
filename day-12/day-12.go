package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

type memo_key struct {
	records    string
	conditions []int
}

func StringifyMemoKey(key memo_key) string {
	return fmt.Sprintf("%s:%v", key.records, key.conditions)
}

var memo = make(map[string]int)

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

	key := StringifyMemoKey(memo_key{records: records, conditions: conditions})
	if val, ok := memo[key]; ok {
		return val
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

	memo[key] = result
	return result
}

func Puzzle1(input string) int {
	lines := strings.Split(input, "\n")

	total := 0

	for _, line := range lines {
		if len(line) > 0 {
			elements := strings.Split(line, " ")
			var conditions []int
			for _, condition := range strings.Split(elements[1], ",") {
				c1, _ := strconv.Atoi(condition)
				conditions = append(conditions, c1)
			}

			total += ArrangementCounter(elements[0], conditions)
		}
	}

	return total
}

func Puzzle2(input string) int {
	lines := strings.Split(input, "\n")

	total := 0

	for _, line := range lines {
		if len(line) > 0 {
			elements := strings.Split(line, " ")
			var conditions []int
			for _, condition := range strings.Split(elements[1], ",") {
				c1, _ := strconv.Atoi(condition)
				conditions = append(conditions, c1)
			}

			count_input := strings.Join([]string{strings.Repeat(elements[0]+"?", 5)}, "")

			appended := make([]int, len(conditions))
			copy(appended, conditions)
			for i := 0; i < 4; i++ {
				appended = append(appended, conditions...)
			}

			total += ArrangementCounter(count_input[:len(count_input)-1], appended)
		}
	}

	return total
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

	c.Visit("https://adventofcode.com/2023/day/12/input")
}
