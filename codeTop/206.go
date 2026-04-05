// 206. 反转链表
/*
给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。


提示：
链表中节点的数目范围是 [0, 5000]
-5000 <= Node.val <= 5000
*/
package main

import "lc/pkg"

func reverseList(head *pkg.ListNode) *pkg.ListNode {
	if head == nil {
		return nil
	}
	var first, second, third *pkg.ListNode
	first = head
	second = first.Next
	for second != nil {
		third = second.Next
		second.Next = first
		first = second
		second = third
	}
	head.Next = nil
	return first
}

// 示例 1：
// 输入：head = [1,2,3,4,5]
// 输出：[5,4,3,2,1]
//
// 示例 2：
// 输入：head = [1,2]
// 输出：[2,1]
//
// 示例 3：
// 输入：head = []
// 输出：[]
func main() {
	list := pkg.CreateList()
	reverseHead := reverseList(list)
	pkg.PrintList(reverseHead)

}
