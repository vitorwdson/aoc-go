package year2023

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type MapRange struct {
	SourceStart      int
	DestinationStart int
	Range            int
}
type Map struct {
	Ranges []*MapRange
}
type Almanac = map[string]*Map

func (m *Map) GetDestination(source int) int {
	for _, mapRange := range m.Ranges {
		minSource := mapRange.SourceStart
		maxSource := minSource + mapRange.Range - 1

		if source >= minSource && source <= maxSource {
			offset := source - minSource
			return mapRange.DestinationStart + offset
		}
	}

	return source
}

func GetLocationFromSeed(almanac Almanac, seed int) int {
	soil := almanac["seed-to-soil"].GetDestination(seed)
	fertilizer := almanac["soil-to-fertilizer"].GetDestination(soil)
	water := almanac["fertilizer-to-water"].GetDestination(fertilizer)
	light := almanac["water-to-light"].GetDestination(water)
	temperature := almanac["light-to-temperature"].GetDestination(light)
	humidity := almanac["temperature-to-humidity"].GetDestination(temperature)
	location := almanac["humidity-to-location"].GetDestination(humidity)

	return location
}

func Day5() {
	file, err := os.Open("./inputs/2023/day5.input")
	if err != nil {
		panic(err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	if !scanner.Scan() {
		panic("Empty input")
	}
	firstLine := scanner.Text()

	seeds := []int{}
	for _, numberString := range strings.Split(firstLine[7:], " ") {
		number, err := strconv.Atoi(numberString)
		if err != nil {
			panic(err)
		}

		seeds = append(seeds, number)
	}

	// Ignore blank line
	if !scanner.Scan() {
		panic("Invalid input")
	}
	scanner.Text()

	almanac := Almanac{}
	mapKey := ""

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		if strings.Contains(line, " map:") {
			mapKey = line[:strings.Index(line, " ")]
			almanac[mapKey] = &Map{}

			continue
		}

		// Destination Start, Source Start, Range Length
		mapLineNumbers := [3]int{}
		for i, numberString := range strings.Split(line, " ") {
			number, err := strconv.Atoi(numberString)
			if err != nil {
				panic(err)
			}

			mapLineNumbers[i] = number
		}

		almanac[mapKey].Ranges = append(almanac[mapKey].Ranges, &MapRange{
			DestinationStart: mapLineNumbers[0],
			SourceStart:      mapLineNumbers[1],
			Range:            mapLineNumbers[2],
		})
	}

	lowestLocationPlain := int(math.Inf(1)) - 1
	for _, seed := range seeds {
		location := GetLocationFromSeed(almanac, seed)

		if location < lowestLocationPlain {
			lowestLocationPlain = location
		}
	}
	fmt.Println("Lowest Location (with plain seed list):", lowestLocationPlain)

	lowestLocationRange := int(math.Inf(1)) - 1
	for j := 0; j < len(seeds); j += 2 {
		minSeed := seeds[j]
		maxSeed := minSeed + seeds[j+1]

		for seed := minSeed; seed < maxSeed; seed++ {
			location := GetLocationFromSeed(almanac, seed)

			if location < lowestLocationRange {
				lowestLocationRange = location
			}
		}
	}
	fmt.Println("Lowest Location (with seed ranges):", lowestLocationRange)
}
