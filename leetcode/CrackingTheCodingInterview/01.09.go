// 01.09. 字符串轮转
/*
字符串轮转。给定两个字符串s1和s2，请编写代码检查s2是否为s1旋转而成（比如，waterbottle是erbottlewat旋转后的字符串）。
提示：
字符串长度在[0, 100000]范围内。
说明:
你能只调用一次检查子串的方法吗？
*/
package main

import (
	"fmt"
	. "lc/pkg"
	"strings"
)

func isFlipedString(s1 string, s2 string) bool {
	ls1, ls2 := len(s1), len(s2)
	if s1 == "" && s2 == "" {
		return true
	}
	if ls1 != ls2 {
		return false
	}
	for i := 0; i < ls1; i++ {
		if s1[i] == s2[0] && s1 == s2[ls1-i:]+s2[0:ls1-i] {
			return true
		}
	}
	return false
}

// 最优解法, 只需比较一次字符串(查找子串)
// 将 s1 追加到 s1 后, s1 + s1 包含了所有 s1 的可能的旋转字符串
// 这种情况下, 只需要查看 s2 是否为 s1 + s1 的子串即可
func isFlipedString2(s1 string, s2 string) bool {
	return len(s1) == len(s2) && strings.Contains(s1+s1, s2)
}

// 示例 1：
// 输入：s1 = "waterbottle", s2 = "erbottlewat"
// 输出：True
//
// 示例 2：
// 输入：s1 = "aa", s2 = "aba"
// 输出：False
func main() {
	fmt.Println("Input s1:")
	s1 := CreateString()
	fmt.Println("Input s2:")
	s2 := CreateString()
	fmt.Println(isFlipedString(s1, s2))
	fmt.Println(isFlipedString2(s1, s2))
}
