// 21. 合并两个有序链表
/*
将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

提示：
两个链表的节点数目范围是 [0, 50]
-100 <= Node.val <= 100
l1 和 l2 均按 非递减顺序 排列
*/
package main

import "lc/pkg"

// 示例 1：
// 输入：l1 = [1,2,4], l2 = [1,3,4]
// 输出：[1,1,2,3,4,4]
//
// 示例 2：
// 输入：l1 = [], l2 = []
// 输出：[]
//
// 示例 3：
// 输入：l1 = [], l2 = [0]
// 输出：[0]
func mergeTwoLists(list1 *pkg.ListNode, list2 *pkg.ListNode) *pkg.ListNode {
	dummy := &pkg.ListNode{}
	cur := dummy
	curList1, curList2 := list1, list2
	for curList1 != nil && curList2 != nil {
		if curList1.Val < curList2.Val {
			cur.Next = curList1
			curList1 = curList1.Next
		} else {
			cur.Next = curList2
			curList2 = curList2.Next
		}
		cur = cur.Next
	}
	if curList1 != nil {
		cur.Next = curList1
	}
	if curList2 != nil {
		cur.Next = curList2
	}
	return dummy.Next
}
func main() {
	list1 := pkg.CreateList()
	list2 := pkg.CreateList()
	pkg.PrintList(mergeTwoLists(list1, list2))
}
