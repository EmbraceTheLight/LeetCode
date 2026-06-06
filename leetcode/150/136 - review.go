// 136. 只出现一次的数字
/*
给你一个非空整数数组 nums ，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
你必须设计并实现线性时间复杂度的算法来解决此问题，且该算法只使用常量额外空间。

提示：
1 <= nums.length <= 3 * 10^4
-3 * 10^4 <= nums[i] <= 3 * 10^4
除了某个元素只出现一次以外，其余每个元素均出现两次。
*/
package main

import (
	"fmt"
	"lc/pkg"
)

// * 本题核心:
// * 1. 相同数字异或结果为 0
// * 2. 异或满足交换律和结合律
func singleNumberR(nums []int) int {
	var ans int
	for i := 0; i < len(nums); i++ {
		ans ^= nums[i]
	}
	return ans
}

// 示例 1：
// 输入：nums = [2,2,1]
// 输出：1
//
// 示例 2：
// 输入：nums = [4,1,2,1,2]
// 输出：4
//
// 示例 3：
// 输入：nums = [1]
// 输出：1
func main() {
	nums := pkg.CreateSlice[int]()
	fmt.Println(singleNumberR(nums))
}
