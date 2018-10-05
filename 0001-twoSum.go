package main

func twoSum(nums []int, target int) []int {
	var result []int
	flag := false

	if len(nums) == 1 {
		return nil
	}

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				flag = true
				result = append(result, i)
				result = append(result, j)
				break
			}
		}
		if flag {
			break
		}
	}

	if len(result) == 0 {
		return nil
	}
	return result
}

func twoSumFast(nums []int, target int) []int {
	temp := make(map[int]int, len(nums))
	for k, v := range nums {
		if _, ok := temp[target-v]; !ok {
			temp[v] = k
		} else {
			return []int{temp[target-v], k}
		}
	}
	return nil
}
