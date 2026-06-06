// 92. 反转链表 II
/*
给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right 。请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表 。
示例 1：
输入：head = [1,2,3,4,5], left = 2, right = 4
输出：[1,4,3,2,5]

示例 2：
输入：head = [5], left = 1, right = 1
输出：[5]

提示：
链表中节点数目为 n
1 <= n <= 500
-500 <= Node.val <= 500
1 <= left <= right <= n

进阶： 你可以使用一趟扫描完成反转吗？
*/
package main

import (
	"fmt"
	"lc/pkg"
)

func reverseBetween(head *pkg.ListNode, left int, right int) *pkg.ListNode {
	if left == right {
		return head
	}
	count := 1
	var pre, cur, next *pkg.ListNode   // 反转过程中的临时节点
	var start, end *pkg.ListNode       // start: 要反转的部分前一个节点 end: 要反转的部分最后一个节点
	var newHead, newTail *pkg.ListNode // 反转完毕后反转部分的头/尾指针

	dummy := &pkg.ListNode{Next: head}
	pre = dummy
	cur = head
	next = cur.Next
	for {
		if count < left {
			count++
			pre = cur
			cur = next
			next = next.Next
		} else if count == left {
			start = pre   // 记录带反转链表前一格节点
			newTail = cur // 记录反转链表尾节点

			pre = cur
			cur = next
			next = next.Next
			count++
		} else {
			cur.Next = pre

			pre = cur
			cur = next
			next = next.Next

			count++
			if count == right+1 {
				newHead = pre
				end = cur
				start.Next = newHead
				newTail.Next = end
				break
			}

		}
	}
	return dummy.Next
}

// Test Case1: [1,2,3,4,5] left=2 right=4	Output: [1,4,3,2,5]
// Test Case2: [5] 		   left=1 right=1	Output: [5]
// Test Case3: [1,2,3,4,5] left=1 right=5	Output: [5,4,3,2,1]
func main() {
	var left, right int
	fmt.Scan(&left, &right)
	head := pkg.CreateList()
	pkg.PrintList(reverseBetween(head, left, right))
}
