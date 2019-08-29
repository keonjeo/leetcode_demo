package main

import (
	"fmt"
)

func findTheDifference(s string, t string) byte {
	firstSum, secondSum := 0, 0
	for _, m := range s {
		firstSum += int(m)
	}
	for _, n := range t {
		secondSum += int(n)
	}

	return byte(secondSum - firstSum)
}

func main() {
	a := "abcd"
	b := "abcdf"
	result := findTheDifference(a, b)
	fmt.Println(result)
}
