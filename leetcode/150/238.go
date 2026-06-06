// 238. 除自身以外数组的乘积
/*

代码
测试用例
测试用例
测试结果
238. 除自身以外数组的乘积
已解答
中等
相关标签
premium lock icon
相关企业
提示
给你一个整数数组 nums，返回 数组 answer ，其中 answer[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积 。
题目数据 保证 数组 nums之中任意元素的全部前缀元素和后缀的乘积都在  32 位 整数范围内。
请 不要使用除法，且在 O(n) 时间复杂度内完成此题。

示例 1:
输入: nums = [1,2,3,4]
输出: [24,12,8,6]

示例 2:
输入: nums = [-1,1,0,-3,3]
输出: [0,0,9,0,0]

提示：

2 <= nums.length <= 10^5
-30 <= nums[i] <= 30
输入 保证 数组 answer[i] 在 32 位 整数范围内

进阶：你可以在 O(1) 的额外空间复杂度内完成这个题目吗？（ 出于对空间复杂度分析的目的，输出数组 不被视为 额外空间。）
*/
package main

import (
	"fmt"
	"lc/pkg"
)

func productExceptSelf(nums []int) []int {
	n := len(nums)
	left := make([]int, n)
	right := make([]int, n)
	left[0] = nums[0]
	right[n-1] = nums[n-1]
	for i := 1; i < n; i++ {
		left[i] = left[i-1] * nums[i]
	}
	for i := n - 2; i >= 0; i-- {
		right[i] = right[i+1] * nums[i]
	}

	var ans = make([]int, n)
	ans[0] = right[1]
	ans[n-1] = left[n-2]
	for i := 1; i < n-1; i++ {
		ans[i] = left[i-1] * right[i+1]
	}
	return ans
}

// Test Data 1: 4  1 2 3 4  expected output: 24 12 8 6
// Test Data 2: 5 -1 1 0 -3 3 expected output: 0 0 9 0 0
func main() {
	nums := pkg.CreateSlice[int]()
	fmt.Println(productExceptSelf(nums))
}
