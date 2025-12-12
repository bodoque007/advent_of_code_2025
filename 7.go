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
	visitedSplits := make(TupleSet)

	for len(beams) > 0 {
		last := len(beams) - 1
		b := beams[last]
		beams = beams[:last]
		nextRow := b.Row + 1
		if nextRow < heightBoard {
			if board[nextRow][b.Col] == '^' {
				beams = append(beams, Coordinate{
					nextRow,
					b.Col + 1,
				}, Coordinate{
					nextRow,
					b.Col - 1,
				})
				split := Coordinate{
					nextRow,
					b.Col,
				}
				if _, exists := visitedSplits[split]; !exists {
					splitCount++
					visitedSplits[split] = struct{}{}
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
