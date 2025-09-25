/*
1218. 最长定差子序列
给你一个整数数组 arr 和一个整数 difference，请你找出并返回 arr 中最长等差子序列的长度，该子序列中相邻元素之间的差等于 difference 。
子序列 是指在不改变其余元素顺序的情况下，通过删除一些元素或不删除任何元素而从 arr 派生出来的序列。

示例 1：
输入：arr = [1,2,3,4], difference = 1
输出：4
解释：最长的等差子序列是 [1,2,3,4]。

示例 2：
输入：arr = [1,3,5,7], difference = 1
输出：1
解释：最长的等差子序列是任意单个元素。

示例 3：
输入：arr = [1,5,7,8,5,3,4,2,1], difference = -2
输出：4
解释：最长的等差子序列是 [7,5,3,1]。

提示：

1 <= arr.length <= 10^5
-10^4 <= arr[i], difference <= 10^4
*/
package main

import (
	"fmt"
	"lc/pkg"
)

func longestSubsequence(arr []int, difference int) int {
	/*
		n := len(arr)
		ans := 0
		dp := make([]int, n) //dp[i]：以arr[i]结尾的最长等差子序列长度（arr[i]必须被选取）
		for i := 0; i < n; i++ {
			dp[i] = 1
			for j := i - 1; j >= 0; j-- {
				if arr[i]-arr[j] == difference {
					dp[i] = max(dp[i], dp[j]+1) //下标 j 取满足 j<i 且 arr[j]=arr[i]−d 的所有下标的最大值即可，不用再往前找了。因此接一个break
					break                       //可以用map优化
				}
			}
			ans = max(ans, dp[i])
		}
		return ans
	*/
	/*数据长度最大是10^5,O(N^2)会超时，因此要找更快的解法*/
	dp := make(map[int]int) //FIXME：dp[i]表示以i结尾的最长等差子序列长度，而不是arr[i]，使用map来存储，就不用像上面的方法来遍历j了，这样就可以降低时间复杂度
	ans := 0
	for i := 0; i < len(arr); i++ {
		dp[arr[i]] = dp[arr[i]-difference] + 1
		ans = max(ans, dp[arr[i]])
	}
	return ans
}

func main() {
	var diff int
	fmt.Println("Input difference:")
	fmt.Scan(&diff)
	arr := pkg.CreateIntSlice[int]()
	fmt.Println(longestSubsequence(arr, diff))

}
