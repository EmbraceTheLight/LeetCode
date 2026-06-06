// 236. 二叉树的最近公共祖先
/*
给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。
百度百科中最近公共祖先的定义为：“对于有根树 T 的两个节点 p、q，最近公共祖先表示为一个节点 x，
满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”


示例 1：
输入：root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
输出：3
解释：节点 5 和节点 1 的最近公共祖先是节点 3 。

示例 2：
输入：root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
输出：5
解释：节点 5 和节点 4 的最近公共祖先是节点 5 。因为根据定义最近公共祖先节点可以为节点本身。

示例 3：
输入：root = [1,2], p = 1, q = 2
输出：1

提示：
树中节点数目在范围 [2, 10^5] 内。
-10^9 <= Node.val <= 10^9
所有 Node.val 互不相同 。
p != q
p 和 q 均存在于给定的二叉树中。
*/
package main

import (
	"fmt"
	"lc/pkg"
)

func lowestCommonAncestorR(root, p, q *pkg.TreeNode) *pkg.TreeNode {
	path1, path2 := make([]*pkg.TreeNode, 0), make([]*pkg.TreeNode, 0)
	tracePath(root, p.Val, &path1)
	tracePath(root, q.Val, &path2)
	mp := make(map[*pkg.TreeNode]bool)
	for i := 0; i < len(path1); i++ {
		mp[path1[i]] = true
	}

	for i := 0; i < len(path2); i++ {
		if mp[path2[i]] {
			return path2[i]
		}
	}
	return root
}
func lowestCommonAncestorR2(root, p, q *pkg.TreeNode) *pkg.TreeNode {
	// 节点为 nil, 或自己是自己的祖先
	if root == nil || root == q || root == p {
		return root
	}

	left := lowestCommonAncestorR(root.Left, p, q)
	right := lowestCommonAncestorR(root.Right, p, q)

	// * p,q 位于当前节点两侧子树, 则当前节点为最近公共祖先
	// * 若 p 在 q 的子树中(或 q 在 p 的子树中), 则找到 p 或 q 直接返回即可
	if left != nil && right != nil {
		return root
	}

	if left != nil {
		return left
	}
	if right != nil {
		return right
	}
	return nil
}

// tracePath 寻找并记录节点所在路径, 返回 true 表示找到了 val 对应的目标节点
func tracePath(root *pkg.TreeNode, val int, mp *[]*pkg.TreeNode) bool {
	if root == nil {
		return false
	}
	if root.Val == val {
		*mp = append(*mp, root)
		return true
	}

	res := tracePath(root.Left, val, mp) || tracePath(root.Right, val, mp)
	if res {
		*mp = append(*mp, root)
		return true
	}
	return false
}

func findNode(val int, root *pkg.TreeNode) *pkg.TreeNode {
	if root.Val == val {
		return root
	}

	if root.Left != nil {
		tmp1 := findNode(val, root.Left)
		if tmp1 != nil {
			return tmp1
		}
	}
	if root.Right != nil {
		tmp2 := findNode(val, root.Right)
		if tmp2 != nil {
			return tmp2
		}
	}
	return nil
}

// Test Case1: p=5, q=1	 [3,5,1,6,2,0,8,null,null,7,4]	Output: 3
// Test Case2: p=5, q=4	 [3,5,1,6,2,0,8,null,null,7,4]	Output: 5
// Test Case3: p=1, q=2	 [1,2]	Output: 1
func main() {
	var p, q int
	fmt.Println("Input p,q:")
	fmt.Scan(&p, &q)
	root := pkg.CreateTree()
	pNode := findNode(p, root)
	qNode := findNode(q, root)
	fmt.Printf("%+v", lowestCommonAncestorR(root, pNode, qNode))
}
