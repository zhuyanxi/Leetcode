package main

import (
	"fmt"
)

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return climbStairs(n-1) + climbStairs(n-2)
}

func climbStairsFast(n int) int {
	steps := []int{1, 1}
	for i := 0; i < n; i++ {
		steps[0], steps[1] = steps[1], steps[0]+steps[1]
	}
	return steps[0]
}

func testClimbStairs() {
	n := 120
	fmt.Println(climbStairsFast(n))
}
