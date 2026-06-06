// 13. 罗马数字转整数
/*
罗马数字包含以下七种字符: I， V， X， L，C，D 和 M。

字符          数值
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
例如， 罗马数字 2 写做 II ，即为两个并列的 1 。12 写做 XII ，即为 X + II 。 27 写做  XXVII, 即为 XX + V + II 。

通常情况下，罗马数字中小的数字在大的数字的右边。但也存在特例，
例如 4 不写做 IIII，而是 IV。
数字 1 在数字 5 的左边，所表示的数等于大数 5 减小数 1 得到的数值 4 。
同样地，数字 9 表示为 IX。这个特殊的规则只适用于以下六种情况：
I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。
C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。
给定一个罗马数字，将其转换成整数。

示例 1:
输入: s = "III"
输出: 3

示例 2:
输入: s = "IV"
输出: 4

示例 3:
输入: s = "IX"
输出: 9

示例 4:
输入: s = "LVIII"
输出: 58
解释: L = 50, V= 5, III = 3.

示例 5:
输入: s = "MCMXCIV"
输出: 1994
解释: M = 1000, CM = 900, XC = 90, IV = 4.


提示：
1 <= s.length <= 15
s 仅含字符 ('I', 'V', 'X', 'L', 'C', 'D', 'M')
题目数据保证 s 是一个有效的罗马数字，且表示整数在范围 [1, 3999] 内
题目所给测试用例皆符合罗马数字书写规则，不会出现跨位等情况。
IL 和 IM 这样的例子并不符合题目要求，49 应该写作 XLIX，999 应该写作 CMXCIX。
*/
package main

import "fmt"

// 最暴力的解法
func romanToInt(s string) int {
	var ans int
	mp := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	n := len(s)
	for i := 0; i < n; {
		if i+1 < n {
			if s[i] == 'I' && s[i+1] == 'V' {
				ans += 4
				i += 2
				continue
			} else if s[i] == 'I' && s[i+1] == 'X' {
				ans += 9
				i += 2
				continue
			} else if s[i] == 'X' && s[i+1] == 'L' {
				ans += 40
				i += 2
				continue
			} else if s[i] == 'X' && s[i+1] == 'C' {
				ans += 90
				i += 2
				continue
			} else if s[i] == 'C' && s[i+1] == 'D' {
				ans += 400
				i += 2
				continue
			} else if s[i] == 'C' && s[i+1] == 'M' {
				ans += 900
				i += 2
				continue
			}
		}
		ans += mp[s[i]]
		i++
	}
	return ans
}

// 思路：第 s[i] 比第 s[i+1] 表示的数字小时，用右边大的数字减去左边小的数字。比如：IV = V - I = 5 - 1 = 4 。
func romanToIntBetter(s string) int {
	var ans int
	mp := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	n := len(s)
	for i := 0; i < n; {
		if i+1 < n {
			if mp[s[i]] < mp[s[i+1]] {
				ans += mp[s[i+1]] - mp[s[i]]
				i += 2
				continue
			}
		}
		ans += mp[s[i]]
		i++
	}
	return ans
}

// Test Case 1: III   	Output: 3
// Test Case 2: IV    	Output: 4
// Test Case 3: IX    	Output: 9
// Test Case 4: LVIII 	Output: 58
// Test Case 5: MCMXCIV Output: 1994
func main() {
	var s string
	fmt.Scan(&s)
	fmt.Println(romanToIntBetter(s))
}
