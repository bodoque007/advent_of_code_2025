package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func getMaxFromIndex(s string, i int) byte {
	if i >= len(s)-1 {
		return 0
	}

	start := i + 1
	maxVal := s[start]

	for k := start + 1; k < len(s); k++ {
		if s[k] > maxVal {
			maxVal = s[k]
		}
	}

	return maxVal
}

func maxTwoWithPosition(s string) (max1, max2 byte, pos1, pos2 int) {
	max1, max2 = '0'-1, '0'-1
	pos1, pos2 = -1, -1
	for i := 0; i < len(s); i++ {
		n := s[i]
		if n > max1 {
			max2 = max1
			pos2 = pos1

			pos1 = i
			max1 = n

		} else if n > max2 {
			max2 = n
			pos2 = i
		}
	}
	return max1, max2, pos1, pos2
}

func main() {
	count := 0
	file, _ := os.Open("3.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		result := ""
		currentPos := 0
		for i := 0; i < 12; i++ {
			digitsNeededAfter := 12 - i - 1
			searchBound := len(line) - digitsNeededAfter
			currMax := byte(0)

			for j := currentPos; j < searchBound; j++ {
				if currMax < line[j] {
					currMax = line[j]
					currentPos = j
				}
			}
			result += string(line[currentPos])
			currentPos++

		}
		numRes, _ := strconv.Atoi(result)
		fmt.Println(numRes)
		count += numRes
	}
	fmt.Println(count)
}
