package main

import (
	"fmt"
)

func TestDominantIndex() {
	s := []int{1}
	fmt.Println(dominantIndex(s))
}

func dominantIndex(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	max := 0
	secondMax := 0
	maxValue := 0

	for i := 1; i < len(nums); i++ {
		if nums[i] >= nums[max] {
			max = i
		}
	}
	maxValue = nums[max]

	s := append(nums[:max], nums[max+1:]...)
	for i := 1; i < len(s); i++ {
		if s[i] >= s[secondMax] {
			secondMax = i
		}
	}

	if s[secondMax]<<1 <= maxValue {
		return max
	}

	return -1
}

func dominantIndexFast(nums []int) int {
	first, second := 0, 0
	maxIndex := -0
	for i, v := range nums {
		if v > second {
			second = v
			if second > first {
				first, second = second, first
				maxIndex = i
			}
		}
	}

	if first >= second<<1 {
		return maxIndex
	}
	return -1
}
