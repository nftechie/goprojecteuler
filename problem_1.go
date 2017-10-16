/* 
	Project Euler
	Problem 1 - Multiples of 3 and 5
	Alex Wormuth (roamingtechie)
	10/15/2017

	Answer submitted: 233168
*/ 

package main

import "fmt"

func main() {
	sum := 0

	for i := 0; i < 1000; i++ {
		if i % 3 == 0 || i % 5 == 0 {
			sum += i
		}
	}

	fmt.Println(sum)
}

