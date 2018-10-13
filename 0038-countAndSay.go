package main

import (
	"fmt"
)

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	prev := countAndSay(n - 1)
	return countAndSaySingle(prev)
}

func countAndSaySingle(prev string) string {
	var barr []byte
	count := '1'
	for i := 0; i < len(prev); i++ {
		if i != len(prev)-1 {
			if prev[i] == prev[i+1] {
				count++
			} else {
				barr = append(barr, byte(count))
				barr = append(barr, prev[i])
				count = '1'
			}
		} else {
			barr = append(barr, byte(count))
			barr = append(barr, prev[i])
		}
	}
	return string(barr)
}

func testCountAndSay() {
	fmt.Println(countAndSay(5))
	//fmt.Println(countAndSaySingle("11122311222"))
}
