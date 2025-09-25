/*
440.字典序的第K小数字
给定整数 n 和 k，返回  [1, n] 中字典序第 k 小的数字。

示例 1:
输入: n = 13, k = 2
输出: 10
解释: 字典序的排列是 [1, 10, 11, 12, 13, 2, 3, 4, 5, 6, 7, 8, 9]，所以第二小的数字是 10。

示例 2:
输入: n = 1, k = 1
输出: 1

提示:
1 <= k <= n <= 10^9
*/
package main

import (
	"fmt"
)

//type SmallRootQueue440 struct {
//	que []string
//	pos int
//}
//
//func NewPrioriQueue() *SmallRootQueue440 {
//	return &SmallRootQueue440{
//		que: make([]string, 1),
//		pos: 0,
//	}
//}
//func (s *SmallRootQueue440) Len() int { return len(s.que) }
//func (s *SmallRootQueue440) Insert(val string) {
//	s.que = append(s.que, val)
//	s.pos++
//	if s.pos == 1 {
//		return
//	} else {
//		var i = s.pos
//		//上浮
//		for i/2 >= 1 && s.que[i/2] > s.que[i] {
//			s.que[i], s.que[i/2] = s.que[i/2], s.que[i]
//			i = i / 2
//		}
//	}
//}
//
//func (s *SmallRootQueue440) Top() string {
//	return s.que[1]
//}
//func (s *SmallRootQueue440) Pop() {
//	s.que[1], s.que[s.pos] = s.que[s.pos], s.que[1]
//	s.que = s.que[:len(s.que)-1] //移除堆尾元素
//	s.pos--
//	var i = 1
//	for i*2+1 <= s.pos {
//		if s.que[i] < s.que[i*2+1] && s.que[i] < s.que[i*2] {
//			break
//		} else { //将s.que[i]与s.que[i*2],s.que[i*2+1]中的较小者交换
//			if s.que[i*2] < s.que[i*2+1] {
//				s.que[i], s.que[i*2] = s.que[i*2], s.que[i]
//				i = i * 2
//			} else {
//				s.que[i], s.que[i*2+1] = s.que[i*2+1], s.que[i]
//				i = i*2 + 1
//			}
//		}
//	}
//
//	if i*2 <= s.pos { //s.que[i]有左节点无右节点
//		if s.que[i] > s.que[i*2] {
//			s.que[i], s.que[i*2] = s.que[i*2], s.que[i]
//		}
//	}
//}
//func findKthNumber(n int, k int) int {
//	srq := NewPrioriQueue()
//	for i := 1; i <= n; i++ {
//		s := strconv.Itoa(i)
//		srq.Insert(s)
//	}
//	for i := 0; i < k-1; i++ {
//		srq.Pop()
//	}
//	ans, _ := strconv.Atoi(srq.Top())
//	return ans
//}

// 字典树每个节点最多有10个直接子节点
// 返回某个字典树节点的总子节点数
func getNumOfChildren(n int, cur int) (nums int) {
	left, right := cur, cur
	for left <= n {
		nums += min(right, n) - left + 1 //累加当前子节点数
		left *= 10                       //下一层最左子节点
		right = right*10 + 9             //下一层最右子节点
	}
	return
}
func findKthNumber(n int, k int) int {
	var cur int = 1 //当前节点为1
	k--             //已经遍历了当前节点(即字典树的根节点)
	for k > 0 {
		nums := getNumOfChildren(n, cur)
		if nums <= k { //当前节点的总子节点数小于目标数
			k -= nums //更新k，即只需搜寻剩下节点的第(k-nums)小的数字
			cur++     //不在当前节点及其子节点中，转而搜寻其右兄弟节点
		} else {
			cur = cur * 10 //更新当前节点为其最左子节点(记为n)
			k--            //减去当前节点n
		}
	}
	return cur
}

func main() {
	var n, k int
	fmt.Println("Input n:")
	fmt.Scan(&n)
	fmt.Println("Input k:")
	fmt.Scan(&k)
	fmt.Println(findKthNumber(n, k))
}
