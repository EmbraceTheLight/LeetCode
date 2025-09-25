package main

import (
	"fmt"
	"lc/pkg"
)

func reverseList25(head *pkg.ListNode) (headNode, tail *pkg.ListNode) {

	var pre *pkg.ListNode = nil
	var first, second = head, head.Next

	for second != nil {
		first.Next = pre
		pre = first
		first = second
		second = second.Next
		if second == nil {
			first.Next = pre
		}
	}

	tail = head
	headNode = first
	return
}
func reverseKGroupR(head *pkg.ListNode, k int) *pkg.ListNode {
	var ans *pkg.ListNode
	var first, second, pre *pkg.ListNode = head, head, nil
	var ok bool

	for second != nil {
		var next *pkg.ListNode
		for i := 0; i < k-1; i++ {
			second = second.Next
			if second == nil {
				break
			}
		}
		if second != nil {
			next = second.Next
			second.Next = nil
		} else { //剩下的节点不足k个，结束反转
			pre.Next = first
			break
		}

		h, t := reverseList25(first)

		if pre != nil {
			pre.Next = h
		}
		pre = t

		if !ok { // 将返回值设置为第一个反转后的头节点
			ans = h
			ok = true
		}

		first, second = next, next
	}

	return ans
}
func main() {
	fmt.Println("Input k:")
	var k int
	fmt.Scan(&k)
	head := pkg.CreateList()
	pkg.PrintList(reverseKGroupR(head, k))
}
