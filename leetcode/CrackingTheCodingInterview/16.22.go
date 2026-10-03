// 16.22 兰顿蚂蚁🐜
/*
一只蚂蚁坐在由白色和黑色方格构成的无限网格上。开始时，网格全白，蚂蚁面向右侧。每行走一步，蚂蚁执行以下操作。
(1) 如果在白色方格上，则翻转方格的颜色，向右(顺时针)转 90 度，并向前移动一个单位。
(2) 如果在黑色方格上，则翻转方格的颜色，向左(逆时针方向)转 90 度，并向前移动一个单位。

编写程序来模拟蚂蚁执行的前 K 个动作，并返回最终的网格。
网格由数组表示，每个元素是一个字符串，代表网格中的一行，黑色方格由 'X' 表示，
白色方格由 '_' 表示，蚂蚁所在的位置由 'L', 'U', 'R', 'D' 表示，
分别表示蚂蚁 左、上、右、下 的朝向。只需要返回能够包含蚂蚁走过的所有方格的最小矩形。

说明：
K <= 100000
*/
package main

import (
	"fmt"
	"strings"
)

type point1622 struct {
	x, y int
}
type langtonsAnt struct {
	pos                    point1622
	face                   byte
	minX, minY, maxX, maxY int
	pointsMap              map[point1622]byte
}

func newLantonAnt() *langtonsAnt {
	return &langtonsAnt{
		pos:       point1622{},
		face:      'R',
		minX:      0,
		minY:      0,
		maxX:      0,
		maxY:      0,
		pointsMap: make(map[point1622]byte),
	}
}

func (ant *langtonsAnt) move() {
	ant.toggleFace()
	ant.togglePointColor()
	switch ant.face {
	case 'R':
		ant.pos.x = ant.pos.x + 1
		ant.updateRectBoundary()
	case 'L':
		ant.pos.x = ant.pos.x - 1
		ant.updateRectBoundary()
	case 'D':
		ant.pos.y = ant.pos.y - 1
		ant.updateRectBoundary()
	case 'U':
		ant.pos.y = ant.pos.y + 1
		ant.updateRectBoundary()
	default:
	}
}

// updateRectBoundary 更新蚂蚁移动的最小矩形边界
func (ant *langtonsAnt) updateRectBoundary() {
	pos := ant.pos
	ant.minX = min(ant.minX, pos.x)
	ant.minY = min(ant.minY, pos.y)
	ant.maxX = max(ant.maxX, pos.x)
	ant.maxY = max(ant.maxY, pos.y)
}

func (ant *langtonsAnt) toggleFace() {
	isBlack := ant.isBlack(ant.pos)
	switch ant.face {
	case 'R':
		if isBlack == true {
			ant.face = 'U'
		} else {
			ant.face = 'D'
		}
	case 'L':
		if isBlack == true {
			ant.face = 'D'
		} else {
			ant.face = 'U'
		}
	case 'U':
		if isBlack == true {
			ant.face = 'L'
		} else {
			ant.face = 'R'
		}
	case 'D':
		if isBlack == true {
			ant.face = 'R'
		} else {
			ant.face = 'L'
		}
	}
}

// togglePointColor 翻转区块颜色
func (ant *langtonsAnt) togglePointColor() {
	if ant.isBlack(ant.pos) == true {
		ant.pointsMap[ant.pos] = '_'
	} else {
		ant.pointsMap[ant.pos] = 'X'
	}
}

// isBlack 当前 pos 位置颜色
func (ant *langtonsAnt) isBlack(pos point1622) bool {
	if _, ok := ant.pointsMap[pos]; !ok {
		return false
	}
	return ant.pointsMap[pos] == 'X'
}

func (ant *langtonsAnt) Rectangle() []string {
	rows := ant.maxY - ant.minY + 1
	ans := make([]string, rows)
	for i := ant.maxY; i >= ant.minY; i-- {
		var sb strings.Builder
		for j := ant.minX; j <= ant.maxX; j++ {
			if ant.pos.x == j && ant.pos.y == i {
				sb.WriteByte(ant.face)
			} else if ant.isBlack(point1622{x: j, y: i}) == true {
				sb.WriteByte('X')
			} else {
				sb.WriteByte('_')
			}
		}
		ans[ant.maxY-i] = sb.String()
	}
	return ans
}
func printKMoves(K int) []string {
	lantonAnt := newLantonAnt()
	for i := 0; i < K; i++ {
		lantonAnt.move()
	}
	return lantonAnt.Rectangle()
}

// 示例 1：
// 输入：0
// 输出：["R"]
//
// 示例 2：
// 输入：2
// 输出：
// [
// "_X",
// "LX"
// ]
//
// 示例 3：
// 输入：5
// 输出：
// [
// "_U",
// "X_",
// "XX"
// ]
func main() {
	var k int
	fmt.Println("Input k:")
	fmt.Scan(&k)
	fmt.Println(printKMoves(k))
}
