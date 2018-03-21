/* 
	Project Euler
	Problem 7 - 10001st Prime
	Alex Wormuth (roamingtechie)
	10/15/2017

	Answer submitted: 104743
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
	primeCounter := 0
	i := 2

	for {
		if isPrime(i) {
			primeCounter++

			if primeCounter == 10001 {
				fmt.Println(i)
				break
			}
		}

		i++
	}
}

