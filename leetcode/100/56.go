/*
56. 合并区间
以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。
示例 1：
输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
输出：[[1,6],[8,10],[15,18]]
解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].

示例 2：
输入：intervals = [[1,4],[4,5]]
输出：[[1,5]]
解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。

提示：

1 <= intervals.length <= 10^4
intervals[i].length == 2
0 <= starti <= endi <= 10^4
*/
package main

import (
	"fmt"
	"sort"
)

type Slices [][]int

func (s *Slices) Len() int {
	return len(*s)
}
func (s *Slices) Less(i int, j int) bool {
	return (*s)[i][0] < (*s)[j][0]
}
func (s *Slices) Swap(i, j int) {
	(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
}
func merge(intervals Slices) [][]int {
	if len(intervals) < 2 {
		return intervals
	}
	sort.Sort(&intervals)
	var ans [][]int
	for i := 0; i < len(intervals); {
		j := i + 1
		for j < len(intervals) && intervals[i][1] >= intervals[j][0] {
			if intervals[i][1] < intervals[j][1] {
				intervals[i][1] = intervals[j][1]
			}
			j++
		}
		//if j == len(intervals) {
		//	j = j - 1
		//}
		ans = append(ans, intervals[i])
		i = j
	}
	return ans
}
func main() {
	var in int
	var cnt = 1
	fmt.Println("Creating Slice... Input -1 to quit")
	fmt.Scan(&in)
	var Sli = make([][]int, 0)
	for in != -1 {
		fmt.Println("Input Slice ", cnt, ":")
		var subSli = make([]int, 2)
		for i := 0; i < 2; i++ {
			subSli[i] = in
			fmt.Scan(&in)
		}
		cnt++
		Sli = append(Sli, subSli)
	}
	fmt.Printf("%v", merge(Sli))
}
