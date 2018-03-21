/* 
	Project Euler
	Problem 4 - Largest Palindrome Product
	Alex Wormuth (awormuth)
	10/15/2017

	Answer submitted: 906609
*/ 

package main

import (
	"fmt"
	"strconv"
)

// isPalindrome returns true if the given integer is a palindrome.
func isPalindrome(n int) bool {
	s := strconv.Itoa(n)

	for i, j := 0, len(s) - 1; i < j; i, j = i + 1, j - 1 {
		if s[i] != s[j] {
			return false
		}
	}

	return true
}

func main() {
	largestPalindrome := 0
	n1, n2 := 0, 0

	for i := 999; i > 99; i-- {
		for j := 999; j > 99; j-- {
			if isPalindrome(i * j) && (i * j) > largestPalindrome {
				largestPalindrome = i * j
				n1, n2 = i, j
			}
		}
	}

	fmt.Println(n1, "x", n2)
	fmt.Println(largestPalindrome)
}

