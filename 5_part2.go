package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Node struct {
	Start int
	End   int
	Max   int // The maximum 'End' value in this node's entire subtree
	Left  *Node
	Right *Node
}

type Interval struct {
	Start int
	End   int
}

func main() {
	file, _ := os.Open("5.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	intervals := []*Interval{}

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		parts := strings.Split(line, "-")

		left, _ := strconv.Atoi(parts[0])
		right, _ := strconv.Atoi(parts[1])
		interval := &Interval{
			left,
			right,
		}
		intervals = append(intervals, interval)
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})

	currentInterval := intervals[0]
	count := 0

	for _, intv := range intervals[1:] {
		if intv.Start > currentInterval.End {
			count += currentInterval.End - currentInterval.Start + 1
			currentInterval = intv
		} else if intv.Start <= currentInterval.End {
			currentInterval.End = max(currentInterval.End, intv.End)
		}
	}
	count += currentInterval.End - currentInterval.Start + 1
	fmt.Println(count)
}
