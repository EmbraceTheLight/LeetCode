// 04.09. 二叉搜索树序列
/*
从左向右遍历一个数组，通过不断将其中的元素插入树中可以逐步地生成一棵二叉搜索树。
给定一个由不同节点组成的二叉搜索树 root，输出所有可能生成此树的数组。

提示：
二叉搜索树中的节点数在 [0, 1000] 的范围内
1 <= 节点值 <= 106
用例保证符合要求的数组数量不超过 5000
*/
package main

import (
	"fmt"
	. "lc/pkg"
)

func BSTSequences(root *TreeNode) [][]int {
	return bstSequences0409(root)
}

func bstSequences0409(root *TreeNode) [][]int {
	// 递归内部把空树看成一种“空序列”，方便和另一侧子树正常交织。
	if root == nil {
		return [][]int{{}}
	}

	leftSeqs := bstSequences0409(root.Left)
	rightSeqs := bstSequences0409(root.Right)
	ans := make([][]int, 0)
	// 当前根节点必须先插入，后面才能插入左右子树中的节点。
	prefix := []int{root.Val}

	for _, left := range leftSeqs {
		for _, right := range rightSeqs {
			weave0409(left, right, prefix, &ans)
		}
	}

	return ans
}

func weave0409(left, right, prefix []int, ans *[][]int) {
	// 只剩一侧时，剩余节点的相对顺序已经固定，直接追加即可。
	if len(left) == 0 || len(right) == 0 {
		cur := append([]int{}, prefix...)
		cur = append(cur, left...)
		cur = append(cur, right...)
		*ans = append(*ans, cur)
		return
	}

	// 分别尝试从左序列、右序列取下一个节点；不能打乱各自内部顺序。
	prefix = append(prefix, left[0])
	weave0409(left[1:], right, prefix, ans)
	prefix = prefix[:len(prefix)-1]

	prefix = append(prefix, right[0])
	weave0409(left, right[1:], prefix, ans)
}

// 示例 1：
// 输入：root = [2,1,3]
// 输出：[[2,1,3],[2,3,1]]
// 解释：数组 [2,1,3]、[2,3,1] 均可以通过从左向右遍历元素插入树中形成以下二叉搜索树
//
//	  2
//	 / \
//	1   3
//
// 示例 2：
// 输入：root = [4,1,null,null,3,2]
// 输出：[[4,1,3,2]]
func main() {
	fmt.Println(BSTSequences(CreateTree()))
}
