// CC3 多少个点位于同一直线
/*
描述
对于给定的n个位于同一二维平面上的点，求最多能有多少个点位于同一直线上
*/
package main

import (
	"fmt"
	"lc/pkg"
	"math"
)

type Point struct {
	X int
	Y int
}

func maxPoints(points []*Point) int {
	// write code here
	var ans int
	if len(points) == 0 {
		return 0
	}
	for i := 0; i < len(points); i++ {
		same := 0 // 与"原点"重合的点计数 原点见 lc hot.150 No.149
		maxCount := 0
		mp := make(map[float64]int)
		var k float64
		for j := i + 1; j < len(points); j++ {
			if points[j].X == points[i].X && points[j].Y == points[i].Y {
				same++
				continue
			}
			k = math.MaxInt32
			if points[j].X-points[i].X != 0 {
				k = float64(points[j].Y-points[i].Y) / float64(points[j].X-points[i].X)
			}
			mp[k]++
			if maxCount < mp[k] {
				maxCount = mp[k]
			}
		}
		ans = max(ans, maxCount+same)
	}

	return ans + 1
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 示例1
// 输入: [(0,0),(0,1)]
// 返回值：2
//
// 示例2
// 输入: [(2,3),(3,3),(-5,3)]
// 返回值: 3
func main() {
	input := pkg.CreateSlice2D[int]()
	points := make([]*Point, 0)
	for _, point := range input {
		points = append(points, &Point{point[0], point[1]})
	}
	fmt.Println(maxPoints(points))
}
