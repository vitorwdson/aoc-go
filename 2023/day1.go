package year2023

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func isDigit(c rune) bool {
	return c >= '0' && c <= '9'
}

func Day1() {
	file, err := os.Open("./inputs/2023/day1.input")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	total := 0
	for scanner.Scan() {
		line := scanner.Text()

		var firstDigit rune
		var lastDigit rune

		for _, c := range line {
			if isDigit(c) {
				firstDigit = c
				break
			}
		}

		for i := range line {
			c := rune(line[len(line)-1-i])

			if isDigit(c) {
				lastDigit = c
				break
			}
		}

		calibrationString := string(firstDigit) + string(lastDigit)
		calibrationValue, err := strconv.Atoi(calibrationString)
		if err != nil {
			panic(err)
		}

		total += calibrationValue
	}

	fmt.Println(total)
}
