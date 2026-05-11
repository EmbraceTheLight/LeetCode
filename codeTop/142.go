// 142. 环形链表 II
/*
给定一个链表的头节点  head ，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，
评测系统内部使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。
如果 pos 是 -1，则在该链表中没有环。注意：pos 不作为参数进行传递，仅仅是为了标识链表的实际情况。
不允许修改 链表。
*/
package main

import (
	"fmt"
	"lc/pkg"
)

func detectCycle(head *pkg.ListNode) *pkg.ListNode {
	/*
			思路: 快慢指针
			假定无环节点数量为 a, fast 指针与 slow 指针第一次相遇时 slow 指针走了 s1 = a + b 个节点,
			fast 指针走了 s2 = a + b + n(b + c) 个节点, 其中 b + c 为环的长度, n 为 fast 指针在环中走的圈数
			则由于快指针移速是慢指针的两倍, 所以有 s2 = 2s1
			即 a + b + n(b + c) = 2(a + b) ==> n(b + c) = a + b ==> a = (n - 1)b + nc ==> a = c + (n - 1)(b + c)
			由于 b + c 为环的长度, 所以 a = c + (n - 1)(b + c)。
			此时将 fast 指针指向头节点, 然后两个指针以相同的速度移动, 这样一来, fast 走 a 个节点, slow 走 c 个节点,
		  	即使是 n - 1 > 0 的情况, 无非就是 slow 节点从入环点多转了几圈, 最终还是会回到入环点与 fast 相遇
	*/
	slow, fast := head, head
	for {
		if fast == nil || fast.Next == nil {
			return nil
		}
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			fast = head
			break
		}
	}
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return fast
}

// 示例 1：
// 输入：head = [3,2,0,-4], pos = 1
// 输出：返回索引为 1 的链表节点
// 解释：链表中有一个环，其尾部连接到第二个节点。
//
// 示例 2：
// 输入：head = [1,2], pos = 0
// 输出：返回索引为 0 的链表节点
// 解释：链表中有一个环，其尾部连接到第一个节点。
//
// 示例 3：
// 输入：head = [1], pos = -1
// 输出：返回 null
// 解释：链表中没有环。
func main() {
	fmt.Println("Input pos:")
	var pos int
	fmt.Scan(&pos)
	head := pkg.CreateList()
	tmp, cur := head, head
	cnt := 0
	for cur.Next != nil {
		if cnt == pos {
			tmp = cur
		}
		cur = cur.Next
	}
	cur.Next = tmp
	fmt.Println(detectCycle(head) == tmp)
}
