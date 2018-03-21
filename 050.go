/* 
	Project Euler
	Problem 50 - Consecutive Prime Sum
	Alex Wormuth (awormuth)
	10/15/2017

	Answer submitted: 997651

	TODO: Create a more efficient solution. Current solution is brute froce.
*/ 

package main

import (
	"fmt"
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
	primes := make(map[int]bool, 0)
	primesArr := make([]int, 0)
	longestSum := 0
	longestDiff := 0

	for i := 0; i < 1000000; i++ {
		if isPrime(i) {
			primes[i] = true
			primesArr = append(primesArr, i)
		}
	}

	for i := 0; i < len(primesArr); i++ {
		currSum := 0

		for j := i; j < len(primesArr); j++ {
			currSum += primesArr[j]

			if _, ok := primes[currSum]; ok {
				if j - i > longestDiff {
					longestDiff = j - i + 1
					longestSum = currSum
				}
			}
		}
	}

	fmt.Println(longestDiff)
	fmt.Println(longestSum)
}

