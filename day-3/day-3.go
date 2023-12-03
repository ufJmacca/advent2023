package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"

	"github.com/gocolly/colly"
)

type MatrixMatch struct {
	Row      int
	ColStart int
	ColEnd   int
}

type GearMatch struct {
	Row int
	Col int
}

func isSymbol(r rune) bool {
	return !unicode.IsLetter(r) && !unicode.IsNumber(r) && r != rune("."[0])
}

func EngineSchematic(input string) int {
	lines := strings.Split(input, "\n")

	var matrix [][]string

	for _, line := range lines {
		elements := strings.Split(line, "")
		matrix = append(matrix, elements)
	}

	rows := len(matrix)
	cols := len(matrix[0])

	regex := regexp.MustCompile("[0-9]+")
	adjacentNumbers := make(map[MatrixMatch]struct{})

	for i := 0; i < rows; i++ {
		matches := regex.FindAllStringIndex(strings.Join(matrix[i], ""), -1)

		for j := 0; j < cols; j++ {
			if unicode.IsNumber(rune(matrix[i][j][0])) {
				for x := -1; x <= 1; x++ {
					for y := -1; y <= 1; y++ {
						newX := i + x
						newY := j + y

						if newX >= 0 && newX < rows && newY >= 0 && newY < cols && !(x == 0 && y == 0) {
							if isSymbol(rune(matrix[newX][newY][0])) {
								for _, match := range matches {
									startIndex := match[0]
									endIndex := match[1]

									if j >= startIndex && j < endIndex {
										coords := MatrixMatch{Row: i, ColStart: startIndex, ColEnd: endIndex}
										adjacentNumbers[coords] = struct{}{}
									}
								}
							}
						}
					}
				}
			}
		}
	}

	sumSchematics := 0

	for key, _ := range adjacentNumbers {
		intAdjacent, _ := strconv.Atoi(strings.Join(matrix[key.Row][key.ColStart:key.ColEnd], ""))
		sumSchematics += intAdjacent
	}

	return sumSchematics
}

func GearSchematic(input string) int {
	lines := strings.Split(input, "\n")

	var matrix [][]string

	for _, line := range lines {
		elements := strings.Split(line, "")
		matrix = append(matrix, elements)
	}

	rows := len(matrix)
	cols := len(matrix[0])

	regex := regexp.MustCompile("[0-9]+")
	adjacentNumbers := make(map[GearMatch][]MatrixMatch)

	for i := 0; i < rows; i++ {
		matches := regex.FindAllStringIndex(strings.Join(matrix[i], ""), -1)

		for j := 0; j < cols; j++ {
			if unicode.IsNumber(rune(matrix[i][j][0])) {
				for x := -1; x <= 1; x++ {
					for y := -1; y <= 1; y++ {
						newX := i + x
						newY := j + y

						if newX >= 0 && newX < rows && newY >= 0 && newY < cols && !(x == 0 && y == 0) {
							if matrix[newX][newY] == "*" {
								for _, match := range matches {
									startIndex := match[0]
									endIndex := match[1]

									if j >= startIndex && j < endIndex {
										gearCoords := GearMatch{Row: newX, Col: newY}
										coords := MatrixMatch{Row: i, ColStart: startIndex, ColEnd: endIndex}

										if _, ok := adjacentNumbers[gearCoords]; !ok {
											adjacentNumbers[gearCoords] = []MatrixMatch{}
										}
										adjacentNumbers[gearCoords] = append(adjacentNumbers[gearCoords], coords)
									}
								}
							}
						}
					}
				}
			}
		}
	}

	sumGearRatios := 0

	for _, gear := range adjacentNumbers {
		uniqueMatrixMatches := make(map[MatrixMatch]struct{})
		for _, matches := range gear {
			uniqueMatrixMatches[matches] = struct{}{}
		}
		if len(uniqueMatrixMatches) == 2 {
			result := 1

			for location, _ := range uniqueMatrixMatches {
				intAdjacent, _ := strconv.Atoi(strings.Join(matrix[location.Row][location.ColStart:location.ColEnd], ""))
				result *= intAdjacent
			}
			sumGearRatios += result
		}
	}

	return sumGearRatios
}

func main() {
	c := colly.NewCollector()

	// Sets cookie from environment variable
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("cookie", os.Getenv("COOKIE"))
	})

	c.OnResponse(func(r *colly.Response) {
		input := string(r.Body)
		lastNewlineIndex := strings.LastIndex(input, "\n")
		input = input[:lastNewlineIndex]

		partNumbers := EngineSchematic(input)
		fmt.Println(partNumbers)

		gearRatios := GearSchematic(input)
		fmt.Println(gearRatios)
	})

	c.Visit("https://adventofcode.com/2023/day/3/input")
}
