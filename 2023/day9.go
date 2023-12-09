package year2023

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CalculateExtrapolatedDifference(history []int) int {
	differences := []int{}
	allZeros := true

	for i := 1; i < len(history); i++ {
		difference := history[i] - history[i-1]

		differences = append(differences, difference)
		allZeros = allZeros && difference == 0
	}

	if allZeros {
		return 0
	}

	lastValue := differences[len(differences)-1]
	return lastValue + CalculateExtrapolatedDifference(differences)
}

func CalculateExtrapolatedValue(history []int) int {
	if len(history) == 0 {
		return 0
	}

	lastValue := history[len(history)-1]
	return lastValue + CalculateExtrapolatedDifference(history)
}

func CalculateBackwardsExtrapolatedDifference(history []int) int {
	differences := []int{}
	allZeros := true

	for i := len(history) - 1; i > 0; i-- {
		difference := history[i] - history[i-1]

		differences = append(differences, 0)
		copy(differences[1:], differences)
		differences[0] = difference

		allZeros = allZeros && difference == 0
	}

	if allZeros {
		return 0
	}

	firstValue := differences[0]
	return firstValue - CalculateBackwardsExtrapolatedDifference(differences)
}

func CalculateBackwardsExtrapolatedValue(history []int) int {
	if len(history) == 0 {
		return 0
	}

	firstValue := history[0]
	return firstValue - CalculateBackwardsExtrapolatedDifference(history)
}

func Day9() {
	file, err := os.Open("./inputs/2023/day9.input")
	if err != nil {
		panic(err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	totalExtrapolatedValue := 0
	totalBackwardsExtrapolatedValue := 0
	for scanner.Scan() {
		line := scanner.Text()

		history := []int{}
		for _, numberString := range strings.Split(line, " ") {
			number, err := strconv.Atoi(numberString)
			if err != nil {
				panic(err)
			}

			history = append(history, number)
		}

		extrapolatedValue := CalculateExtrapolatedValue(history)
		totalExtrapolatedValue += extrapolatedValue

		backwardsExtrapolatedValue := CalculateBackwardsExtrapolatedValue(history)
		totalBackwardsExtrapolatedValue += backwardsExtrapolatedValue
	}

	fmt.Println("Sum of extrapolated values:", totalExtrapolatedValue)
	fmt.Println("Sum of extrapolated values (backwards):", totalBackwardsExtrapolatedValue)

}
