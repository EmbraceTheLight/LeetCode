// 56. 合并区间
/*
以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi]。
请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。


示例 1：
输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
输出：[[1,6],[8,10],[15,18]]
解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].

示例 2：
输入：intervals = [[1,4],[4,5]]
输出：[[1,5]]
解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。

示例 3：
输入：intervals = [[4,7],[1,4]]
输出：[[1,7]]
解释：区间 [1,4] 和 [4,7] 可被视为重叠区间。

提示：

1 <= intervals.length <= 10^4
intervals[i].length == 2
0 <= starti <= endi <= 10^4
*/
package main

import (
	"fmt"
	"lc/pkg"
	"sort"
)

func merge56R(intervals [][]int) [][]int {
	var ans [][]int
	n := len(intervals)
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]

	})

	left, right := intervals[0][0], intervals[0][1]
	for i := 1; i < n; i++ {
		nLeft, nRight := intervals[i][0], intervals[i][1]
		if nLeft > right {
			ans = append(ans, []int{left, right})
			left, right = nLeft, nRight
		} else if nLeft <= right && right < nRight {
			right = nRight
		} else if right > nRight {
			continue
		}
	}
	ans = append(ans, []int{left, right})
	return ans
}

// Test Case1:
/*
[[1,3]
,[2,6]
,[8,10]
,[15,18]]
*/
// Output: [[1,6],[8,10],[15,18]]

// Test Case2:
/*
[[1,4]
,[4,5]]
*/
// Output: [[1,5]]

// Test Case3:
/*
[[4,7]
,[1,4]]
*/
// Output: [[1,7]]
func main() {
	intervals := pkg.CreateSlice2D[int]()
	fmt.Println(merge56R(intervals))
}
