// 25. K 个一组翻转链表
/*
给你链表的头节点 head ，每 k 个节点一组进行翻转，请你返回修改后的链表。
k 是一个正整数，它的值小于或等于链表的长度。如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。

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
	dummy := &pkg.ListNode{Next: head}
	tail := dummy
Outer:
	for tail != nil {
		tmp := head
		pre := head
		for i := 0; i < k; i++ {
			if head == nil {
				tail.Next = tmp
				break Outer
			}
			pre = head
			head = head.Next
		}
		pre.Next = nil
		newHead, newTail := reverseLinklist(tmp)
		tail.Next = newHead
		tail = newTail
	}
	return dummy.Next
}

func reverseLinklist(head *pkg.ListNode) (newHead, newTail *pkg.ListNode) {
	var first, second, third *pkg.ListNode = head, head.Next, nil
	newTail = first
	for second != nil {
		third = second.Next
		second.Next = first
		first = second
		second = third
	}
	newTail.Next = nil
	return first, newTail
}

// 示例 1：
// 输入：head = [1,2,3,4,5], k = 2
// 输出：[2,1,4,3,5]
//
// 示例 2：
// 输入：head = [1,2,3,4,5], k = 3
// 输出：[3,2,1,4,5]
func main() {
	var k int
	fmt.Println("Input k:")
	fmt.Scan(&k)
	pkg.PrintList(reverseKGroup(pkg.CreateList(), k))
}
