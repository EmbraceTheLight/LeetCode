/*
1144. 二叉树的前序遍历
给你二叉树的根节点 root ，返回它节点值的 前序 遍历。
*/

package main

import (
	"fmt"
	"lc/100/pkg"
)

type stack struct {
	sli []*pkg.TreeNode
	top *pkg.TreeNode //指向栈顶,即切片尾部
	len int
}

func (s *stack) Len() int { return s.len }
func (s *stack) Pop() *pkg.TreeNode {
	ret := s.top
	s.sli = s.sli[:s.len-1]
	s.len--
	s.top = s.sli[s.len-1]
	return ret
}
func (s *stack) Push(node *pkg.TreeNode) {
	s.sli = append(s.sli, node)
	s.top = node
	s.len++
}

func inorderTraversal(root *pkg.TreeNode) []int {
	ret := make([]int, 0)
	var s stack

	node := root
	s.Push(root)
	for s.Len() > 0 {
		ret = append(ret, node.Val)
		for node.Left != nil {
			s.Push(node.Left)
			ret = append(ret, node.Left.Val)
			node = node.Left
		}

		for s.Len() > 0 {
			n := s.Pop()
			if n.Right != nil {
				ret = append(ret, n.Right.Val)
			}
		}
	}
	return ret
}
func main()
	head := pkg.CreateTree()
	fmt.Println(inorderTraversal(head))
}
