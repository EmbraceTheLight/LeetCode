package main

import (
	"fmt"
	"lc/pkg"
	"sort"
)

func changeR(amount int, coins []int) int {
	sort.Slice(coins, func(i, j int) bool { return coins[i] < coins[j] })
	var dp = make([]int, amount+1) //dp[i]表示 组合成面值 i 的组合数
	dp[0] = 1
	// ! 错误解法。这么做会有重复计数。比如有面值 1，2 两种硬币，要组成金额 3，这种解法就会包含 1 + 2 和 2 + 1 两种情况。
	//for i := 1; i <= amount; i++ {
	//	for j := 0; j < len(coins); j++ {
	//		if coins[j] <= i {
	//			dp[i] += dp[i-coins[j]] // 先组合成面值 i-coins[j]，再加上面值 coins[j] ，凑成面值 i 的组合数
	//		} else {
	//			break
	//		}
	//	}
	//}

	// 外层循环遍历硬币面值，内层循环遍历金额。求的是加入该面值硬币可达到的金额的组合数。
	// 这种解法相比于上面的错误解法，不会有重复计数。因为这是向组合中追加新面额。
	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			dp[j] += dp[j-coins[i]]
		}
	}
	return dp[amount]
}
func main() {
	var amount int
	fmt.Scan(&amount)
	coins := pkg.CreateSlice[int]()
	fmt.Println(changeR(amount, coins))
}
