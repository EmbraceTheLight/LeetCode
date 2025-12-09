// 25. K 个一组翻转链表
/*
给你链表的头节点 head ，每 k 个节点一组进行翻转，请你返回修改后的链表。
k 是一个正整数，它的值小于或等于链表的长度。如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。

示例 1：
输入：head = [1,2,3,4,5], k = 2
输出：[2,1,4,3,5]

示例 2：
输入：head = [1,2,3,4,5], k = 3
输出：[3,2,1,4,5]

提示：
链表中的节点数目为 n
1 <= k <= n <= 5000
0 <= Node.val <= 1000

进阶：你可以设计一个只用 O(1) 额外内存空间的算法解决此问题吗？
*/
package main

import (
	"fmt"
	"lc/pkg"
)

func reverseKGroup(head *pkg.ListNode, k int) *pkg.ListNode {
	if k == 1 {
		return head
	}

	count := 0

	for tmp := head; tmp != nil; tmp = tmp.Next {
		count++
	}

	groups := count / k
	dummy := &pkg.ListNode{Next: head}
	pre := dummy
	next := dummy.Next
	for i := 0; i < groups; i++ {
		head, next = reverse(next, k)
		pre.Next = head
		pre = next
		next = next.Next
	}
	return dummy.Next
}

// 翻转从 head 开始的 k 个节点, 返回头节点和尾节点的位置
func reverse(head *pkg.ListNode, k int) (*pkg.ListNode, *pkg.ListNode) {
	newTail := head
	var pre *pkg.ListNode
	cur := head
	next := cur.Next
	count := 0
	for count < k {
		cur.Next = pre

		pre = cur
		cur = next
		if next != nil {
			next = next.Next
		}
		count++
	}

	newTail.Next = cur
	newHead := pre
	return newHead, newTail
}

// Test Case1: [1,2,3,4,5] k=2	Output: [2,1,4,3,5]
// Test Case2: [1,2,3,4,5] k=3	Output: [3,2,1,4,5]
func main() {
	var k int
	fmt.Println("Input k:")
	fmt.Scan(&k)

	head := pkg.CreateList()
	pkg.PrintList(reverseKGroup(head, k))
}
