package main

import (
	"fmt"
)

func testPivotIndex() {
	nums := []int{1, 7, 3, 6, 5, 6}
	fmt.Println(pivotIndex(nums))
}

func pivotIndex(nums []int) int {
	for i := 0; i < len(nums); i++ {
		left := sumSlice(nums[:i])
		right := sumSlice(nums[i+1:])
		if left == right {
			return i
		}
	}
	return -1
}

func sumSlice(s []int) int {
	var sum int
	for i := 0; i < len(s); i++ {
		sum += s[i]
	}
	return sum
}

func pivotIndexFast(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}

	sum2 := 0
	for i, v := range nums {
		if sum2<<1+v == sum {
			return i
		}
		sum2 += v
	}
	return -1
}
