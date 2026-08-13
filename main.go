package main

import (
	"fmt"
	. "lc/pkg"
)


// 示例 1：
// 输入：[[0,0,0],[0,1,0],[0,0,0]]
// 输出：[[0,0],[0,1],[0,2],[1,2],[2,2]]
func main() {
	fmt.Println("Input nums:")
	nums := CreateSlice[int]()
	fmt.Println(findMagicIndex(nums))
}

func findTarget(root *TreeNode, target int) *TreeNode {
	if root == nil {
		return nil
	}
	node := root
	for node != nil {
		if node.Val < target {
			node = node.Right
		} else if node.Val > target {
			node = node.Left
		} else {
			return node
		}
	}
	return nil
}
