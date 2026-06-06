package main

import "fmt"

func dfs11(nums []int, tmp []int, step, limit int, ans *[][]int) {
	if len(tmp) == limit {
		if len(tmp) < limit {
			return
		}
		c := make([]int, len(tmp))
		copy(c, tmp)
		*ans = append(*ans, c)
		return
	}
	for i := step; i < len(nums); i++ {
		if step+(len(nums)-i) < limit {
			break
		}
		tmp = append(tmp, nums[i])
		dfs11(nums, tmp, i+1, limit, ans)
		tmp = tmp[:len(tmp)-1]
	}
}
func solution(nums []int, k int) [][]int {
	var ans = make([][]int, 0)
	dfs11(nums, []int{}, 0, k, &ans)
	return ans
}
func main() {
	var k int
	fmt.Println("Input k:")
	fmt.Scan(&k)
	fmt.Println("input -1 to stop:")
	var nums []int = make([]int, 0)
	var in int
	fmt.Scan(&in)
	for in != -1 {
		nums = append(nums, in)
		fmt.Scan(&in)

	}
	fmt.Println("ans is:")
	fmt.Println(solution(nums, k))
}
