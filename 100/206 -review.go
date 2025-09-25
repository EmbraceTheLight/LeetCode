package main

import (
	"lc/pkg"
)

func reverseListR(head *pkg.ListNode) *pkg.ListNode {
	if head == nil {
		return nil
	}
	first, second := head, head.Next
	var pre *pkg.ListNode = nil
	for second != nil {
		first.Next = pre

		pre = first
		first = second
		second = second.Next
		if second == nil {
			first.Next = pre
			break
		}
	}

	return first
}
func main() {
	head := pkg.CreateList()
	pkg.PrintList(reverseListR(head))
}
