package year2024

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Day1() {
	f, err := os.Open("./inputs/2024/day1.input")
	if err != nil {
		return
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	firstCol := []int{}
	secondCol := []int{}

	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Split(line, "   ")
		if len(values) != 2 {
			panic("invalid length")
		}

		v1, err := strconv.Atoi(values[0])
		if err != nil {
			log.Fatalf("invalid integer %s", values[0])
		}

		v2, err := strconv.Atoi(values[1])
		if err != nil {
			log.Fatalf("invalid integer %s", values[1])
		}

		firstCol = append(firstCol, v1)
		secondCol = append(secondCol, v2)
	}

	slices.Sort(firstCol)
	slices.Sort(secondCol)

	if len(firstCol) != len(secondCol) {
		log.Fatalf("invalid list lengths %d - %d", len(firstCol), len(secondCol))
	}

	repetitions := make(map[int]int)
	var sum int

	for i := 0; i < len(firstCol); i++ {
		l := firstCol[i]
		r := secondCol[i]

		d := r - l
		if d < 0 {
			d = d * -1
		}

		sum += d

		repetitions[r]++
	}

	var similarity int
	for _, l := range firstCol {
		similarity += l * repetitions[l]
	}

	fmt.Println("Part1: ", sum)
	fmt.Println("Part2: ", similarity)
}
