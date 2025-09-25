package pkg

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func PrintList(res *ListNode) {
	for ; res != nil; res = res.Next {
		fmt.Printf("%d ", res.Val)
	}
	println()
}

func CreateList() *ListNode {
	var tmp int
	head := new(ListNode)
	node := head
	head.Next = node
	for {
		fmt.Scan(&tmp)
		if tmp == -1 {
			break
		}
		node.Next = new(ListNode)
		node.Next.Val = tmp
		node = node.Next
	}
	return head.Next
}
