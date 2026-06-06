// 373. 查找和最小的 K 对数字
/*
给定两个以 非递减顺序排列 的整数数组 nums1 和 nums2 , 以及一个整数 k 。
定义一对值 (u,v)，其中第一个元素来自 nums1，第二个元素来自 nums2 。
请找到和最小的 k 个数对 (u1,v1),  (u2,v2)  ...  (uk,vk) 。

提示:
1 <= nums1.length, nums2.length <= 10^5
-10^9 <= nums1[i], nums2[i] <= 10^9
nums1 和 nums2 均为 升序排列
1 <= k <= 10^4
k <= nums1.length * nums2.length
*/
package main

import (
	"container/heap"
	"fmt"
	"lc/pkg"
)

// * 本题难点在于要找的第 k 小的数对, 随着 k 增大而需要比较的数对的数量也就越大
// * 如最小的数对就是 [nums1[0], nums2[0]], 第二小的数对就是 [nums1[0], nums2[1]] 或 [nums1[1], nums2[0]], 第 k 小的数对要比较的数对就更多了
// * 所有数对和有一个二维矩阵的关系, 即 nums1[0] + nums2[0 ... n - 1] 构成第一行, nums1[1] + nums2[0 ... n - 1] 构成第二行, 以此类推
// * 每列数对和逐行递增, 每行数对和逐列递增
// * 思路: 可以使用小根堆保存下标对(i, j), 比较规则为 nums1[i] + nums2[j] 越小越好
// * 初始时, 将 (0, 0) 放入小根堆, 然后不断从堆顶取出下标对 (i, j), 将候选者下标对 (i, j + 1) 和 (i + 1, j) 放入堆中(其他情况都显然比 nums1[i] + nums2[j + 1] 和 nums1[i + 1] + nums2[j] 要大),
// * 直到找到 k 个数对
// * 也可以理解为沿着矩阵主对角线寻找候选
// * 这种做法需要一个 map 来记录已经访问过的下标对, 避免重复访问
func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	var ans [][]int
	visited := make(map[int]map[int]bool)
	pq := &priorityQueue373{}
	heap.Init(pq)
	heap.Push(pq, &pair{i: 0, j: 0, sum: nums1[0] + nums2[0]})
	for k > 0 {
		top := heap.Pop(pq).(*pair)
		if _, ok := visited[top.i]; !ok {
			visited[top.i] = map[int]bool{}
		}
		if visited[top.i][top.j] { // 跳过重复的 (i, j) 对
			continue
		}
		visited[top.i][top.j] = true
		k--
		ans = append(ans, []int{nums1[top.i], nums2[top.j]})
		if top.i+1 < len(nums1) {
			candidate1 := &pair{i: top.i + 1, j: top.j, sum: nums1[top.i+1] + nums2[top.j]}
			heap.Push(pq, candidate1)
		}
		if top.j+1 < len(nums2) {
			candidate2 := &pair{i: top.i, j: top.j + 1, sum: nums1[top.i] + nums2[top.j+1]}
			heap.Push(pq, candidate2)
		}

	}
	return ans
}

type pair struct {
	i, j int
	sum  int
}

type priorityQueue373 []*pair

func (q *priorityQueue373) Len() int {
	return len(*q)
}

func (q *priorityQueue373) Less(i, j int) bool {
	return (*q)[i].sum < (*q)[j].sum
}

func (q *priorityQueue373) Swap(i, j int) {
	(*q)[i], (*q)[j] = (*q)[j], (*q)[i]
}

func (q *priorityQueue373) Push(x any) {
	*q = append(*q, x.(*pair))
}

func (q *priorityQueue373) Pop() any {
	ret := (*q)[q.Len()-1]
	*q = (*q)[:q.Len()-1]
	return ret
}

// 示例 1:
// 输入: nums1 = [1,7,11], nums2 = [2,4,6], k = 3
// 输出: [[1,2],[1,4],[1,6]]
// 解释: 返回序列中的前 3 对数：
//
//	[1,2],[1,4],[1,6],[7,2],[7,4],[11,2],[7,6],[11,4],[11,6]
//
// 示例 2:
// 输入: nums1 = [1,1,2], nums2 = [1,2,3], k = 2
// 输出: [[1,1],[1,1]]
// 解释: 返回序列中的前 2 对数：
//
//	[1,1],[1,1],[1,2],[2,1],[1,2],[2,2],[1,3],[1,3],[2,3]
func main() {
	var k int
	fmt.Println("Input k:")
	fmt.Scan(&k)
	fmt.Println("Input nums1, nums2:")
	fmt.Println(kSmallestPairs(pkg.CreateSlice[int](), pkg.CreateSlice[int](), k))
}
