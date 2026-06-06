package main

import (
	"fmt"
	"lc/pkg"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *pkg.ListNode, n int) *pkg.ListNode {
	var dummy = &pkg.ListNode{}
	dummy.Next = head
	var start, end = head, head
	for i := 0; i < n; i++ {
		end = end.Next
	}
	var pre = dummy
	for end != nil {
		pre = start
		start = start.Next
		end = end.Next
	}
	pre.Next = start.Next
	return dummy.Next
}
func main() {
	var n int
	fmt.Scan(&n)

	root := pkg.CreateList()
	fmt.Println(removeNthFromEnd(root, n))
}
