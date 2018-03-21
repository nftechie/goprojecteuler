/* 
	Project Euler
	Problem 6 - Sum Square Difference
	Alex Wormuth (awormuth)
	10/15/2017

	Answer submitted: 25164150
*/ 

package main

import (
	"fmt"
)

func main() {
	sumOfSquares := 0
	squareOfSums := 0

	for i := 0; i <= 100; i++ {
		sumOfSquares += i * i
		squareOfSums += i
	}

	squareOfSums = squareOfSums*squareOfSums

	delta := squareOfSums - sumOfSquares

	fmt.Println(delta)
}

