/*
646. 最长数对链
给你一个由 n 个数对组成的数对数组 pairs ，其中 pairs[i] = [lefti, righti] 且 lefti < righti 。
现在，我们定义一种 跟随 关系，当且仅当 b < c 时，数对 p2 = [c, d] 才可以跟在 p1 = [a, b] 后面。我们用这种形式来构造 数对链 。
找出并返回能够形成的 最长数对链的长度 。
你不需要用到所有的数对，你可以以任何顺序选择其中的一些数对来构造。

示例 1：
输入：pairs = [[1,2], [2,3], [3,4]]
输出：2
解释：最长的数对链是 [1,2] -> [3,4] 。

示例 2：
输入：pairs = [[1,2],[7,8],[4,5]]
输出：3
解释：最长的数对链是 [1,2] -> [4,5] -> [7,8] 。

提示：

n == pairs.length
1 <= n <= 1000
-1000 <= lefti < righti <= 1000
*/
package main

import (
	"fmt"
	"math"
	"sort"
	"lc/pkg"
)

func findLongestChain(pairs [][]int) int {
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i][1] != pairs[j][1] {
			return pairs[i][1] < pairs[j][1]
		} else {
			return pairs[i][0] < pairs[j][0]
		}
	})
	n := len(pairs)
	dp := make([]int, n)
	ans := 0
	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := i - 1; j >= 0; j-- {
			if pairs[j][1] < pairs[i][0] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		ans = max(ans, dp[i])
	}
	return ans
}

// findLongestChainBinary 贪心
func findLongestChainBinary(pairs [][]int) int {
	//按照右端点进行升序排序,这样就可以选择右端点最小的，为后续查找提供更多选择空间
	//FIXME：我们只关心右端点
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][1] < pairs[j][1]
	})
	ans := 0
	cur := math.MinInt32 //初始为最小值，可以在后面接入第一个数对,表示目前右端点的值
	for i := 0; i < len(pairs); i++ {
		if cur < pairs[i][0] {
			cur = pairs[i][1]
			ans++
		}
	}
	return ans
}
func main() {
	pairs := pkg.CreateIntSlice2[int]()
	fmt.Println(findLongestChain(pairs))
}
