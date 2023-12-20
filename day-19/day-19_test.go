package main

import "testing"

type puzzle_input_1 struct {
	input  string
	result int
}

type part_evaluation struct {
	part      map[string]int
	workflows map[string][]string
	result    int
}

func TestPuzzle1(t *testing.T) {
	test_data := []puzzle_input_1{
		{`px{a<2006:qkq,m>2090:A,rfg}
pv{a>1716:R,A}
lnx{m>1548:A,A}
rfg{s<537:gd,x>2440:R,A}
qs{s>3448:A,lnx}
qkq{x<1416:A,crn}
crn{x>2662:A,R}
in{s<1351:px,qqz}
qqz{s>2770:qs,m<1801:hdj,R}
gd{a>3333:R,R}
hdj{m>838:A,pv}

{x=787,m=2655,a=1222,s=2876}
{x=1679,m=44,a=2067,s=496}
{x=2036,m=264,a=79,s=2244}
{x=2461,m=1339,a=466,s=291}
{x=2127,m=1623,a=2188,s=1013}`, 19114},
	}

	for _, datum := range test_data {
		result := Puzzle1(datum.input)

		if result != datum.result {
			t.Errorf("Puzzle1(%s) FAILED - Expected %d Got %d\n", datum.input, datum.result, result)
		} else {
			t.Logf("Puzzle1(%s) PASSED", datum.input)
		}
	}
}

func TestPartEvaluation(t *testing.T) {
	test_data := []part_evaluation{
		{map[string]int{"x": 787, "m": 2655, "a": 1222, "s": 2876}, map[string][]string{
			"px":  {"a<2006:qkq", "m>2090:A", "rfg"},
			"pv":  {"a>1716:R", "A"},
			"lnx": {"m>1548:A", "A"},
			"rfg": {"s<537:gd", "x>2440:R", "A"},
			"qs":  {"s>3448:A", "lnx"},
			"qkq": {"x<1416:A", "crn"},
			"crn": {"x>2662:A", "R"},
			"in":  {"s<1351:px", "qqz"},
			"qqz": {"s>2770:qs", "m<1801:hdj", "R"},
			"gd":  {"a>3333:R", "R"},
			"hdj": {"m>838:A", "pv"},
		}, 7540},
		{map[string]int{"x": 1679, "m": 44, "a": 2067, "s": 496}, map[string][]string{
			"px":  {"a<2006:qkq", "m>2090:A", "rfg"},
			"pv":  {"a>1716:R", "A"},
			"lnx": {"m>1548:A", "A"},
			"rfg": {"s<537:gd", "x>2440:R", "A"},
			"qs":  {"s>3448:A", "lnx"},
			"qkq": {"x<1416:A", "crn"},
			"crn": {"x>2662:A", "R"},
			"in":  {"s<1351:px", "qqz"},
			"qqz": {"s>2770:qs", "m<1801:hdj", "R"},
			"gd":  {"a>3333:R", "R"},
			"hdj": {"m>838:A", "pv"},
		}, 0},
		{map[string]int{"x": 2036, "m": 264, "a": 79, "s": 2244}, map[string][]string{
			"px":  {"a<2006:qkq", "m>2090:A", "rfg"},
			"pv":  {"a>1716:R", "A"},
			"lnx": {"m>1548:A", "A"},
			"rfg": {"s<537:gd", "x>2440:R", "A"},
			"qs":  {"s>3448:A", "lnx"},
			"qkq": {"x<1416:A", "crn"},
			"crn": {"x>2662:A", "R"},
			"in":  {"s<1351:px", "qqz"},
			"qqz": {"s>2770:qs", "m<1801:hdj", "R"},
			"gd":  {"a>3333:R", "R"},
			"hdj": {"m>838:A", "pv"},
		}, 4623},
		{map[string]int{"x": 2461, "m": 1339, "a": 466, "s": 291}, map[string][]string{
			"px":  {"a<2006:qkq", "m>2090:A", "rfg"},
			"pv":  {"a>1716:R", "A"},
			"lnx": {"m>1548:A", "A"},
			"rfg": {"s<537:gd", "x>2440:R", "A"},
			"qs":  {"s>3448:A", "lnx"},
			"qkq": {"x<1416:A", "crn"},
			"crn": {"x>2662:A", "R"},
			"in":  {"s<1351:px", "qqz"},
			"qqz": {"s>2770:qs", "m<1801:hdj", "R"},
			"gd":  {"a>3333:R", "R"},
			"hdj": {"m>838:A", "pv"},
		}, 0},
		{map[string]int{"x": 2127, "m": 1623, "a": 2188, "s": 1013}, map[string][]string{
			"px":  {"a<2006:qkq", "m>2090:A", "rfg"},
			"pv":  {"a>1716:R", "A"},
			"lnx": {"m>1548:A", "A"},
			"rfg": {"s<537:gd", "x>2440:R", "A"},
			"qs":  {"s>3448:A", "lnx"},
			"qkq": {"x<1416:A", "crn"},
			"crn": {"x>2662:A", "R"},
			"in":  {"s<1351:px", "qqz"},
			"qqz": {"s>2770:qs", "m<1801:hdj", "R"},
			"gd":  {"a>3333:R", "R"},
			"hdj": {"m>838:A", "pv"},
		}, 6951},
	}

	for _, datum := range test_data {
		result := PartEvaluation(datum.part, []string{}, datum.workflows)

		if result != datum.result {
			t.Errorf("InstructionDecode(%v, %v) FAILED - Expected %d Got %d\n", datum.part, datum.workflows, datum.result, result)
		} else {
			t.Logf("InstructionDecode(%v, %v) PASSED", datum.part, datum.workflows)
		}
	}
}
