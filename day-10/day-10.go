package main

import (
	"fmt"
	"os"
	"reflect"
	"sort"
	"strings"

	"github.com/gocolly/colly"
	"github.com/peterstace/simplefeatures/geom"
)

type coords struct {
	x int
	y int
}

type next_path_steps struct {
	steps  int
	coords coords
}

func equalStructSlicesIgnoreOrder(s1, s2 []coords) bool {
	if len(s1) != len(s2) {
		return false
	}

	s1Sorted := make([]coords, len(s1))
	copy(s1Sorted, s1)
	s2Sorted := make([]coords, len(s2))
	copy(s2Sorted, s2)

	sort.Slice(s1Sorted, func(i, j int) bool {
		if s1Sorted[i].x != s1Sorted[j].x {
			return s1Sorted[i].x < s1Sorted[j].x
		}
		return s1Sorted[i].y < s1Sorted[j].y
	})
	sort.Slice(s2Sorted, func(i, j int) bool {
		if s2Sorted[i].x != s2Sorted[j].x {
			return s2Sorted[i].x < s2Sorted[j].x
		}
		return s2Sorted[i].y < s2Sorted[j].y
	})

	return reflect.DeepEqual(s1Sorted, s2Sorted)
}

func TileConnections(tile string) []coords {
	switch tile {
	case "|":
		return []coords{
			{-1, 0},
			{1, 0},
		}
	case "-":
		return []coords{
			{0, -1},
			{0, 1},
		}
	case "L":
		return []coords{
			{-1, 0},
			{0, 1},
		}
	case "J":
		return []coords{
			{-1, 0},
			{0, -1},
		}
	case "7":
		return []coords{
			{1, 0},
			{0, -1},
		}
	case "F":
		return []coords{
			{1, 0},
			{0, 1},
		}
	default:
		return []coords{}
	}
}

func TileType(input_coords []coords, start_location coords) string {
	var test_coords []coords

	for _, co := range input_coords {
		test_coords = append(test_coords, coords{x: co.x - start_location.x, y: co.y - start_location.y})
	}

	if equalStructSlicesIgnoreOrder(test_coords, []coords{
		{-1, 0},
		{1, 0},
	}) {
		return "|"
	} else if equalStructSlicesIgnoreOrder(test_coords, []coords{
		{0, -1},
		{0, 1},
	}) {
		return "-"
	} else if equalStructSlicesIgnoreOrder(test_coords, []coords{
		{-1, 0},
		{0, 1},
	}) {
		return "L"
	} else if equalStructSlicesIgnoreOrder(test_coords, []coords{
		{-1, 0},
		{0, -1},
	}) {
		return "J"
	} else if equalStructSlicesIgnoreOrder(test_coords, []coords{
		{1, 0},
		{0, -1},
	}) {
		return "7"
	} else if equalStructSlicesIgnoreOrder(test_coords, []coords{
		{1, 0},
		{0, 1},
	}) {
		return "F"
	} else {
		return ""
	}
}

func FindStart(input [][]string) coords {
	rows := len(input)
	cols := len(input[0])

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if input[i][j] == "S" {
				return coords{x: i, y: j}
			}
		}
	}
	return coords{}
}

func FindConnectedTiles(grid [][]string, location coords) []coords {
	rows := len(grid)
	cols := len(grid[0])
	var connected_tiles []coords

	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			newX := x + location.x
			newY := y + location.y

			if newX >= 0 && newX < rows && newY >= 0 && newY < cols {
				if (x == 0 || y == 0) && !(x == 0 && y == 0) {
					target_tile_connections := TileConnections(grid[newX][newY])
					if len(target_tile_connections) > 0 {
						for _, target_tile_connection := range target_tile_connections {
							if x+target_tile_connection.x == 0 && y+target_tile_connection.y == 0 {
								connected_tiles = append(connected_tiles, coords{x: newX, y: newY})
							}
						}
					}
				}

			}
		}
	}
	return connected_tiles
}

func ContainsStruct(tiles []coords, t coords) bool {
	for _, tile := range tiles {
		if tile == t {
			return true
		}
	}
	return false
}

