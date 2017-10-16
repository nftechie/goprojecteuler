/* 
	Project Euler
	Problem 16 - Power Digit Sum
	Alex Wormuth (roamingtechie)
	10/15/217

	Answer submitted: 1366
*/ 

package main

import (
	"fmt"
    "strconv"
)

type Digit struct {
    Value int
    Carry int
}

func printArr(arr []Digit) {
    for i := len(arr) - 1; i >= 0 ; i-- {
        fmt.Print(arr[i].Value)
    }

    fmt.Println("")
}

func calcSum(arr []Digit) int {
    sum := 0

    for i := 0; i < len(arr); i++ {
        sum += arr[i].Value
    }

    return sum
}

func main() {
    arr := make([]Digit, 1)

    arr[0] = Digit{1, 0}

    for i := 1; i <= 1000; i++ {
        for j := 0; j < len(arr); j++ {
            if (2 * arr[j].Value) + arr[j].Carry > 9 {
                s := strconv.Itoa((2 * arr[j].Value) + arr[j].Carry)
                s0, _ := strconv.ParseInt(string(s[0]), 10, 64)
                s1, _ := strconv.ParseInt(string(s[1]), 10, 64)

                arr[j].Value = int(s1)

                if j == len(arr) - 1 {
                    arr = append(arr, Digit{1, 0})
                    break
                } 

                arr[j + 1].Carry = int(s0)
            } else {
                arr[j].Value = 2 * arr[j].Value + arr[j].Carry
            }
        }

        // Clear carries.
        for j := 0; j < len(arr); j++ {
            arr[j].Carry = 0
        }
    }

    printArr(arr)

    sum := calcSum(arr)

    fmt.Println(sum)
}

