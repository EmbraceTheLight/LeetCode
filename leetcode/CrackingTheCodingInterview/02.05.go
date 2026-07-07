// 02.05. 链表求和
/*
给定两个用链表表示的整数，每个节点包含一个数位。
这些数位是反向存放的，也就是个位排在链表首部。
编写函数对这两个整数求和，并用链表形式返回结果。
*/
package main

import (
	"fmt"
	. "lc/pkg"
)

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1Len, l2Len := 0, 0
	for cur := l1; cur != nil; cur = cur.Next {
		l1Len++
	}
	for cur := l2; cur != nil; cur = cur.Next {
		l2Len++
	}
	cur := l1
	if l1Len < l2Len {
		cur = l2
	}
	cur1, cur2 := l1, l2
	pre := cur
	carry := 0
	for cur1 != nil && cur2 != nil {
		sum := cur1.Val + cur2.Val + carry
		carry = sum / 10
		cur.Val = sum % 10
		pre = cur
		cur = cur.Next
		cur1 = cur1.Next
		cur2 = cur2.Next
	}

	for cur != nil {
		sum := cur.Val + carry
		carry = sum / 10
		cur.Val = sum % 10
		pre = cur
		cur = cur.Next
	}
	if carry == 1 {
		pre.Next = &ListNode{Val: 1}
	}
	if l1Len < l2Len {
		return l2
	}
	return l1
}

// 示例1：
// 输入：[7,1,6](7 -> 1 -> 6) + [5,9,2](5 -> 9 -> 2)，即617 + 295
// 输出：2 -> 1 -> 9，即912
//
// 进阶：思考一下，假设这些数位是正向存放的，又该如何解决呢?
// 示例：
// 输入：[6,1,7](6 -> 1 -> 7) + [2,9,5](2 -> 9 -> 5)，即617 + 295
// 输出：9 -> 1 -> 2，即912
func main() {
	fmt.Println("Input L1:")
	l1 := CreateList()
	fmt.Println("Input L2:")
	l2 := CreateList()
	PrintList(addTwoNumbers(l1, l2))
}
