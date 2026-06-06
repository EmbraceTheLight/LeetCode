// 530. 二叉搜索树的最小绝对差
/*
给你一个二叉搜索树的根节点 root ，返回 树中任意两不同节点值之间的最小差值 。
差值是一个正数，其数值等于两值之差的绝对值。

示例 1：
输入：root = [4,2,6,1,3]
输出：1

示例 2：
输入：root = [1,0,48,null,null,12,49]
输出：1

提示：
树中节点的数目范围是 [2, 10^4]
0 <= Node.val <= 10^5

注意：本题与 783 https://leetcode.cn/problems/minimum-distance-between-bst-nodes/ 相同
*/
package main

import (
	"fmt"
	"lc/pkg"
	"math"
)

func getMinimumDifference(root *pkg.TreeNode) int {
	// 方法 1
	ans := math.MaxInt
	//dfs530(root, &ans)

	// 方法 2
	//sli := make([]int, 0)
	//dfs530_2(root, &sli)
	//for i := 1; i < len(sli); i++ {
	//	ans = min(ans, sli[i]-sli[i-1])
	//}

	// 方法 3, 不使用切片, 直接一遍遍历完成
	var dfs530_3 func(root *pkg.TreeNode, ans *int)
	pre := -1 // pre 可以看作方法 2 中的 sli[i -1], 而 root.Val 则可以看作 sli[i]
	dfs530_3 = func(root *pkg.TreeNode, ans *int) {
		if root == nil {
			return
		}
		dfs530_3(root.Left, ans)
		if pre != -1 {
			*ans = min(*ans, root.Val, pre)
		}
		pre = root.Val
		dfs530_3(root.Right, ans)
	}
	dfs530_3(root, &ans)
	return ans
}

func dfs530(root *pkg.TreeNode, ans *int) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return root.Val
	}

	dfs530(root.Left, ans)
	dfs530(root.Right, ans)

	// 左子树最大值、右子树最小值分别与当前节点值比较, 取最小值并更新
	cur := root.Left
	if cur != nil {
		for ; cur.Right != nil; cur = cur.Right {
		}
		*ans = min(*ans, int(math.Abs(float64(root.Val-cur.Val))))
	}

	cur = root.Right
	if cur != nil {
		for ; cur.Left != nil; cur = cur.Left {
		}
		*ans = min(*ans, int(math.Abs(float64(root.Val-cur.Val))))
	}
	return root.Val
}

func dfs530_2(root *pkg.TreeNode, ans *[]int) {
	if root == nil {
		return
	}
	// 使用二叉搜索树性质, 中序遍历得到的一定是升序结果
	dfs530_2(root.Left, ans)
	*ans = append(*ans, root.Val)
	dfs530_2(root.Right, ans)
	return
}

// Test Case1:	[4,2,6,1,3]		Output: 1
// Test Case2:	[1,0,48,null, null,12,49]	Output: 1
// Test Case3:	[236,104,701,null,227,null,911]	Output: 9
// Test Case4:  [0,null,2236,1277,2776,519]	Output: 519
func main() {
	root := pkg.CreateTree()
	fmt.Println(getMinimumDifference(root))
}
