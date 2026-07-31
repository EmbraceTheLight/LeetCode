// 05.04. 下一个数
/*
下一个数。给定一个正整数，找出与其二进制表达式中1的个数相同且大小最接近的那两个数（一个略大，一个略小）。

提示：
num 的范围在[1, 2147483647]之间；
如果找不到前一个或者后一个满足条件的正数，那么输出 -1。
*/
package main

import "fmt"

func findClosedNumbers(num int) []int {
	var ans []int
	// 下一个 1 相同的更大值
	// 对于形如 ...00111001100 这样的二进制数,
	// 需要从低位到高位找到第一批 1 中最高的那一位, 并将其左移一位, 即 ...00111001100 -> ...00111010100
	// 然后将高位的 1 向低位的 0 转移
	// 如果无法左移, 则返回 -1
	idx := findBitIdx(num, 1)
	if idx == -1 || idx == 30 { // idx == 31, 说明除最高符号位外, num 所有位全部为 1 (2^32 - 1), 此时不会有比 num 更大的数
		ans = append(ans, -1)
	} else {
		setOne := idx + 1
		tmp := num
		tmp = tmp &^ (1 << idx)   // 将 tmp idx 位置零
		tmp = tmp | (1 << setOne) // 将 tmp idx 位置 1
		// 收集位于 idx 右侧的 1, 并将其移动至更低位
		countOnes := 0
		for i := idx - 1; i >= 0; i-- {
			if (tmp>>i)&1 == 1 {
				countOnes++
				tmp = tmp &^ (1 << i)
			}
		}
		for i := 0; countOnes > 0; i, countOnes = i+1, countOnes-1 {
			tmp = tmp | (1 << i)
		}
		ans = append(ans, tmp)
	}

	// 上一个 1 相同的更小值
	// 对于形如 ...00111001100 这样的二进制数,
	// 需要从低位到高位找到第一批 0 中最高的那一位, 并将其左移一位, 即 ...00111001100 -> ...00111001010
	// 如果无法左移, 则返回 -1
	idx = findBitIdx(num, 0)
	if idx == -1 || idx == 31 { // idx == 31, 说明 num 所有位全部为 0, 此时不会有比 num 更小的数
		ans = append(ans, -1)
	} else {
		setZero := idx + 1
		tmp := num
		tmp = tmp &^ (1 << setZero) // 将 setZero 位置零
		tmp = tmp | (1 << idx)      // 将 idx 位置 1
		// 收集位于 idx 右侧的 1 移动靠近至 idx 的更高位
		countOnes := 0
		for i := idx - 1; i >= 0; i-- {
			if (tmp>>i)&1 == 1 {
				countOnes++
				tmp = tmp &^ (1 << i)
			}
		}
		for i := idx - 1; countOnes > 0; i, countOnes = i-1, countOnes-1 {
			tmp = tmp | (1 << i)
		}
		ans = append(ans, tmp)
	}
	return ans
}

// findBitIdx 寻找 32 位整数 num 中第一批 bit 的最高位索引
// bit 为 0 或 1.
// 索引从 0 开始, 顺序: 由低位到高位. 返回 -1 表示 num 中没有找到 bit
// 示例1: num: 10 (01001), bit: 1 ==> 0
// 示例2: num: 10 (01001), bit: 0 ==> 2
func findBitIdx(num int, bit int) int {
	bitIdx := 0
	flag := false
	hasBit := false // num 中是否有目标 bit
	for num != 0 {
		if num&1 == bit { // 遇到第一批目标 bit, 将 flag 置为 false
			flag = true
			hasBit = true
		} else if flag == true { // 若 flag == true, 说明第一批 bit 已遍历完, bitIdx - 1 即为所求
			break
		}
		num = num >> 1
		bitIdx++
	}
	if hasBit == false {
		return -1
	}

	return bitIdx - 1
}

// 示例1:
//
//	输入: 67    (0100 0011)
//	输出: [69 (0100 0101), 56 (0011 1000)]
//
// 示例 2：
//
//	输入：num = 2（或者0b10）
//	输出：[4, 1] 或者（[0b100, 0b1]）
//
// 示例 3：
//
//	输入：num = 1
//	输出：[2, -1]
func main() {
	fmt.Println("Input num:")
	var num int
	fmt.Scan(&num)
	fmt.Println(findClosedNumbers(num))
}
