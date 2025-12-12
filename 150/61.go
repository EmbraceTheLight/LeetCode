// 61. 旋转链表
/*
给你一个链表的头节点 head ，旋转链表，将链表每个节点向右移动 k 个位置。
示例 1：
输入：head = [1,2,3,4,5], k = 2
输出：[4,5,1,2,3]

示例 2：
输入：head = [0,1,2], k = 4
输出：[2,0,1]

提示：
链表中节点的数目在范围 [0, 500] 内
-100 <= Node.val <= 100
0 <= k <= 2 * 10^9
*/
package main

import (
	"fmt"
	"lc/pkg"
)

func rotateRight(head *pkg.ListNode, k int) *pkg.ListNode {
	count := 0
	if head == nil {
		return head
	}
	for tmp := head; tmp != nil; tmp = tmp.Next {
		count++
	}

	k = k % count
	if k == 0 {
		return head
	}

	dummy := &pkg.ListNode{Next: head}
	newHead, newTail := findNewHeadNewTail(head, k)

	// 处理指针关系, 注意处理顺序
	var tmp *pkg.ListNode
	for tmp = newHead; tmp.Next != nil; tmp = tmp.Next {
	}

	// 先处理新的头部链表的尾指针, 使其指向原链表的头部
	tmp.Next = dummy.Next

	// 处理 dummy, 令其指向新的头部
	dummy.Next = newHead

	// 处理新的尾部链表的尾指针, 使其指向 nil
	newTail.Next = nil

	return dummy.Next
}

// findNewHeadNewTail 找原链表的倒数第 k 个节点, 返回它和它前面的节点, 它本身就是旋转链表后新链表的头节点, 而它的前驱节点就是旋转链表后新链表的尾节点
func findNewHeadNewTail(head *pkg.ListNode, k int) (*pkg.ListNode, *pkg.ListNode) {
	slow, fast := head, head
	var preSlow *pkg.ListNode
	for i := 0; i < k; i++ {
		fast = fast.Next
	}

	for fast != nil {
		preSlow = slow
		slow = slow.Next
		fast = fast.Next
	}

	return slow, preSlow
}

// Test Case1:	[1,2,3,4,5] k = 2	Output: [4,5,1,2,3]
// Test Case2:	[0,1,2] k = 4	Output: [2,0,1]
func main() {
	var k int
	fmt.Println("input k:")
	fmt.Scan(&k)

	head := pkg.CreateList()
	pkg.PrintList(rotateRight(head, k))
}
