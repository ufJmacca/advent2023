package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strings"

	"github.com/gocolly/colly"
)

type module struct {
	module_type string
	name        string
	memory      map[string]string
	outputs     []string
}

type queue_item struct {
	origin string
	target string
	pulse  string
}

type queue struct {
	items []queue_item
}

func (q *queue) Enqueue(item queue_item) {
	q.items = append(q.items, item)
}

func (q *queue) Dequeue() queue_item {
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

func Puzzle1(input string) int {
	line_regex := regexp.MustCompile(`(.*?) -> (.*?)\n`)
	matches := line_regex.FindAllStringSubmatch(input, -1)

	modules := make(map[string]module)
	var broadcast_targets []string

	for _, match := range matches {
		if match[1] == "broadcaster" {
			broadcast_targets = append(broadcast_targets, strings.Split(match[2], ", ")...)
		} else {
			module_type := match[1][0]
			name := match[1][1:]
			modules[name] = module{
				module_type: string(module_type),
				name:        name,
				outputs:     strings.Split(match[2], ", "),
				memory:      make(map[string]string),
			}
		}
	}

	for _, module := range modules {
		if module.module_type == "%" {
			module.memory["self"] = "off"
		}

		for _, output := range module.outputs {
			if _, ok := modules[output]; ok && modules[output].module_type == "&" {
				modules[output].memory[module.name] = "low"
			}
		}
	}

	low := 0
	high := 0

	for i := 0; i < 1000; i++ {
		low += 1
		p_queue := queue{}

		for _, target := range broadcast_targets {
			p_queue.Enqueue(queue_item{
				origin: "broadcaster",
				target: target,
				pulse:  "low",
			})
		}

		for len(p_queue.items) > 0 {
			element := p_queue.Dequeue()

			if element.pulse == "low" {
				low += 1
			} else {
				high += 1
			}

			if _, ok := modules[element.target]; !ok {
				continue
			}

			module := modules[element.target]

			if module.module_type == "%" {
				if element.pulse == "low" {
					if module.memory["self"] == "on" {
						module.memory["self"] = "off"
					} else {
						module.memory["self"] = "on"
					}

					var outgoing string
					if module.memory["self"] == "on" {
						outgoing = "high"
					} else {
						outgoing = "low"
					}

					for _, output := range module.outputs {
						p_queue.Enqueue(queue_item{
							origin: module.name,
							target: output,
							pulse:  outgoing,
						})
					}
				}
			} else {
				module.memory[element.origin] = element.pulse
				outgoing := "low"
				for _, memory := range module.memory {
					if memory != "high" {
						outgoing = "high"
					}
				}
				for _, x := range module.outputs {
					p_queue.Enqueue(queue_item{
						origin: module.name,
						target: x,
						pulse:  outgoing,
					})
				}
			}

		}
	}

	return low * high
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

func Puzzle2(input string) int {
	line_regex := regexp.MustCompile(`(.*?) -> (.*?)\n`)
	matches := line_regex.FindAllStringSubmatch(input, -1)

	modules := make(map[string]module)
	var broadcast_targets []string

	for _, match := range matches {
		if match[1] == "broadcaster" {
			broadcast_targets = append(broadcast_targets, strings.Split(match[2], ", ")...)
		} else {
			module_type := match[1][0]
			name := match[1][1:]
			modules[name] = module{
				module_type: string(module_type),
				name:        name,
				outputs:     strings.Split(match[2], ", "),
				memory:      make(map[string]string),
			}
		}
	}

	var feed string
	for _, module := range modules {
		if module.module_type == "%" {
			module.memory["self"] = "off"
		}

		for _, output := range module.outputs {
			if _, ok := modules[output]; ok && modules[output].module_type == "&" {
				modules[output].memory[module.name] = "low"
			}
		}

		if slices.Contains(module.outputs, "rx") {
			feed = module.name
		}
	}

	cycle_lengths := make(map[string]int)
	presses := 0
	for len(cycle_lengths) <= 3 {
		presses += 1
		p_queue := queue{}

		for _, target := range broadcast_targets {
			p_queue.Enqueue(queue_item{
				origin: "broadcaster",
				target: target,
				pulse:  "low",
			})
		}

		for len(p_queue.items) > 0 {
			element := p_queue.Dequeue()

			if _, ok := modules[element.target]; !ok {
				continue
			}

			module := modules[element.target]

			if module.name == feed && element.pulse == "high" {
				if _, ok := cycle_lengths[element.origin]; !ok {
					cycle_lengths[element.origin] = presses
				}
			}

			if module.module_type == "%" {
				if element.pulse == "low" {
					if module.memory["self"] == "on" {
						module.memory["self"] = "off"
					} else {
						module.memory["self"] = "on"
					}

					var outgoing string
					if module.memory["self"] == "on" {
						outgoing = "high"
					} else {
						outgoing = "low"
					}

					for _, output := range module.outputs {
						p_queue.Enqueue(queue_item{
							origin: module.name,
							target: output,
							pulse:  outgoing,
						})
					}
				}
			} else {
				module.memory[element.origin] = element.pulse
				outgoing := "low"
				for _, memory := range module.memory {
					if memory != "high" {
						outgoing = "high"
					}
				}
				for _, x := range module.outputs {
					p_queue.Enqueue(queue_item{
						origin: module.name,
						target: x,
						pulse:  outgoing,
					})
				}
			}

		}
	}

	var cycles []int
	for _, cycle := range cycle_lengths {
		cycles = append(cycles, cycle)
	}

	return LCM(cycles)
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

	c.Visit("https://adventofcode.com/2023/day/20/input")
}
