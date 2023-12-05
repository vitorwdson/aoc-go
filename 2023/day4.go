package year2023

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/vitorwdson/aoc-go/utils"
)

func Day4() {
	file, err := os.Open("./inputs/2023/day4.input")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	totalPoints := 0
	cardCopies := map[int]int{}
	lastCard := 0

	for scanner.Scan() {
		line := scanner.Text() + " "

		colonIndex := strings.Index(line, ":")
		pipeIndex := strings.Index(line, "|")

		numberString := ""
		for _, c := range line[:colonIndex] {
			if !utils.IsDigit(c) {
				continue
			}
			numberString += string(c)
		}

		cardNumber, err := strconv.Atoi(numberString)
		if err != nil {
			panic(err)
		}

		lastCard = cardNumber
		numberString = ""

		winningNumbers := map[int]bool{}
		for _, c := range line[colonIndex+1 : pipeIndex] {
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

			winningNumbers[number] = true
			numberString = ""
		}

		numberString = ""
		matchCount := 0
		for _, c := range line[pipeIndex+1:] {
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

			_, match := winningNumbers[number]
			if match {
				matchCount++
			}

			numberString = ""
		}

		if matchCount != 0 {
			totalPoints += int(math.Pow(2, float64(matchCount-1)))
		}

		cardCopies[cardNumber]++
		for i := 1; i <= matchCount; i++ {
			cardCopies[cardNumber+i] += cardCopies[cardNumber]
		}

	}

	totalCards := 0
	for k, v := range cardCopies {
		if k > lastCard {
			continue
		}

		totalCards += v
	}

	fmt.Println("Total points:", totalPoints)
	fmt.Println("Total scratchcards:", totalCards)
}
