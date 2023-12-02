package year2023

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/vitorwdson/aoc-go/utils"
)

var maxCubes = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func Day2() {
	file, err := os.Open("inputs/2023/day2.input")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	idSum := 0
	for scanner.Scan() {
		line := scanner.Text()

		colonIndex := strings.Index(line, ":")
		idString := line[5:colonIndex]

		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		subsets := line[colonIndex+2:] + ";"

		var amountString string
		var color string
		isValid := true

		for _, c := range subsets {
			if c == ' ' {
				continue
			}

			if c == ';' || c == ',' {
				amount, err := strconv.Atoi(amountString)
				if err != nil {
					panic(err)
				}

				if amount > maxCubes[color] {
					isValid = false
					break
				}

				amountString = ""
				color = ""
				continue
			}

			if utils.IsDigit(c) {
				amountString += string(c)
				color = ""

				continue
			}

			color += string(c)
		}

		if isValid {
			idSum += id
		}

	}

	fmt.Println(idSum)
}
