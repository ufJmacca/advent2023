package main

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/peterstace/simplefeatures/geom"
)

type directions struct {
	direction string
	steps     int
	colour    string
}

func Puzzle1(input string) int {
	lines := strings.Split(input, "\n")
	pattern := `(\w) (\d+) \((.*?)\)`
	regex := regexp.MustCompile(pattern)

	var dig_instructions []directions
	var flat_sequence []float64

	current_location := []float64{0, 0}

	flat_sequence = append(flat_sequence, current_location...)

	for _, line := range lines {
		if len(line) > 0 {
			matches := regex.FindAllStringSubmatch(line, -1)
			direction := matches[0][1]
			steps, _ := strconv.Atoi(matches[0][2])
			colour := matches[0][3]
			dig_instructions = append(dig_instructions, directions{direction, steps, colour})

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
			flat_sequence = append(flat_sequence, current_location...)
		}
	}

	seq := geom.NewSequence(flat_sequence, geom.DimXY)

	var polygon_lines []geom.LineString
	polygon_lines = append(polygon_lines, geom.NewLineString(seq))

	polygon := geom.NewPolygon(polygon_lines)
	cnt := 0

	for i := 0; i < 9; i++ {
		for j := 0; j < 6; j++ {
			point := geom.NewPoint(geom.Coordinates{XY: geom.XY{X: float64(i), Y: float64(j)}, Type: geom.DimXY})
			contains, _ := geom.Within(point.AsGeometry(), polygon.AsGeometry())
			if contains {
				cnt++
			}
		}
	}

	return cnt + int(polygon.Boundary().Length())
}
