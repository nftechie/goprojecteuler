/* 
	Project Euler
	Problem 10 - Summation of Primes
	Alex Wormuth (awormuth)
	10/15/2017

	Answer submitted: 142913828922
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
	sum := 0

	for i := 0; i < 2000000; i++ {
		if isPrime(i) {
			sum += i
		}
	}

	fmt.Println(sum)
}

