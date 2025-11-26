// 57. 插入区间
/*
给你一个 无重叠的 ，按照区间起始端点排序的区间列表 intervals，其中 intervals[i] = [starti, endi] 表示第 i 个区间的开始和结束，并且 intervals 按照 starti 升序排列。
同样给定一个区间 newInterval = [start, end] 表示另一个区间的开始和结束。
在 intervals 中插入区间 newInterval，使得 intervals 依然按照 starti 升序排列，且区间之间不重叠（如果有必要的话，可以合并区间）。
返回插入之后的 intervals。
注意 你不需要原地修改 intervals。你可以创建一个新数组然后返回它。

示例 1：
输入：intervals = [[1,3],[6,9]], newInterval = [2,5]
输出：[[1,5],[6,9]]

示例 2：
输入：intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
输出：[[1,2],[3,10],[12,16]]
解释：这是因为新的区间 [4,8] 与 [3,5],[6,7],[8,10] 重叠。

提示：
0 <= intervals.length <= 10^4
intervals[i].length == 2
0 <= starti <= endi <= 10^5
intervals 根据 starti 按 升序 排列
newInterval.length == 2
0 <= start <= end <= 10^5
*/
package main

import (
	"fmt"
	"lc/pkg"
)

func insert(intervals [][]int, newInterval []int) [][]int {
	n := len(intervals)
	if n == 0 {
		return [][]int{newInterval}
	}
	var ans [][]int
	left, right := intervals[0][0], intervals[0][1]
	i := 0

	// 可以使用二分法查找插入位置
	for ; i < n; i++ {
		if newInterval[0] <= intervals[i][0] {
			break
		}
	}

	newIntervals := make([][]int, n+1)
	for j := 0; j < i; j++ {
		newIntervals[j] = intervals[j]
	}
	newIntervals[i] = newInterval
	for j := i; j < n; j++ {
		newIntervals[j+1] = intervals[j]
	}

	left, right = newIntervals[0][0], newIntervals[0][1]
	for j := 1; j < n+1; j++ {
		if right < newIntervals[j][0] {
			ans = append(ans, []int{left, right})
			left, right = newIntervals[j][0], newIntervals[j][1]
		} else if right < newIntervals[j][1] {
			right = newIntervals[j][1]
		}
	}
	ans = append(ans, []int{left, right})
	return ans
}

// Test Case1:
/*
newInterval = [2,5]

intervals:
[[1,3]
,[6,9]]
^D

*/
// Output: [[1,5],[6,9]]

// Test Case2:
/*
newInterval = [4,8]

intervals:
[[1,2]
,[3,5]
,[6,7]
,[8,10]
,[12,16]]
^D

*/
// Output: [[1,2],[3,10],[12,16]]
func main() {
	newInterval := pkg.CreateSlice[int]()
	intervals := pkg.CreateSlice2D[int]()
	fmt.Println(insert(intervals, newInterval))
}
