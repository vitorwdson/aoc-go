package year2023

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func CalculateSteps(network map[string][2]string, instructions string, start string, part2 bool) int {
	steps := 0
	found := false

	currNode := start
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
			if part2 {
				if newNode[len(newNode)-1] == 'Z' {
					found = true
					break
				}
			} else {
				if newNode == "ZZZ" {
					found = true
					break
				}
			}

			currNode = newNode
		}
	}

	return steps
}

func GreatestCommonDivision(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LeastCommonMultiple(values ...int) int {
	if len(values) < 2 {
		return 1
	}

	a := values[0]
	b := values[1]

	result := a * b / GreatestCommonDivision(a, b)

	for _, v := range values[2:] {
		result = LeastCommonMultiple(result, v)
	}

	return result
}

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

	part1Steps := CalculateSteps(network, instructions, "AAA", false)
	fmt.Println("Steps required (AAA to ZZZ):", part1Steps)

	part2StepsList := []int{}
	for node := range network {
		if node[len(node)-1] == 'A' {
			part2StepsList = append(
				part2StepsList,
				CalculateSteps(
					network,
					instructions,
					node,
					true,
				),
			)
		}
	}

	part2Steps := LeastCommonMultiple(part2StepsList...)
	fmt.Println("Steps required (All A's to all Z's):", part2Steps)
}
