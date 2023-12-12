package year2023

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func CalculateGroups(records string) []int {
	groups := []int{}

	for _, g := range strings.Split(records, ".") {
		if len(g) > 0 && g[0] == '#' {
			groups = append(groups, len(g))
		}
	}

	return groups
}

func CalculateArrangementCount(records string, groups []int) int {
	unknownIndex := strings.Index(records, "?")
	if unknownIndex == -1 {
		recordsGroups := CalculateGroups(records)
		if slices.Equal(recordsGroups, groups) {
			return 1
		}

		return 0
	}

	operationalRecords := strings.Replace(records, "?", ".", 1)
	count := CalculateArrangementCount(operationalRecords, groups)

	damagedRecords := strings.Replace(records, "?", "#", 1)
	count += CalculateArrangementCount(damagedRecords, groups)

	return count
}

func Day12() {
	file, err := os.Open("./inputs/2023/day12.input")
	if err != nil {
		panic(err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	totalArrangementCount := 0
	for scanner.Scan() {
		line := scanner.Text()

		lineParts := strings.Split(line, " ")
		if len(lineParts) != 2 {
			panic("Invalid input")
		}

		records := lineParts[0]
		groups := []int{}

		for _, groupString := range strings.Split(lineParts[1], ",") {
			group, err := strconv.Atoi(groupString)
			if err != nil {
				panic(err)
			}

			groups = append(groups, group)
		}

		arrangementCount := CalculateArrangementCount(records, groups)
		totalArrangementCount += arrangementCount
	}

	fmt.Println("Sum of all arrangement counts:", totalArrangementCount)
}
