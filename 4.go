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

	for r := 0; r < height; r++ {
		for c := 0; c < width; c++ {
			cur := grid[r][c]
			if cur != '@' {
				continue
			}
			adjacentRolls := 0
			for i := r - 1; i <= r+1; i++ {
				for j := c - 1; j <= c+1; j++ {
					if (i == r && j == c) || i < 0 || i >= height || j < 0 || j >= width {
						continue
					}
					if grid[i][j] == '@' {
						adjacentRolls++
					}
					if adjacentRolls >= 4 {
						break
					}
				}
			}
			if adjacentRolls < 4 {
				accessible++
			}
		}
	}
	fmt.Println(accessible)
}
