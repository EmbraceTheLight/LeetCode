// 167. 两数之和 II - 输入有序数组
/*
给你一个下标从 1 开始的整数数组 numbers ，该数组已按 非递减顺序排列，
请你从数组中找出满足相加之和等于目标数 target 的两个数。如果设这两个数分别是 numbers[index1] 和 numbers[index2]，
则 1 <= index1 < index2 <= numbers.length 。
以长度为 2 的整数数组 [index1, index2] 的形式返回这两个整数的下标 index1 和 index2。
你可以假设每个输入 只对应唯一的答案 ，而且你 不可以 重复使用相同的元素。

你所设计的解决方案必须只使用常量级的额外空间。

示例 1：
输入：numbers = [2,7,11,15], target = 9
输出：[1,2]
解释：2 与 7 之和等于目标数 9 。因此 index1 = 1, index2 = 2 。返回 [1, 2]。

示例 2：
输入：numbers = [2,3,4], target = 6
输出：[1,3]
解释：2 与 4 之和等于目标数 6 。因此 index1 = 1, index2 = 3 。返回 [1, 3]。

示例 3：
输入：numbers = [-1,0], target = -1
输出：[1,2]
解释：-1 与 0 之和等于目标数 -1 。因此 index1 = 1, index2 = 2 。返回 [1, 2] 。

提示：

2 <= numbers.length <= 3 * 10^4
-1000 <= numbers[i] <= 1000
numbers 按 非递减顺序 排列
-1000 <= target <= 1000
仅存在一个有效答案
*/
package main

import (
	"fmt"
	"lc/pkg"
)

func twoSum(numbers []int, target int) []int {
	n := len(numbers)
	left, right := 0, n-1
	ans := []int{-1, -1}
	for left < right {
		s := numbers[left] + numbers[right]
		if s == target {
			ans[0], ans[1] = left+1, right+1 // 因为数组下标从 1 开始
			break
		} else if s < target {
			left++
		} else if s > target {
			right--
		}
	}
	return ans
}

// Test Case 1: target = 9, numbers = [2,7,11,15]	Output: [1,2]
// Test Case 2: target = 6, numbers = [2,3,4]	Output: [1,3]
// Test Case 3: target = -1, numbers = [-1,0]	Output: [1,2]
func main() {
	fmt.Println("Input target")
	var target int
	fmt.Scan(&target)
	fmt.Println("Input numbers")
	numbers := pkg.CreateSlice[int]()
	fmt.Println(twoSum(numbers, target))
}
