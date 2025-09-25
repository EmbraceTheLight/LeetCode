//给你两个按 非递减顺序 排列的整数数组 nums1 和 nums2，另有两个整数 m 和 n ，分别表示 nums1 和 nums2 中的元素数目。
//请你 合并 nums2 到 nums1 中，使合并后的数组同样按 非递减顺序 排列。

// 注意：最终，合并后数组不应由函数返回，而是存储在数组 nums1 中。为了应对这种情况，nums1 的初始长度为 m + n，其中前 m 个元素表示应合并的元素，后 n 个元素为 0 ，应忽略。nums2 的长度为 n 。
//
// 示例 1：
// 输入：nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
// 输出：[1,2,2,3,5,6]
// 解释：需要合并 [1,2,3] 和 [2,5,6] 。
// 合并结果是 [1,2,2,3,5,6] ，其中斜体加粗标注的为 nums1 中的元素。
//
// 示例 2：
// 输入：nums1 = [1], m = 1, nums2 = [], n = 0
// 输出：[1]
// 解释：需要合并 [1] 和 [] 。
// 合并结果是 [1] 。
//
// 示例 3：
// 输入：nums1 = [0], m = 0, nums2 = [1], n = 1
// 输出：[1]
// 解释：需要合并的数组是 [] 和 [1] 。
// 合并结果是 [1] 。
// 注意，因为 m = 0 ，所以 nums1 中没有元素。nums1 中仅存的 0 仅仅是为了确保合并结果可以顺利存放到 nums1 中。
package main

import "fmt"

//func merge(nums1 []int, m int, nums2 []int, n int)  {
//	i := 0
//	j := 0
//	if m == 0 {
//		copy(nums1, nums2)
//	}
//	if n == 0 {
//		return
//	}
//
//	for i < m && j < n {
//		if nums1[i] > nums2[j] {
//			copy(nums1[i+1:m+1], nums1[i:m])
//			nums1[i] = nums2[j]
//			j++
//			m++
//		} else {
//			i++
//		}
//	}
//	if j < n {
//		nums1 = append(nums1, nums2[j:]...)
//	}
//	return
//}

// 逆向指针，时间，空间复杂度最小
// 将最大的排到nums1尾部
func merge(nums1 []int, m int, nums2 []int, n int) {
	tail := m + n - 1 //指向nums1尾部可用空间
	p1 := m - 1       //nums1索引
	p2 := n - 1       //nums2索引
	if m == 0 {
		copy(nums1, nums2)
	}
	if n == 0 {
	}
	for p1 >= 0 && p2 >= 0 {
		if nums1[p1] > nums2[p2] {
			nums1[tail] = nums1[p1]
			p1--
			tail--
		} else {
			nums1[tail] = nums2[p2]
			p2--
			tail--
		}
	}
	if p2 >= 0 {
		for ; p2 >= 0; p2-- {
			nums1[tail] = nums2[p2]
			tail--
		}
	}
}
func merge2(nums1 []int, m int, nums2 []int, n int) {
	idx := m + n - 1
	i, j := m-1, n-1
	for ; i >= 0 && j >= 0; idx-- {
		if nums1[i] > nums2[j] {
			nums1[idx] = nums1[i]
			i--
		} else {
			nums1[idx] = nums2[j]
			j--
		}
	}

	for i >= 0 {
		nums1[idx] = nums1[i]
		i--
		idx--
	}
	for j >= 0 {
		nums1[idx] = nums2[j]
		j--
		idx--
	}
}
func main() {
	var m int
	var n int

	//fmt.Println("Input m:")
	//fmt.Scanf("%d\n", &m)
	//fmt.Println("Input n:")
	//fmt.Scanf("%d\n", &n)
	//nums1 := make([]int, m+n)
	//nums2 := make([]int, n)
	////var tmp int
	//
	//for i := 0; i < m; i++ {
	//	fmt.Scan(&nums1[i])
	//}
	//for i := 0; i < n; i++ {
	//	fmt.Scan(&nums2[i])
	//	nums1 = append(nums1, 0)
	//}
	//merge(nums1, m, nums2, n)
	//for i := 0; i < m+n; i++ {
	//	fmt.Printf("%d ", nums1[i])
	//}

	fmt.Println("Input m:")
	fmt.Scanf("%d\n", &m)
	fmt.Println("Input n:")
	fmt.Scanf("%d\n", &n)

	nums11 := make([]int, m+n)
	nums22 := make([]int, n)

	for i := 0; i < m; i++ {
		fmt.Scan(&nums11[i])
	}
	for i := 0; i < n; i++ {
		fmt.Scan(&nums22[i])
		nums11 = append(nums11, 0)
	}
	merge2(nums11, n, nums22, n)
	for i := 0; i < m+n; i++ {
		fmt.Printf("%d ", nums11[i])
	}
}
