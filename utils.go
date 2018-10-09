package main

import "fmt"

func isNumber(b byte) bool {
	return b >= 48 && b <= 57
}

func isLowerAlphabet(b byte) bool {
	return b >= 97 && b <= 122
}

func isUpperAlphabet(b byte) bool {
	return b >= 65 && b <= 90
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
