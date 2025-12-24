package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Helper to calculate the result of a block
func solveBlock(numbers []int, op rune) int {
	if len(numbers) == 0 {
		return 0
	}
	res := numbers[0]
	for _, n := range numbers[1:] {
		switch op {
		case '*':
			res *= n
		case '+':
			res += n
		}
	}
	return res
}

func main() {
	file, _ := os.Open("6.txt")
	defer file.Close()

	var matrix [][]rune
	scanner := bufio.NewScanner(file)

	// We read the lines as they are, we do NOT strip based on empty characters
	for scanner.Scan() {
		matrix = append(matrix, []rune(scanner.Text()))
	}

	grandTotal := 0
	width := len(matrix[0])
	height := len(matrix)

	var currentNums []int
	var currentOp rune

	for col := 0; col < width; col++ {
		isSeparator := true
		colDigits := ""
		var foundOp rune

		for row := 0; row < height; row++ {
			char := matrix[row][col]

			if char != ' ' {
				isSeparator = false

				// Bottom row is always the operator
				if row == height-1 {
					foundOp = char
				} else {
					// Keep track of column digits (top-bottom) to make up the number once we finish the column.
					colDigits += string(char)
				}
			}
		}

		if isSeparator {
			// We solve the operation and restore the counters for the next operation.
			grandTotal += solveBlock(currentNums, currentOp)
			currentNums = []int{}
			currentOp = 0
		} else {
			// We're still inside the operation
			if foundOp != 0 {
				currentOp = foundOp
			}
			if colDigits != "" {
				val, _ := strconv.Atoi(colDigits)
				currentNums = append(currentNums, val)
			}
		}
	}

	// Input does not end in an empty column, so we need to do the very last operation here.
	if len(currentNums) > 0 {
		grandTotal += solveBlock(currentNums, currentOp)
	}

	fmt.Println(grandTotal)
}
