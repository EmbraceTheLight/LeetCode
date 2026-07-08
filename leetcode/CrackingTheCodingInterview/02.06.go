// 面试题 02.06. 回文链表
/*
编写一个函数，检查输入的链表是否是回文的。

进阶：
你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？
*/
package main

import (
	"fmt"
	. "lc/pkg"
	"strings"
)

func isPalindrome(head *ListNode) bool {
	var sb1 strings.Builder
	var sb2 strings.Builder
	for cur := head; cur != nil; cur = cur.Next {
		sb1.WriteByte(byte(cur.Val + '0'))
	}
	newHead := reverseList(head)
	for cur := newHead; cur != nil; cur = cur.Next {
		sb2.WriteByte(byte(cur.Val + '0'))
	}
	return sb1.String() == sb2.String()
}
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var pre, next *ListNode
	cur := head
	for cur != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

// 示例 1：
// 输入： [1,2](1->2)
// 输出： false
//
// 示例 2：
// 输入： [1,2,2,1](1->2->2->1)
// 输出： true
func main() {
	head := CreateList()
	fmt.Println(isPalindrome(head))
}
