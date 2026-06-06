//移除元素
/*
给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素，并返回移除后数组的新长度。
不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并 原地 修改输入数组。
元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。

示例 1：
输入：nums = [3,2,2,3], val = 3
输出：2, nums = [2,2]
解释：函数应该返回新的长度 2, 并且 nums 中的前两个元素均为 2。你不需要考虑数组中超出新长度后面的元素。
例如，函数返回的新长度为 2 ，而 nums = [2,2,3,3] 或 nums = [2,2,0,0]，也会被视作正确答案。

示例 2：
输入：nums = [0,1,2,2,3,0,4,2], val = 2
输出：5, nums = [0,1,3,0,4]
解释：函数应该返回新的长度 5, 并且 nums 中的前五个元素为 0, 1, 3, 0, 4。
注意这五个元素可为任意顺序。你不需要考虑数组中超出新长度后面的元素。
*/
package main

import "fmt"

func removeElement(nums []int, val int) int {
	l := len(nums)
	var cnt = 0 //统计移除的元素数
	if l == 1 && nums[0] == val {
		return 0
	}
	p := 0
	tail := l - 1
	for p <= tail {
		for tail >= 0 && nums[tail] == val { //从尾部扫描
			tail--
			cnt++
		}
		for p <= tail && nums[p] != val {
			p++
		}
		if p <= tail { //tail更新后可能小于等于p，要进行检查
			nums[p], nums[tail] = nums[tail], nums[p] //交换，把要删除的元素交换到切片尾部
			cnt++                                     //更新移除元素数
			p++                                       //更新头部索引p
			tail--                                    //注意更新尾指针索引，否则cnt会重复记录
		}
	}
	return l - cnt
}

func main() {
	nums := make([]int, 0)
	var val int
	fmt.Println("Input Value:")
	fmt.Scanf("%d \n", &val)

	var tmp int
	for {
		fmt.Scan(&tmp)
		if tmp == -1 {
			break
		}
		nums = append(nums, tmp)
	}
	fmt.Scan(&nums)
	a := removeElement(nums, val)
	fmt.Println("return Value:", a)
	for i := 0; i < a; i++ {
		fmt.Printf("%d ", nums[i])

	}
}
