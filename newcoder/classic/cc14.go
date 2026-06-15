// CC14 出现一次的数字ii
/*
描述
现在有一个整数类型的数组，数组中只有一个元素只出现一次，其余元素都出现三次。你需要找出只出现一次的元素
数据范围:  数组长度满足 0<n≤4000 ，数组中每个元素的值满足 0≤val≤2147483648
进阶: 空间复杂度 O(1)， 时间复杂度 O(n)
*/
package main

import (
	"fmt"
	. "lc/pkg"
)

func singleNumber(A []int) int {
	// write code here
	var ans int
	var cntMp [32]int
	for i := 0; i < len(A); i++ {
		tmp := 1
		for j := 0; j < 32; j++ {
			cntMp[j] += (A[i] & tmp)
			tmp <<= 1
		}
	}
	for i := 0; i < 32; i++ {
		ans |= (cntMp[i] % 3) << i
	}
	return ans
}

// 示例1
// 输入: [0,0,0,5]
// 返回值: 5
//
// 示例2
// 输入: [0]
// 返回值: 0
func main() {

	nums := CreateSlice[int]()
	fmt.Println(singleNumber(nums))
}
