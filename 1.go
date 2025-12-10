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
			current = (current - num) % 100
		case 'R':
			current = (current + num) % 100
		default:
			log.Fatalf("invalid first character %q in line %q", dir, line)
		}
		if current == 0 {
			password++
		}
	}
	fmt.Println(password)

}
