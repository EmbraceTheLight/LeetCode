// 02.01. 移除重复节点
/*
编写代码，移除未排序链表中的重复节点。保留最开始出现的节点。

提示：
链表长度在[0, 20000]范围内。
链表元素在[0, 20000]范围内。

进阶：
如果不得使用临时缓冲区，该怎么解决？
*/
package main

import . "lc/pkg"

func removeDuplicateNodes(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode
	mp := make(map[int]bool)
	for cur != nil {
		if mp[cur.Val] {
			tmp := cur
			for tmp != nil && mp[tmp.Val] == true {
				tmp = tmp.Next
			}
			pre.Next = tmp
		}
		mp[cur.Val] = true
		pre = cur
		cur = cur.Next
	}
	return head
}

// 示例1：
//
//	输入：[1, 2, 3, 3, 2, 1]
//	输出：[1, 2, 3]
//
// 示例2：
//
//	输入：[1, 1, 1, 1, 2]
//	输出：[1, 2]
func main() {
	head := CreateList()
	PrintList(removeDuplicateNodes(head))
}
