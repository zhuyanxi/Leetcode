package main

import (
	"fmt"
)

func isPalindromeInt(x int) bool {
	if x < 0 {
		return false
	}
	tmp := x
	result := 0
	ys := 0
	for x > 0 {
		ys = x % 10
		result = result*10 + ys
		x = x / 10
	}
	if result != tmp {
		return false
	}
	return true
}

func testIsPalindromeInt() {
	a := 1001
	fmt.Println(isPalindromeInt(a))
}
