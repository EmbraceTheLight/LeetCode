// 05.03. 翻转数位
/*
给定一个32位整数 num，你可以将一个数位从0变为1。请编写一个程序，找出你能够获得的最长的一串1的长度。
*/
package main

import "fmt"

func reverseBits(num int) int {
	var ans int
	bits := exactBits0503(num)
	idx := 0
	for ; idx < 32; idx++ {
		count := 1
		for idx < 32 && bits[idx] == 1 {
			count++
			idx++
		}
		if idx+1 < 32 && bits[idx+1] == 1 {
			j := idx + 1
			//count++ // 将 bits[idx] 从 0 翻转为 1
			for j < 32 && bits[j] == 1 {
				j++
				count++
			}
		}
		ans = max(ans, min(32, count))
	}
	return ans
}
func exactBits0503(num int) [32]int {
	var ret [32]int
	idx := 31
	for idx >= 0 {
		ret[idx] = (num >> (31 - idx)) & 1
		idx--
	}
	return ret
}

// 示例 1：
// 输入: num = 1775(11011101111)
// 输出: 8
//
// 示例 2：
// 输入: num = 7(0111)
// 输出: 4

func main() {
	fmt.Println("Input num:")
	var num int
	fmt.Scan(&num)
	fmt.Println(reverseBits(num))
}
