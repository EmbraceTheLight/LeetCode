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
	// * 思路1: 遍历每个点, 以其为原点, 其他点到它组成的直线的斜率
	// * 斜率相同的点, 它们一定在同一条直线上
	// * 这样一来, 就不需要考虑截距的问题了, 只考虑斜率即可
	for i := 0; i < len(points); i++ {
		// 临时 map, 统计以 points[i] 作为坐标原点, 其他点与它组成的直线的斜率的次数
		mp := make(map[float64]int) // key: 斜率  value: 该斜率出现的次数.
		x1, y1 := points[i][0], points[i][1]
		for j := i + 1; j < len(points); j++ {
			x2, y2 := points[j][0], points[j][1]
			// 这里注意: 考虑斜率为无穷的情况, 即 x1 == x2, 此时设置 k 的默认值为 math.MaxFloat64
			// 如果不设置 k 的默认值, 最终对于同一条直线上的点如 (0,0) (0,1) (0,-1), 会得到 -inf 和 +inf 两个"斜率", 导致结果错误
			k := math.MaxFloat64
			if x2-x1 != 0 {
				k = float64(y2-y1) / float64(x2-x1)
			}
			mp[k]++
			ans = max(ans, mp[k]) // ans 最后记得加一, 因为直线上的点包含"原点"本身
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
