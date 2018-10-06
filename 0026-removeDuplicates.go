//删除排序数组中的重复项
package main

import (
	"fmt"
)

// 双指针法
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}
	fmt.Println(nums[:i+1])
	return i + 1
}

func removeDuplicatesFast(nums []int) int {
	pmap := make(map[int]int)
	for _, v := range nums {
		if _, ok := pmap[v]; !ok {
			pmap[v] = len(pmap)
			nums[pmap[v]] = v
		}
	}
	return len(pmap)
}
