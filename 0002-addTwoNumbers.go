package main

import (
	"fmt"

	datastructure "github.com/zhuyanxi/Leetcode/DataStructure"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
给定两个非空链表来表示两个非负整数。位数按照逆序方式存储，它们的每个节点只存储单个数字。将两数相加返回一个新的链表。

你可以假设除了数字 0 之外，这两个数字都不会以零开头。

示例：

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807
*/
func addTwoNumbers(l1 *datastructure.ListNode, l2 *datastructure.ListNode) *datastructure.ListNode {
	var overnum, sum, v1, v2 int
	var ss []int
	var result *datastructure.ListNode
	result = nil
	for {
		if l1 == nil && l2 == nil {
			break
		}
		if l1 == nil {
			v1 = 0
		} else {
			v1 = l1.Val
		}
		if l2 == nil {
			v2 = 0
		} else {
			v2 = l2.Val
		}
		sum = overnum + v1 + v2

		if sum >= 10 {
			overnum = sum / 10
			sum = sum % 10
		} else {
			overnum = 0
		}
		ss = append(ss, sum)

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}

	}
	if overnum > 0 {
		ss = append(ss, overnum)
	}
	for i := len(ss) - 1; i >= 0; i-- {
		result = &datastructure.ListNode{ss[i], result}
	}
	return result
}

func testAddTwoNumbers() {
	m := &datastructure.ListNode{2, nil}
	m = &datastructure.ListNode{3, m}
	m = &datastructure.ListNode{4, m}

	n := &datastructure.ListNode{9, nil}
	n = &datastructure.ListNode{4, n}

	s := addTwoNumbers(m, n)
	for s != nil {
		fmt.Println(s.Val)
		s = s.Next
	}

}
