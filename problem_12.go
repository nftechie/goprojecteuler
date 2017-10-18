/* 
	Project Euler
	Problem 12 - Highly Divisible Triangular Number
	Alex Wormuth (roamingtechie)
	10/17/2017

	Answer submitted: 76576500

	// TODO: Implement more efficient solution.
	// Idea -> store computed factors in hash map and reference for future values.
*/ 

package main

import (
	"fmt"
)

var factors map[int]int

func numFactors(n int) int {
	numFactors := 1

	for i := n / 2 + 1; i >= 1; i-- {
		if n % i == 0 {
			// if val, ok := factors[i]; ok {
			// 	numFactors += val
			// 	break
			// } else {
				numFactors++
			// }
		}
	}

	factors[n] = numFactors

	return numFactors
}

func main() {
	factors = make(map[int]int, 0)

	highestTriangleNumber := 0
	currTriangleNum := 0

	for i := 1; i < 100000; i++ {
		currTriangleNum += i

		if nf := numFactors(currTriangleNum); nf > 500 {
			fmt.Println(nf)
			highestTriangleNumber = currTriangleNum
			break
		}
	}

	fmt.Println(highestTriangleNumber)
	fmt.Println(factors)
}
