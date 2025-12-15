package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Point struct {
	id      int
	x, y, z int
}

type Edge struct {
	u, v   int
	distSq int
	x, x2  int
}

type UnionFind struct {
	parent []int
	size   []int
}

func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = 1
	}
	return &UnionFind{parent: parent, size: size}
}

func (uf *UnionFind) Find(i int) int {
	if uf.parent[i] != i {
		uf.parent[i] = uf.Find(uf.parent[i])
	}
	return uf.parent[i]
}

func (uf *UnionFind) Union(i, j int) {
	rootI := uf.Find(i)
	rootJ := uf.Find(j)

	if rootI != rootJ {
		if uf.size[rootI] < uf.size[rootJ] {
			rootI, rootJ = rootJ, rootI
		}
		uf.parent[rootJ] = rootI
		uf.size[rootI] += uf.size[rootJ]
	}
}

func main() {
	file, _ := os.Open("8.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var points []Point
	idCounter := 0

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		z, _ := strconv.Atoi(parts[2])
		points = append(points, Point{id: idCounter, x: x, y: y, z: z})
		idCounter++
	}

	var edges []Edge
	n := len(points)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			p1, p2 := points[i], points[j]
			// Use distance squared, it won't change the sorting of the edges.
			distSq := (p1.x-p2.x)*(p1.x-p2.x) +
				(p1.y-p2.y)*(p1.y-p2.y) +
				(p1.z-p2.z)*(p1.z-p2.z)
			edges = append(edges, Edge{u: i, v: j, distSq: distSq, x: p1.x, x2: p2.x})
		}
	}

	// Sort edges by distance
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].distSq < edges[j].distSq
	})

	// We only care about the top 1000 edges
	uf := NewUnionFind(n)
	limit := len(edges)

	edgesCount := 0
	// We need to break out of the loop as soon as we know the current graph is a tree (N-1 edges). When we add the n-1th edge, just perform the x's multiplication of each end
	for i := 0; i < limit; i++ {
		if uf.Find(edges[i].u) != uf.Find(edges[i].v) {
			uf.Union(edges[i].u, edges[i].v)
			edgesCount++
			if edgesCount == n-1 {
				fmt.Println(edges[i].x * edges[i].x2)
				break
			}
		}

	}

	sizeMap := make(map[int]int)
	for i := 0; i < n; i++ {
		root := uf.Find(i)
		sizeMap[root] = uf.size[root]
	}

	var sizes []int
	for _, size := range sizeMap {
		sizes = append(sizes, size)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(sizes)))

	product := 1
	count := 0
	for _, s := range sizes {
		product *= s
		count++
		if count == 3 {
			break
		}
	}
	fmt.Println(product)
}
