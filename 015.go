/* 
	Project Euler
	Problem 15 - Lattice Paths
	Alex Wormuth (awormuth)
	03/21/2018

	Answer submitted: 837799
*/ 

package main

import (
	"fmt"
)

func collatz(n int, count int) int {
	if n == 1 {
		return count
	} else {
		count++

		if n % 2 == 0 {
			return collatz(n / 2, count)
		} else {
			return collatz(3 * n + 1, count)
		}
	}
}

func main() {
	longestChain := 0
	startingNum := 0

	for i := 1; i < 1000000; i++ {
		if col := collatz(i, 1); col > longestChain {
			longestChain = col
			startingNum = i
		}
	}

	fmt.Println(startingNum)
}

