package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/gocolly/colly"
)

type path_nodes struct {
	R string
	L string
}

func PathSolver(navigation string, nodes map[string]path_nodes, start_node string) int {
	navigation_length := len(navigation)

	current_node := nodes[start_node]

	steps := 0
	for current_node != nodes["ZZZ"] {
		fields := map[string]string{
			"R": current_node.R,
			"L": current_node.L,
		}

		next_step := string(navigation[steps%navigation_length])
		current_node = nodes[fields[next_step]]
		steps += 1
	}

	return steps
}

func PathSolver2(navigation string, nodes map[string]path_nodes, start_node string) int {
	navigation_length := len(navigation)

	current_node := start_node

	steps := 0
	for string(current_node[len(current_node)-1]) != "Z" {
		fields := map[string]string{
			"R": nodes[current_node].R,
			"L": nodes[current_node].L,
		}

		next_step := string(navigation[steps%navigation_length])
		current_node = fields[next_step]
		steps += 1
	}

	return steps
}

func PathMap(input string) int {
	nodes := make(map[string]path_nodes)
	lines := strings.Split(input, "\n")

	navigation := lines[0]

	re := regexp.MustCompile(`(\w+) = \((\w+), (\w+)\)`)

	for _, line := range lines[2:] {
		if len(line) > 0 {
			matches := re.FindAllStringSubmatch(line, -1)
			nodes[matches[0][1]] = path_nodes{R: matches[0][3], L: matches[0][2]}
		}
	}

	return PathSolver(navigation, nodes, "AAA")
}

func GhostPathMap(input string) int {
	nodes := make(map[string]path_nodes)
	lines := strings.Split(input, "\n")

	navigation := lines[0]

	re := regexp.MustCompile(`(\w+) = \((\w+), (\w+)\)`)

	var current_nodes []map[string]path_nodes

	for _, line := range lines[2:] {
		if len(line) > 0 {
			matches := re.FindAllStringSubmatch(line, -1)
			node_name := matches[0][1]

			node := path_nodes{
				R: matches[0][3],
				L: matches[0][2],
			}

			nodes[node_name] = node

			if node_name[len(node_name)-1] == 'A' {
				startingNode := map[string]path_nodes{
					node_name: node,
				}
				current_nodes = append(current_nodes, startingNode)
			}
		}
	}
	fmt.Println(current_nodes)

	var start_node_steps []int
	for _, node := range current_nodes {
		for index, current_node := range node {
			fmt.Println(index)
			fmt.Println(current_node)
			start_node_steps = append(start_node_steps, PathSolver2(navigation, nodes, index))
		}
	}
	fmt.Println(start_node_steps)

	return LCM(start_node_steps)
}

func AllFinished(input []map[string]path_nodes) bool {
	for _, value := range input {
		for key := range value {
			if string(key[len(key)-1]) != "Z" {
				return false
			}
		}
	}
	return true
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(input []int) int {
	result := input[0] * input[1] / GCD(input[0], input[1])

	for i := 0; i < len(input[2:]); i++ {
		list := append([]int{result}, input[2:]...)
		result = LCM(list)
	}

	return result
}

func main() {
	c := colly.NewCollector()

	// Sets cookie from environment variable
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("cookie", os.Getenv("COOKIE"))
	})

	c.OnResponse(func(r *colly.Response) {
		inputs := string(r.Body)

		steps_required_1 := PathMap(inputs)
		fmt.Println(steps_required_1)

		steps_requires_2 := GhostPathMap(inputs)
		fmt.Println(steps_requires_2)
	})

	c.Visit("https://adventofcode.com/2023/day/8/input")
}
