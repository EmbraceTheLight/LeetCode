// CC15 出现一次的数字
/*
描述
现在有一个整数类型的数组，数组中素只有一个元素只出现一次，其余的元素都出现两次。
数据范围: 0<n≤4000 ， 数组中每个值满足 0≤val≤4000
进阶: 空间复杂度 O(1) ，时间复杂度 O(n)
*/
package main

import (
	"fmt"
	. "lc/pkg"
)

func singleNumberCC15(A []int) int {
	// write code here
	var ans int = A[0]
	for i := 1; i < len(A); i++ {
		ans = ans ^ A[i]
	}
	return ans
}

// 示例1
// 输入: [1,0,1]
// 返回值: 0
//
// 示例2
// 输入: [0]
// 返回值: 0
func main() {
	nums := CreateSlice[int]()
	fmt.Println(singleNumberCC15(nums))
}
