package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
	"github.com/gocolly/colly"
)

type line struct {
	hand string
	bet  int
}

type hand struct {
	Index     int
	Hand      string
	Bet       int
	HandType  int
	HandScore int
}

func characterValue(input string) int {
	switch input {
	case `A`:
		return 12
	case `K`:
		return 11
	case `Q`:
		return 10
	case `J`:
		return 9
	case `T`:
		return 8
	case `9`:
		return 7
	case `8`:
		return 6
	case `7`:
		return 5
	case `6`:
		return 4
	case `5`:
		return 3
	case `4`:
		return 2
	case `3`:
		return 1
	case `2`:
		return 0
	}

	return 0
}

func characterValue2(input string) int {
	switch input {
	case `A`:
		return 13
	case `K`:
		return 12
	case `Q`:
		return 11
	case `T`:
		return 10
	case `9`:
		return 9
	case `8`:
		return 8
	case `7`:
		return 7
	case `6`:
		return 6
	case `5`:
		return 5
	case `4`:
		return 4
	case `3`:
		return 3
	case `2`:
		return 2
	case `J`:
		return 1
	}

	return 0
}

func sortByValue(input map[rune]int) []rune {
	keys := make([]rune, 0, len(input))

	for key := range input {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		return input[keys[i]] > input[keys[j]]
	})

	return keys
}

func HandType(hand string) int {
	counts := make(map[rune]int)

	for _, char := range hand {
		counts[char]++
	}

	sortedKeys := sortByValue(counts)

	for _, value := range sortedKeys {
		if counts[value] == 5 {
			return 6
		} else if counts[value] == 4 {
			return 5
		} else if counts[value] == 3 && len(counts) == 2 {
			return 4
		} else if counts[value] == 3 {
			return 3
		} else if counts[value] == 2 && len(counts) == 3 {
			return 2
		} else if counts[value] == 2 {
			return 1
		} else {
			return 0
		}
	}

	return 0
}

func sortByKeyValue(input map[int][]string) []int {
	var keys []int

	for key := range input {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		return keys[i] > keys[j]
	})

	return keys
}

func BestHand(hand string) string {
	alternative_cards := []string{
		`A`,
		`K`,
		`Q`,
		`T`,
		`9`,
		`8`,
		`7`,
		`6`,
		`5`,
		`4`,
		`3`,
		`2`,
	}

	re := regexp.MustCompile(`[J]`)

	j_matches := re.FindAllStringSubmatchIndex(hand, -1)

	// fmt.Println(hand)
	// fmt.Println(j_matches)

	current_score := HandType(strings.ReplaceAll(hand, "J", ""))
	better_hands := make(map[int][]string)

	for _, match := range j_matches {
		for _, key := range alternative_cards {
			new_hand_runes := []rune(hand)
			new_hand_runes[match[0]] = rune(key[0])
			new_hand := string(new_hand_runes)

			new_score := HandType(new_hand)

			fmt.Printf("Hand %s has score %d Old hand %s has score %d\n", new_hand, new_score, hand, current_score)

			if new_score > current_score {
				better_hands[new_score] = append(better_hands[new_score], new_hand)
			}
		}
	}

	fmt.Println(better_hands)

	if len(better_hands) > 0 {
		// fmt.Println(better_hands)

		sorted_best_hands := sortByKeyValue(better_hands)

		for _, hand := range better_hands[sorted_best_hands[0]] {
			fmt.Println("recursion?")
			fmt.Println(hand)
			return BestHand(hand)
		}
	}

	return hand
}

func HandType2(hand string) int {
	counts := make(map[rune]int)
	best_hand := BestHand(hand)

	fmt.Println(hand)
	fmt.Println(best_hand)

	for _, char := range best_hand {
		counts[char]++
	}

	sortedKeys := sortByValue(counts)

	for _, value := range sortedKeys {
		if counts[value] == 5 {
			return 6
		} else if counts[value] == 4 {
			return 5
		} else if counts[value] == 3 && len(counts) == 2 {
			return 4
		} else if counts[value] == 3 {
			return 3
		} else if counts[value] == 2 && len(counts) == 3 {
			return 2
		} else if counts[value] == 2 {
			return 1
		} else {
			return 0
		}
	}

	return 0
}

func HandScore(hand string) int {
	score := 0

	for index, char := range hand {
		zeroString := "1"
		for i := index; i < 5; i++ {
			zeroString += "000"
		}

		multiple, _ := strconv.Atoi(zeroString)

		score += characterValue(string(char)) * multiple
	}

	return score
}

func HandScore2(hand string) int {
	score := 0

	for index, char := range hand {
		zeroString := "1"
		for i := index; i < 5; i++ {
			zeroString += "00"
		}

		multiple, _ := strconv.Atoi(zeroString)

		score += characterValue2(string(char)) * multiple
	}

	return score
}

func Puzzle1(input string) int {
	lines := strings.Split(input, "\n")

	var hands []hand

	re := regexp.MustCompile(`(.{5}) (\d+)`)

	for index, line := range lines {
		if len(line) > 0 {
			matches := re.FindAllStringSubmatch(line, -1)
			bet, _ := strconv.Atoi(matches[0][2])
			hands = append(hands, hand{Index: index, Hand: matches[0][1], Bet: bet, HandType: HandType(matches[0][1]), HandScore: HandScore(matches[0][1])})
		}
	}

	df := dataframe.LoadStructs(hands)
	sorted := df.Arrange(
		dataframe.Sort("HandType"),
		dataframe.Sort("HandScore"),
	)

	idx := 0

	nc := sorted.Rapply(func(s series.Series) series.Series {
		idx++
		bet, _ := s.Elem(2).Int()
		return series.Ints(bet * idx)
	})

	sorted = sorted.Mutate(nc.Col("X0")).Rename("winnings", "X0")

	return int(sorted.Col("winnings").Sum())
}

func Puzzle2(input string) int {
	lines := strings.Split(input, "\n")

	var hands []hand

	re := regexp.MustCompile(`(.{5}) (\d+)`)

	for index, line := range lines {
		if len(line) > 0 {
			matches := re.FindAllStringSubmatch(line, -1)
			bet, _ := strconv.Atoi(matches[0][2])
			hands = append(hands, hand{Index: index, Hand: matches[0][1], Bet: bet, HandType: HandType2(matches[0][1]), HandScore: HandScore2(matches[0][1])})
		}
	}

	df := dataframe.LoadStructs(hands)
	sorted := df.Arrange(
		dataframe.Sort("HandType"),
		dataframe.Sort("HandScore"),
	)

	idx := 0

	nc := sorted.Rapply(func(s series.Series) series.Series {
		idx++
		bet, _ := s.Elem(2).Int()
		return series.Ints(bet * idx)
	})

	sorted = sorted.Mutate(nc.Col("X0")).Rename("winnings", "X0")

	f, _ := os.Create("output.csv")
	sorted.WriteCSV(f)

	return int(sorted.Col("winnings").Sum())
}

func main() {
	c := colly.NewCollector()

	// Sets cookie from environment variable
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("cookie", os.Getenv("COOKIE"))
	})

	c.OnResponse(func(r *colly.Response) {
		inputs := string(r.Body)

		sum_winnings := Puzzle1(inputs)
		fmt.Println(sum_winnings)

		sum_winnings_2 := Puzzle2(inputs)
		fmt.Println(sum_winnings_2)
	})

	c.Visit("https://adventofcode.com/2023/day/7/input")
}
