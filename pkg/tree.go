package pkg

import (
	"fmt"
	"math"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type TreeQueue struct {
	queue []*TreeNode
}

func (q *TreeQueue) Len() int { return len(q.queue) }
func (q *TreeQueue) Pop() {
	q.queue = q.queue[1:]
}
func (q *TreeQueue) Top() *TreeNode {
	ret := q.queue[0]
	return ret
}
func (q *TreeQueue) Push(node *TreeNode) {
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
	var queue TreeQueue
	var LoR bool //指示向左节点或右节点插入数据

	tmp := input[0]
	val, _ := strconv.Atoi(tmp)
	head.Val = val
	queue.Push(head)

	for i := 1; i < len(input); i++ {
		tmp = input[i]
		if queue.Len() > 0 {
			node = queue.Top()
			if LoR == true {
				queue.Pop() //队列pop操作
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

type TreeNodeInfo struct {
	node  *TreeNode
	level int
	id    int
}

type TreeNodeInfoQueue struct {
	queue []*TreeNodeInfo
}

func (t *TreeNodeInfoQueue) Len() int {
	return len(t.queue)
}

func (t *TreeNodeInfoQueue) Push(nodeInfo *TreeNodeInfo) {
	t.queue = append(t.queue, nodeInfo)
}

func (t *TreeNodeInfoQueue) Pop() {
	t.queue = t.queue[1:]
}

func (t *TreeNodeInfoQueue) Top() *TreeNodeInfo {
	return t.queue[0]
}

func PrintTreeByLevel(head *TreeNode) {
	q := new(TreeNodeInfoQueue)
	var level = 1

	headInfo := &TreeNodeInfo{node: head, level: 1, id: 1}
	q.Push(headInfo)
	tmp := []*TreeNodeInfo{headInfo}

	for q.Len() > 0 {
		infoNode := q.Top()
		if infoNode.level != level {
			printLevel(tmp)
			level++
			tmp = make([]*TreeNodeInfo, 0)
		}

		if infoNode.node.Left != nil {
			q.Push(&TreeNodeInfo{node: infoNode.node.Left, level: level + 1, id: 2 * infoNode.id})
		}
		if infoNode.node.Right != nil {
			q.Push(&TreeNodeInfo{node: infoNode.node.Right, level: level + 1, id: 2*infoNode.id + 1})
		}

		tmp = append(tmp, infoNode)
		q.Pop()
	}
	if len(tmp) > 0 {
		printLevel(tmp)
	}
}
func printLevel(nodes []*TreeNodeInfo) {
	level := nodes[0].level
	start := int(math.Pow(2, float64(level-1)))
	end := int(math.Pow(2, float64(level)) - 1)
	fmt.Printf("level %d: ", level)
	for i := start; i <= end; i++ {
		if nodes[0].id == i {
			fmt.Printf("%d ", nodes[0].node.Val)
			nodes = nodes[1:]
		} else {
			fmt.Printf("null ")
		}
	}

	fmt.Println()
}
