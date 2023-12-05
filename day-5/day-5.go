package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

type range_map struct {
	destination_range_start int
	source_range_start      int
	range_length            int
}

func LowestLocation(input string) int {
	seeds_string := regexp.MustCompile(`seeds: (.*?)$`)
	seeds_re := regexp.MustCompile(`\d+`)
	map_title := regexp.MustCompile(`(.*?) map\:`)
	map_data := regexp.MustCompile(`(\d+) (\d+) (\d+)`)

	lines := strings.Split(input, "\n")

	maps := make(map[string][]range_map)
	previous_map_title := ""
	var seeds []int

	for _, line := range lines {
		if len(seeds_string.FindAllString(line, -1)) > 0 {
			seeds_matches := seeds_re.FindAllString(line, -1)

			for _, seeds_str := range seeds_matches {
				seed, _ := strconv.Atoi(seeds_str)
				seeds = append(seeds, seed)
			}
		} else if len(map_title.FindAllString(line, -1)) > 0 {
			matches := map_title.FindAllStringSubmatch(line, -1)
			for _, match := range matches {
				previous_map_title = match[1]
			}
		} else if len(map_data.FindAllString(line, -1)) > 0 {
			matches := map_data.FindAllStringSubmatch(line, -1)
			for _, match := range matches {
				destination, _ := strconv.Atoi(match[1])
				source, _ := strconv.Atoi(match[2])
				rang, _ := strconv.Atoi(match[3])
				maps[previous_map_title] = append(maps[previous_map_title], range_map{destination_range_start: destination, source_range_start: source, range_length: rang})
			}
		}
	}

	min_location := 100 * 100 * 100 * 100 * 100

	for _, seed := range seeds {
		seed_to_soil := SourceToDestination(maps["seed-to-soil"], seed)
		soil_to_fertilizer := SourceToDestination(maps["soil-to-fertilizer"], seed_to_soil)
		fertilizer_to_water := SourceToDestination(maps["fertilizer-to-water"], soil_to_fertilizer)
		water_to_light := SourceToDestination(maps["water-to-light"], fertilizer_to_water)
		light_to_temperature := SourceToDestination(maps["light-to-temperature"], water_to_light)
		temperature_to_humidity := SourceToDestination(maps["temperature-to-humidity"], light_to_temperature)
		humidity_to_location := SourceToDestination(maps["humidity-to-location"], temperature_to_humidity)
		if humidity_to_location < min_location {
			min_location = humidity_to_location
		}
	}

	return min_location
}

func LowestLocationRange(input string) int {
	seeds_string := regexp.MustCompile(`seeds: (.*?)$`)
	seeds_re := regexp.MustCompile(`(\d+) (\d+)`)
	map_title := regexp.MustCompile(`(.*?) map\:`)
	map_data := regexp.MustCompile(`(\d+) (\d+) (\d+)`)

	lines := strings.Split(input, "\n")

	maps := make(map[string][]range_map)
	previous_map_title := ""
	var seeds []int

	for _, line := range lines {
		if len(seeds_string.FindAllString(line, -1)) > 0 {
			seeds_matches := seeds_re.FindAllStringSubmatch(line, -1)

			for _, seeds_str := range seeds_matches {
				seed, _ := strconv.Atoi(seeds_str[1])
				seed_range, _ := strconv.Atoi(seeds_str[2])
				for i := 0; i < seed_range; i++ {
					seeds = append(seeds, seed+i)
				}
			}
		} else if len(map_title.FindAllString(line, -1)) > 0 {
			matches := map_title.FindAllStringSubmatch(line, -1)
			for _, match := range matches {
				previous_map_title = match[1]
			}
		} else if len(map_data.FindAllString(line, -1)) > 0 {
			matches := map_data.FindAllStringSubmatch(line, -1)
			for _, match := range matches {
				destination, _ := strconv.Atoi(match[1])
				source, _ := strconv.Atoi(match[2])
				rang, _ := strconv.Atoi(match[3])
				maps[previous_map_title] = append(maps[previous_map_title], range_map{destination_range_start: destination, source_range_start: source, range_length: rang})
			}
		}
	}

	min_location := 100 * 100 * 100 * 100 * 100

	for _, seed := range seeds {
		seed_to_soil := SourceToDestination(maps["seed-to-soil"], seed)
		soil_to_fertilizer := SourceToDestination(maps["soil-to-fertilizer"], seed_to_soil)
		fertilizer_to_water := SourceToDestination(maps["fertilizer-to-water"], soil_to_fertilizer)
		water_to_light := SourceToDestination(maps["water-to-light"], fertilizer_to_water)
		light_to_temperature := SourceToDestination(maps["light-to-temperature"], water_to_light)
		temperature_to_humidity := SourceToDestination(maps["temperature-to-humidity"], light_to_temperature)
		humidity_to_location := SourceToDestination(maps["humidity-to-location"], temperature_to_humidity)
		if humidity_to_location < min_location {
			min_location = humidity_to_location
		}
	}

	return min_location
}

func SourceToDestination(range_map []range_map, input int) int {
	for _, datum := range range_map {
		if datum.source_range_start <= input && input < datum.source_range_start+datum.range_length {
			return datum.destination_range_start + (input - datum.source_range_start)
		}
	}
	return input
}

func main() {
	c := colly.NewCollector()

	// Sets cookie from environment variable
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("cookie", os.Getenv("COOKIE"))
	})

	c.OnResponse(func(r *colly.Response) {
		inputs := string(r.Body)

		min_location := LowestLocation(inputs)
		fmt.Println(min_location)

		min_location_range := LowestLocationRange(inputs)
		fmt.Println(min_location_range)
	})

	c.Visit("https://adventofcode.com/2023/day/5/input")
}
