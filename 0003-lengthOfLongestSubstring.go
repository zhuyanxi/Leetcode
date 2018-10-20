package main

import (
	"fmt"
)

/*
给定一个字符串，找出不含有重复字符的最长子串的长度。

示例 1:

输入: "abcabcbb"
输出: 3
解释: 无重复字符的最长子串是 "abc"，其长度为 3。
示例 2:

输入: "bbbbb"
输出: 1
解释: 无重复字符的最长子串是 "b"，其长度为 1。
示例 3:

输入: "pwwkew"
输出: 3
解释: 无重复字符的最长子串是 "wke"，其长度为 3。
     请注意，答案必须是一个子串，"pwke" 是一个子序列 而不是子串。
*/
func lengthOfLongestSubstring(s string) int {
	longestLen := 0
	i := 0
	var lArr []byte
	for {
		for j := i; j < len(s); j++ {
			if !elementInSlice(lArr, s[j]) {
				lArr = append(lArr, s[j])
			} else {
				i++
				break
			}
		}
		if len(lArr) > longestLen {
			longestLen = len(lArr)
		}
		lArr = []byte{}
		if longestLen >= len(s)-i {
			break
		}
	}
	return longestLen
}

func elementInSlice(s []byte, ch byte) bool {
	for _, v := range s {
		if v == ch {
			return true
		}
	}
	return false
}

func lengthOfLongestSubstringFast(s string) int {
	lastOccurred := make(map[byte]int)
	start := 0
	maxLength := 0
	for i, ch := range []byte(s) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}

func testLengthOfLongestSubstring() {
	s := "pwwkew"
	fmt.Println(lengthOfLongestSubstring(s))
}
