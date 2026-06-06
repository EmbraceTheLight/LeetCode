/*
207. 课程表
你这个学期必须选修 numCourses 门课程，记为 0 到 numCourses - 1 。
在选修某些课程之前需要一些先修课程。 先修课程按数组 prerequisites 给出，其中 prerequisites[i] = [ai, bi] ，表示如果要学习课程 ai 则 必须 先学习课程  bi 。
例如，先修课程对 [0, 1] 表示：想要学习课程 0 ，你需要先完成课程 1 。
请你判断是否可能完成所有课程的学习？如果可以，返回 true ；否则，返回 false 。

示例 1：
输入：numCourses = 2, prerequisites = [[1,0]]
输出：true
解释：总共有 2 门课程。学习课程 1 之前，你需要完成课程 0 。这是可能的。

示例 2：
输入：numCourses = 2, prerequisites = [[1,0],[0,1]]
输出：false
解释：总共有 2 门课程。学习课程 1 之前，你需要先完成 课程 0 ；并且学习课程 0 之前，你还应先完成课程 1 。这是不可能的。

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
	"lc/100/pkg"
)

func canFinish(numCourses int, prerequisites [][]int) bool {
	numEdges := len(prerequisites)
	edge := make(map[int][]int)            ///key标识某门课，value为该课程指向的课程，即key是value所有课程的先修课
	var indegree = make([]int, numCourses) //入度表

	for i := 0; i < numEdges; i++ {
		edge[prerequisites[i][1]] = append(edge[prerequisites[i][1]], prerequisites[i][0])
		indegree[prerequisites[i][0]]++
	}

	var q = make([]int, 0)
	for i := 0; i < numCourses; i++ { //将入度为0的节点加入队列中
		if indegree[i] == 0 {
			q = append(q, i)
		}
	}

	var visit = make(map[int]bool)
	//bfs
	for len(q) > 0 {
		v := q[0]
		for _, vv := range edge[v] {
			if !visit[vv] { //该课程未被访问过
				indegree[vv]-- //邻接节点入度-1
				if indegree[vv] == 0 {
					q = append(q, vv)
				}
			}
		}
		numCourses--
		q = q[1:]
		visit[v] = true //将该课程标识为访问过
	}
	if numCourses == 0 { //所有课程都排好了
		return true
	} else { //有冲突的课程
		return false
	}

}
func main() {
	fmt.Println("Input numCourses :")
	var nc int
	fmt.Scan(&nc)
	sli := pkg.CreateIntSlice2()
	fmt.Println(canFinish(nc, sli))
}
