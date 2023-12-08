package year2023

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

var stampStrength = map[rune]int{
	'A': 13,
	'K': 12,
	'Q': 11,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
	'J': 1,
}

type Hand struct {
	cards        string
	bid          int
	typeStrength int
}

func GetTypeStrength(cards string) int {
	labelCount := map[rune]int{}

	for _, c := range cards {
		labelCount[c]++
	}

	hasThree := false
	hasTwo := false

	for _, v := range labelCount {
		switch v {
		case 5:
			return 7
		case 4:
			return 6
		case 3:
			if hasTwo {
				return 5
			}
			hasThree = true
		case 2:
			if hasThree {
				return 5
			}
			if hasTwo {
				return 3
			}
			hasTwo = true
		case 1:
			continue
		}
	}

	if hasThree {
		return 4
	}

	if hasTwo {
		return 2
	}

	return 1
}

func GetJokerTypeStrength(cards string) int {
	if strings.Index(cards, "J") == -1 {
		return GetTypeStrength(cards)
	}

	strength := 0
	for c := range stampStrength {
		if c == 'J' {
			continue
		}

		tmpCards := strings.ReplaceAll(cards, "J", string(c))
		tmpStrength := GetTypeStrength(tmpCards)

		if tmpStrength > strength {
			strength = tmpStrength
		}
	}

	return strength
}

func CompareHands(hand1, hand2 *Hand) int {
	if hand1.typeStrength > hand2.typeStrength {
		return 1
	}
	if hand1.typeStrength < hand2.typeStrength {
		return -1
	}

	for i := 0; i < len(hand1.cards); i++ {
		c1 := rune(hand1.cards[i])
		c2 := rune(hand2.cards[i])

		if c1 == c2 {
			continue
		}

		if stampStrength[c1] > stampStrength[c2] {
			return 1
		} else {
			return -1
		}
	}

	return 0
}

func Day7() {
	file, err := os.Open("./inputs/2023/day7.input")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	hands := []*Hand{}
	for scanner.Scan() {
		line := scanner.Text()

		spaceIndex := strings.Index(line, " ")
		cards := line[:spaceIndex]
		bidString := line[spaceIndex+1:]

		bid, err := strconv.Atoi(bidString)
		if err != nil {
			panic(err)
		}

		hands = append(hands, &Hand{
			cards:        cards,
			bid:          bid,
			typeStrength: GetJokerTypeStrength(cards),
		})
	}

	slices.SortFunc(hands, CompareHands)

	totalWinnings := 0
	for i, hand := range hands {
		totalWinnings += (i + 1) * hand.bid
	}

	fmt.Println("Total Winnings:", totalWinnings)
}
