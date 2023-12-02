package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

// See if Game is possible by conparing cubes pulled against actual bag contents
func GamePossible(input string) (int, string) {
	actualBagContents := make(map[string]int)

	actualBagContents["red"] = 12
	actualBagContents["green"] = 13
	actualBagContents["blue"] = 14

	game_index_pattern := regexp.MustCompile(`\d+`)
	colour_count_pattern := regexp.MustCompile(`(?P<count>\d+) (?P<colour>\w+)`)

	split := strings.Split(input, ":")

	game_index, _ := strconv.Atoi(game_index_pattern.FindAllString(split[0], 1)[0])

	split_games := strings.Split(split[1], ";")

	for _, datum := range split_games {
		split_game := strings.Split(datum, ",")

		for _, game_datum := range split_game {
			matches := colour_count_pattern.FindStringSubmatch(game_datum)

			i, _ := strconv.Atoi(matches[1])

			if i > actualBagContents[matches[2]] {
				return game_index, "impossible"
			}
		}
	}

	return game_index, "possible"
}

// Identifing the minimum count of cubes for each colour in a bag to match observered cube extraction
func MinimumPossibleCubes(input string) map[string]int {
	minimumBageContents := make(map[string]int)

	minimumBageContents["red"] = 0
	minimumBageContents["green"] = 0
	minimumBageContents["blue"] = 0

	colour_count_pattern := regexp.MustCompile(`(?P<count>\d+) (?P<colour>\w+)`)

	split := strings.Split(input, ":")

	split_games := strings.Split(split[1], ";")

	for _, datum := range split_games {
		split_game := strings.Split(datum, ",")

		for _, game_datum := range split_game {
			matches := colour_count_pattern.FindStringSubmatch(game_datum)

			i, _ := strconv.Atoi(matches[1])

			if i > minimumBageContents[matches[2]] {
				minimumBageContents[matches[2]] = i
			}
		}
	}

	return minimumBageContents
}

func BagPower(input map[string]int) int {
	power := 1

	for _, item := range input {
		power *= item
	}

	return power
}

func main() {
	c := colly.NewCollector()

	// Sets cookie from environment variable
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("cookie", os.Getenv("COOKIE"))
	})

	c.OnResponse(func(r *colly.Response) {
		inputs := strings.Split(string(r.Body), "\n")

		possibleIndexSum := 0
		totalBagPower := 0

		for _, datum := range inputs {
			if len(datum) != 0 {
				// Part 1
				index, outcome := GamePossible(datum)
				if outcome == "possible" {
					possibleIndexSum += index
				}

				// Part 2
				totalBagPower += BagPower(MinimumPossibleCubes(datum))
			}
		}

		fmt.Println(possibleIndexSum)
		fmt.Println(totalBagPower)
	})

	c.Visit("https://adventofcode.com/2023/day/2/input")
}
