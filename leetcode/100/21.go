/*
21. 合并两个有序链表
将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
*/
package main

import (
	"fmt"
)

type ListNode21 struct {
	Val  int
	Next *ListNode21
}

func mergeTwoLists(list1 *ListNode21, list2 *ListNode21) *ListNode21 {
	switch {
	case list1 == nil && list2 == nil:
		return nil
	case list1 == nil && list2 != nil:
		return list2
	case list1 != nil && list2 == nil:
		return list1
	}
	head1 := list1
	pre1 := head1
	tmp2 := list2

	//处理list1的第一个节点，为了使得pre1在list1前一个节点
	if list2.Val < list1.Val {
		tmp2 = list2.Next //保存list2的下一个位置，因为要更改它的指向
		list2.Next = list1
		pre1 = list2
		head1 = list2
		list2 = tmp2
	} else {
		list1 = list1.Next
	}
	for list1 != nil && list2 != nil {
		if list1 != nil && list1.Val <= list2.Val {
			pre1 = list1
			list1 = list1.Next
		} else if list2 != nil && list1.Val > list2.Val {
			tmp2 = list2.Next
			list2.Next = list1
			pre1.Next = list2
			pre1 = list2 //使得pre1一直是list1的前一个节点
			list2 = tmp2
		}
	}
	if list2 != nil { //此时，list1==nil,pre1为list1尾结点
		pre1.Next = list2
	}
	return head1
}
func (l *ListNode21) Insert(val int) (tail *ListNode21) {
	node := new(ListNode21)
	node.Val = val
	l.Next = node
	tail = node
	return
}
func printList(res *ListNode21) {
	for ; res != nil; res = res.Next {
		fmt.Printf("%d ", res.Val)
	}
}
func main() {
	fmt.Println("Input List1:")
	var tmp int
	head := new(ListNode21)
	fmt.Scan(&tmp)
	head.Val = tmp
	next := head
	for {
		fmt.Scan(&tmp)
		if tmp == -1 {
			break
		}
		next = next.Insert(tmp)
	}
	printList(head)
	fmt.Println()

	fmt.Println("Input List2:")
	head2 := new(ListNode21)
	fmt.Scan(&tmp)
	head2.Val = tmp
	next = head2
	for {
		fmt.Scan(&tmp)
		if tmp == -1 {
			break
		}
		next = next.Insert(tmp)
	}
	printList(head2)
	fmt.Println()

	printList(mergeTwoLists(head, head2))

}
