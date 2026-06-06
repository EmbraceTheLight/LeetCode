// 230. 二叉搜索树中第 K 小的元素
/*
给定一个二叉搜索树的根节点 root ，和一个整数 k ，请你设计一个算法查找其中第 k 小的元素（k 从 1 开始计数）。

示例 1：
输入：root = [3,1,4,null,2], k = 1
输出：1

示例 2：
输入：root = [5,3,6,2,4,null,null,1], k = 3
输出：3

提示：

树中的节点数为 n 。
1 <= k <= n <= 10^4
0 <= Node.val <= 10^4

进阶：如果二叉搜索树经常被修改（插入/删除操作）并且你需要频繁地查找第 k 小的值，你将如何优化算法？
*/
package main

import (
	"fmt"
	"lc/pkg"
)

func kthSmallest(root *pkg.TreeNode, k int) int {
	ans := -1
	var count int
	dfs230(root, &count, &ans, k)
	return ans
}
func dfs230(root *pkg.TreeNode, count, ans *int, k int) {
	if root == nil || *count >= k {
		return
	}
	dfs230(root.Left, count, ans, k)
	(*count)++
	if *count == k {
		*ans = root.Val
		return
	}
	dfs230(root.Right, count, ans, k)
}

// Test Case1: k = 1	[3,1,4,null,2] 	Output: 1
// Test Case2: k = 3	[5,3,6,2,4, null, null, 1] 	Output: 3

func main() {
	var k int
	fmt.Println("Input k:")
	fmt.Scan(&k)
	root := pkg.CreateTree()
	fmt.Println(kthSmallest(root, k))
}
