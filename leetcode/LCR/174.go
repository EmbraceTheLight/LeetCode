/*
LCR 174. 寻找二叉搜索树中的目标节点
某公司组织架构以二叉搜索树形式记录，节点值为处于该职位的员工编号。请返回第 cnt 大的员工编号。
*/
package main

import (
	"fmt"
	"lc/100/pkg"
)

//func getSlice(node *pkg.TreeNode, sli *[]int) {
//	if node == nil {
//		return
//	}
//	getSlice(node.Left, sli)
//	*sli = append(*sli, node.Val)
//	getSlice(node.Right, sli)
//}
//func findTargetNode(root *pkg.TreeNode, cnt int) int {
//	sli := make([]int, 0)
//	getSlice(root, &sli)
//	fmt.Printf("%#v", sli)
//	l := len(sli)
//
//	return sli[l-cnt]
//}

// 思路：求第n大，即为求倒序中序遍历的第n个元素
func dfs(node *pkg.TreeNode, cnt *int, ans *int) {
	if node == nil || *cnt == 0 {
		return
	}
	dfs(node.Right, cnt, ans)
	(*cnt)--
	if *cnt == 0 {
		*ans = node.Val
		return
	}
	dfs(node.Left, cnt, ans)
}
func findTargetNode(root *pkg.TreeNode, cnt int) int {
	var ans int
	dfs(root, &cnt, &ans)
	return ans

}
func main() {
	println("Input cnt:")
	var tmp int
	fmt.Scan(&tmp)

	root := pkg.CreateTree()

	println(findTargetNode(root, tmp))
}
