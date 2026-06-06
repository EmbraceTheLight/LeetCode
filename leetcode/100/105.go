/*
105. 从前序与中序遍历序列构造二叉树
给定两个整数数组 preorder 和 inorder ，其中 preorder 是二叉树的先序遍历， inorder 是同一棵树的中序遍历，请构造二叉树并返回其根节点。
1 <= preorder.length <= 3000
inorder.length == preorder.length
-3000 <= preorder[i], inorder[i] <= 3000
preorder 和 inorder 均 无重复 元素
inorder 均出现在 preorder
preorder 保证 为二叉树的前序遍历序列
inorder 保证 为二叉树的中序遍历序列
*/
package main

import (
	"fmt"
	"lc/100/pkg"
)

var mp105 map[int]int //键:inorder的值,值：inorder的索引

func initMap(inorder []int) {
	mp105 = make(map[int]int)
	l := len(inorder)
	for i := 0; i < l; i++ {
		mp105[inorder[i]] = i
	}
}

// 后四个参数全为闭区间
func dfs105(preorder []int, inorder []int, pstart, pend, istart, iend int) *pkg.TreeNode {
	if len(mp105) == 0 {
		initMap(inorder)
	}
	if pstart > pend || istart > iend {
		return nil
	}
	var node = new(pkg.TreeNode)
	idx := mp105[preorder[pstart]] //得到在inorder中对应值的索引
	node.Val = preorder[pstart]

	prelen := idx - istart
	node.Left = dfs105(preorder, inorder, pstart+1, pstart+prelen, istart, idx-1)
	node.Right = dfs105(preorder, inorder, pstart+prelen+1, pend, idx+1, iend)
	return node
}
func buildTree(preorder []int, inorder []int) *pkg.TreeNode {
	return dfs105(preorder, inorder, 0, len(preorder)-1, 0, len(inorder)-1)
}

func main() {
	preorder := makeSlice()
	inorder := makeSlice()
	root := buildTree(preorder, inorder)
	pkg.PrintTreeByLevel(root)
}
func makeSlice() []int {
	fmt.Println("Making Slice...")
	fmt.Println("Input Value,-1 to quit:")
	sli := make([]int, 0)
	var tmp int
	fmt.Scan(&tmp)
	for tmp != -1 {
		sli = append(sli, tmp)
		fmt.Scan(&tmp)
	}
	return sli
}
