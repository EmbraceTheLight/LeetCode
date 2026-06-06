// 137. 只出现一次的数字 II
/*
给你一个整数数组 nums ，除某个元素仅出现 一次 外，其余每个元素都恰出现 三次 。请你找出并返回那个只出现了一次的元素。
你必须设计并实现线性时间复杂度的算法且使用常数级空间来解决此问题。

提示：
1 <= nums.length <= 3 * 10^4
-2^31 <= nums[i] <= 2^31 - 1
nums 中，除某个元素仅出现 一次 外，其余每个元素都恰出现 三次
*/
package main

import (
	"fmt"
	"lc/pkg"
	"math/bits"
)

// * 思路: 统计每个数字每一位中 1 出现的次数, 最后将它对 3 取模, 如果得到 1, 说明只出现一次的数字的这一位为 1, 否则为 0
func singleNumber137(nums []int) int {
	// 这里明确使用 int32, 否则 int 在 64 位机器上会被当作 int64
	var ret int32
	for i := 31; i >= 0; i-- {
		bit := int32(0)
		for j := 0; j < len(nums); j++ {
			bit += (int32(nums[j]) >> i) & 1
		}
		bit = bit % 3
		ret = ret | (bit << i)
	}
	return int(ret)
}

// * 状态机解法. 思路:
// * 设 ones 为只出现一次的 1 的位的集合, twos 为出现两次的 1 的位的集合, 如果某一位的 1 出现了三次, 则该位置零, 重新计数
func singleNumber137StateMachine(nums []int) int {
	// 示例: 有一组数, 它们的某一位分别为: 0 1 0 1 1 1 0, 那么 ones 和 twos 该位置值变化如下:
	//  |1 出现次数|	   |ones|	  |twos|
	//   0			    0			0
	//   1			    1			0
	//   1			    1			0
	//   2			    0			1
	//   3			    0			0		// 这里将 ones 和 twos 对应位置置零, 重新开始计数
	//   4			    1			0
	//   4			    1			0
	var ones, twos int32
	for _, num := range nums {
		// &^ 按位清除, 即按位与 & + 按位取反 ^,
		// 这里的意思是 twos 中对应位已经为 1 了,
		// 那么需要重置, 即将其取反为 0, 再与 ones 计算后的结果相与, 将 ones 对应位清零
		// 下面同理.
		ones = (ones ^ int32(num)) & ^twos
		twos = (twos ^ int32(num)) & ^ones
	}
	return int(ones)
}

// 示例 1：
// 输入：nums = [2,2,3,2]
// 输出：3
//
// 示例 2：
// 输入：nums = [0,1,0,1,0,1,99]
// 输出：99
func main() {

	fmt.Println(bits.UintSize)
	nums := pkg.CreateSlice[int]()
	fmt.Println(singleNumber137(nums))
	fmt.Println(singleNumber137StateMachine(nums))
}
