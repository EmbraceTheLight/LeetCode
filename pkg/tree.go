package pkg

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type treeQueue struct {
	queue []*TreeNode
}

func (q *treeQueue) len() int { return len(q.queue) }
func (q *treeQueue) pop() {
	q.queue = q.queue[1:]
}
func (q *treeQueue) top() *TreeNode {
	ret := q.queue[0]
	return ret
}
func (q *treeQueue) Push(node *TreeNode) {
	q.queue = append(q.queue, node)
}

func CreateTree() *TreeNode {
	fmt.Println("Enter 'tree' string (e.g. [1, 2, null, 3]): ")
	input := CreateSlice[string]()
	if len(input) == 0 || input[0] == "null" {
		return nil
	}

	head := new(TreeNode)
	node := head
	var queue treeQueue
	var LoR bool //指示向左节点或右节点插入数据

	tmp := input[0]
	val, _ := strconv.Atoi(tmp)
	head.Val = val
	queue.Push(head)

	for i := 1; i < len(input); i++ {
		tmp = input[i]
		if queue.len() > 0 {
			node = queue.top()
			if LoR == true {
				queue.pop() //队列pop操作
			}

			if LoR == false { //向左节插入数据
				LoR = true         //之后该向右节点插入数据
				if tmp != "null" { //要插入的数据值不是null
					val, _ := strconv.Atoi(tmp)
					node.Left = new(TreeNode)
					node = node.Left
					node.Val = val
					queue.Push(node)
				} else {
					node.Left = nil
				}
			} else {
				LoR = false //复原，下一个应该向左节点插入数据
				if tmp != "null" {
					val, _ := strconv.Atoi(tmp)
					node.Right = new(TreeNode)
					node = node.Right
					node.Val = val
					queue.Push(node)
				} else {
					node.Right = nil
				}
			}

		}
	}
	return head
}

func PrintTreeByLevel(head *TreeNode) {
	var queue treeQueue
	var nextLevelHead = head //下一层头节点
	var a bool               //表示是否已确认下一层的起始节点
	var level = 0
	queue.Push(head)
	for queue.len() > 0 {
		node := queue.top()
		if node == nextLevelHead {
			level++
			a = false //寻找下一层头节点
			fmt.Println()
			fmt.Printf("level%d: ", level)
		}

		if node.Left != nil {
			queue.Push(node.Left)
		}
		if node.Right != nil {
			queue.Push(node.Right)
		}

		if !a {
			if node.Left != nil {
				nextLevelHead = node.Left
				a = true
			} else if node.Right != nil {
				nextLevelHead = node.Right
				a = true
			}
		}
		fmt.Printf("%d ", node.Val)
		queue.pop()
	}
}
