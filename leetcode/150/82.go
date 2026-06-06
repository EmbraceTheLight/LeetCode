// 82. 删除排序链表中的重复元素 II
/*
给定一个已排序的链表的头 head，删除原始链表中所有重复数字的节点，只留下不同的数字 。返回已排序的链表。

示例 1：
输入：head = [1,2,3,3,4,4,5]
输出：[1,2,5]

示例 2：
输入：head = [1,1,1,2,3]
输出：[2,3]

提示：
链表中节点数目在范围 [0, 300] 内
-100 <= Node.val <= 100
题目数据保证链表已经按升序 排列
*/
package main

import "lc/pkg"

func deleteDuplicates(head *pkg.ListNode) *pkg.ListNode {
	dummy := &pkg.ListNode{}
	cur := dummy
	slow, fast := head, head
	for slow != nil {
		for fast.Next != nil && fast.Next.Val == fast.Val {
			fast = fast.Next
		}

		fast = fast.Next

		if slow.Next != fast {
			slow = fast
		} else {
			cur.Next = slow
			cur = slow
			cur.Next = nil // 防止最后几个节点是重复元素的情况, 这种情况不应该保留这些节点
			slow = fast
		}
	}
	return dummy.Next
}

// Test Case1:	[1,2,3,3,4,4,5]		Output: [1,2,5]
// Test Case2:	[1,1,1,2,3]			Output: [2,3]
func main() {
	head := pkg.CreateList()
	pkg.PrintList(deleteDuplicates(head))
}
