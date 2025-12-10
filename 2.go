package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
)

func sumArithmetic(start, end int) int {
	if start > end {
		return 0
	}
	count := end - start + 1
	// Just Gauss formula: (a_1 + a_n) * n / 2. Sums every number in the interval [start, end]
	return (start + end) * count / 2
}

// sumXXUpTo calculates the SUM of all "XX" numbers between 1 and n
func sumXXUpTo(upperBound int) int {
	if upperBound < 0 {
		return 0
	}
	s := strconv.Itoa(upperBound)
	digits := len(s)
	totalSum := 0

	// All smaller lengths fit nicely (if upper bound has 2m digits, any m digit X will work, we calcualate the sum of all of them)
	for length := 2; length < digits; length += 2 {
		m := length / 2

		// Range of X is 10**(m-1) to 10**m - 1
		startX := int(math.Pow(10, float64(m-1)))
		endX := int(math.Pow(10, float64(m))) - 1

		// The number form is X * (10^m + 1).
		// for example m=2: X * 101. (12 * 101 = 1212)
		factor := int(math.Pow(10, float64(m))) + 1

		sumOfX := sumArithmetic(startX, endX)
		totalSum += sumOfX * factor
	}

  // DANGEROUS CASE: Once we want to create numbers with 2m digits, we must check if the first half of the upperBound concatenated with itself is NOT greater than the upperBound 
  // (otherwise, we'd for example count 4242 when the upperbound is 4230 or 4300).
	if digits%2 == 0 {
		m := digits / 2
		startX := int(math.Pow(10, float64(m-1)))

		firstHalf, _ := strconv.Atoi(s[:m])
		candidateVal, _ := strconv.Atoi(fmt.Sprintf("%d%d", firstHalf, firstHalf))

		var endX int
		if candidateVal <= upperBound {
			endX = firstHalf // We can include the prefix+prefix, the first half.
		} else {
			endX = firstHalf - 1 // Prefix is too big, go for the immediate smaller X.
		}

		factor := int(math.Pow(10, float64(m))) + 1
		sumOfX := sumArithmetic(startX, endX)
		totalSum += sumOfX * factor
	}

	return totalSum
}

func main() {
	content, _ := os.ReadFile("2.txt")

	line := strings.TrimSpace(string(content))
	parts := strings.Split(line, ",")

	var globalSum int64
	var wg sync.WaitGroup

  // Solution is simple. We're looking for patterns N = XX. We must see how many such Ns exist in a given interval. We generate them based on generating the Xs instead of the full N. 
  // Given any N with 2m digits, any XX is composed of an X of m digits. Precisely, all of these are all possible integers of m digits. Say, N has 4 digits, X = 10,11,12,13,14... will create every possible pattern XX.
	for _, s := range parts {
		wg.Add(1)

		go func(interval string) {
			defer wg.Done()

			lowerStr, upperStr, found := strings.Cut(interval, "-")
			if !found {
				return
			}

			lower, _ := strconv.Atoi(lowerStr)
			upper, _ := strconv.Atoi(upperStr)
			intervalSum := sumXXUpTo(upper) - sumXXUpTo(lower-1)

			// Atomic add, because we are using go routines and running for each interval concurrently.
			atomic.AddInt64(&globalSum, int64(intervalSum))

		}(s)
	}

	wg.Wait()
	fmt.Println("Total Sum:", globalSum)
}
