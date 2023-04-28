package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("inputs/day2.input")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	shapeScoreMap := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
		"X": 1,
		"Y": 2,
		"Z": 3,
	}

	totalScore := 0
	for scanner.Scan() {
		line := scanner.Text()

		opponentChoice := string(line[0])
		yourChoice := string(line[2])

		opponentShapeScore := shapeScoreMap[opponentChoice]
		yourShapeScore := shapeScoreMap[yourChoice]

		var outcomeScore int
		if yourShapeScore == opponentShapeScore {
			outcomeScore = 3
		} else if opponentShapeScore%3 == yourShapeScore-1 {
			outcomeScore = 6
		} else {
			outcomeScore = 0
		}

		totalScore += yourShapeScore + outcomeScore
	}

	fmt.Printf("Your total score is %d\n", totalScore)
}
