package year2022

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func Day1() {
	file, err := os.Open("./inputs/2022/day1.input")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	totalPerElf := 0
	topElves := []int{0, 0, 0}

	for scanner.Scan() {
		line := scanner.Text()
		calories, err := strconv.Atoi(line)
		if err == nil {
			totalPerElf += calories
		} else {
			topElves = append(topElves, totalPerElf)
			sort.Ints(topElves)
			topElves = topElves[1:]
			totalPerElf = 0
		}
	}

	topElvesSum := topElves[0] + topElves[1] + topElves[2]
	fmt.Printf("The elf carrying the most calories is carrying %d calories\n", topElves[2])
	fmt.Printf("The top elves are carrying %d calories combined\n", topElvesSum)
}
