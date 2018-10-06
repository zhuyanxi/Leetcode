package main

import (
	"fmt"
	"math"
	"strconv"
)

func reverse(x int) int {
	if x > math.MaxInt32 || x < math.MinInt32 {
		return 0
	}

	isMinus := false
	if x < 0 {
		isMinus = true
		x = -x
	}

	xStr := strconv.Itoa(x)

	result := 0
	flag := float64(0)
	for i := range xStr {
		j, _ := strconv.Atoi(string(xStr[i]))
		result += int(math.Pow(10, flag)) * j
		if result > math.MaxInt32 || result < math.MinInt32 {
			return 0
		}
		flag++
		fmt.Println(j)
	}

	if isMinus {
		return -result
	}

	return result
}

func reverseFast(x int) int {
	result := 0
	absX := int(math.Abs(float64(x)))

	for absX > 0 {
		tmp := result*10 + absX%10
		if tmp/10 != result {
			return 0
		}
		result = tmp
		absX = absX / 10
	}
	if result > math.MaxInt32 || result < math.MinInt32 {
		return 0
	}
	if x < 0 {
		return -result
	}
	return result
}
