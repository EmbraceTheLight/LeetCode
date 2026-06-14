// CC10 判断链表中是否有环
/*
描述
判断给定的链表中是否有环。如果有环则返回true，否则返回false。
数据范围: 链表长度 0≤n≤10000，链表中任意节点的值满足 ∣val∣<=100000
要求: 空间复杂度 O(1)，时间复杂度 O(n)

输入分为两部分，第一部分为链表，第二部分代表是否有环，然后将组成的head头结点传入到函数里面。
-1代表无环，其它的数字代表有环，这些参数解释仅仅是为了方便读者自测调试。实际在编程时读入的是链表的头节点。
例如输入{3,2,0,-4},1时，环的入口结点为从头结点开始的第1个结点（注: 头结点为第0个结点），所以输出true。
*/
package main

import (
	"fmt"
	. "lc/pkg"
)

func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}
	if fast == nil || fast.Next == nil {
		return false
	}
	return true
}

// 示例1
// 输入: [3,2,0,-4],1
// 返回值: true
// 说明: 第一部分{3,2,0,-4}代表一个链表，第二部分的1表示，-4到位置1（注: 头结点为位置0），即-4->2存在一个链接，组成传入的head为一个带环的链表，返回true
//
// 示例2
// 输入: [1],-1
// 返回值: false
// 说明:
// 第一部分{1}代表一个链表，-1代表无环，组成传入head为一个无环的单链表，返回false
//
// 示例3
// 输入:
// [-1,-7,7,-4,19,6,-9,-5,-2,-5],6
// 返回值: true
func main() {
	fmt.Println("Input idx:")
	var idx int
	fmt.Scan(&idx)
	list := CreateList()
	tmp := list
	if idx != -1 {
		for ; idx > 0; idx-- {
			tmp = tmp.Next
		}
		tail := list
		for tail.Next != nil {
			tail = tail.Next
		}
		tail.Next = tmp
	}
	fmt.Println(hasCycle(list))
}
