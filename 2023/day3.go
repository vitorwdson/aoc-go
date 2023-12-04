package year2023

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/vitorwdson/aoc-go/utils"
)

type PartNumber struct {
	number int
	line   int
	col    int
	len    int
}

type Symbol struct {
	char rune
	line int
	col  int
}

func Day3() {
	file, err := os.Open("./inputs/2023/day3.input")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	partNumbers := []PartNumber{}
	symbols := []Symbol{}

	lineNumber := 0
	for scanner.Scan() {
		line := scanner.Text() + "."

		numberString := ""
		for col, c := range line {
			if utils.IsDigit(c) {
				numberString += string(c)
				continue
			}

			if numberString != "" {
				number, err := strconv.Atoi(numberString)
				if err != nil {
					panic(err)
				}

				numberLen := len(numberString)
				partNumbers = append(partNumbers, PartNumber{
					number: number,
					line:   lineNumber,
					col:    col - numberLen,
					len:    numberLen,
				})

				numberString = ""
			}

			if c == '.' {
				continue
			}

			symbols = append(symbols, Symbol{
				char: c,
				line: lineNumber,
				col:  col,
			})
		}

		lineNumber++
	}

	totalPartNumber := 0
	for _, partNumber := range partNumbers {
		isValid := false

		for _, symbol := range symbols {
			validLine := symbol.line >= partNumber.line-1 && symbol.line <= partNumber.line+1
			if !validLine {
				continue
			}

			validCol := symbol.col >= partNumber.col-1 && symbol.col <= partNumber.col + partNumber.len
			if !validCol {
				continue
			}

			isValid = true
			break
		}

		if isValid {
			totalPartNumber += partNumber.number
		}
	}

	fmt.Println("Sum of part numbers:", totalPartNumber)
}
