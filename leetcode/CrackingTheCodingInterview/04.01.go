// 04.01. 节点间通路
/*
节点间通路。给定有向图，设计一个算法，找出两个节点之间是否存在一条路径。
提示：
节点数量n在[0, 1e5]范围内。
节点编号大于等于 0 小于 n。
图中可能存在自环和平行边。
*/
package main

import (
	"fmt"
	. "lc/pkg"
)

func findWhetherExistsPath(n int, graph [][]int, start int, target int) bool {
	graphMp := make([][]int, n)
	for i := 0; i < len(graph); i++ {
		from, to := graph[i][0], graph[i][1]
		graphMp[from] = append(graphMp[from], to)
	}
	visited := map[int]bool{} // 记录已经访问过的节点, 表示由 start 出发, 能到达的点的集合
	for _, i := range graphMp[start] {
		visited[i] = true
	}
	queue := make([]int, 0)
	for i := range visited {
		queue = append(queue, i)
	}
	for len(queue) > 0 {
		i := queue[0]
		queue = queue[1:]
		if i == target {
			return true
		}
		for _, j := range graphMp[i] {
			if j == i || visited[j] == true {
				continue
			}
			visited[j] = true
			queue = append(queue, j)
		}
	}
	return false
}

// 示例 1：
//
//	输入：n = 3, graph = [[0, 1], [0, 2], [1, 2], [1, 2]], start = 0, target = 2
//	输出：true
//
// 示例 2：
//
//	输入：n = 5, graph = [[0, 1], [0, 2], [0, 4], [0, 4], [0, 1], [1, 3], [1, 4], [1, 3], [2, 3], [3, 4]], start = 0, target = 4
//	输出：true
func main() {
	var n, start, target int
	fmt.Println("Input n, start, target:")
	fmt.Scan(&n, &start, &target)
	fmt.Println("Input grid:")
	grid := CreateSlice2D[int]()
	fmt.Println(findWhetherExistsPath(n, grid, start, target))

}
