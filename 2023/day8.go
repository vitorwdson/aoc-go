package year2023

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Day8() {
	file, err := os.Open("./inputs/2023/day8.input")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	if !scanner.Scan() {
		panic("Invalid input")
	}
	instructions := scanner.Text()

	network := map[string][2]string{}
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		parts := strings.Split(line, " = ")
		if len(parts) != 2 {
			panic("Invalid input")
		}

		node := parts[0]
		elementsString := parts[1]
		elements := strings.Split(elementsString[1:len(elementsString)-1], ", ")
		if len(parts) != 2 {
			panic("Invalid input")
		}

		network[node] = ([2]string)(elements)
	}

	steps := 0
	found := false

	currNode := "AAA"
	for !found {
		for _, c := range instructions {
			steps++

			var direction int
			if c == 'L' {
				direction = 0
			} else if c == 'R' {
				direction = 1
			} else {
				panic("Invalid input")
			}

			newNode := network[currNode][direction]
			if newNode == "ZZZ" {
				found = true
				break
			}

			currNode = newNode
		}
	}

	fmt.Println("Steps required:", steps)
}
