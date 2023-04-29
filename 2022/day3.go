package year2022

import (
	"bufio"
	"fmt"
	"os"
)

func Day3() {
	file, err := os.Open("./inputs/2022/day3.input")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	totalItemPriorities := 0
	for scanner.Scan() {
		line := scanner.Text()
		compartmentSize := len(line) / 2

		firstCompartment := line[:compartmentSize]
		secondCompartment := line[compartmentSize:]

		firstCompartmentSet := map[rune]bool{}
		for _, char := range firstCompartment {
			firstCompartmentSet[char] = true
		}

		var sharedItem rune
		for _, char := range secondCompartment {
			if firstCompartmentSet[char] {
				sharedItem = char
				break
			}
		}

		if sharedItem == 0 {
			continue
		}

		var itemPriority int
		if sharedItem <= 90 {
			itemPriority = int(sharedItem) - 65 + 27
		} else {
			itemPriority = int(sharedItem) - 97 + 1
		}

		totalItemPriorities += itemPriority
	}

	fmt.Printf("The sum of the shared item priorities is %d\n", totalItemPriorities)
}
