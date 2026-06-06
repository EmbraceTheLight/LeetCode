// 23. 合并 K 个升序链表
/*
给你一个链表数组，每个链表都已经按升序排列。
请你将所有链表合并到一个升序链表中，返回合并后的链表。

示例 1：
输入：lists = [[1,4,5],[1,3,4],[2,6]]
输出：[1,1,2,3,4,4,5,6]
解释：链表数组如下：
[
  1->4->5,
  1->3->4,
  2->6
]
将它们合并到一个有序链表中得到。
1->1->2->3->4->4->5->6

示例 2：
输入：lists = []
输出：[]

示例 3：
输入：lists = [[]]
输出：[]

提示：
k == lists.length
0 <= k <= 10^4
0 <= lists[i].length <= 500
-10^4 <= lists[i][j] <= 10^4
lists[i] 按 升序 排列
lists[i].length 的总和不超过 10^4
*/
package main

import (
	"lc/pkg"
)

func mergeKLists(lists []*pkg.ListNode) *pkg.ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists)%2 != 0 {
		lists = append(lists, nil)
	}
	return divide23(lists)

}
func divide23(lists []*pkg.ListNode) *pkg.ListNode {

	if len(lists) == 1 {
		return lists[0]
	}
	mid := len(lists) / 2
	list1 := divide23(lists[:mid])
	list2 := divide23(lists[mid:])
	return mergeTwoList23(list1, list2)
}
func mergeTwoList23(list1, list2 *pkg.ListNode) *pkg.ListNode {
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
	} else if list2 != nil {
		cur.Next = list2
	}
	return dummy.Next
}

// Test Case1: [[1,4,5],[1,3,4],[2,6]]	Output: [1,1,2,3,4,4,5,6]
// Test Case2: []	Output: []
func main() {
	input := pkg.CreateSlice2D[int]()
	var lists []*pkg.ListNode
	for _, nums := range input {
		lists = append(lists, makeList(nums))
	}
	ans := mergeKLists(lists)
	pkg.PrintList(ans)
}
func makeList(nums []int) *pkg.ListNode {
	dummy := &pkg.ListNode{}
	cur := dummy
	for _, n := range nums {
		cur.Next = &pkg.ListNode{
			Val: n,
		}
		cur = cur.Next
	}
	return dummy.Next
}
