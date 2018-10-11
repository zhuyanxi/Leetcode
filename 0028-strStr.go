package main

import (
	"fmt"
)

//给定一个 haystack 字符串和一个 needle 字符串，
//在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。
//如果不存在，则返回  -1。
//当 needle 是空字符串时返回 0 。

//使用传统每个字符进行比较
//需要考虑的因素，如果出现没有匹配成功，则重置target，从target+1开始重新进行匹配
func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	flag := false
	target := -1
	for i := 0; i < len(haystack); i++ {
		if !flag {
			if needle[0] == haystack[i] {
				flag = true
				target = i
			}
		}
		if flag {
			if len(haystack)-target < len(needle) {
				return -1
			}
			if haystack[i] != needle[i-target] {
				flag = false
				i = target
				target = -1
			}
			if target != -1 {
				if i-target+1 == len(needle) {
					break
				}
			}
		}

	}
	return target
}

//使用内置的slice方式比较
func strStrFast(haystack string, needle string) int {
	hlen, nlen := len(haystack), len(needle)
	for i := 0; i <= hlen-nlen; i++ {
		if haystack[i:i+nlen] == needle {
			return i
		}
	}
	return -1
}

func testStrStr() {
	//"mississippi"
	//"issip"
	haystack := "mississippi"
	needle := "sipp"
	fmt.Println(strStrFast(haystack, needle))
}
