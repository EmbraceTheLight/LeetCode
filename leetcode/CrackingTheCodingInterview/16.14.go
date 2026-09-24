// 16.14. 最佳直线
/*
给定一个二维平面及平面上的 N 个点列表Points，其中第i个点的坐标为Points[i]=[Xi,Yi]。请找出一条直线，其通过的点的数目最多。
设穿过最多点的直线所穿过的全部点编号从小到大排序的列表为S，
你仅需返回[S[0],S[1]]作为答案，若有多条直线穿过了相同数量的点，则选择S[0]值较小的直线返回，S[0]相同则选择S[1]值较小的直线返回。

提示：
2 <= len(Points) <= 300
len(Points[i]) = 2
*/
package main

import (
	"fmt"
	. "lc/pkg"
	"math"
	"sort"
)

type line1614 struct {
	k, b float64
}
type point1614 struct {
	x, y float64
	idx  int
}

func bestLine(points [][]int) []int {
	mp := make(map[line1614][]point1614)
	visit := make(map[line1614]map[point1614]bool)
	for i := 0; i < len(points); i++ {
		point1 := point1614{x: float64(points[i][0]), y: float64(points[i][1]), idx: i}
		for j := i + 1; j < len(points); j++ {
			line := line1614{}
			point2 := point1614{x: float64(points[j][0]), y: float64(points[j][1]), idx: j}
			if point1.x == point2.x {
				k := math.Inf(1)
				b := point1.x
				line = line1614{k: k, b: b}
				if _, ok := visit[line]; !ok || visit[line][point2] == false {
					if !ok {
						visit[line] = make(map[point1614]bool)
					}
					visit[line][point2] = true
					mp[line] = append(mp[line], point2)
				}
			} else {
				k := (point2.y - point1.y) / (point2.x - point1.x)
				b := point1.y - k*point1.x
				line = line1614{k: k, b: b}
				if _, ok := visit[line]; !ok || visit[line][point2] == false {
					if !ok {
						visit[line] = make(map[point1614]bool)
					}
					visit[line][point2] = true
					mp[line] = append(mp[line], point2)
				}
			}
			if _, ok := visit[line]; !ok || visit[line][point1] == false {
				if !ok {
					visit[line] = make(map[point1614]bool)
				}
				visit[line][point1] = true
				mp[line] = append(mp[line], point1)
			}
		}
	}
	var tmp []point1614
	for _, v := range mp {
		sort.Slice(v, func(i, j int) bool {
			return v[i].idx < v[j].idx
		})
		if len(v) > len(tmp) {
			tmp = v
		}
		if len(v) == len(tmp) {
			if tmp[0].idx > v[0].idx || (tmp[0].idx == v[0].idx && tmp[1].idx > v[1].idx) {
				tmp = v
			}
		}
	}
	return []int{tmp[0].idx, tmp[1].idx}
}

// 示例：
// 输入： [[0,0],[1,1],[1,0],[2,0]]
// 输出： [0,2]
// 解释： 所求直线穿过的3个点的编号为[0,2,3]
func main() {
	points := CreateSlice2D[int]()
	fmt.Println(bestLine(points))
}
