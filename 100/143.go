/*
143. 重排链表
给定一个单链表 L 的头节点 head ，单链表 L 表示为：

L0 → L1 → … → Ln - 1 → Ln
请将其重新排列后变为：

L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换
*/
package main

import (
	"fmt"
)

type ListNode143 struct {
	Val  int
	Next *ListNode143
}

// 反转链表
func rvsList(n *ListNode143) *ListNode143 {
	if n == nil || n.Next == nil {
		return n
	}
	head := n
	var slow *ListNode143
	var fast *ListNode143
	nxt := n.Next.Next
	slow = n
	fast = n.Next
	for fast != nil {

		fast.Next = slow

		slow = fast
		fast = nxt
		if fast == nil {
			break
		}
		nxt = nxt.Next
	}
	head.Next = nil
	return slow
}
func printList143(res *ListNode143) {
	for ; res != nil; res = res.Next {
		fmt.Printf("%d ", res.Val)
	}
}
func reorderList(head *ListNode143) *ListNode143 {
	if head == nil || head.Next == nil {
		return head
	}
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	//slow为链表中点
	head2 := rvsList(slow.Next) //反转中点之后的链表部分
	slow.Next = nil
	start := head
	for head2 != nil {
		tmps := start.Next
		tmph2 := head2.Next

		head2.Next = tmps
		start.Next = head2
		start = tmps
		head2 = tmph2
	}
	return head
}
func main() {
	fmt.Println("Input List1:")

	printList143(head.Next)
	fmt.Println()
	//printList143(rvsList(head.Next))
	printList143(reorderList(head.Next))
}
