package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	current := 50
	password := 0
	file, _ := os.Open("1.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}

		dir := line[0]
		numStr := line[1:]

		var num int
		fmt.Sscanf(numStr, "%d", &num)
		switch dir {
		case 'L':
			distToZero := current
			if current == 0 {
				distToZero = 100
			}

			if num >= distToZero {
				password += 1 + (num-distToZero)/100
			}

			current = (100 + (current-num)%100) % 100
		case 'R':
			password += (current + num) / 100
			current = (current + num) % 100
		default:
			log.Fatalf("invalid first character %q in line %q", dir, line)
		}
	}
	fmt.Println(password)

}
