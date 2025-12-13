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

// CreateList 创建一个链表。
func CreateList() *ListNode {
	fmt.Println("Enter list string (e.g. [1, 2, 3]): ")
	nums := CreateSlice[int]()
	if len(nums) == 0 {
		return nil
	}
	head := new(ListNode)
	node := head
	head.Next = node
	for i := 0; i < len(nums); i++ {
		node.Next = new(ListNode)
		node.Next.Val = nums[i]
		node = node.Next
	}
	return head.Next
}
