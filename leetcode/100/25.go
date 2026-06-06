/*
25. K 个一组翻转链表
给你链表的头节点 head ，每 k 个节点一组进行翻转，请你返回修改后的链表。
k 是一个正整数，它的值小于或等于链表的长度。
如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。
*/
package main

import "fmt"

type ListNode25 struct {
	Val  int
	Next *ListNode25
}

/**
 * Definition for singly-linked list.
 * type ListNode25 struct {
 *     Val int
 *     Next *ListNode25
 * }
 */
func reverseKGroup(head *ListNode25, k int) *ListNode25 {
	if k == 1 {
		return head
	}
	var tail *ListNode25 //待反转链表段的末尾
	var end *ListNode25  //下一组待反转的起始节点
	var slow, fast *ListNode25
	var tmp *ListNode25
	end = head
	slow = head
	var flag bool = true
	ts := slow //保存slow初始状态，为了以后改变其指针指向

	for i := 0; i < k; i++ {
		tail = end
		end = end.Next
	}

	if end == nil {
		fast = ts.Next
		for fast != end {
			tmp = fast.Next
			fast.Next = slow
			slow = fast
			fast = tmp
			if fast == nil {
				break
			}
			tmp = fast.Next
		}
		ts.Next = nil
		return tail
	}

Outer:
	for {
		if flag {
			head = tail
			flag = false
		}
		ts = slow

		fast = ts.Next

		for fast != end {
			tmp = fast.Next
			fast.Next = slow
			slow = fast
			fast = tmp
			if fast == nil {
				break
			}
			tmp = fast.Next
		}

		slow = end
		for i := 0; i < k; i++ {
			tail = end
			if end == nil { //不足k个节点
				break Outer
			}
			end = end.Next

		}
		ts.Next = tail
	}
	ts.Next = fast
	return head
}
func printList25(res *ListNode25) {
	for ; res != nil; res = res.Next {
		fmt.Printf("%d ", res.Val)
	}
}
func main() {
	var k int
	println("Input k:")
	fmt.Scan(&k)

	fmt.Println("Input List1:")
	var tmp int
	head := new(ListNode25)
	node := head
	head.Next = node
	for {
		fmt.Scan(&tmp)
		if tmp == -1 {
			break
		}
		node.Next = new(ListNode25)
		node.Next.Val = tmp
		node = node.Next
	}
	printList25(head.Next)
	fmt.Println()

	printList25(reverseKGroup(head.Next, k))
}
