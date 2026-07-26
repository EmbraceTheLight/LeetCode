// 04.10. 检查子树
/*
检查子树。你有两棵非常大的二叉树：T1，有几万个节点；T2，有几万个节点。设计一个算法，判断 T2 是否为 T1 的子树。
如果 T1 有这么一个节点 n，其子树与 T2 一模一样，则 T2 为 T1 的子树，也就是说，从节点 n 处把树砍断，得到的树与 T2 完全相同。

注意：此题相对书上原题略有改动。

提示：
树的节点数目范围为 [0, 20000]。
*/
package main

import (
	"fmt"
	. "lc/pkg"
)

func checkSubTree(t1 *TreeNode, t2 *TreeNode) bool {
	return dfs0410(t1, t2)
}
func dfs0410(node1, t2 *TreeNode) bool {
	if node1 == nil {
		return false
	}
	if node1.Val == t2.Val {
		res := check0410(node1, t2)
		if res == true {
			return true
		}
	}
	return dfs0410(node1.Left, t2) || dfs0410(node1.Right, t2)
}
func check0410(node1, node2 *TreeNode) bool {
	if node1 == nil && node2 == nil {
		return true
	}
	if (node1 == nil && node2 != nil) || (node1 != nil && node2 == nil) || node1.Val != node2.Val {
		return false
	}
	var lRes, rRes bool = true, true
	lRes = check0410(node1.Left, node2.Left)
	rRes = check0410(node1.Left, node2.Left)
	return lRes || rRes
}

// 示例 1：
//
//	输入：t1 = [1, 2, 3], t2 = [2]
//	输出：true
//
// 示例 2：
//
//	输入：t1 = [1, null, 2, 4], t2 = [3, 2]
//	输出：false
func main() {
	t1 := CreateTree()
	t2 := CreateTree()
	fmt.Println(checkSubTree(t1, t2))
}
