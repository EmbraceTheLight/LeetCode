/*
23.合并 K 个升序链表
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

type ListNode23 struct {
	Val  int
	Next *ListNode23
}

func merge23(head1, head2 *ListNode23) *ListNode23 {
	var dummy = &ListNode23{}
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

func mergeKLists(lists []*ListNode23) *ListNode23 {
	for i := 0; i < len(lists)-1; i += 2 {
		lists = append(lists, merge23(lists[i], lists[i+1]))
	}
	return lists[len(lists)-1]
}
func main() {

}
