package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

// Combine first and last digit in given string to form a two digit number
func Calibration(input string) int {
	re := regexp.MustCompile("[0-9]")
	first_digit, _ := strconv.Atoi(re.FindAllString(input, 1)[0])
	last_digit, _ := strconv.Atoi(re.FindAllString(input, -1)[len(re.FindAllString(input, -1))-1])
	return first_digit*10 + last_digit
}

// replace strings of numbers with their number counterparts, BE CAREFUL of overlapping string
// e.g. oneight should be replaced with 18 not 1
func CalibrationString(input string) int {
	r := strings.NewReplacer("oneight", "18", "twone", "21", "eightwo", "82", "one", "1", "two", "2", "three", "3", "four", "4", "five", "5", "six", "6", "seven", "7", "eight", "8", "nine", "9")
	// fmt.Println(r.Replace(input))
	return Calibration(r.Replace(input))
}

func main() {
	c := colly.NewCollector()

	// Sets cookie from environment variable
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("cookie", os.Getenv("COOKIE"))
	})

	c.OnResponse(func(r *colly.Response) {
		inputs := strings.Split(string(r.Body), "\n")

		calibration_sum_part_one := 0
		calibration_sum_part_two := 0

		for _, datum := range inputs {
			if len(datum) != 0 {
				calibration_sum_part_one += Calibration(datum)
				calibration_sum_part_two += CalibrationString(datum)
			}
		}

		fmt.Println(calibration_sum_part_one)
		fmt.Println(calibration_sum_part_two)
	})

	c.Visit("https://adventofcode.com/2023/day/1/input")
}
