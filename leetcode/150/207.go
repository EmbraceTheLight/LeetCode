// 207. 课程表
/*
你这个学期必须选修 numCourses 门课程，记为 0 到 numCourses - 1 。
在选修某些课程之前需要一些先修课程。 先修课程按数组 prerequisites 给出，
其中 prerequisites[i] = [ai, bi] ，表示如果要学习课程 ai 则 必须 先学习课程  bi 。
例如，先修课程对 [0, 1] 表示：想要学习课程 0 ，你需要先完成课程 1 。
请你判断是否可能完成所有课程的学习？如果可以，返回 true ；否则，返回 false 。

示例 1：
输入：numCourses = 2, prerequisites = [[1,0]]
输出：true
解释：总共有 2 门课程。学习课程 1 之前，你需要完成课程 0 。这是可能的。

示例 2：
输入：numCourses = 2, prerequisites = [[1,0],[0,1]]
输出：false
解释：总共有 2 门课程。学习课程 1 之前，你需要先完成课程 0 ；并且学习课程 0 之前，你还应先完成课程 1 。这是不可能的。

提示：
1 <= numCourses <= 2000
0 <= prerequisites.length <= 5000
prerequisites[i].length == 2
0 <= ai, bi < numCourses
prerequisites[i] 中的所有课程对 互不相同
*/
package main

import (
	"fmt"
	"lc/pkg"
)

func canFinish(numCourses int, prerequisites [][]int) bool {
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

	return numCourses == 0
}

// Test Case1: numCourses = 2, prerequisites = [[1,0]]		Output: true
// Test Case2: numCourses = 2, prerequisites = [[1,0],[0,1]]	Output: false
func main() {
	var numCourses int
	fmt.Println("Input num of courses:")
	fmt.Scan(&numCourses)
	fmt.Println("Input prerequisites:")
	prerequisites := pkg.CreateSlice2D[int]()
	fmt.Println(canFinish(numCourses, prerequisites))
}
