package main

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	lenStr := len(strs)
	if lenStr == 0 {
		return ""
	}
	if lenStr == 1 {
		return strs[0]
	}
	if lenStr == 2 {
		return longestPrefixOfTwo(strs[0], strs[1])
	}

	preLeft := longestCommonPrefix(strs[:lenStr/2])
	if longestCommonPrefix(strs[:lenStr/2]) == "" {
		return ""
	}
	preRight := longestCommonPrefix(strs[lenStr/2:])
	if preRight == "" {
		return ""
	}

	return longestPrefixOfTwo(preLeft, preRight)
}

func longestPrefixOfTwo(s1, s2 string) string {
	var pre []byte
	if len(s1) > len(s2) {
		s1, s2 = s2, s1
	}
	for i := 0; i < len(s1); i++ {
		if s1[i] == s2[i] {
			pre = append(pre, s1[i])
		} else {
			break
		}
	}
	if len(pre) == 0 {
		return ""
	}
	return string(pre)
}

func longestCommonPrefixFast(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	s1 := strs[0]

	i := 0

	for ; i < len(s1); i++ {
		char := s1[i]
		for j := 1; j < len(strs); j++ {
			if i < len(strs[j]) && strs[j][i] == char {
				continue
			} else {
				return s1[:i]
			}
		}
	}

	return s1[:i]
}

func testLongestCommonPrefix() {
	//s1 := "zhzuyanxi"
	//s2 := "zhzh"
	//s3 := "zhuxiwang"
	strs := []string{}
	//fmt.Println(longestPrefixOfTwo(s1, s2))
	fmt.Println(longestCommonPrefix(strs))
}
