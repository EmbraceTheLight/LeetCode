// 16.03. 交点
/*
给定两条线段（表示为起点start = {X1, Y1}和终点end = {X2, Y2}），如果它们有交点，请计算其交点，没有交点则返回空值。
要求浮点型误差不超过10^-6。若有多个交点（线段重叠）则返回 X 值最小的点，X 坐标相同则返回 Y 值最小的点。

提示：
坐标绝对值不会超过 2^7
输入的坐标均是有效的二维坐标
*/
package main

import (
	"fmt"
	. "lc/pkg"
)

type point struct {
	x, y float64
}

func intersection(start1 []int, end1 []int, start2 []int, end2 []int) []float64 {
	startPoint1 := point{float64(start1[0]), float64(start1[1])}
	endPoint1 := point{float64(end1[0]), float64(end1[1])}
	startPoint2 := point{float64(start2[0]), float64(start2[1])}
	endPoint2 := point{float64(end2[0]), float64(end2[1])}

	// 特殊情况: k = ∞
	if startPoint1.x == endPoint1.x {
		if startPoint2.x == endPoint2.x {
			checkPoint := point{startPoint1.x, max(min(startPoint1.y, endPoint1.y), min(startPoint2.y, endPoint2.y))} // 选择两条线段最小的两个纵坐标的最大值
			if startPoint1.x != startPoint2.x || check(checkPoint, startPoint1, endPoint1) == false || check(checkPoint, startPoint2, endPoint2) == false {
				return []float64{}
			} else {
				return []float64{checkPoint.x, checkPoint.y}
			}
		}
		k, b := getkb(startPoint2, endPoint2)
		y := k*startPoint1.x + b
		checkPoint := point{x: startPoint1.x, y: y}
		if check(checkPoint, startPoint1, endPoint1) == false || check(checkPoint, startPoint2, endPoint2) == false {
			return []float64{}
		} else {
			return []float64{startPoint1.x, y}
		}
	}
	if startPoint2.x == endPoint2.x {
		k, b := getkb(startPoint1, endPoint1)
		y := k*startPoint2.x + b
		checkPoint := point{x: startPoint2.x, y: y}
		if check(checkPoint, startPoint1, endPoint1) == false || check(checkPoint, startPoint2, endPoint2) == false {
			return []float64{}
		} else {
			return []float64{startPoint2.x, y}
		}
	}
	k1, b1 := getkb(startPoint1, endPoint1)
	k2, b2 := getkb(startPoint2, endPoint2)
	if k1 == k2 && b1 != b2 { // 平行
		return []float64{}
	} else if k1 == k2 && b1 == b2 { // 重合
		minX := max(min(startPoint1.x, endPoint1.x), min(startPoint2.x, endPoint2.x)) // 选择两条线段最小的两个横坐标的最大值
		minY := k1*minX + b1
		checkPoint := point{x: minX, y: minY}
		if check(checkPoint, startPoint1, endPoint1) == false || check(checkPoint, startPoint2, endPoint2) == false {
			return []float64{}
		}
		return []float64{checkPoint.x, checkPoint.y}
	}

	// 相交
	x := (b1 - b2) / (k2 - k1)
	y := k1*x + b1
	checkPoint := point{x: x, y: y}
	if check(checkPoint, startPoint1, endPoint1) == false || check(checkPoint, startPoint2, endPoint2) == false {
		return []float64{}
	}
	return []float64{x, y}
}

// 返回斜率和截距
func getkb(start, end point) (float64, float64) {
	k := (end.y - start.y) / (end.x - start.x)
	b := end.y - k*end.x
	return k, b
}

// 检查交点是否位于线段内, x, y 为两直线交点
func check(checkPoint, start, end point) bool {
	return checkPoint.x >= min(start.x, end.x) && checkPoint.x <= max(start.x, end.x) && checkPoint.y >= min(start.y, end.y) && checkPoint.y <= max(start.y, end.y)
}

// 示例 1：
// 输入：
// line1 = {0, 0}, {1, 0}
// line2 = {1, 1}, {0, -1}
// 输出： {0.5, 0}
//
// 示例 2：
// 输入：
// line1 = {0, 0}, {3, 3}
// line2 = {1, 1}, {2, 2}
// 输出： {1, 1}
//
// 示例 3：
// 输入：
// line1 = {0, 0}, {1, 1}
// line2 = {1, 0}, {2, 1}
// 输出： {}，两条线段没有交点
func main() {
	fmt.Println("Input start1, end1:")
	start1 := CreateSlice[int]()
	end1 := CreateSlice[int]()
	fmt.Println("Input start2, end2:")
	start2 := CreateSlice[int]()
	end2 := CreateSlice[int]()
	fmt.Println(intersection(start1, end1, start2, end2))
}
