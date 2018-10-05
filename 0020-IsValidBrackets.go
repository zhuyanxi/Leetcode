package main

import (
	"fmt"

	datastructure "github.com/zhuyanxi/Leetcode/DataStructure"
)

//'('，')'，'{'，'}'，'['，']'

func TestIsValidBrackets() {
	s := "]{"
	fmt.Println(IsValidBrackets(s))
}

func IsValidBrackets(s string) bool {
	brackets := map[string]string{
		")": "(",
		"}": "{",
		"]": "[",
	}

	var n datastructure.Stack
	var j string
	for i := range s {
		j = string(s[i])
		if j == "(" || j == "{" || j == "[" {
			n.Push(j)
		} else {
			if n.Peek() == nil {
				return false
			}
			if n.Peek().(string) == brackets[j] {
				n.Pop()
			} else {
				return false
			}
		}
	}
	if n.Len() > 0 {
		return false
	}
	return true
}
