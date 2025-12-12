package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func sumSlice(numbers []int) int {
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	return sum
}

func main() {
	file, _ := os.Open("6.txt")
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	// We store the lines first, because we'll just accumulate each result of a column in a slice, instead of having to save the entire table in memory.
	// To accumulate properly, we need to know the operands first in advance before scanning each table.
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Reads operands and stores them in a slice.
	opLine := lines[len(lines)-1]
	fields := strings.Fields(opLine)

	ops := make([]rune, len(fields))
	for i, f := range fields {
		ops[i] = rune(f[0])
	}

	numericLines := lines[:len(lines)-1]
	firstFields := strings.Fields(numericLines[0])
	numCols := len(firstFields)

	vals := make([]int, numCols)

	// Store initial values on the vals slice. vals[i] = int in the ith column of the first row.
	for i, f := range firstFields {
		n, _ := strconv.Atoi(f)
		vals[i] = n
	}

	// Accumulate
	for _, line := range numericLines[1:] {
		fields := strings.Fields(line)
		for i, f := range fields {
			n, _ := strconv.Atoi(f)
			switch ops[i] {
			case '*':
				vals[i] *= n
			case '+':
				vals[i] += n
			}
		}
	}
	fmt.Println(sumSlice(vals))
}
