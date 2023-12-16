package year2023

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func IsArrangementValid(records string, groups []int) bool {
	damaged := strings.Count(records, "#")

	if damaged > 0 && len(groups) == 0 {
		return false
	}

	damagedGroups := strings.FieldsFunc(records, func(r rune) bool {
		return r != '#'
	})
	if len(damagedGroups) != len(groups) {
		return false
	}

	for i, g := range damagedGroups {
		if len(g) != groups[i] {
			return false
		}
	}

	return true
}

func IgnoreValidGroup(records string, group int) string {
	if len(records) < group {
		return ""
	}

	for i := 0; i < group; i++ {
		if records[i] == '.' {
			return ""
		}
	}

	if len(records) == group {
		return ""
	}

	if records[group] == '#' {
		return ""
	}

	return records[group+1:]
}

var cache = map[string]int{}

func CalculateArrangementCount(records string, groups []int) int {
	key := records + fmt.Sprint(groups)
	if count, exists := cache[key]; exists {
		return count
	}

	cache[key] = 0
	unknown := strings.Count(records, "?")
	valid := IsArrangementValid(records, groups)

	if unknown == 0 {
		if valid {
			cache[key] = 1
			return 1
		}

		return 0
	}

	if records[0] == '.' {
		cache[key] = CalculateArrangementCount(records[1:], groups)
		return cache[key]
	}

	count := 0
	if records[0] == '?' {
		count += CalculateArrangementCount(records[1:], groups)
	}

	if len(groups) > 0 {
		newRecords := IgnoreValidGroup(records, groups[0])
		if newRecords != "" {
			count += CalculateArrangementCount(newRecords, groups[1:])
		}
	}

	cache[key] = count
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
	totalUnfoldedArrangementCount := 0
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
		fmt.Print("-----", arrangementCount, "\n\n")

		// for i := 0; i < 5; i++ {
		// 	records += "?" + records
		// 	groups = append(groups, groups...)
		// }
		//
		// unfoldedUrrangementCount := CalculateArrangementCount(records, groups)
		// totalUnfoldedArrangementCount += unfoldedUrrangementCount
	}

	fmt.Println("Sum of all arrangement counts:", totalArrangementCount)
	fmt.Println("Sum of all unfolded arrangement counts:", totalUnfoldedArrangementCount)
}
