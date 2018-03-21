/* 
	Project Euler
	Problem 2 - Even Fibonacci Numbers
	Alex Wormuth (awormuth)
	10/15/2017

	Answer submitted: 4613732
*/ 

package main

import "fmt"

func main() {
	tmp, prev, curr := 1, 1, 2
	sum := 0

	for {
		if curr > 4000000 {
			break
		}

		if curr % 2 == 0 {
			sum += curr
		}

		tmp = curr
		curr += prev
		prev = tmp
	}

	fmt.Println(sum)
}

