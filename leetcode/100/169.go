// 169. 多数元素
/*
给定一个大小为 n 的数组 nums ，返回其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。
你可以假设数组是非空的，并且给定的数组总是存在多数元素。

提示：
n == nums.length
1 <= n <= 5 * 10^4
-10^9 <= nums[i] <= 10^9

输入保证数组中一定有一个多数元素。
进阶：尝试设计时间复杂度为 O(n)、空间复杂度为 O(1) 的算法解决此问题。
*/
package main

import (
	"fmt"
	"lc/pkg"
)

// * 投票法, 当 count = 0 时, 设置当前数字为多数元素 x, 当遇到非 x 的数时, count--, 当遇到 x 时, count++
// * 重复这一过程. 最终得到多数元素. 时间复杂度 O(n), 空间复杂度 O(1)
// * 因为多数元素大于 ⌊ n/2 ⌋, 所以最后 count > 0 的数字一定是多数元素
// * 当 count 变为 0 时, 最多有一半的真正的多数元素被抵消, 剩下的数字中它仍然占有多数
/*
证明: 当 count 变为 0 后, 剩下的数字中“多数元素”仍占有多数
设数组大小为 n, 其中“多数元素”为 x, 其出现总次数为 a, 剩下的数字中“多数元素” x 出现次数为 b,
当遍历了前 c 个数后, count 重新变为 0, 设剩下的数字个数为 d
求: b > d / 2

根据本题中“多数元素”的性质, 有 a > n / 2
设前 c 个数中多数元素出现了 a1 次, 有 0 < a1 <= c / 2, b = a - a1
有 a - a1 > (n - c) / 2  <--  (a 取 n / 2 + 1, a1 取 c / 2, 这种情况下剩余数字中“多数元素”出现的次数最少),
又由于 n = c + d, 所以 d = n - c
即 b > d / 2, 证毕
*/
func majorityElement(nums []int) int {
	n := len(nums)
	count := 0
	ans := nums[0]
	for i := 0; i < n; i++ {
		if count == 0 {
			ans = nums[i]
		}
		if nums[i] == ans {
			count++
		} else {
			count--
		}
	}
	return ans
}

// 示例 1：
// 输入：nums = [3,2,3]
// 输出：3
//
// 示例 2：
// 输入：nums = [2,2,1,1,1,2,2]
// 输出：2
func main() {
	nums := pkg.CreateSlice[int]()
	fmt.Println(majorityElement(nums))
}
