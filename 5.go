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

func (n *Node) isInAnyInterval(x int) bool {
	if x > n.Max {
		return false
	}
	if x >= n.Start && x <= n.End {
		return true
	}
	if n.Left != nil && (n.Left.Max > x && n.Left.isInAnyInterval(x) || x == n.Left.Max) {
		return true
	}
	if n.Right != nil && (x >= n.Start && n.Right.isInAnyInterval(x) || n.Right.Max == x) {
		return true
	}
	return false

}

func (n *Node) updateMax() {
	max := n.End
	if n.Left != nil && n.Left.Max > max {
		max = n.Left.Max
	}
	if n.Right != nil && n.Right.Max > max {
		max = n.Right.Max
	}
	n.Max = max
}

func BuildTree(intervals []*Interval) *Node {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})
	return buildRecursive(intervals)
}

func buildRecursive(intervals []*Interval) *Node {
	if len(intervals) == 0 {
		return nil
	}
	mid := len(intervals) / 2
	root := &Node{
		Start: intervals[mid].Start,
		End:   intervals[mid].End,
	}
	root.Left = buildRecursive(intervals[:mid])
	root.Right = buildRecursive(intervals[mid+1:])
	root.updateMax()
	return root
}

func main() {
	file, _ := os.Open("5.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	intervals := []*Interval{}
	var nums []int

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

	for scanner.Scan() {
		line := scanner.Text()

		n, _ := strconv.Atoi(line)
		nums = append(nums, n)
	}

	tree := BuildTree(intervals)

	count := 0
	for _, n := range nums {
		if tree.isInAnyInterval(n) {
			count++
		}
	}
	fmt.Println(count)
}
