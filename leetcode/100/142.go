/*
142. 环形链表 II
给定一个链表的头节点  head ，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。
为了表示给定链表中的环，评测系统内部使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。
如果 pos 是 -1，则在该链表中没有环。注意：pos 不作为参数进行传递，仅仅是为了标识链表的实际情况。
不允许修改 链表。
*/
package main

type ListNode142 struct {
	Val  int
	Next *ListNode142
}

func detectCycle(head *ListNode142) *ListNode142 {
	slow := head
	fast := head
	if head == nil || head.Next == nil {
		return nil
	}
	for {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == nil || fast.Next == nil {
			return nil
		}
		//两指针第一次相遇时设slow走了x个节点，在环中转了y圈，环有z个节点（含入口节点），除环之外有a个节点
		//此时fast走了s个节点，
		//x=a
		//有s=2x
		//s=x+y*z（因为fast和slow都走了a个节点，fast多出的部分为多绕的圈数）
		//故a=x=y*z，s=a+y*z
		//slow已经走了y*z个节点
		//
		if fast == slow {
			fast = head
			break
		}
	}
	for fast != slow { //fast步长改为1
		fast = fast.Next
		slow = slow.Next
	}
	return fast
}
func main() {

}
