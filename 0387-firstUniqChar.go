package main

import (
	"fmt"
)

// 给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。
func firstUniqChar(s string) int {
	l := len(s)
	if l == 0 {
		return -1
	}
	if l == 1 {
		return 0
	}
	index := 0
	isUniq := true
	for i := 0; i < l; i++ {
		isUniq = true
		for j := 0; j < l; j++ {
			if j != i {
				if s[i] == s[j] {
					isUniq = false
					break
				}
			}

		}
		index = i
		if isUniq {
			break
		}
	}

	if isUniq {
		return index
	}
	return -1
}

func firstUniqCharFast(s string) int {
	n := len(s)
	if n == 1 {
		return 0
	}

	az := make([]int, 26, 26)

	for i := 0; i < n; i++ {
		az[s[i]-'a']++
	}

	for i := 0; i < n; i++ {
		if az[s[i]-'a'] == 1 {
			return i
		}
	}

	return -1
}

func testFirstUniqChar() {
	s := "qqqqrqqrq"
	fmt.Println(firstUniqChar(s))
}
