package year2023

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var pipeConnections = map[string][2][2]int{
	"|": {{0, -1}, {0, 1}},
	"-": {{-1, 0}, {1, 0}},
	"L": {{0, -1}, {1, 0}},
	"J": {{0, -1}, {-1, 0}},
	"7": {{-1, 0}, {0, 1}},
	"F": {{1, 0}, {0, 1}},
}

type Pipe struct {
	char string
	x    int
	y    int
	next *Pipe
	prev *Pipe
	dist int
}

func (p *Pipe) FindConnections(maze [][]string) [2][2]int {
	connections := pipeConnections[p.char]

	return [2][2]int{
		{p.x + connections[0][0], p.y + connections[0][1]},
		{p.x + connections[1][0], p.y + connections[1][1]},
	}
}

func (p *Pipe) FindDifferentConnection(maze [][]string, cmp *Pipe) [2]int {
	connections := p.FindConnections(maze)

	for _, connection := range connections {
		connX := connection[0]
		connY := connection[1]

		if connX != cmp.x || connY != cmp.y {
			return connection
		}
	}

	fmt.Println(connections, cmp.x, cmp.y)

	panic("No connection")
}

func (p *Pipe) PrevConnection(maze [][]string) [2]int {
	return p.FindDifferentConnection(maze, p.next)
}

func (p *Pipe) NextConnection(maze [][]string) [2]int {
	return p.FindDifferentConnection(maze, p.prev)
}

func FindStartChar(
	maze [][]string,
	startX int,
	startY int,
) (string, *Pipe, *Pipe) {
	connectingPositions := [][2]int{}
	var positionsToCheck = [8][2]int{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, 1},
		{1, 1},
		{1, 0},
		{1, -1},
		{0, -1},
	}

	for _, pos := range positionsToCheck {
		x := startX + pos[0]
		y := startY + pos[1]

		if y < 0 || y >= len(maze) {
			continue
		}

		if x < 0 || x >= len(maze[y]) {
			continue
		}

		char := maze[y][x]

		connections, exists := pipeConnections[char]
		if !exists {
			continue
		}

		for _, connection := range connections {
			connX := x + connection[0]
			connY := y + connection[1]

			if connX == startX && connY == startY {
				connectingPositions = append(connectingPositions, [2]int{x, y})
			}
		}
	}

	if len(connectingPositions) != 2 {
		panic("Invalid input")
	}

	startChar := ""
	for char, connections := range pipeConnections {
		match := true

		for _, connection := range connections {
			connX := startX + connection[0]
			connY := startY + connection[1]

			innerMatch := false

			for _, pos := range connectingPositions {
				posX := pos[0]
				posY := pos[1]

				if posX == connX && posY == connY {
					innerMatch = true
					break
				}
			}

			match = match && innerMatch
		}

		if match {
			startChar = char
			break
		}
	}

	nextX := connectingPositions[0][0]
	nextY := connectingPositions[0][1]
	next := &Pipe{
		char: maze[nextY][nextX],
		x:    nextX,
		y:    nextY,
		next: nil,
		prev: nil,
		dist: 1,
	}

	prevX := connectingPositions[1][0]
	prevY := connectingPositions[1][1]
	prev := &Pipe{
		char: maze[prevY][prevX],
		x:    prevX,
		y:    prevY,
		next: nil,
		prev: nil,
		dist: 1,
	}

	return startChar, next, prev
}

func FindPipe(head *Pipe, x int, y int) *Pipe {
	if head.x == x && head.y == y {
		return head
	}

	curr := head.next
	for curr != head {
		if curr.x == x && curr.y == y {
			return curr
		}

		curr = curr.next
	}

	return nil
}

func Day10() {
	file, err := os.Open("./inputs/2023/day10.input")
	if err != nil {
		panic(err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	startX := -1
	startY := 0

	maze := [][]string{}
	for scanner.Scan() {
		line := scanner.Text()
		pipes := strings.Split(line, "")
		maze = append(maze, pipes)

		if startX == -1 {
			startX = strings.Index(line, "S")
		}
		if startX == -1 {
			startY++
		}
	}

	startChar, nextPipe, prevPipe := FindStartChar(maze, startX, startY)
	start := &Pipe{
		char: startChar,
		x:    startX,
		y:    startY,
		next: nextPipe,
		prev: prevPipe,
		dist: 0,
	}

	start.next.prev = start
	start.prev.next = start

	ended := false
	next := start.next
	prev := start.prev
	dist := 1

	for !ended {
		dist++

		nextConnection := next.NextConnection(maze)
		nextChar := maze[nextConnection[1]][nextConnection[0]]

		prevConnection := prev.PrevConnection(maze)
		prevChar := maze[prevConnection[1]][prevConnection[0]]

		sameX := nextConnection[0] == prevConnection[0]
		sameY := nextConnection[1] == prevConnection[1]
		sameChar := nextChar == prevChar

		if sameX && sameY && sameChar {
			endPipe := &Pipe{
				char: nextChar,
				x:    nextConnection[0],
				y:    nextConnection[1],
				next: prev,
				prev: next,
				dist: dist,
			}
			next.next = endPipe
			prev.prev = endPipe

			ended = true
			continue
		}

		next.next = &Pipe{
			char: nextChar,
			x:    nextConnection[0],
			y:    nextConnection[1],
			next: nil,
			prev: next,
			dist: dist,
		}
		prev.prev = &Pipe{
			char: prevChar,
			x:    prevConnection[0],
			y:    prevConnection[1],
			next: prev,
			prev: nil,
			dist: dist,
		}

		next = next.next
		prev = prev.prev
	}

	fmt.Println("Steps from start to farthest position:", dist)

	enclosedTiles := 0
	for y := 0; y < len(maze); y++ {
		vertical := 0
		lastDelta := 0

		for x := 0; x < len(maze[y]); x++ {
			pipe := FindPipe(start, x, y)

			if pipe != nil {
				delta := 0
				if pipe.y > pipe.prev.y || pipe.y < pipe.next.y {
					delta = 1
				} else if pipe.y < pipe.prev.y || pipe.y > pipe.next.y {
					delta = -1
				}

				if delta != 0 && delta != lastDelta {
					vertical++
					lastDelta = delta
				}
			} else if vertical % 2 == 1 {
				enclosedTiles += 1
			}
		}
	}

	fmt.Println("Tiles enclosed by the loop:", enclosedTiles)
}
