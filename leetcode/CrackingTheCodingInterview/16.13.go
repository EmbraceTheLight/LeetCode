// 16.13 平分正方形
/*
给定两个正方形及一个二维平面。请找出将这两个正方形分割成两半的一条直线。假设正方形顶边和底边与 x 轴平行。
每个正方形的数据square包含3个数值，正方形的左下顶点坐标[X,Y] = [square[0],square[1]]，以及正方形的边长square[2]。
所求直线穿过两个正方形会形成4个交点，请返回4个交点形成线段的两端点坐标（两个端点即为4个交点中距离最远的2个点，这2个点所连成的线段一定会穿过另外2个交点）。
2个端点坐标[X1,Y1]和[X2,Y2]的返回格式为{X1,Y1,X2,Y2}，要求若X1 != X2，需保证X1 < X2，否则需保证Y1 <= Y2。

若同时有多条直线满足要求，则选择斜率最大的一条计算并返回（与Y轴平行的直线视为斜率无穷大）。
提示：
square.length == 3
square[2] > 0
*/

package main

import (
	"fmt"
	. "lc/pkg"
	"sort"
)

func cutSquares(square1 []int, square2 []int) []float64 {
	type point1613 struct {
		x, y float64
	}

	// 获取正方形中心点坐标
	getCentralPoint := func(leftBottomPoint point1613, edgeLength int) point1613 {
		return point1613{
			x: leftBottomPoint.x + float64(edgeLength)/2,
			y: leftBottomPoint.y + float64(edgeLength)/2,
		}
	}

	leftBottomPoint1 := point1613{x: float64(square1[0]), y: float64(square1[1])}
	leftBottomPoint2 := point1613{x: float64(square2[0]), y: float64(square2[1])}
	centralPoint1 := getCentralPoint(leftBottomPoint1, square1[2])
	centralPoint2 := getCentralPoint(leftBottomPoint2, square2[2])

	var intersectionPoints []point1613
	// 中心点横坐标重合, 选取斜率最大的, 即与 y 轴平行
	if centralPoint1.x == centralPoint2.x {
		intersectionPoints = append(intersectionPoints,
			point1613{x: centralPoint1.x, y: centralPoint1.y - float64(square1[2])/2},
			point1613{x: centralPoint1.x, y: centralPoint1.y + float64(square1[2])/2},
			point1613{x: centralPoint2.x, y: centralPoint2.y - float64(square2[2])/2},
			point1613{x: centralPoint2.x, y: centralPoint2.y + float64(square2[2])/2})
	} else {
		// 中心点不重合, 则为两中心点所形成的直线
		k := (centralPoint1.y - centralPoint2.y) / (centralPoint1.x - centralPoint2.x)
		b := centralPoint1.y - k*centralPoint1.x

		// 交点落在正方形上下两条边
		if k*leftBottomPoint1.x+b > leftBottomPoint1.y+float64(square1[2]) || k*leftBottomPoint1.x+b < leftBottomPoint1.y {
			intersectionPoints = append(intersectionPoints,
				point1613{x: (leftBottomPoint1.y + float64(square1[2]) - b) / k, y: leftBottomPoint1.y + float64(square1[2])},
				point1613{x: (leftBottomPoint1.y - b) / k, y: leftBottomPoint1.y},
				point1613{x: (leftBottomPoint2.y + float64(square2[2]) - b) / k, y: leftBottomPoint2.y + float64(square2[2])},
				point1613{x: (leftBottomPoint2.y - b) / k, y: leftBottomPoint2.y},
			)
		} else {
			// 交点落在正方形左右两条边
			intersectionPoints = append(intersectionPoints,
				point1613{x: leftBottomPoint1.x, y: k*leftBottomPoint1.x + b},
				point1613{x: leftBottomPoint1.x + float64(square1[2]), y: k*(leftBottomPoint1.x+float64(square1[2])) + b},
				point1613{x: leftBottomPoint2.x, y: k*leftBottomPoint2.x + b},
				point1613{x: leftBottomPoint2.x + float64(square2[2]), y: k*(leftBottomPoint2.x+float64(square2[2])) + b})
		}
	}
	sort.Slice(intersectionPoints, func(i, j int) bool {
		if intersectionPoints[i].x == intersectionPoints[j].x {
			return intersectionPoints[i].y <= intersectionPoints[j].y
		}
		return intersectionPoints[i].x <= intersectionPoints[j].x
	})
	return []float64{intersectionPoints[0].x, intersectionPoints[0].y, intersectionPoints[3].x, intersectionPoints[3].y}
}

func main() {
	square1 := CreateSlice[int]()
	square2 := CreateSlice[int]()
	fmt.Println(cutSquares(square1, square2))
}
