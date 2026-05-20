// 148. 排序链表
/*
给你链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。

提示：
链表中节点的数目在范围 [0, 5 * 10^4] 内
-10^5 <= Node.val <= 10^5

进阶：你可以在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序吗？
*/
package main

import (
	"lc/pkg"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *pkg.ListNode) *pkg.ListNode {
	return divideList(head)
}

func divideList(head *pkg.ListNode) *pkg.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	left, right := head, slow.Next
	slow.Next = nil
	leftList := divideList(left)
	rightList := divideList(right)
	return mergeTwoSortedList(leftList, rightList)
}
func mergeTwoSortedList(head1, head2 *pkg.ListNode) *pkg.ListNode {
	dummy := &pkg.ListNode{}
	cur := dummy

	for head1 != nil && head2 != nil {
		if head1.Val < head2.Val {
			cur.Next = head1
			head1 = head1.Next
		} else {
			cur.Next = head2
			head2 = head2.Next
		}
		cur = cur.Next
	}
	if head1 != nil {
		cur.Next = head1
	}
	if head2 != nil {
		cur.Next = head2
	}
	return dummy.Next
}

// 示例 1：
// 输入：head = [4,2,1,3]
// 输出：[1,2,3,4]
//
// 示例 2：
// 输入：head = [-1,5,3,4,0]
// 输出：[-1,0,3,4,5]
//
// 示例 3：
// 输入：head = []
// 输出：[]
func main() {
	listHead := pkg.CreateList()
	pkg.PrintList(sortList(listHead))
}
