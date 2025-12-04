// 141. 环形链表
/*
给你一个链表的头节点 head ，判断链表中是否有环。
如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。
为了表示给定链表中的环，评测系统内部使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。注意：pos 不作为参数进行传递 。仅仅是为了标识链表的实际情况。
如果链表中存在环 ，则返回 true 。 否则，返回 false 。

示例 1：
输入：head = [3,2,0,-4], pos = 1
输出：true
解释：链表中有一个环，其尾部连接到第二个节点。

示例 2：
输入：head = [1,2], pos = 0
输出：true
解释：链表中有一个环，其尾部连接到第一个节点。

示例 3：
输入：head = [1], pos = -1
输出：false
解释：链表中没有环。

提示：
链表中节点的数目范围是 [0, 104]
-10^5 <= Node.val <= 10^5
pos 为 -1 或者链表中的一个 有效索引 。

进阶：你能用 O(1)（即，常量）内存解决此问题吗？
*/
package main

import (
	"fmt"
	"lc/pkg"
)

func hasCycle(head *pkg.ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			break
		}
	}
	if fast == slow {
		return true
	}
	return false
}

// Test Case1: 1 	[3,2,0,-4]		Output: true
// Test Case2: 0 	[1,2]			Output: true
// Test Case3: -1 	[1]				Output: false
func main() {
	var pos int
	fmt.Println("Enter positive integer: ")
	fmt.Scan(&pos)

	fmt.Println("Enter a int slice to make a list")
	head := pkg.CreateList()

	// 制作环路
	if pos >= 0 {
		var posNode, tail *pkg.ListNode
		tmp := head
		count := 0
		for tmp != nil {
			if count == pos {
				posNode = tmp
			}
			tail = tmp
			tmp = tmp.Next
		}

		tail.Next = posNode
	}

	fmt.Println(hasCycle(head))
}
