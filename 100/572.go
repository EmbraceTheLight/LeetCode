/*
572.另一棵树的子树
给你两棵二叉树 root 和 subRoot 。检验 root 中是否包含和 subRoot 具有相同结构和节点值的子树。如果存在，返回 true ；否则，返回 false 。
二叉树 tree 的一棵子树包括 tree 的某个节点和这个节点的所有后代节点。tree 也可以看做它自身的一棵子树。
root 树上的节点数量范围是 [1, 2000]
subRoot 树上的节点数量范围是 [1, 1000]
-10^4 <= root.val <= 10^4
-10^4 <= subRoot.val <= 10^4
*/
package main

import "lc/100/pkg"

type myQue struct {
	node  *pkg.TreeNode
	index int
}

func check(TreeSli []*myQue, SubTreeSli []*myQue) bool {
	lT := len(TreeSli)
	lsT := len(SubTreeSli)
	if lT != lsT {
		return false
	}
	for i := 0; i < lT; i++ {
		if TreeSli[i].node.Val != SubTreeSli[i].node.Val ||
			TreeSli[i].index != SubTreeSli[i].index {
			return false
		}
	}
	return true
}

func GetNodes(root *pkg.TreeNode, tar int) []*pkg.TreeNode {
	if root == nil {
		return nil
	}

	res := make([]*pkg.TreeNode, 0)
	if root.Val == tar {
		res = append(res, root)
	}
	res = append(res, GetNodes(root.Left, tar)...)
	res = append(res, GetNodes(root.Right, tar)...)

	return res
}

func isSubtree(root *pkg.TreeNode, subRoot *pkg.TreeNode) bool {
	var tar = subRoot.Val
	MayBe := GetNodes(root, tar)

	for _, m := range MayBe {
		var T = make([]*myQue, 0)
		var ST = make([]*myQue, 0)
		var flag = true
		ST = append(ST, &myQue{
			node:  subRoot,
			index: 1,
		})
		T = append(T, &myQue{
			node:  m,
			index: 1,
		})

		for len(T) > 0 {
			var levelsz = len(T)

			//一层一层地对比
			for i := 0; i < levelsz; i++ {
				if T[i].node.Left != nil {
					T = append(T, &myQue{
						node:  T[i].node.Left,
						index: T[i].index * 2,
					})
				}
				if T[i].node.Right != nil {
					T = append(T, &myQue{
						node:  T[i].node.Right,
						index: T[i].index*2 + 1,
					})
				}
				if ST[i].node.Left != nil {
					ST = append(ST, &myQue{
						node:  ST[i].node.Left,
						index: ST[i].index * 2,
					})
				}
				if ST[i].node.Right != nil {
					ST = append(ST, &myQue{
						node:  ST[i].node.Right,
						index: ST[i].index*2 + 1,
					})
				}
			}
			if check(T, ST) == false {
				flag = false
				break
			}
			T = T[levelsz:]
			ST = ST[levelsz:]
		}
		if flag {
			return true
		}

	}
	return false
}
func main() {
	root := pkg.CreateTree()
	subRoot := pkg.CreateTree()
	pkg.PrintTreeByLevel(root)
	pkg.PrintTreeByLevel(subRoot)
	println(isSubtree(root, subRoot))
}
