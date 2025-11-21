// 202. 快乐数
/*
那么这个数就是快乐数。
如果 n 是 快乐数 就返回 true ；不是，则返回 false 。

示例 1：
输入：n = 19
输出：true

解释：
1^2 + 9^2 = 82
8^2 + 2^2 = 68
6^2 + 8^2 = 100
1^2 + 0^2 + 0^2 = 1

示例 2：
输入：n = 2
输出：false

提示：
1 <= n <= 2^31 - 1
*/
package main

import "fmt"

// 快乐数，一定不会出现无限大循环，因为
/*
Digits	Largest		Next
1			9		81
2			99		162
3			999		243
4			9999	324
13	9999999999999	1053
*/
// 这就表示即使是很大的数字，计算之后，也会回到一格比较小的数字范围内。如果有任意一位不是 9，则得到的数字只会更小
// 因此可以使用 map，如果会出现循环的情况，可以及时检测出来。否则一定会得到 1.
func isHappy(n int) bool {
	mp := make(map[int]bool)
	for n != 1 {
		n = next(n)
		if mp[n] == true {
			return false
		} else {
			mp[n] = true
		}
	}
	return true
}

// 优化版，使用双指针，不使用 map，空间复杂度更优
func isHappy2Pointer(n int) bool {
	slow, fast := next(n), next(next(n))
	// 情况 1：有环 则 slow 一定会被 fast 套圈追上
	// 情况 2：无环， slow == fast == 1，循环结束
	for slow != fast {
		slow = next(slow)
		fast = next(next(fast))
	}
	return slow == 1
}

func next(n int) int {
	sum := 0
	for n != 0 {
		num := n % 10
		n /= 10
		sum += num * num
	}
	return sum
}

// Test Case1: 19	Output: true
// Test Case2: 2	Output: false
func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(isHappy(n))
}
