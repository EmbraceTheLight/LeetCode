/*
474. 一和零
给你一个二进制字符串数组 strs 和两个整数 m 和 n 。
请你找出并返回 strs 的最大子集的长度，该子集中 最多 有 m 个 0 和 n 个 1 。
如果 x 的所有元素也是 y 的元素，集合 x 是集合 y 的 子集 。

示例 1：
输入：strs = ["10", "0001", "111001", "1", "0"], m = 5, n = 3
输出：4
解释：最多有 5 个 0 和 3 个 1 的最大子集是 {"10","0001","1","0"} ，因此答案是 4 。
其他满足题意但较小的子集包括 {"0001","1"} 和 {"10","1","0"} 。{"111001"} 不满足题意，因为它含 4 个 1 ，大于 n 的值 3 。

示例 2：
输入：strs = ["10", "0", "1"], m = 1, n = 1
输出：2
解释：最大的子集是 {"0", "1"} ，所以答案是 2 。

提示：
1 <= strs.length <= 600
1 <= strs[i].length <= 100
strs[i] 仅由 '0' 和 '1' 组成
1 <= m, n <= 100
*/
package main

import (
	"fmt"
	"lc/pkg"
)

func findMaxForm(strs []string, m int, n int) int {
	//有两种容量的0-1背包，使用三维dp来解。难点在于dp数组初始化的过程
	length := len(strs)
	nums0 := make([]int, length+1) //nums0[i]：第i个子串包含0的个数
	nums1 := make([]int, length+1) //nums1[i]：第i个子串包含1的个数
	for i := 0; i < length; i++ {
		for _, char := range strs[i] {
			if char == '0' {
				nums0[i+1]++
			} else {
				nums1[i+1]++
			}
		}
	}
	dp := make([][][]int, length+1) //dp[i][j][k]：选择前i个子串，且最多j个0的情况下和最多k个1的情况下的子集最大长度
	for i := 0; i <= length; i++ {
		dp[i] = make([][]int, m+1)
		for j := 0; j <= m; j++ {
			dp[i][j] = make([]int, n+1)
		}
	}

	for i := 1; i <= length; i++ {
		for j := 0; j <= m; j++ {
			for k := 0; k <= n; k++ {
				if nums0[i] > j || nums1[i] > k {
					dp[i][j][k] = dp[i-1][j][k]
				} else {
					dp[i][j][k] = max(dp[i-1][j][k], dp[i-1][j-nums0[i]][k-nums1[i]]+1)
				}
			}
		}
	}

	return dp[length][m][n]
}
func main() {
	//var s string
	//strs := make([]string, 0)
	//fmt.Scan(&s)
	//for s != "\n" {
	//	strs = append(strs, s)
	//	fmt.Scan(&s)
	//}

	var m, n int
	fmt.Println("Input m,n")
	fmt.Scan(&m, &n)
	strs := pkg.CreateSlice[string]()
	fmt.Println(findMaxForm(strs, m, n))
}
