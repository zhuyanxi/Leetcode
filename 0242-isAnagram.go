package main

import (
	"fmt"
)

// 给定两个字符串 s 和 t ，判断 t 是否是 s 的一个字母异位词。
func isAnagram(s string, t string) bool {
	smap := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		smap[s[i]]++
	}
	for i := 0; i < len(t); i++ {
		_, ok := smap[t[i]]
		if ok {
			smap[t[i]]--
			if smap[t[i]] < 0 {
				return false
			}

		} else {
			return false
		}
	}
	for _, v := range smap {
		if v != 0 {
			return false
		}
	}
	return true
}

func testIsAnagram() {
	s := "anagram"
	t := "nagaramm"
	fmt.Println(isAnagram(s, t))
}
