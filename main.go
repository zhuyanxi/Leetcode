package main

import (
	"fmt"
)

func main() {
	// nums := []int{12, 2, 17, 7, 7}
	// target := 9
	// t1 := time.Now()
	// r := twoSumFast(nums, target)
	// t2 := time.Now()
	// fmt.Println("twoSum time:", t2.Sub(t1))
	// fmt.Println(r[0], r[1])

	// x := 320
	// y := 1000000
	// r := x / y
	// fmt.Println(r)
	// t1 := time.Now().UnixNano()
	// fmt.Println(reverse(x))
	// t2 := time.Now().UnixNano()
	// fmt.Println("time:", t2-t1)

	a := make([]int, 0, 5)
	printSlice("a", a)
	a = append(a, 4)
	printSlice("a", a)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
