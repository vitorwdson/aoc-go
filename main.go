package main

import (
	"fmt"
	"os"

	"github.com/vitorwdson/aoc-go/2022"
	year2023 "github.com/vitorwdson/aoc-go/2023"
)

func main() {
	args := os.Args

	if len(args) != 3 {
		fmt.Println("You need to specify the year and day to execute the program")
		os.Exit(1)
	}
		

	year := args[1]
	day := args[2]

	programMap := map[string]map[string]func(){
		"2022": {
			"1": year2022.Day1,
			"2": year2022.Day2,
			"3": year2022.Day3,
		},
		"2023": {
			"1": year2023.Day1,
			"2": year2023.Day2,
			"3": year2023.Day3,
			"4": year2023.Day4,
			"5": year2023.Day5,
			"6": year2023.Day6,
		},
	}

	program := programMap[year][day]
	if program == nil {
		fmt.Println("The specified year-day pair does not exist")
		os.Exit(1)
	}

	program()
}
