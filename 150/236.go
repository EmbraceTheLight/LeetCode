/*
236. 二叉树的最近公共祖先
给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。
百度百科中最近公共祖先的定义为：“对于有根树 T 的两个节点 p、q，最近公共祖先表示为一个节点 x，
满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”
*/
package main


func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	left := lowestCommonAncestor(root.Left, p, q)   //搜索p，q是否在左子树
	right := lowestCommonAncestor(root.Right, p, q) //搜索p，q是否在右子树
	if root == p || root == q {                     //找到p/q了
		return root
	}
	if left == nil && right == nil { //p，q均不在root左右子树中，显然不存在公共节点
		return nil
	}
	if left != nil && right != nil { //p，q分别在左右子树中，root即为公共节点
		return root
	}
	if left != nil && right == nil { //p，q都不在右子树中
		return left
	}
	if left == nil && right != nil { //p，q都不在左子树中
		return right
	}
	return root
}
func main() {

}
