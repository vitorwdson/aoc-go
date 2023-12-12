package year2023

import (
	"bufio"
	"fmt"
	"os"
)

func Day11() {
	file, err := os.Open("./inputs/2023/day11.input")
	if err != nil {
		panic(err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	galaxies := [][2]int{}
	y := 0

	for scanner.Scan() {
		line := scanner.Text()

		hasGalaxy := false
		for x, c := range line {
			if c == '#' {
				galaxies = append(galaxies, [2]int{x, y})
				hasGalaxy = true
			}
		}

		if !hasGalaxy {
			y += 999999
		}
		y++
	}

	existingColumns := map[int]bool{}
	maxColumn := -1

	for _, galaxy := range galaxies {
		gx := galaxy[0]
		if gx > maxColumn {
			maxColumn = gx
		}

		existingColumns[gx] = true
	}

	colOffset := 0
	for i := 0; i <= maxColumn; i++ {
		_, exists := existingColumns[i]
		if exists {
			continue
		}

		for j := 0; j < len(galaxies); j++ {
			gx := galaxies[j][0]

			if gx >= i+colOffset {
				galaxies[j][0] += 999999
			}
		}

		colOffset += 999999
	}

	sumOfLengths := 0
	for i := 0; i < len(galaxies); i++ {
		g1 := galaxies[i]

		for j := i + 1; j < len(galaxies); j++ {
			g2 := galaxies[j]

			xLength := g2[0] - g1[0]
			if xLength < 0 {
				xLength *= -1
			}

			yLength := g2[1] - g1[1]
			if yLength < 0 {
				yLength *= -1
			}

			sumOfLengths += xLength + yLength
		}
	}

	fmt.Println("Sum of the shortest pair lengths:", sumOfLengths)
}
