// 面试题 08.02. 迷路的机器人
/*
设想有个机器人坐在一个网格的左上角，网格 r 行 c 列。机器人只能向下或向右移动，但不能走到一些被禁止的网格（有障碍物）。设计一种算法，寻找机器人从左上角移动到右下角的路径。

网格中的障碍物和空位置分别用 1 和 0 来表示。
返回一条可行的路径，路径由经过的网格的行号和列号组成。左上角为 0 行 0 列。如果没有可行的路径，返回空数组。

示例 1：
输入：[[0,0,0],[0,1,0],[0,0,0]]
输出：[[0,0],[0,1],[0,2],[1,2],[2,2]]
解释：
输入中标粗的位置即为输出表示的路径，即
0 行 0 列（左上角） -> 0 行 1 列 -> 0 行 2 列 -> 1 行 2 列 -> 2 行 2 列（右下角）
说明：r 和 c 的值均不超过 100。
*/
package main

import (
	"fmt"
	. "lc/pkg"
)

func pathWithObstacles(obstacleGrid [][]int) [][]int {
	type point struct {
		x, y int
	}
	if obstacleGrid == nil || len(obstacleGrid) == 0 {
		return nil
	}
	rows, cols := len(obstacleGrid), len(obstacleGrid[0])
	paths := make([][]int, rows+cols-1)
	dp := make([][]bool, rows)
	previous := make([][]point, rows) // previous 记录路径. previous[i][j] 表示到达 (i,j) 的上一个点的坐标
	for i := 0; i < rows; i++ {
		dp[i] = make([]bool, cols)
		previous[i] = make([]point, cols)
	}
	dp[0][0] = obstacleGrid[0][0] == 0
	previous[0][0] = point{-1, -1}
	for i := 1; i < rows; i++ {
		dp[i][0] = dp[i-1][0] && obstacleGrid[i][0] == 0
		if dp[i][0] == true {
			previous[i][0] = point{i - 1, 0}
		}
	}
	for i := 1; i < cols; i++ {
		dp[0][i] = dp[0][i-1] && obstacleGrid[0][i] == 0
		if dp[0][i] == true {
			previous[0][i] = point{0, i - 1}
		}
	}
	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			dp[i][j] = (dp[i][j-1] || dp[i-1][j]) && obstacleGrid[i][j] == 0

			if dp[i][j] == true {
				if dp[i][j-1] == true {
					previous[i][j] = point{i, j - 1}
				} else {
					previous[i][j] = point{i - 1, j}
				}
			}
		}
	}
	if dp[rows-1][cols-1] == false {
		return nil
	}
	for i, j := rows-1, cols-1; i >= 0 && j >= 0; i, j = previous[i][j].x, previous[i][j].y {
		paths[i+j] = []int{i, j}
	}
	return paths
}

// 示例 1：
// 输入：[[0,0,0],[0,1,0],[0,0,0]]
// 输出：[[0,0],[0,1],[0,2],[1,2],[2,2]]
func main() {
	fmt.Println("Input obstacleGrid:")
	obstacleGrid := CreateSlice2D[int]()
	fmt.Println(pathWithObstacles(obstacleGrid))
}
