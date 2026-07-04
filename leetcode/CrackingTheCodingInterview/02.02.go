// 02.02. 返回倒数第 k 个节点
/*
实现一种算法，找出单向链表中倒数第 k 个节点。返回该节点的值。
注意：本题相对原题稍作改动

说明：
给定的 k 保证是有效的。
*/
package main

import (
	"fmt"
	. "lc/pkg"
)

func kthToLast(head *ListNode, k int) int {
	fast := head
	for i := 0; i < k; i++ {
		fast = fast.Next
	}
	for fast != nil {
		head = head.Next
		fast = fast.Next
	}
	return head.Val
}

// 示例：
// 输入： 1->2->3->4->5 ([1,2,3,4,5]) 和 k = 2
// 输出： 4
func main() {
	var k int
	fmt.Println("Input k:")
	fmt.Scan(&k)
	fmt.Println("Input list:")
	head := CreateList()
	fmt.Println(kthToLast(head, k))
}
