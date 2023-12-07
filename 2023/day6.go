package year2023

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/vitorwdson/aoc-go/utils"
)

func GetLineNumbers(line string) []int {
	colonIndex := strings.Index(line, ":")
	numbersPart := line[colonIndex + 1:] + " "

	numberString := ""
	numbers := []int{}

	for _, c := range numbersPart {
		if utils.IsDigit(c) {
			numberString += string(c)
			continue
		}

		if numberString == "" {
			continue
		}

		number, err := strconv.Atoi(numberString)
		if err != nil {
			panic(err)
		}

		numbers = append(numbers, number)
		numberString = ""
	}

	return numbers
}

func Day6() {
	file, err := os.Open("./inputs/2023/day6.input")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	if !scanner.Scan() {
		panic("Invalid input")
	}
	timeLine := scanner.Text()
	times := GetLineNumbers(timeLine)

	if !scanner.Scan() {
		panic("Invalid input")
	}
	distanceLine := scanner.Text()
	distances := GetLineNumbers(distanceLine)

	if len(times) != len(distances) {
		panic("Invalid input")
	}

	total := 1
	for i := 0; i < len(times); i++ {
		time := times[i]
		distance := distances[i]

		numberOfWays := 0
		for t := 0; t <= time; t++ {
			d := (time-t)*t
			if d > distance {
				numberOfWays++
			}
		}

		total *= numberOfWays
	}

	fmt.Println("Total:", total)
}
