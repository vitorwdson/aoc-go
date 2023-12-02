package year2022

import (
	"bufio"
	"fmt"
	"os"
)

func calculateItemPriority(item rune) int {
	var itemPriority int
	if item <= 90 {
		itemPriority = int(item) - 65 + 27
	} else {
		itemPriority = int(item) - 97 + 1
	}

	return itemPriority
}

func Day3() {
	file, err := os.Open("./inputs/2022/day3.input")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	totalItemPriorities := 0
    totalBadgePriorities := 0

    groupItemsSet := map[rune]bool{}
    elfIndexPerGroup := 0
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

		if sharedItem != 0 {
            itemPriority := calculateItemPriority(sharedItem)
			totalItemPriorities += itemPriority
		}


        // Part 2
        elfIndexPerGroup += 1
        
        var groupBadge rune
        if elfIndexPerGroup == 1 {
            for _, char := range line {
                groupItemsSet[char] = false
            }
        } else if elfIndexPerGroup == 2 {
            for _, char := range line {
                if _, exists := groupItemsSet[char]; exists {
                    groupItemsSet[char] = true 
                }
            }
        } else {
            for _, char := range line {
                if groupItemsSet[char] {
                    groupBadge = char
                    break
                }
            }

            elfIndexPerGroup = 0
            groupItemsSet = make(map[rune]bool)
        }

        if groupBadge != 0 {
            badgePriority := calculateItemPriority(groupBadge)
            totalBadgePriorities += badgePriority
        }
	}

	fmt.Printf("The sum of the shared item priorities is %d\n", totalItemPriorities)
	fmt.Printf("The sum of the badge item priorities is %d\n", totalBadgePriorities)
}
