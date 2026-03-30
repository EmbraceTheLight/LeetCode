// 75. 颜色分类
/*
给定一个包含红色、白色和蓝色、共 n 个元素的数组 nums ，原地 对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。
我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。
必须在不使用库内置的 sort 函数的情况下解决这个问题。

提示：
n == nums.length
1 <= n <= 300
nums[i] 为 0、1 或 2

进阶：
你能想出一个仅使用常数空间的一趟扫描算法吗？
*/
package main

import (
	"fmt"
	"lc/pkg"
)

func sortColors(nums []int) {
	n := len(nums)
	i, j, k := 0, n-1, n-1 // i 从左向右, j,k 从右向左. nums[i] 下一个非 0 数. nums[j] 下一个等于 0 的数. nums[k] 下一个等于 1 的数
	for ; i < n && nums[i] == 0; i++ {
	}
	for ; j >= 0 && nums[j] != 0; j-- {
	}
	for ; k >= 0 && nums[k] != 1; k-- {
	}
	updateIJK := func() {
		i++
		for ; j >= 0 && nums[j] != 0; j-- {
		}
		for ; k >= 0 && nums[k] != 1; k-- {
		}
	}

	// 思路: 将左侧较大的数 nums[i] 与右侧较小的数 nums[j] nums[k] 进行交换.
	// 递增 i, 直到 i 右侧没有可以交换的数字 i >= j && i >= k
	// 若有可以交换的数字, 交换, 并递增 i, 递减 j, k 直到找到下一个满足要求的数
	for i < n && (i < j || i < k) {
		if nums[i] == 1 { // 当前遍历到的数字为 1
			if i < j { // 在 i 右侧的数字中还有等于 0, 将二者交换
				nums[i], nums[j] = nums[j], nums[i]
				if j >= k { // 交换后, nums[j] == 1, 如果此时 j >= k, 则将 k 更新为 j, 表示这是下一个要交换的值为 1 的数
					k = j
				}
			} else { // 在 i 右侧的数字中已经没有等于 0 的数字, nums[i] 已经是最小的数字. 所以不用交换, 遍历下一个的数字
				updateIJK()
			}
		} else if nums[i] == 2 { // 当前遍历到的数字为 2
			if i < k { // 在 i 右侧的数字中还有等于 1, 将二者交换
				nums[i], nums[k] = nums[k], nums[i]
			} else { // 在 i 右侧的数字中已经没有等于 1 的数字, 但是还有等于 0 的数字, 将二者交换
				nums[i], nums[j] = nums[j], nums[i]
			}
		} else { // nums[i] == 0 遍历下一个数字
			updateIJK()
		}
	}

}

// 示例 1：
// 输入：nums = [2,0,2,1,1,0]
// 输出：[0,0,1,1,2,2]
//
// 示例 2：
// 输入：nums = [2,0,1]
// 输出：[0,1,2]
func main() {
	nums := pkg.CreateSlice[int]()
	sortColors(nums)
	fmt.Println(nums)
}
