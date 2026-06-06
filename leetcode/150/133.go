// 133. 克隆图
/*
给你无向 连通 图中一个节点的引用，请你返回该图的 深拷贝（克隆）。
图中的每个节点都包含它的值 val（int） 和其邻居的列表（list[Node]）。

class Node {
    public int val;
    public List<Node> neighbors;
}

测试用例格式：
简单起见，每个节点的值都和它的索引相同。例如，第一个节点值为 1（val = 1），第二个节点值为 2（val = 2），以此类推。该图在测试用例中使用邻接列表表示。
邻接列表 是用于表示有限图的无序列表的集合。每个列表都描述了图中节点的邻居集。
给定节点将始终是图中的第一个节点（值为 1）。你必须将 给定节点的拷贝 作为对克隆图的引用返回。


示例 1：
输入：adjList = [[2,4],[1,3],[2,4],[1,3]]
输出：[[2,4],[1,3],[2,4],[1,3]]
解释：
图中有 4 个节点。
节点 1 的值是 1，它有两个邻居：节点 2 和 4 。
节点 2 的值是 2，它有两个邻居：节点 1 和 3 。
节点 3 的值是 3，它有两个邻居：节点 2 和 4 。
节点 4 的值是 4，它有两个邻居：节点 1 和 3 。

示例 2：
输入：adjList = [[]]
输出：[[]]
解释：输入包含一个空列表。该图仅仅只有一个值为 1 的节点，它没有任何邻居。

示例 3：
输入：adjList = []
输出：[]
解释：这个图是空的，它不含任何节点。

这张图中的节点数在 [0, 100] 之间。
1 <= Node.val <= 100
每个节点值 Node.val 都是唯一的，
图中没有重复的边，也没有自环。
图是连通图，你可以从给定节点访问到所有节点。
*/
package main

import "lc/pkg"

type Node133 struct {
	Val       int
	Neighbors []*Node133
}

func cloneGraph(node *Node133) *Node133 {
	if node == nil {
		return nil
	}
	visited := make(map[int]bool)    // 记录已经克隆过的节点, 防止重复克隆导致死循环
	cloned := make(map[int]*Node133) // 因题目明确说明每个 Node 的 Val 都是唯一的, 所以用 map 存储已克隆的节点, key 为 Val, value 为 克隆后的节点地址
	queue := []*Node133{node}
	var curNode, ret *Node133
	for len(queue) > 0 {
		originNode := queue[0]
		queue = queue[1:]
		if visited[originNode.Val] {
			continue
		}
		visited[originNode.Val] = true
		curNode = copyNode(originNode, cloned) // 克隆当前节点
		if ret == nil {
			ret = curNode
		}
		for _, n := range originNode.Neighbors {
			curNode.Neighbors = append(curNode.Neighbors, copyNode(n, cloned))
		}
		queue = append(queue, originNode.Neighbors...)
	}
	return ret
}

// copyNode 克隆节点, 仅克隆 Val, 并初始化 Neighbors
// 如果该节点已经克隆过, 则直接返回克隆后的节点
func copyNode(originNode *Node133, clonedMap map[int]*Node133) *Node133 {
	if clonedMap[originNode.Val] != nil {
		return clonedMap[originNode.Val]
	}
	ret := &Node133{Val: originNode.Val, Neighbors: make([]*Node133, 0)}
	clonedMap[originNode.Val] = ret
	return ret
}

// 简洁形式, 只用了一个 map 存储已克隆的节点, 并用一个队列来存储待克隆的节点, 避免了死循环
func cloneGraphBetter(node *Node133) *Node133 {
	if node == nil {
		return nil
	}
	visited := make(map[*Node133]*Node133) // visited 记录已经克隆过的节点, 防止重复克隆导致死循环. key 为原节点, value 为克隆后的节点

	visited[node] = &Node133{Val: node.Val, Neighbors: make([]*Node133, 0)} // 克隆起始节点
	queue := []*Node133{node}
	for len(queue) > 0 {
		curNode := queue[0]
		queue = queue[1:]

		// 遍历当前节点的邻居
		for _, n := range curNode.Neighbors {
			if _, ok := visited[n]; !ok { // 当前邻居节点尚未被克隆, 则克隆该节点
				visited[n] = &Node133{Val: n.Val, Neighbors: make([]*Node133, 0)}
				queue = append(queue, n) // 邻居节点未被克隆, 则将其加入队列, 通过这种方式避免死循环
			}
			visited[curNode].Neighbors = append(visited[curNode].Neighbors, visited[n]) // 注意, 这里向当前克隆后的节点的邻居列表中添加克隆后的邻居节点, 而不是原节点的邻居节点
		}
	}
	return visited[node]
}

// Test Case1:	[[2,4],[1,3],[2,4],[1,3]]	[[2,4],[1,3],[2,4],[1,3]]
// Test Case2:	[[]]	[[]]
// Test Case3:	[]	[]
func main() {
	slices := pkg.CreateSlice2D[int]()
	cloneGraph(makeGraph(slices))
}

func makeGraph(slices [][]int) *Node133 {
	rows := len(slices)
	cols := len(slices[rows-1])
	mp := map[int]*Node133{}
	for i := 0; i < rows; i++ {
		curNode := copyNode(&Node133{Val: i + 1, Neighbors: make([]*Node133, 0)}, mp)
		for j := 0; j < cols; j++ {
			curNode.Neighbors = append(curNode.Neighbors, copyNode(&Node133{Val: slices[i][j], Neighbors: make([]*Node133, 0)}, mp))
		}
	}
	return mp[1]
}
