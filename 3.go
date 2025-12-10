package main

import (
	"bufio"
	"fmt"
	"os"
)
// TODO: Comment properly. Approach is simple, though. Keep track of the two maximums and of which is the leftmost one and which is the rightmost one in the input line and separate in cases:
// First max > second max: return firstmax + secondmax (+ is concatenation)
// Second max > first max: if secondmax is last element, return firstmax + secondmax. If not, return secondmax + whatever is max to the right side of the secondmax
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

	toInt := func(b byte) int {
		return int(b - '0')
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		max1, max2, pos1, pos2 := maxTwoWithPosition(line)
		if pos1 > pos2 {
			max1, max2 = max2, max1
			pos1, pos2 = pos2, pos1
		}

		if max2 > max1 {
			if pos2 == len(line)-1 {
				count += (toInt(max1) * 10) + toInt(max2) // Concatenates, we know max1 is a sole digit.
			} else {
				nextVal := getMaxFromIndex(line, pos2)
				count += (toInt(max2) * 10) + toInt(nextVal)
			}
		} else {
			count += (toInt(max1) * 10) + toInt(max2)
		}
	}
	fmt.Println(count)
}
