package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Coordinate struct {
	Row int
	Col int
}

type TupleSet map[Coordinate]struct{}

func main() {
	file, _ := os.Open("7.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	first := scanner.Text()

	startingBeamCol := strings.Index(first, "S")
	beams := []Coordinate{
		{0, startingBeamCol},
	}

	board := make([]string, 0)

	for scanner.Scan() {
		line := scanner.Text()
		board = append(board, line)
	}

	heightBoard := len(board)
	splitCount := 0
	visited := make(TupleSet)

	for len(beams) > 0 {
		last := len(beams) - 1
		b := beams[last]
		beams = beams[:last]
		if _, ok := visited[b]; ok {
			continue
		}
		visited[b] = struct{}{}

		nextRow := b.Row + 1
		if nextRow < heightBoard {
			if board[nextRow][b.Col] == '^' {
				if b.Col+1 >= 0 && b.Col+1 < len(board[nextRow]) {
					beams = append(beams, Coordinate{nextRow, b.Col + 1})
				}
				if b.Col-1 >= 0 && b.Col-1 < len(board[nextRow]) {
					beams = append(beams, Coordinate{nextRow, b.Col - 1})
				}
				split := Coordinate{
					nextRow,
					b.Col,
				}
				if _, exists := visited[split]; !exists {
					splitCount++
				}
			} else {
				beams = append(beams, Coordinate{
					nextRow,
					b.Col,
				})
			}
		}
	}
	fmt.Println(splitCount)
}
