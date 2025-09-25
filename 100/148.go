/*
148. 排序链表
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

import (
	"fmt"
)

type ListNode148 struct {
	Val  int
	Next *ListNode148
}

func mergeList(head1, head2 *ListNode148) *ListNode148 {
	var dummy = &ListNode148{}
	var cur = dummy
	for head1 != nil && head2 != nil {
		if head1.Val <= head2.Val {
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

// 自底向上合并链表，无需递归。使得空间复杂度O(1)
func sortList(head *ListNode148) *ListNode148 {
	var count = 0
	for cur := head; cur != nil; cur = cur.Next {
		count++
	}
	ans := &ListNode148{Next: head}

	for sublistLen := 1; sublistLen < count; sublistLen = sublistLen * 2 {
		var pre, cur = ans, ans.Next
		for cur != nil {
			var first, tmpf = cur, cur
			for i := 1; i < sublistLen && tmpf.Next != nil; i++ {
				tmpf = tmpf.Next //确定第一段链表终点
			}
			var second, tmps = tmpf.Next, tmpf.Next
			tmpf.Next = nil //将第一段链表尾部置空

			for i := 1; i < sublistLen && tmps != nil && tmps.Next != nil; i++ {
				tmps = tmps.Next //确定第二段链表终点
			}
			if tmps != nil {
				cur = tmps.Next //保存第二段链表后的头部
				tmps.Next = nil //将第一段链表尾部置空
			} else {
				cur = nil
			}

			pre.Next = mergeList(first, second)
			for pre.Next != nil {
				pre = pre.Next
			}
		}

	}
	return ans.Next
}
func main() {
	var sli1 = []int{-1, 5, 3, 4, 0}
	var preHead1 = &ListNode148{}
	var cur = preHead1
	for _, v := range sli1 {
		cur.Next = new(ListNode148)
		cur = cur.Next
		cur.Val = v
	}
	for iter := sortList(preHead1.Next); iter != nil; iter = iter.Next {
		fmt.Printf("%d ", iter.Val)
	}
	fmt.Println()

	var sli2 = []int{4, 2, 1, 3}
	var preHead2 = &ListNode148{}
	cur = preHead2
	for _, v := range sli2 {
		cur.Next = new(ListNode148)
		cur = cur.Next
		cur.Val = v
	}
	for iter := sortList(preHead2.Next); iter != nil; iter = iter.Next {
		fmt.Printf("%d ", iter.Val)
	}
	fmt.Println()
}
