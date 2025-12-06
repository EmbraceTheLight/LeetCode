// 21. 合并两个有序链表
/*
将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

示例 1：
输入：l1 = [1,2,4], l2 = [1,3,4]
输出：[1,1,2,3,4,4]

示例 2：
输入：l1 = [], l2 = []
输出：[]

示例 3：
输入：l1 = [], l2 = [0]
输出：[0]

提示：
两个链表的节点数目范围是 [0, 50]
-100 <= Node.val <= 100
l1 和 l2 均按 非递减顺序 排列
*/
package main

import "lc/pkg"

func mergeTwoLists(list1 *pkg.ListNode, list2 *pkg.ListNode) *pkg.ListNode {
	dummy := &pkg.ListNode{Next: list1}
	preList1 := dummy
	for list1 != nil && list2 != nil {
		var start, end *pkg.ListNode
		if list1.Val > list2.Val {
			start = list2
			for list2 != nil && list1.Val > list2.Val {
				end = list2
				list2 = list2.Next
			}
			preList1.Next = start
			preList1 = end
			end.Next = list1
		} else {
			preList1 = list1
			list1 = list1.Next
		}

	}

	if list2 != nil {
		preList1.Next = list2
	}
	return dummy.Next
}

// Test Case1:	[1,2,4]  [1,3,4]		Output: [1,1,2,3,4,4]
// Test Case2:	[]		[]		Output: []
// Test Case3:	[]		[0]		Output: [0]
// Test Case4:	[1,2,4]	[1,3,3,4,5,6]		Output: [1,1,2,3,3,4,4,5,6,6]
// Test Case4:	[1,3,3,4,5,6] [1,2,4]		Output: [1,1,2,3,3,4,4,5,6,6]
// Test Case5:  [5]	 [1,2,4]				Output: [1,2,4,5]
func main() {
	list1 := pkg.CreateList()
	list2 := pkg.CreateList()
	pkg.PrintList(mergeTwoLists(list1, list2))
}
