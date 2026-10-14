// 17.09. 第 k 个数
/*
有些数的素因子只有 3，5，7，请设计一个算法找出第 k 个数。
注意，不是必须有这些素因子，而是必须不包含其他的素因子。
例如，前几个数按顺序应该是 1，3，5，7，9，15，21。
*/
package main

import "fmt"

// getKthMagicNumber 返回第 k 个仅含素因子 3、5、7 的正整数。
// 输入：k 从 1 开始计数，因此 k 必须大于等于 1，且 1 被视为第一个数。
// 算法同时维护乘以 3、5、7 得到的三条递增候选序列，每轮取三者中的最小值。
// 时间复杂度为 O(k)，空间复杂度为 O(k)。
func getKthMagicNumber(k int) int {
	// dp[i] 表示按升序排列的第 i 个目标数。
	// 这里使用从 1 开始的下标，所以申请 k+1 个位置，并令 dp[1]=1。
	dp := make([]int, k+1)
	dp[1] = 1

	// idx3、idx5、idx7 分别指向 dp 中尚未用于生成下一个候选值的位置：
	// dp[idx3]*3、dp[idx5]*5、dp[idx7]*7 是三条候选序列当前最小的未处理元素。
	var idx3, idx5, idx7 int = 1, 1, 1

	// dp[1] 已经初始化为 1，因此从第 2 个数开始生成，直到得到第 k 个数。
	for i := 2; i <= k; i++ {
		// 任意一个新的目标数，都可以由之前某个目标数乘以 3、5 或 7 得到。
		// 由于 dp 本身有序，所以每个索引向后移动时，对应的候选序列也保持递增。
		next3, next5, next7 := dp[idx3]*3, dp[idx5]*5, dp[idx7]*7

		// 三条候选序列中最小的未处理值，就是整体序列的下一个数。
		next := min(next3, next5, next7)
		dp[i] = next

		// 必须分别使用 if，而不能使用 else if，因为同一个数可能由不同候选序列生成。
		// 例如 15 既等于 5*3，也等于 3*5；dp 中只添加一次 15，
		// 但产生 15 的 idx3 和 idx5 都必须后移，否则下一轮还会再次生成 15。
		if next == next3 {
			idx3++
		}
		if next == next5 {
			idx5++
		}
		if next == next7 {
			idx7++
		}
	}
	// dp 使用 1-based 下标，因此第 k 个数保存在 dp[k]。
	return dp[k]
}

// 示例 1：
// 输入：k = 5
// 输出：9
func main() {
	var k int
	fmt.Scan(&k)
	fmt.Println(getKthMagicNumber(k))
}
