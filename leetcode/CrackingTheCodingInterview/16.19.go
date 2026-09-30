// 16.19. 水域大小
/*
你有一个用于表示一片土地的整数矩阵land，该矩阵中每个点的值代表对应地点的海拔高度。若值为0则表示水域。由垂直、水平或对角连接的水域为池塘。
池塘的大小是指相连接的水域的个数。编写一个方法来计算矩阵中所有池塘的大小，返回值需要从小到大排序。

提示：
0 < len(land) <= 1000
0 < len(land[i]) <= 1000
*/
package main

import (
	"fmt"
	. "lc/pkg"
	"slices"
)

func pondSizes(land [][]int) []int {
	var ans []int
	for i := 0; i < len(land); i++ {
		for j := 0; j < len(land[0]); j++ {
			if land[i][j] == 0 {
				var sum int
				dfs1619(land, i, j, &sum)
				ans = append(ans, sum)
			}
		}
	}
	slices.Sort(ans)
	return ans
}

func dfs1619(land [][]int, x, y int, sum *int) {
	if x < 0 || x >= len(land) || y < 0 || y >= len(land[0]) || land[x][y] != 0 {
		return
	}
	land[x][y] = -1
	*sum = *sum + 1
	dfs1619(land, x-1, y, sum)
	dfs1619(land, x+1, y, sum)
	dfs1619(land, x, y-1, sum)
	dfs1619(land, x, y+1, sum)
	dfs1619(land, x+1, y+1, sum)
	dfs1619(land, x-1, y+1, sum)
	dfs1619(land, x-1, y-1, sum)
	dfs1619(land, x+1, y-1, sum)
}

// 示例：
// 输入：[[0,2,1,0],[0,1,0,1],[1,1,0,1],[0,1,0,1]]
// 输出：[1,2,4]
func main() {
	land := CreateSlice2D[int]()
	fmt.Println(pondSizes(land))
}
