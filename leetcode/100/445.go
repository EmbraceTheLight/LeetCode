/*
445. 两数相加 II
给你两个 非空 链表来代表两个非负整数。数字最高位位于链表开始位置。
它们的每个节点只存储一位数字。将这两数相加会返回一个新的链表。
你可以假设除了数字 0 之外，这两个数字都不会以零开头。
输入：l1 = [7,2,4,3], l2 = [5,6,4]
输出：[7,8,0,7]
*/
package main

import (
	"fmt"
	"lc/100/pkg"
)

func rvsList445(head *pkg.ListNode) *pkg.ListNode {
	var pre, cur *pkg.ListNode
	pre = nil
	cur = head
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre
}

func addTwoNumbers445(l1 *pkg.ListNode, l2 *pkg.ListNode) *pkg.ListNode {
	h1 := rvsList445(l1)
	h2 := rvsList445(l2)

	var cur *pkg.ListNode
	var c int
	for h1 != nil && h2 != nil {
		addv := h1.Val + h2.Val + c
		if addv >= 10 {
			addv -= 10
			c = 1
		} else {
			c = 0
		}
		newNode := new(pkg.ListNode)
		newNode.Val = addv
		newNode.Next = cur
		cur = newNode
		h1 = h1.Next
		h2 = h2.Next
	}

	for h1 != nil {
		addv := h1.Val + c
		if addv >= 10 {
			addv -= 10
			c = 1
		} else {
			c = 0
		}
		newNode := new(pkg.ListNode)
		newNode.Val = addv
		newNode.Next = cur
		cur = newNode

		h1 = h1.Next
	}
	for h2 != nil {
		addv := h2.Val + c
		if addv >= 10 {
			addv -= 10
			c = 1
		} else {
			c = 0
		}
		newNode := new(pkg.ListNode)
		newNode.Val = addv
		newNode.Next = cur
		cur = newNode
		h2 = h2.Next

	}
	if c != 0 {
		newNode := new(pkg.ListNode)
		newNode.Val = 1
		newNode.Next = cur
		cur = newNode
	}
	return cur
}

func main() {
	fmt.Println("Input List1:")
	head1 := pkg.CreateList()
	pkg.PrintList(head1)

	fmt.Println("Input List2:")
	head2 := pkg.CreateList()
	pkg.PrintList(head2)

	pkg.PrintList(addTwoNumbers445(head1, head2))
	fmt.Println()
}
