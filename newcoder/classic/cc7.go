// CC7 求二叉树的前序遍历
/*
描述
求给定的二叉树的前序遍历。
例如：
给定的二叉树为 [1,null,2,3],
返回：[1,2,3].
备注；用递归来解这道题很简单，你可以给出迭代的解法么？
*/
package main

import (
	"fmt"
	. "lc/pkg"
)

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	ans := make([]int, 0)
	stk := make([]*TreeNode, 0)
	stk = append(stk, root)
	ans = append(ans, root.Val)
	cur := root
	for len(stk) != 0 {
		if cur.Left != nil {
			stk = append(stk, cur.Left)
			cur = cur.Left
			ans = append(ans, cur.Val)
		} else {
			top := stk[len(stk)-1]
			stk = stk[:len(stk)-1]
			if top.Right != nil {
				stk = append(stk, top.Right)
				cur = top.Right
				ans = append(ans, cur.Val)
			}
		}
	}
	return ans
}

// 示例1
// 输入: [1,null,2,3]
// 返回值: [1,2,3]
func main() {
	tree := CreateTree()
	fmt.Println(preorderTraversal(tree))
}
