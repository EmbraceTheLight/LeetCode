// 92. 反转链表 II
/*
给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right。
请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表 。

提示：
链表中节点数目为 n
1 <= n <= 500
-500 <= Node.val <= 500
1 <= left <= right <= n
*/
package main

import (
	"fmt"
	"lc/pkg"
)

func reverseBetween(head *pkg.ListNode, left int, right int) *pkg.ListNode {
	dummy := &pkg.ListNode{Next: head}
	pre, cur := dummy, head
	var leftNode *pkg.ListNode
	count := 1 // cur "下标"
	for {
		if count == left {
			leftNode = cur
			break
		}
		pre = cur
		cur = cur.Next
		count++
	}

	for {
		if count == right {
			break
		}
		cur = cur.Next
		count++
	}
	tail := cur.Next
	cur.Next = nil

	h, t := reverse(leftNode)
	pre.Next = h
	t.Next = tail
	return dummy.Next
}

func reverse(headNode *pkg.ListNode) (head *pkg.ListNode, tail *pkg.ListNode) {
	cur, next := headNode, headNode.Next

	var pre *pkg.ListNode
	var tmp *pkg.ListNode
	for next != nil {
		tmp = next.Next
		cur.Next = pre
		next.Next = cur
		pre = cur
		cur = next
		next = tmp
	}
	return cur, headNode
}

// 示例 1：
// 输入：head = [1,2,3,4,5], left = 2, right = 4
// 输出：[1,4,3,2,5]
//
// 示例 2：
// 输入：head = [5], left = 1, right = 1
// 输出：[5]
func main() {
	var left, right int
	fmt.Println("Input left, right:")
	fmt.Scan(&left, &right)
	prices := pkg.CreateList()
	pkg.PrintList(reverseBetween(prices, left, right))
}
