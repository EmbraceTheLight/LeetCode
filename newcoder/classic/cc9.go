// CC9 链表中环的入口结点
/*
给一个长度为n链表，若其中包含环，请找出该链表的环的入口结点，否则，返回null。
数据范围:n≤10000 1<=结点值<=10000
要求: 空间复杂度O(1) 时间复杂度 O(n)

例如，输入 [1,2],[3,4,5]时
输入描述:
输入分为2段，第一段是入环前的链表部分，第二段是链表环的部分，后台会根据第二段是否为空将这两段组装成一个无环或者有环单链表
返回值描述:
返回链表的环的入口结点即可，我们后台程序会打印这个结点对应的结点值；若没有，则返回对应编程语言的空结点即可。

说明: 环的部分只有一个结点，所以返回该环形链表入口结点，后台程序打印该结点对应的结点值，即2
*/
package main

import (
	"fmt"
	. "lc/pkg"
)

func EntryNodeOfLoop(pHead *ListNode) *ListNode {
	slow, fast := pHead, pHead
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}
	if fast == nil || fast.Next == nil {
		return nil
	}
	tmp := pHead
	for tmp != slow {
		tmp = tmp.Next
		slow = slow.Next
	}
	return tmp
}

// 示例1
// 输入: [1,2], [3,4,5]
// 返回值: 3
// 说明: 返回环形链表入口结点，我们后台程序会打印该环形链表入口结点对应的结点值，即3
//
// 示例2
// 输入: [1],[]
// 返回值: "null"
// 说明: 没有环，返回对应编程语言的空结点，后台程序会打印"null"
//
// 示例3
// 输入: [],[2]
// 返回值: 2
func main() {
	list1 := CreateList()
	list2 := CreateList()
	cur := list1
	for cur.Next != nil {
		cur = cur.Next
	}

	tail2 := list2
	for tail2.Next != nil {
		tail2 = tail2.Next
	}
	tail2.Next = list2
	cur.Next = list2
	fmt.Println(EntryNodeOfLoop(list1).Val)
}
