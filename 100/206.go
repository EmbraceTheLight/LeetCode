/*
206. 反转链表
链表中节点的数目范围是 [0, 5000]
-5000 <= Node.val <= 5000
*/
package main

import "fmt"

type ListNode206 struct {
	Val  int
	Next *ListNode206
}

func (l *ListNode206) Insert(val int) (tail *ListNode206) {
	node := new(ListNode206)
	node.Val = val
	l.Next = node
	tail = node
	return
}

func reverseList(head *ListNode206) *ListNode206 {
	var slow, fast, tmp *ListNode206
	slow = nil
	fast = head
	for fast != nil {
		tmp = fast.Next
		fast.Next = slow
		slow = fast
		fast = tmp
	}
	return slow
}
func main() {
	fmt.Println("Input Value:")
	var tmp int
	head := new(ListNode206)
	next := head
	for {
		fmt.Scan(&tmp)
		if tmp == -1 {
			break
		}
		next = next.Insert(tmp)
	}

	for p := reverseList(head); p != nil; p = p.Next {
		fmt.Printf("%d ", p.Val)
	}
	fmt.Println()
}
