// 149. 直线上最多的点数
/*
给你一个数组 points ，其中 points[i] = [xi, yi] 表示 X-Y 平面上的一个点。求最多有多少个点在同一条直线上。

提示：
1 <= points.length <= 300
points[i].length == 2
-10^4 <= xi, yi <= 10^4
points 中的所有点 互不相同
*/
package main

import (
	"fmt"
	"lc/pkg"
	"math"
)

type point struct {
	x int
	y int
}

func maxPoints(points [][]int) int {
	var ans int

	for i := 0; i < len(points); i++ {
		mp := make(map[float64]int) // key: 斜率  value: 该斜率出现的次数
		x1, y1 := points[i][0], points[i][1]
		for j := i + 1; j < len(points); j++ {
			x2, y2 := points[j][0], points[j][1]
			k := math.MaxFloat64
			if x2-x1 != 0 {
				k = float64(y2-y1) / float64(x2-x1)
			}
			mp[k]++
			ans = max(ans, mp[k])
		}

	}
	ans++
	return ans
}

// calculateKB 返回两点组成的直线的斜率 + 截距
func calculateKB(x1, y1, x2, y2 int) (k float64, b float64) {
	return float64(y2-y1) / float64(x2-x1), float64(y1) - float64(x1)*float64(y2-y1)/float64(x2-x1)
}

// 示例 1：
// 输入：points = [[1,1],[2,2],[3,3]]
// 输出：3
//
// 示例 2：
// 输入：points = [[1,1],[3,2],[5,3],[4,1],[2,3],[1,4]]
// 输出：4
func main() {
	points := pkg.CreateSlice2D[int]()
	fmt.Println(maxPoints(points))
}
