package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	grid := make([][]byte, 0)
	accessible := 0
	file, _ := os.Open("4.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Bytes()
		row := append([]byte(nil), line...)
		grid = append(grid, row)
	}

	height := len(grid)
	width := len(grid[0])
	d := [8][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	for r := 0; r < height; r++ {
		for c := 0; c < width; c++ {
			cur := grid[r][c]
			if cur != '@' {
				continue
			}
			adjacentRolls := 0

			for _, direction := range d {
				i, j := r+direction[0], c+direction[1]
				if i < 0 || i >= height || j < 0 || j >= width {
					continue
				}

				if grid[i][j] == '@' {
					adjacentRolls++
				}
			}

			if adjacentRolls < 4 {
				accessible++
			}
		}
	}
	fmt.Println(accessible)
}
