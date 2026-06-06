// 23. 合并 K 个升序链表
/*
给你一个链表数组，每个链表都已经按升序排列。
请你将所有链表合并到一个升序链表中，返回合并后的链表。

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
	"container/heap"
	"fmt"
	"lc/pkg"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type nodeList []*pkg.ListNode

func (n *nodeList) Len() int {
	return len(*n)
}
func (n *nodeList) Less(childIdx, parentIdx int) bool {
	return (*n)[childIdx].Val < (*n)[parentIdx].Val
}
func (n *nodeList) Swap(i, j int) {
	(*n)[i], (*n)[j] = (*n)[j], (*n)[i]
}
func (n *nodeList) Push(x any) {
	node := x.(*pkg.ListNode)
	*n = append(*n, node)
}
func (n *nodeList) Pop() any {
	top := (*n)[len(*n)-1]
	*n = (*n)[:len(*n)-1]
	return top
}
func mergeKLists(lists []*pkg.ListNode) *pkg.ListNode {
	var nl nodeList
	dummy := &pkg.ListNode{}
	cur := dummy
	heap.Init(&nl)
	for _, list := range lists {
		if list != nil {
			heap.Push(&nl, list)
			list = list.Next

		}
	}
	for nl.Len() > 0 {
		top := heap.Pop(&nl).(*pkg.ListNode)
		if top.Next != nil {
			heap.Push(&nl, top.Next)
		}
		cur.Next = top
		cur = cur.Next
	}

	cur.Next = nil
	return dummy.Next
}

// 示例 1：
// 输入：lists = [[1,4,5],[1,3,4],[2,6]]
// 输出：[1,1,2,3,4,4,5,6]
// 解释：链表数组如下：
// [
//
//	1->4->5,
//	1->3->4,
//	2->6
//
// ]
// 将它们合并到一个有序链表中得到。
// 1->1->2->3->4->4->5->6
//
// 示例 2：
// 输入：lists = []
// 输出：[]
//
// 示例 3：
// 输入：lists = [[]]
// 输出：[]
func main() {
	var n int
	fmt.Println("input n:")
	fmt.Scan(&n)
	lists := make([]*pkg.ListNode, 0)
	for i := 0; i < n; i++ {
		lists = append(lists, pkg.CreateList())
	}
	pkg.PrintList(mergeKLists(lists))
}
