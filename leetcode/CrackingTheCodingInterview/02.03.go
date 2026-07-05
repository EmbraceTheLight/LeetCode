// 面试题 02.03. 删除中间节点
/*
若链表中的某个节点，既不是链表头节点，也不是链表尾节点，则称其为该链表的「中间节点」。
假定已知链表的某一个中间节点，请实现一种算法，将该节点从链表中删除。
例如，传入节点 c（位于单向链表 a->b->c->d->e->f 中），将其删除后，剩余链表为 a->b->d->e->f
*/
package main

import (
	"fmt"
	. "lc/pkg"
)

func deleteNode(node *ListNode) {
	cur := node
	pre := cur
	next := node.Next
	for next != nil {
		cur.Val = next.Val
		pre = cur
		cur = cur.Next
		next = next.Next
	}
	pre.Next = nil
}

// 最优解法: O(1)
func deleteNode2(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

// 示例：
// 输入：节点 5 （位于单向链表 4->5->1->9 ([4, 5, 1, 9]) 中）
// 输出：不返回任何数据，从链表中删除传入的节点 5，使链表变为 4->1->9
func main() {
	fmt.Println("Input target node val:")
	var target int
	fmt.Scan(&target)
	fmt.Println("Input list:")
	head := CreateList()
	cur := head
	for ; cur != nil; cur = cur.Next {
		if cur.Val == target {
			break
		}
	}
	deleteNode(cur)
	PrintList(head)
}
