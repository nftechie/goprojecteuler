/* 
	Project Euler
	Problem 3 - Largest Prime Factor
	Alex Wormuth (awormuth)
	10/15/2017

	Answer submitted: 6857
*/ 

package main

import (
	"fmt"
	"math"
)

// isPrime returns true if given integer, n, is prime. Otherwise, false.
func isPrime(n int) bool {
    if n <= 1 {
        return false
    } else if n <= 3 {
        return true
    } else if n % 2 == 0 || n % 3 == 0 {
        return false
    }
    
    i := 5

    for i * i <= n {
    	if n % i == 0 || n % (i + 2) == 0 {
            return false
    	}

       	i += 6
    }

    return true
}

func main() {
	largestPrimeFactor := 1
	num := math.Sqrt(600851475143)

	for i := 1; i < int(num); i++ {
		if isPrime(i) {
			if 600851475143 % i == 0 {
				largestPrimeFactor = i
			}
		}
	}

	fmt.Println(largestPrimeFactor)
}

