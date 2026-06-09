// CC4 链表排序
/*
描述
在O(n log n)的时间内使用常数级空间复杂度对链表进行排序。
*/
package main

import "lc/pkg"

func sortList(head *pkg.ListNode) *pkg.ListNode {
	// write code here
	if head == nil {
		return nil
	}
	return dfscc4(head)
}
func dfscc4(head *pkg.ListNode) *pkg.ListNode {
	if head.Next == nil {
		return head
	}
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	nextHead := slow.Next
	slow.Next = nil
	list1 := dfscc4(head)
	list2 := dfscc4(nextHead)
	return merge2List(list1, list2)
}
func merge2List(head1, head2 *pkg.ListNode) *pkg.ListNode {
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

// 示例1
// 输入: [30,20,40]
// 返回值: [20,30,40]
func main() {
	head := pkg.CreateList()
	pkg.PrintList(sortList(head))
}
