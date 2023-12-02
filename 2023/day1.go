package year2023

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var numbers = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
	"1":     "1",
	"2":     "2",
	"3":     "3",
	"4":     "4",
	"5":     "5",
	"6":     "6",
	"7":     "7",
	"8":     "8",
	"9":     "9",
	"0":     "0",
}

type FoundDigit struct {
	digit string
	index int
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

		var digits []*FoundDigit
		var firstDigit *FoundDigit
		var lastDigit *FoundDigit

		for k := range numbers {
			offset := 0
			for {
				i := strings.Index(line[offset:], k)
				if i == -1 {
					break
				}

				digits = append(digits, &FoundDigit{
					digit: k,
					index: offset + i,
				})
				offset += i + 1
			}
		}

		if len(digits) == 0 {
			continue
		}

		for _, digit := range digits {
			if firstDigit == nil || digit.index < firstDigit.index {
				firstDigit = digit
			}
			if lastDigit == nil || digit.index > lastDigit.index {
				lastDigit = digit
			}
		}

		calibrationValueString := numbers[firstDigit.digit] + numbers[lastDigit.digit]
		calibrationValue, err := strconv.Atoi(calibrationValueString)
		if err != nil {
			panic(err)
		}

		total += calibrationValue
	}

	fmt.Println(total)
}
