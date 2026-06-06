package main

import (
	"fmt"
	"lc/pkg"
	"sort"
)

//func merge(intervals [][]int) [][]int {
//	if len(intervals) == 1 {
//		return intervals
//	}
//	sort.Slice(intervals, func(i, j int) bool {
//		if intervals[i][0] == intervals[j][0] {
//			return intervals[i][1] < intervals[j][1]
//		}
//		return intervals[i][0] < intervals[j][0]
//	})
//
//	var ans = make([][]int, 0)
//	first, second := 0, 1
//	for ; second < len(intervals); second++ {
//		var firstLeft, firstRight, secondLeft, secondRight = intervals[first][0], intervals[first][1], intervals[second][0], intervals[second][1]
//		if firstRight >= secondLeft {
//			for second < len(intervals) && firstRight >= secondLeft {
//				if firstRight <= secondRight {
//					firstRight = secondRight
//				}
//				second++
//				if second == len(intervals) {
//					break
//				}
//				secondLeft, secondRight = intervals[second][0], intervals[second][1]
//			}
//			secondLeft, secondRight = intervals[second-1][0], intervals[second-1][1]
//			ans = append(ans, []int{firstLeft, firstRight})
//		} else {
//			ans = append(ans, []int{firstLeft, firstRight})
//		}
//		first = second
//	}
//
//	for ; first < len(intervals); first++ {
//		ans = append(ans, intervals[first])
//	}
//	return ans
//}

func mergeR(intervals [][]int) [][]int {
	if len(intervals) == 1 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})
	var ans = make([][]int, 0)

	for i := 0; i < len(intervals); {
		j := i + 1
		for j < len(intervals) && intervals[i][1] >= intervals[j][0] {
			if intervals[i][1] < intervals[j][1] {
				intervals[i][1] = intervals[j][1]
			}
			j++
		}
		ans = append(ans, []int{intervals[i][0], intervals[i][1]})
		i = j
	}
	return ans
}

func main() {
	intervals := pkg.CreateSlice2D[int]()
	fmt.Println(mergeR(intervals))
}
