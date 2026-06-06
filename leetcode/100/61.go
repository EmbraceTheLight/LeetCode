/*
61. 旋转链表
给你一个链表的头节点 head ，旋转链表，将链表每个节点向右移动 k 个位置
输入：head = [1,2,3,4,5], k = 2
输出：[4,5,1,2,3]

链表中节点的数目在范围 [0, 500] 内
-100 <= Node.val <= 100
0 <= k <= 2 * 10^9
*/

package main

import (
	"fmt"
	"lc/100/pkg"
)

func rotateRight(head *pkg.ListNode, k int) *pkg.ListNode {
	if head == nil || k == 0 {
		return head
	}
	var ret = head
	var swap = head
	var cur = head

	var tail *pkg.ListNode

	var length int
	for cur != nil {
		cur = cur.Next
		length++
	}
	var k1 = k % length

	tail = head
	for tail.Next != nil {
		tail = tail.Next
	}

	tail.Next = head

	for i := 0; i < length-k1-1; i++ {
		swap = swap.Next
	}

	ret = swap.Next
	swap.Next = nil
	return ret
}

func main() {
	var k int
	fmt.Println("input k:")
	fmt.Scan(&k)

	println("Input values:")
	head := pkg.CreateList()

	pkg.PrintList(rotateRight(head, k))
}
