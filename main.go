package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello")

	//primes := [6]int{2, 3, 3, 7, 7, 13}
	// s := primes[1:5]
	// printSlice("s", s)

	//datastructure.TestMinStack()
	//TestIsValidBrackets()
	// l := sumSlice(primes[:2])
	// fmt.Println(l)
	// r := sumSlice(primes[6:])
	// fmt.Println(r)
	//testPivotIndex()
	//TestDominantIndex()
	//removeDuplicates(primes[:])
	//fmt.Println(reverseString("A man, a plan, a canal: Panama"))
	//testFirstUniqChar()
	testIsAnagram()
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func addBinary(a string, b string) string {
	return ""
}
