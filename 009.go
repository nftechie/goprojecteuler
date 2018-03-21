/* 
	Project Euler
	Problem 9 - Special Pythagorean Triplet
	Alex Wormuth (awormuth)
	10/15/2017

	Answer submitted: 31875000
*/ 

package main

import (
	"fmt"
	"math"
)

func isPythagTriplet(a, b, c float64) bool {
	if a < b {
		if a * a + b * b == c {
			return true
		}
	}

	return false
}

func main() {
	for c := 1.0; ; c++ {
		if math.Sqrt(c) == math.Floor(math.Sqrt(c)) {
			for i := 1.0; i < math.Sqrt(c); i++ {
				for j := i; j < math.Sqrt(c); j++ {
					if isPythagTriplet(i, j, c) {
						if i + j + math.Sqrt(c) == 1000 {
							fmt.Println(i * j * math.Sqrt(c))
							return
						}
					}
				}
			}
		}
	}
}

