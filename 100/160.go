/*
160. 相交链表
给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表不存在相交节点，返回 null 。
listA 中节点数目为 m
listB 中节点数目为 n
1 <= m, n <= 3 * 10^4
1 <= Node.val <= 10^5
0 <= skipA <= m
0 <= skipB <= n
如果 listA 和 listB 没有交点，intersectVal 为 0
如果 listA 和 listB 有交点，intersectVal == listA[skipA] == listB[skipB]
*/
package main

type ListNode160 struct {
	Val  int
	Next *ListNode160
}

//	func getIntersectionNode(headA, headB *ListNode160) *ListNode160 {
//		ha := headA
//		hb := headB
//		mp := make(map[int][]*ListNode160)
//		for ha = headA; ha != nil; ha = ha.Next {
//			mp[ha.Val] = append(mp[ha.Val], ha)
//		}
//		for hb = headB; hb != nil; hb = hb.Next {
//			if sli, ok := mp[hb.Val]; ok {
//				for _, v := range sli {
//					if hb == v {
//						return hb
//					}
//				}
//			}
//		}
//		return nil
//	}

// 让ha，hb从各自链表开始，走到nil就从另一条链表开始走，消除了链表之间的长度差
// 这样无论链表是否有相交，ha、hb最终一定会指向同一个地方

func getIntersectionNode(headA, headB *ListNode160) *ListNode160 {
	ha := headA
	hb := headB
	for {
		if ha == hb { //先判断，因为两个链表若无相交，一定会同时指向nil
			return ha
		}

		ha = ha.Next
		hb = hb.Next

		if ha == nil {
			ha = headB
		}
		if hb == nil {
			hb = headA
		}

	}
}
func main() {

}
