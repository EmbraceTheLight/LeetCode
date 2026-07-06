// 02.04. 分割链表
/*
给你一个链表的头节点 head 和一个特定值 x，请你对链表进行分隔，
使得所有 小于 x 的节点都出现在 大于或等于 x 的节点之前。
你不需要 保留 每个分区中各节点的初始相对位置。

提示：
链表中节点的数目在范围 [0, 200] 内
-100 <= Node.val <= 100
-200 <= x <= 200
*/
package main

import (
	"fmt"
	. "lc/pkg"
)

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}
	nodeList := make([]*ListNode, 0)

	for head != nil {
		nodeList = append(nodeList, head)
		head = head.Next
	}

	i, j := 0, len(nodeList)-1
	for i < j {
		if nodeList[i].Val >= x && nodeList[j].Val < x {
			nodeList[i].Val, nodeList[j].Val = nodeList[j].Val, nodeList[i].Val
		} else if nodeList[i].Val >= x && nodeList[j].Val >= x {
			j--
		} else if nodeList[i].Val < x && nodeList[j].Val >= x {
			i++
			j--
		} else if nodeList[i].Val < x && nodeList[j].Val < x {
			i++
		}
	}
	return nodeList[0]
}

// 示例 1：
// 输入：head = [1,4,3,2,5,2], x = 3
// 输出：[1,2,2,4,3,5]
//
// 示例 2：
// 输入：head = [2,1], x = 2
// 输出：[1,2]
func main() {
	var x int
	fmt.Println("Input x:")
	fmt.Scan(&x)
	fmt.Println("Input List:")
	head := CreateList()
	PrintList(partition(head, x))
}
