// CC5 链表的插入排序
/*
描述
使用插入排序对链表进行排序。
*/
package main

import (
	. "lc/pkg"
	"math"
)

func insertionSortList(head *ListNode) *ListNode {
	// write code here
	if head == nil {
		return nil
	}
	dummy := &ListNode{Val: math.MinInt32}
	newHead := dummy
	cur := head
	for cur != nil {
		tmp := newHead
		pre := tmp
		for tmp != nil {
			if tmp.Val > cur.Val {
				break
			}
			pre = tmp
			tmp = tmp.Next
		}
		newNode := &ListNode{Val: cur.Val, Next: tmp}
		pre.Next = newNode
		cur = cur.Next
	}
	return dummy.Next
}

// 示例1
// 输入: [30,20,40]
// 返回值: [20,30,40]
func main() {
	head := CreateList()
	PrintList(insertionSortList(head))
}
