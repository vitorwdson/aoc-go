package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
    file, err := os.Open("inputs/day1.input")
    if err != nil {
        panic(err)
    }

    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)

    totalPerElf := 0
    max := 0

    for scanner.Scan() {
        line := scanner.Text()
        calories, err := strconv.Atoi(line)
        if err == nil {
            totalPerElf += calories
        } else {
            if max < totalPerElf {
                max = totalPerElf
            }
            totalPerElf = 0
        }
    }

    fmt.Printf("The elf carrying the most calories is carrying %d calories", max)
}
