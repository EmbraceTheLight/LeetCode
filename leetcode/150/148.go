// 148. 排序链表
/*
给你链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。

示例 1：
输入：head = [4,2,1,3]
输出：[1,2,3,4]

示例 2：
输入：head = [-1,5,3,4,0]
输出：[-1,0,3,4,5]

示例 3：
输入：head = []
输出：[]

提示：
链表中节点的数目在范围 [0, 5 * 10^4] 内
-10^5 <= Node.val <= 10^5
进阶：你可以在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序吗？
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
func sortList(head *pkg.ListNode) *pkg.ListNode {
	if head == nil {
		return nil
	}
	end := head
	for end.Next != nil {
		end = end.Next
	}
	return divide(head, end)
}

// 分治 + 合并两个有序链表
// 只有一个元素的链表认为是有序的
func divide(start, end *pkg.ListNode) *pkg.ListNode {
	if start.Next == nil {
		return start
	}
	slow, fast := start, start
	for fast != end && fast.Next != end {
		slow = slow.Next
		fast = fast.Next.Next
	}
	lEnd, rStart := slow, slow.Next
	slow.Next = nil
	list1 := divide(start, lEnd)
	list2 := divide(rStart, end)
	return mergeTwoSortedList(list1, list2)
}
func mergeTwoSortedList(list1, list2 *pkg.ListNode) *pkg.ListNode {
	dummy := &pkg.ListNode{}
	cur := dummy
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			cur.Next = list1
			list1 = list1.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
		}
		cur = cur.Next
	}
	if list1 != nil {
		cur.Next = list1
	}
	if list2 != nil {
		cur.Next = list2
	}
	return dummy.Next
}

// Test Case1:	[4,2,1,3]	Output:	[1,2,3,4]
// Test Case2:	[-1,5,3,4,0]	Output:	[-1,0,3,4,5]
// Test Case3:	[]	Output:	[]
func main() {
	head := pkg.CreateList()
	pkg.PrintList(sortList(head))
}
