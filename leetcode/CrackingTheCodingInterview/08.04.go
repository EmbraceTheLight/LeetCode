// 08.04. 幂集
/*
  	编写一种方法，返回某集合的所有子集。集合中 不包含重复的元素。
	说明：解集不能包含重复的子集。
*/
package main

import (
	"fmt"
	. "lc/pkg"
)

func subsets(nums []int) [][]int {
	ans := make([][]int, 0)
	ans = append(ans, []int{})
	dfs0804(nums, 0, []int{}, &ans)
	return ans
}

func dfs0804(nums []int, idx int, cur []int, ans *[][]int) {
	for i := idx; i < len(nums); i++ {
		cur = append(cur, nums[i])
		*ans = append(*ans, append([]int{}, cur...))
		dfs0804(nums, i+1, cur, ans)
		cur = cur[:len(cur)-1]
	}
}

// 示例：
//
//	输入：nums = [1,2,3]
//	输出：[[3], [1], [2], [1,2,3], [1,3], [2,3], [1,2], []]
func main() {
	nums := CreateSlice[int]()
	fmt.Println(subsets(nums))
}
