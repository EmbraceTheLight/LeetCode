// 86. 分隔链表
/*
给你一个链表的头节点 head 和一个特定值 x ，请你对链表进行分隔，使得所有 小于 x 的节点都出现在 大于或等于 x 的节点之前。
你应当 保留 两个分区中每个节点的初始相对位置。

示例 1：
输入：head = [1,4,3,2,5,2], x = 3
输出：[1,2,2,4,3,5]

示例 2：
输入：head = [2,1], x = 2
输出：[1,2]


提示：
链表中节点的数目在范围 [0, 200] 内
-100 <= Node.val <= 100
-200 <= x <= 200
*/
package main

import (
	"fmt"
	"lc/pkg"
)

// 创建子链表, 不会增加空间复杂度, 这一点可以关注一下
func partition(head *pkg.ListNode, x int) *pkg.ListNode {
	left := &pkg.ListNode{}
	right := &pkg.ListNode{}

	curLeft := left
	curRight := right
	for ; head != nil; head = head.Next {
		if head.Val < x {
			curLeft.Next = head
			curLeft = curLeft.Next
		} else {
			curRight.Next = head
			curRight = curRight.Next
		}
	}
	curLeft.Next = right.Next
	curRight.Next = nil
	return left.Next
}

// Test Case1:	[1,4,3,2,5,2] x=3	Output:[1,2,2,4,3,5]
// Test Case2:	[2,1] x=2	Output:[1,2]
func main() {
	var x int
	fmt.Println("Input x:")
	fmt.Scan(&x)
	head := pkg.CreateList()
	pkg.PrintList(partition(head, x))
}
