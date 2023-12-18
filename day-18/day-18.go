package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"

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

func processChunk(minRow, maxRow, minCol, maxCol int, polygon geom.Polygon, cnt *int, wg *sync.WaitGroup, mutex *sync.Mutex) {
	defer wg.Done()

	chunkCnt := 0

	for i := minRow; i < maxRow; i++ {
		for j := minCol; j < maxCol; j++ {
			point := geom.NewPoint(geom.Coordinates{XY: geom.XY{X: float64(i), Y: float64(j)}, Type: geom.DimXY})
			contains, _ := geom.Within(point.AsGeometry(), polygon.AsGeometry())
			if contains {
				chunkCnt++
			}
		}
	}

	fmt.Println(chunkCnt)

	// Safely update the shared count variable
	// Using a mutex to avoid data races
	mutex.Lock()
	*cnt += chunkCnt
	mutex.Unlock()
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
	num_works := 31

	var wg sync.WaitGroup
	var mutex sync.Mutex

	cnt := 0
	rows_per_worker := (max_row - min_row) / num_works

	for i := 0; i < num_works; i++ {
		wg.Add(1)

		go func(worker_id int) {
			defer wg.Done()

			start_row := min_row + (worker_id * rows_per_worker)
			end_row := start_row + rows_per_worker

			if worker_id == num_works-1 {
				end_row = max_row
			}
			processChunk(start_row, end_row, min_col, max_col, polygon, &cnt, &wg, &mutex)
		}(i)
	}

	wg.Wait()

	fmt.Printf("count - %d\n", cnt)
	fmt.Printf("boundary length - %d\n", int(polygon.Boundary().Length()))
	fmt.Printf("total - %d", cnt+int(polygon.Boundary().Length()))

	return cnt + int(polygon.Boundary().Length())
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

	c.Visit("https://adventofcode.com/2023/day/18/input")
}
