package main

import "testing"

type almanac struct {
	input  string
	result int
}

type source_to_destination_map struct {
	range_map []range_map
	input     int
	result    int
}

func TestLowestLocation(t *testing.T) {
	test_data := []almanac{
		{`seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4`, 35},
	}

	for _, datum := range test_data {
		result := LowestLocation(datum.input)

		if result != datum.result {
			t.Errorf("LowestLocation(%s) FAILED - Expected %d Got %d\n", datum.input, datum.result, result)
		} else {
			t.Logf("LowestLocation(%s) PASSED", datum.input)
		}
	}
}

func TestLowestLocationRange(t *testing.T) {
	test_data := []almanac{
		{`seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4`, 46},
	}

	for _, datum := range test_data {
		result := LowestLocationRange(datum.input)

		if result != datum.result {
			t.Errorf("LowestLocationRange(%s) FAILED - Expected %d Got %d\n", datum.input, datum.result, result)
		} else {
			t.Logf("LowestLocationRange(%s) PASSED", datum.input)
		}
	}
}

func TestSourceToDestination(t *testing.T) {
	test_data := []source_to_destination_map{
		{
			range_map: []range_map{
				{destination_range_start: 50, source_range_start: 98, range_length: 2},
				{destination_range_start: 52, source_range_start: 50, range_length: 48},
			},
			input:  79,
			result: 81,
		}, {
			range_map: []range_map{
				{destination_range_start: 50, source_range_start: 98, range_length: 2},
				{destination_range_start: 52, source_range_start: 50, range_length: 48},
			},
			input:  14,
			result: 14,
		}, {
			range_map: []range_map{
				{destination_range_start: 50, source_range_start: 98, range_length: 2},
				{destination_range_start: 52, source_range_start: 50, range_length: 48},
			},
			input:  55,
			result: 57,
		}, {
			range_map: []range_map{
				{destination_range_start: 50, source_range_start: 98, range_length: 2},
				{destination_range_start: 52, source_range_start: 50, range_length: 48},
			},
			input:  13,
			result: 13,
		},
	}

	for _, datum := range test_data {
		result := SourceToDestination(datum.range_map, datum.input)

		if result != datum.result {
			t.Errorf("SourceToDestination(%d) FAILED - Expected %d Got %d\n", datum.input, datum.result, result)
		} else {
			t.Logf("SourceToDestination(%d) PASSED", datum.input)
		}
	}
}
