// 210. 课程表 II
/*
现在你总共有 numCourses 门课需要选，记为 0 到 numCourses - 1。给你一个数组 prerequisites ，其中 prerequisites[i] = [ai, bi] ，表示在选修课程 ai 前 必须 先选修 bi 。

例如，想要学习课程 0 ，你需要先完成课程 1 ，我们用一个匹配来表示：[0,1] 。
返回你为了学完所有课程所安排的学习顺序。可能会有多个正确的顺序，
你只要返回 任意一种 就可以了。如果不可能完成所有课程，返回 一个空数组 。

示例 1：
输入：numCourses = 2, prerequisites = [[1,0]]
输出：[0,1]
解释：总共有 2 门课程。要学习课程 1，你需要先完成课程 0。因此，正确的课程顺序为 [0,1].

示例 2：
输入：numCourses = 4, prerequisites = [[1,0],[2,0],[3,1],[3,2]]
输出：[0,2,1,3]
解释：总共有 4 门课程。要学习课程 3，你应该先完成课程 1 和课程 2。并且课程 1 和课程 2 都应该排在课程 0 之后。
因此，一个正确的课程顺序是 [0,1,2,3] 。另一个正确的排序是 [0,2,1,3] 。

示例 3：
输入：numCourses = 1, prerequisites = []
输出：[0]

提示：
1 <= numCourses <= 2000
0 <= prerequisites.length <= numCourses * (numCourses - 1)
prerequisites[i].length == 2
0 <= ai, bi < numCourses
ai != bi
所有[ai, bi] 互不相同
*/
package main

import (
	"fmt"
	"lc/pkg"
)

func findOrder(numCourses int, prerequisites [][]int) []int {
	indegree := make([]int, numCourses) // indegree[i]: 课程 i 的入度, 表示有多少个课程依赖课程 i
	//graph := make(map[int]map[int]bool) // graph 记录课程 i 被哪些课程依赖. 如 graph[i][j] = true 表示课程 j 依赖课程 i
	graph := make(map[int][]int) // graph 课程依赖信息. key 为课程 i, value 为依赖课程 i 的所有课程
	for i := 0; i < len(prerequisites); i++ {
		course, depend := prerequisites[i][0], prerequisites[i][1]
		graph[depend] = append(graph[depend], course)
		indegree[course]++
	}

	set := make([]int, 0) // 记录入度为 0 的课程, 即没有任何依赖的课程
	for i := 0; i < len(indegree); i++ {
		if indegree[i] == 0 {
			set = append(set, i)
		}
	}

	res := make([]int, 0) // 课程学习路径

	// 课程学习路径的生成, 如果存在环路, 则环路包含的节点的入度一定大于 0, 会被忽略掉, 不会出现在 res 中
	// 因此可以根据 res 的长度判断是否所有课程都被学习, 是否可能完成所有课程的学习
	for len(set) > 0 {
		course := set[0]
		res = append(res, course)
		set = set[1:]
		for _, v := range graph[course] {
			indegree[v]--         // 课程 v 的入度减 1, 表示需要学习的前置课程减 1
			if indegree[v] == 0 { // 没有前置课程的课程加入 set
				set = append(set, v)
			}
		}
		numCourses--
	}
	if numCourses > 0 {
		return []int{}
	}
	return res
}

// Test Case1: numCourses = 2, prerequisites = [[1,0]]	Output: [0,1]
// Test Case2: numCourses = 4, prerequisites = [[1,0],[2,0],[3,1],[3,2]]	Output: [0,2,1,3]
// Test Case3: numCourses = 1, prerequisites = []	Output: [0]
func main() {
	var numCourses int
	fmt.Println("Input num of courses:")
	fmt.Scan(&numCourses)
	fmt.Println("Input prerequisites:")
	prerequisites := pkg.CreateSlice2D[int]()
	fmt.Println(findOrder(numCourses, prerequisites))
}