func Puzzle1(input string) int {
	lines := strings.Split(input, "\n")

	var matrix [][]string

	for _, line := range lines {
		if len(line) > 0 {
			elements := strings.Split(line, "")
			matrix = append(matrix, elements)
		}
	}

	var visited_tiles []coords
	var unvisited_tiles []next_path_steps

	unvisited_tiles = append(unvisited_tiles, next_path_steps{steps: 0, coords: FindStart(matrix)})

	max_step := 0

	for len(unvisited_tiles) > 0 {
		current_tile := unvisited_tiles[0]
		visited_tiles = append(visited_tiles, current_tile.coords)

		for _, tile := range FindConnectedTiles(matrix, current_tile.coords) {
			if !ContainsStruct(visited_tiles, tile) {
				unvisited_tiles = append(unvisited_tiles, next_path_steps{steps: current_tile.steps + 1, coords: tile})
			}
		}
		if current_tile.steps > max_step {
			max_step = current_tile.steps
		}
		unvisited_tiles = unvisited_tiles[1:]
	}

	return max_step
}

func CoordsInSlice(slice []coords, cell coords) bool {
	for _, item := range slice {
		if item == cell {
			return true
		}
	}
	return false
}

func ConnectedTiles(node coords, grid [][]string) []coords {
	tile_connections := TileConnections(grid[node.x][node.y])

	var output []coords
	for _, connection := range tile_connections {
		output = append(output, coords{x: node.x + connection.x, y: node.y + connection.y})
	}

	return output
}

func dfs(node coords, visited map[coords]bool, grid [][]string, visited_order *[]coords) {
	visited[node] = true
	*visited_order = append(*visited_order, node)

	for _, neighbor := range ConnectedTiles(node, grid) {
		if !visited[neighbor] {
			dfs(neighbor, visited, grid, visited_order)
		}
	}
}

// TODO - Need to scrap BFS search for a DFS search to follow the loop in one
// direction all the way till it hits the starting square again.
func Puzzle2(input string) int {
	lines := strings.Split(input, "\n")

	var matrix [][]string

	for _, line := range lines {
		if len(line) > 0 {
			elements := strings.Split(line, "")
			matrix = append(matrix, elements)
		}
	}

	var visited_tiles []coords
	var unvisited_tiles []next_path_steps

	unvisited_tiles = append(unvisited_tiles, next_path_steps{steps: 0, coords: FindStart(matrix)})

	max_step := 0

	for len(unvisited_tiles) > 0 {
		current_tile := unvisited_tiles[0]
		visited_tiles = append(visited_tiles, current_tile.coords)

		for _, tile := range FindConnectedTiles(matrix, current_tile.coords) {
			if !ContainsStruct(visited_tiles, tile) {
				unvisited_tiles = append(unvisited_tiles, next_path_steps{steps: current_tile.steps + 1, coords: tile})
			}
		}
		if current_tile.steps > max_step {
			max_step = current_tile.steps
		}
		unvisited_tiles = unvisited_tiles[1:]
	}

	start_tile := FindStart(matrix)
	matrix[start_tile.x][start_tile.y] = TileType(FindConnectedTiles(matrix, FindStart(matrix)), FindStart(matrix))

	var visited_order []coords
	visited := make(map[coords]bool)
	dfs(start_tile, visited, matrix, &visited_order)
	visited_order = append(visited_order, start_tile)

	var new_loop [][]float64

	for _, tile := range visited_order {
		new_loop = append(new_loop, []float64{float64(tile.x), float64(tile.y)})
	}

	var flat_sequence []float64
	for _, point := range new_loop {
		flat_sequence = append(flat_sequence, point...)
	}

	seq := geom.NewSequence(flat_sequence, geom.DimXY)

	var polygon_lines []geom.LineString

	polygon_lines = append(polygon_lines, geom.NewLineString(seq))

	polygon := geom.NewPolygon(polygon_lines)

	rows := len(matrix)
	cols := len(matrix[0])
	var new_matrix [][]string

	for i := 0; i < rows; i++ {
		new_matrix = append(new_matrix, make([]string, cols))
	}

	var inside []geom.Point

	for x, row := range matrix {
		for y := range row {
			xy := geom.XY{X: float64(x), Y: float64(y)}
			coords := geom.Coordinates{XY: xy, Type: geom.DimXY}
			point := geom.NewPoint(coords)
			contains, _ := geom.Within(point.AsGeometry(), polygon.AsGeometry())
			if contains {
				inside = append(inside, point)
				new_matrix[x][y] = "1"
			} else {
				new_matrix[x][y] = "."
			}
		}
	}

	return len(inside)
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

	c.Visit("https://adventofcode.com/2023/day/10/input")
}
