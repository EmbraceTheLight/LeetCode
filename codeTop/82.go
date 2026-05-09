// 82. 删除排序链表中的重复元素 II
/*
给定一个已排序的链表的头 head ， 删除原始链表中所有重复数字的节点，只留下不同的数字 。返回 已排序的链表 。

提示：
链表中节点数目在范围 [0, 300] 内
-100 <= Node.val <= 100
题目数据保证链表已经按升序 排列
*/
package main

import "lc/pkg"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *pkg.ListNode) *pkg.ListNode {
	dummy := &pkg.ListNode{Next: head}
	pre, cur := dummy, head
	for cur != nil {
		if cur.Next != nil && cur.Val == cur.Next.Val {
			for cur.Next != nil && cur.Val == cur.Next.Val {
				cur = cur.Next
			}
			pre.Next = cur.Next
		} else {
			pre = cur
		}

		cur = cur.Next
	}
	return dummy.Next
}

// 示例 1：
// 输入：head = [1,2,3,3,4,4,5]
// 输出：[1,2,5]
//
// 示例 2：
// 输入：head = [1,1,1,2,3]
// 输出：[2,3]
func main() {
	head := pkg.CreateList()
	pkg.PrintList(deleteDuplicates(head))
}
