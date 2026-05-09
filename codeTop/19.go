// 19. 删除链表的倒数第 N 个结点
/*
给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。

提示：
链表中结点的数目为 sz
1 <= sz <= 30
0 <= Node.val <= 100
1 <= n <= sz


进阶：你能尝试使用一趟扫描实现吗？
*/
package main

import (
	"fmt"
	"lc/pkg"
)

func removeNthFromEnd(head *pkg.ListNode, n int) *pkg.ListNode {
	dummy := &pkg.ListNode{Next: head}
	slow, fast := dummy, head
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}

// 示例 1：
// 输入：head = [1,2,3,4,5], n = 2
// 输出：[1,2,3,5]
//
// 示例 2：
// 输入：head = [1], n = 1
// 输出：[]
//
// 示例 3：
// 输入：head = [1,2], n = 1
// 输出：[1]
func main() {
	fmt.Println("Input n:")
	var n int
	fmt.Scan(&n)
	head := pkg.CreateList()
	pkg.PrintList(removeNthFromEnd(head, n))
}
