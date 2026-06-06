/*
2. 两数相加
给你两个 非空 的链表，表示两个非负的整数。
它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
请你将两个数相加，并以相同形式返回一个表示和的链表。
你可以假设除了数字 0 之外，这两个数都不会以 0 开头。

输入：l1 = [2,4,3], l2 = [5,6,4]
输出：[7,0,8]
解释：342 + 465 = 807.
*/
package main

import (
	"fmt"
)

type ListNode2 struct {
	Val  int
	Next *ListNode2
}

func printList2(res *ListNode2) {
	for ; res != nil; res = res.Next {
		fmt.Printf("%d ", res.Val)
	}
}
func (l *ListNode2) Insert(val int) (tail *ListNode2) {
	node := new(ListNode2)
	node.Val = val
	l.Next = node
	tail = node
	return
}

func addTwoNumbers(l1 *ListNode2, l2 *ListNode2) *ListNode2 {
	var carry int
	var ret = new(ListNode2)
	var node = ret
	//var pre = node
	ret.Next = node
	for l1 != nil && l2 != nil {
		res := l1.Val + l2.Val + carry
		if res >= 10 {
			carry = 1
			res %= 10
		} else {
			carry = 0
		}
		node.Next = new(ListNode2)
		node = node.Next
		node.Val = res
		//pre = node
		l1 = l1.Next
		l2 = l2.Next
	}
	for l1 != nil {
		res := l1.Val + carry
		if res >= 10 {
			carry = 1
			res %= 10
		} else {
			carry = 0
		}
		node.Next = new(ListNode2)
		node = node.Next
		node.Val = res

		l1 = l1.Next
	}
	for l2 != nil {
		res := l2.Val + carry
		if res >= 10 {
			carry = 1
			res %= 10
		} else {
			carry = 0
		}
		node.Next = new(ListNode2)
		node = node.Next
		node.Val = res

		l2 = l2.Next
	}
	if carry != 0 {
		node.Next = new(ListNode2)
		node = node.Next
		node.Val = carry

	}
	return ret.Next
}
func main() {
	fmt.Println("Input List1:")
	var tmp int
	head := new(ListNode2)
	fmt.Scan(&tmp)
	head.Val = tmp
	next := head
	for {
		fmt.Scan(&tmp)
		if tmp == -1 {
			break
		}
		next = next.Insert(tmp)
	}
	printList2(head)
	fmt.Println()

	fmt.Println("Input List2:")
	head2 := new(ListNode2)
	fmt.Scan(&tmp)
	head2.Val = tmp
	next = head2
	for {
		fmt.Scan(&tmp)
		if tmp == -1 {
			break
		}
		next = next.Insert(tmp)
	}
	printList2(head2)
	fmt.Println()

	printList2(addTwoNumbers(head, head2))
}
