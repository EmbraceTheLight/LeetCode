// 201. 数字范围按位与
/*
给你两个整数 left 和 right ，表示区间 [left, right] ，返回此区间内所有数字 按位与 的结果（包含 left 、right 端点）

提示：
0 <= left <= right <= 2^31 - 1
*/
package main

import (
	"fmt"
	"math"
)

func rangeBitwiseAnd(left int, right int) int {
	if left == right {
		return left
	}
	// 思路: 求 left 和 right 的二进制, 从高位到低位, 直到找到第一个不同的位, 该位及其之后的位全部清零
	// 它前面的位保留不变, 即为所求
	getBits := func(num int) []int {
		var ret []int = make([]int, 32) // ret 下标从大到小对应二进制从左到右
		for i := 0; i <= 31; i++ {
			ret[i] = (num >> i) & 1
		}
		return ret
	}
	leftBits, rightBits := getBits(left), getBits(right)
	diff := 0
	for i := 31; i >= 0; i-- {
		if leftBits[i] != rightBits[i] {
			diff = i
			break
		}
	}

	var ans int
	for i := 31; i > diff; i-- {
		if leftBits[i] == 1 {
			ans += int(math.Pow(2, float64(i)))
		}
	}
	return ans
}

// * 思路: 直接逐位比较 left 和 right 的位即可, 找出最高不等位。
// 具体实现: 对 left 和 right 进行移位, 直到二者相等, 说明有了相同的前缀, 也即找到了最高不等位
// 总体思路与上面的一致, 但是简洁多了
func rangeBitwiseAndBest(left int, right int) int {
	shift := 0
	for left < right {
		left >>= 1
		right >>= 1
		shift++
	}
	return left << shift
}

// 示例 1：
// 输入：left = 5, right = 7
// 输出：4
//
// 示例 2：
// 输入：left = 0, right = 0
// 输出：0
//
// 示例 3：
// 输入：left = 1, right = 2147483647
// 输出：0
//
// 示例 4:
// 输入：left = 1, right = 1
// 输出：0
func main() {
	var left, right int
	fmt.Println("Input left, right:")
	fmt.Scan(&left, &right)
	fmt.Println(rangeBitwiseAnd(left, right))
	fmt.Println(rangeBitwiseAndBest(left, right))
}
