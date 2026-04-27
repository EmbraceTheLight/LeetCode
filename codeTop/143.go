// 143. 重排链表
/*
给定一个单链表 L 的头节点 head ，单链表 L 表示为：
L0 → L1 → … → Ln - 1 → Ln

请将其重新排列后变为：
L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

提示：
链表的长度范围为 [1, 5 * 10^4]
1 <= node.val <= 1000
*/
package main

import "lc/pkg"

func reorderList(head *pkg.ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	mid, fast := head, head
	for fast != nil && fast.Next != nil {
		mid = mid.Next
		fast = fast.Next.Next
	}
	stk := make([]*pkg.ListNode, 0)
	for mid != nil {
		stk = append(stk, mid)
		mid = mid.Next
	}
	cur := head
	next := cur.Next
	for len(stk) > 0 {
		top := stk[len(stk)-1]
		stk = stk[:len(stk)-1]
		top.Next = next
		cur.Next = top
		cur = next
		next = cur.Next
	}
	if next != nil {
		next.Next = nil
	}
}

// 示例 1：
// 输入：head = [1,2,3,4]
// 输出：[1,4,2,3]
//
// 示例 2：
// 输入：head = [1,2,3,4,5]
// 输出：[1,5,2,4,3]
func main() {
	list := pkg.CreateList()
	reorderList(list)
	pkg.PrintList(list)
}
