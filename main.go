package main

import (
	"fmt"
	"github.com/vitorwdson/aoc-go/2022"
	"os"
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
		},
	}

	program := programMap[year][day]
	if program == nil {
		fmt.Println("The specified year-day pair does not exist")
		os.Exit(1)
	}

	program()
}
