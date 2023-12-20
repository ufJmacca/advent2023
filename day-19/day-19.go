package main

import (
	"strings"

	"github.com/PaesslerAG/gval"
)

func PartEvaluation(part map[string]int, workflow []string, workflows map[string][]string) int {

	if len(workflow) == 0 {
		workflow = workflows["in"]
	}

	if workflow[0] == "A" {
		sum := 0
		for _, value := range part {
			sum += value
		}
		return sum
	}

	if workflow[0] == "R" {
		return 0
	}

	for _, flow := range workflow {
		criteria := strings.Split(flow, ":")
		if len(criteria) > 1 {

			value, _ := gval.Evaluate(criteria[0], part)

			if value == true {
				if strings.Contains("AR", criteria[1]) {
					return PartEvaluation(part, []string{criteria[1]}, workflows)
				}
				return PartEvaluation(part, workflows[criteria[1]], workflows)
			} else {
				return PartEvaluation(part, workflow[1:], workflows)
			}
		} else {
			return PartEvaluation(part, workflows[criteria[0]], workflows)
		}
	}

	return 0
}

func Puzzle1(input string) int {

	return 0
}
