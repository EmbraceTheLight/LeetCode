package main

import (
	"fmt"
	"lc/pkg"
)

type nodeWithNum struct {
	node  *pkg.TreeNode
	level int
	num   int
}

func widthOfBinaryTreeR(root *pkg.TreeNode) int {
	var ans = 1
	nodes := make([]*nodeWithNum, 0)
	nodes = append(nodes, &nodeWithNum{root, 1, 1})

	var level = 1
	var firstNode, lastNode *nodeWithNum = nodes[0], nodes[0] // firstNode: 一层中最左侧节点，lastNode: 一层中最右侧节点
	for len(nodes) > 0 {
		node := nodes[0]
		nodes = nodes[1:]
		if node.level > level {
			level = node.level
			if lastNode != nil {
				ans = max(ans, lastNode.num-firstNode.num+1)
			}

			firstNode = node
		}

		lastNode = node

		if node.node.Left != nil {
			nodes = append(nodes, &nodeWithNum{node.node.Left, node.level + 1, node.num * 2})
		}
		if node.node.Right != nil {
			nodes = append(nodes, &nodeWithNum{node.node.Right, node.level + 1, node.num*2 + 1})
		}

	}

	if lastNode != nil && firstNode != nil {
		ans = max(ans, lastNode.num-firstNode.num+1)
	}
	return ans
}
func main() {
	root := pkg.CreateTree()
	fmt.Println(widthOfBinaryTreeR(root))

}
