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

	drawShapeScoreMap := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
	}
	winShapeScoreMap := map[string]int{
		"A": 2,
		"B": 3,
		"C": 1,
	}
	lossShapeScoreMap := map[string]int{
		"A": 3,
		"B": 1,
		"C": 2,
	}

	outcomeScoreMap := map[string]int{
		"X": 0,
		"Y": 3,
		"Z": 6,
	}

	totalScore := 0
	for scanner.Scan() {
		line := scanner.Text()

		opponentChoice := string(line[0])
		desiredOutcome := string(line[2])

		var shapeScore int
		switch desiredOutcome {
		case "X":
			shapeScore = lossShapeScoreMap[opponentChoice]
		case "Y":
			shapeScore = drawShapeScoreMap[opponentChoice]
		case "Z":
			shapeScore = winShapeScoreMap[opponentChoice]
		}

		outcomeScore := outcomeScoreMap[desiredOutcome]
		totalScore += shapeScore + outcomeScore
	}

	fmt.Printf("Your total score is %d\n", totalScore)
}
