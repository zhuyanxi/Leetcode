package main

import (
	"fmt"
	"time"
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
	//testIsAnagram()
	//testIsPalindrome()
	//testMyAtoi()
	//testStrStr()
	//testCountAndSay()
	//testLongestCommonPrefix()
	t1 := time.Now()
	testThreeSum()
	elapsed := time.Since(t1)
	fmt.Println("Execute time: ", elapsed)
}
