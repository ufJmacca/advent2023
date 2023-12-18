package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
	"github.com/peterstace/simplefeatures/geom"
)

func Puzzle1(input string) int {
	lines := strings.Split(input, "\n")
	pattern := `(\w) (\d+) \((.*?)\)`
	regex := regexp.MustCompile(pattern)

	var flat_sequence []float64

	current_location := []float64{0, 0}

	flat_sequence = append(flat_sequence, current_location...)

	min_row := 0
	max_row := 0
	min_col := 0
	max_col := 0

	for _, line := range lines {
		if len(line) > 0 {
			matches := regex.FindAllStringSubmatch(line, -1)
			direction := matches[0][1]
			steps, _ := strconv.Atoi(matches[0][2])

			switch direction {
			case "U":
				current_location[0] = current_location[0] + (float64(steps) * -1)
			case "D":
				current_location[0] = current_location[0] + float64(steps)
			case "R":
				current_location[1] = current_location[1] + float64(steps)
			case "L":
				current_location[1] = current_location[1] + (float64(steps) * -1)
			}
			min_row = min(min_row, int(current_location[0]))
			max_row = max(max_row, int(current_location[0]))
			min_col = min(min_col, int(current_location[1]))
			max_col = max(max_col, int(current_location[1]))

			flat_sequence = append(flat_sequence, current_location...)
		}
	}

	seq := geom.NewSequence(flat_sequence, geom.DimXY)

	var polygon_lines []geom.LineString
	polygon_lines = append(polygon_lines, geom.NewLineString(seq))

	polygon := geom.NewPolygon(polygon_lines)

	boundary_points := countBoundaryLatticePoints(seq)

	interior := int(polygon.Area()) - (boundary_points / 2) + 1

	fmt.Println(interior)
	fmt.Println(boundary_points)

	cnt := 0

	for i := min_row; i < max_row; i++ {
		for j := min_col; j < max_col; j++ {
			point := geom.NewPoint(geom.Coordinates{XY: geom.XY{X: float64(i), Y: float64(j)}, Type: geom.DimXY})
			contains, _ := geom.Within(point.AsGeometry(), polygon.AsGeometry())
			if contains {
				cnt++
			}
		}
	}

	return cnt + int(polygon.Boundary().Length())
}

func InstructionDecode(hexa string) (string, int) {
	steps, _ := strconv.ParseInt(hexa[1:6], 16, 64)
	var direction string
	switch hexa[6:] {
	case "0":
		direction = "R"
	case "1":
		direction = "D"
	case "2":
		direction = "L"
	case "3":
		direction = "U"
	}
	return direction, int(steps)
}

func countBoundaryLatticePoints(polygon_lines geom.Sequence) int {
	// Iterate through the line strings representing the edges of the polygon
	boundaryPoints := 0
	fmt.Println(polygon_lines.Length())
	for i := 0; i < polygon_lines.Length()-1; i++ {

		// Check each segment of the line string
		// Count points with integer coordinates along the line segment
		start_x := polygon_lines.Get(i).X
		start_y := polygon_lines.Get(i).Y
		end_x := polygon_lines.Get(i + 1).X
		end_y := polygon_lines.Get(i + 1).Y

		// Count lattice points along the line segment (using a simple count, not considering fractional points)
		boundaryPoints += gcd(abs(int(end_x-start_x)), abs(int(end_y-start_y)))
	}
	return boundaryPoints
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Puzzle2(input string) int {
	lines := strings.Split(input, "\n")
	pattern := `(\w) (\d+) \((.*?)\)`
	regex := regexp.MustCompile(pattern)

	var flat_sequence []float64

	current_location := []float64{0, 0}

	flat_sequence = append(flat_sequence, current_location...)

	min_row := 0
	max_row := 0
	min_col := 0
	max_col := 0

	for _, line := range lines {
		if len(line) > 0 {
			matches := regex.FindAllStringSubmatch(line, -1)
			direction, steps := InstructionDecode(matches[0][3])

			switch direction {
			case "U":
				current_location[0] = current_location[0] + (float64(steps) * -1)
			case "D":
				current_location[0] = current_location[0] + float64(steps)
			case "R":
				current_location[1] = current_location[1] + float64(steps)
			case "L":
				current_location[1] = current_location[1] + (float64(steps) * -1)
			}
			min_row = min(min_row, int(current_location[0]))
			max_row = max(max_row, int(current_location[0]))
			min_col = min(min_col, int(current_location[1]))
			max_col = max(max_col, int(current_location[1]))

			flat_sequence = append(flat_sequence, current_location...)
		}
	}

	seq := geom.NewSequence(flat_sequence, geom.DimXY)

	var polygon_lines []geom.LineString
	polygon_lines = append(polygon_lines, geom.NewLineString(seq))

	polygon := geom.NewPolygon(polygon_lines)

	boundary_points := countBoundaryLatticePoints(seq)

	interior := int(polygon.Area()) - (boundary_points / 2) + 1

	fmt.Println(interior)
	fmt.Println(boundary_points)

	return interior + boundary_points
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

	c.Visit("https://adventofcode.com/2023/day/18/input")
}
