package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/gocolly/colly"
)

func GamePoints(wins int) int {
	if wins == 0 {
		return 0
	} else if wins == 1 {
		return 1
	} else {
		points := 1
		for i := 0; i < wins-1; i++ {
			points *= 2
		}
		return points
	}
}

func GameResult(input string) (int, int) {
	re := regexp.MustCompile(`.*?\: (.*?) \| (.*?)$`)

	winning_numbers := strings.Fields(re.FindStringSubmatch(input)[1])
	our_numbers := strings.Fields(re.FindStringSubmatch(input)[2])

	win_count := 0

	for _, our_number := range our_numbers {
		for _, winning_number := range winning_numbers {
			if our_number == winning_number {
				win_count += 1
			}
		}
	}

	return GamePoints(win_count), win_count
}

func GameResultWithDuplicates(input string) int {
	lines := strings.Split(input, "\n")

	// re := regexp.MustCompile(`.*?(\d+)\:.*?$`)

	results := make(map[int]int)

	for i := range lines {
		results[i+1] = 1
	}

	for index, line := range lines {
		map_i := index + 1
		games_to_play := results[map_i]

		for i := 0; i < games_to_play; i++ {
			_, wins := GameResult(line)

			for i := 0; i < wins; i++ {
				results[map_i+i+1] += 1
			}
		}

	}

	total_games_played := 0

	for _, games := range results {
		total_games_played += games
	}

	return total_games_played
}

func main() {
	c := colly.NewCollector()

	// Sets cookie from environment variable
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("cookie", os.Getenv("COOKIE"))
	})

	c.OnResponse(func(r *colly.Response) {
		inputs := strings.Split(string(r.Body), "\n")
		input := inputs[:len(inputs)-1]

		game_results_scores := 0

		for _, datum := range input {
			points, _ := GameResult(datum)
			game_results_scores += points
		}

		fmt.Println(game_results_scores)

		input_2 := string(r.Body)
		lastNewlineIndex := strings.LastIndex(input_2, "\n")
		input_2 = input_2[:lastNewlineIndex]

		with_duplicates := GameResultWithDuplicates(input_2)
		fmt.Println(with_duplicates)
	})

	c.Visit("https://adventofcode.com/2023/day/4/input")
}
