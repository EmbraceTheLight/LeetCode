// CC6 二叉树的后序遍历
/*
描述
用递归的方法对给定的二叉树进行后序遍历。
例如：给定的二叉树为[1,null,2,3],
返回[3,2,1].
*/
package main

import (
	"fmt"
	. "lc/pkg"
)

func postorderTraversal(root *TreeNode) []int {
	ans := make([]int, 0)
	dfscc5(root, &ans)
	return ans
}
func dfscc5(treeNode *TreeNode, ans *[]int) {
	if treeNode == nil {
		return
	}
	dfscc5(treeNode.Left, ans)
	dfscc5(treeNode.Right, ans)
	*ans = append(*ans, treeNode.Val)
}

// 示例1
// 输入: [1,null,2,3]
// 返回值: [3,2,1]
func main() {
	root := CreateTree()
	fmt.Println(postorderTraversal(root))
}
