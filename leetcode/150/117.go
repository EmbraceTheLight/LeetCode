// 117. 填充每个节点的下一个右侧节点指针 II
/*
给定一个二叉树：

struct Node {
  int Val;
  Node *Left;
  Node *Right;
  Node *Next;
}
填充它的每个 Next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 Next 指针设置为 NULL 。
初始状态下，所有 Next 指针都被设置为 NULL 。

示例 1：
输入：root = [1,2,3,4,5,null,7]
输出：[1,#,2,3,#,4,5,7,#]
解释：给定二叉树如图 A 所示，你的函数应该填充它的每个 Next 指针，以指向其下一个右侧节点，如图 B 所示。序列化输出按层序遍历顺序（由 Next 指针连接），'#' 表示每层的末尾。

示例 2：
输入：root = []
输出：[]


提示：
树中的节点数在范围 [0, 6000] 内
-100 <= Node.val <= 100

进阶：
你只能使用常量级额外空间。
使用递归解题也符合要求，本题中递归程序的隐式栈空间不计入额外空间复杂度
*/
package main

import (
	"lc/pkg"
)

type Node117 struct {
	Val   int
	Left  *Node117
	Right *Node117
	Next  *Node117
}

func connect(root *Node117) *Node117 {
	// 主要思想：利用上一层已经填充好的Next指针，填充当前层的Next指针
	var createConnect func(parent *Node117)
	createConnect = func(parent *Node117) {
		for {
			var nextLevelHead *Node117
			var cur *Node117
			for parent != nil {
				for parent != nil && nextLevelHead == nil {
					nextLevelHead = getChild(parent)
					cur = nextLevelHead
					if nextLevelHead == nil {
						parent = parent.Next
					}
				}
				if cur == nil {
					return
				}
				if parent.Left != nil && cur != parent.Left {
					cur.Next = parent.Left
					cur = cur.Next
				}

				if parent.Right != nil && cur != parent.Right {
					cur.Next = parent.Right
					cur = cur.Next
				}
				parent = parent.Next
			}

			if nextLevelHead == nil {
				return
			}
			parent = nextLevelHead
		}
	}
	createConnect(root)
	return root
}

func getChild(parent *Node117) *Node117 {
	if parent.Left != nil {
		return parent.Left
	}
	return parent.Right
}

func createTree() *Node117 {
	tree := pkg.CreateTree()
	var dfs func(root *pkg.TreeNode) *Node117
	dfs = func(root *pkg.TreeNode) *Node117 {
		if root == nil {
			return nil
		}
		node := &Node117{
			Val: root.Val,
		}
		node.Left = dfs(root.Left)
		node.Right = dfs(root.Right)
		return node
	}
	return dfs(tree)
}

// Test Case1: [1,2,3,4,5,null,7]	Output: [1,#,2,3,#,4,5,7,#]
// Test Case2: []	Output: []
// * 注：本题暂时没有本地打印函数
func main() {
	tree := createTree()
	connect(tree)
}
