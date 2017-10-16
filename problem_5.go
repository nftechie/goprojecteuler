/* 
	Project Euler
	Problem 5 - Smallest Multiple
	Alex Wormuth (roamingtechie)
	10/15/2017

	Answer submitted: 232792560
*/ 

package main

import (
	"fmt"
)

// isMultiple returns true if the given number is evenly divisible by numbers 1 through 20.
func isMultiple(n int) bool {
	for i := 1; i <= 20; i++ {
		if n % i != 0 {
			return false
		}
	}

	return true
}

func main() {
	for i := 1; ; i++ {
		if isMultiple(i) {
			fmt.Println(i)
			break
		}
	}
}

