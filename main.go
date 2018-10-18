package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("hello")

	defer timeCost(time.Now())
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
	//testThreeSum()
	//testIsPalindromeInt()
	testClimbStairs()

	//elasp := time.Since(t1)
	//fmt.Println("Execute time: ", elasp)
}

func timeCost(start time.Time) {
	terminal := time.Since(start)
	fmt.Println("Execute time: ", terminal)
}
