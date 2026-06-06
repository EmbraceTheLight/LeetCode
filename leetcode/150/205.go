// 205. 同构字符串
/*
给定两个字符串 s 和 t ，判断它们是否是同构的。
如果 s 中的字符可以按某种映射关系替换得到 t ，那么这两个字符串是同构的。
每个出现的字符都应当映射到另一个字符，同时不改变字符的顺序。不同字符不能映射到同一个字符上，相同字符只能映射到同一个字符上，字符可以映射到自己本身。

示例 1:
输入：s = "egg", t = "add"
输出：true

示例 2：
输入：s = "foo", t = "bar"

输出：false

示例 3：
输入：s = "paper", t = "title"
输出：true

提示：
1 <= s.length <= 5 * 10^4
t.length == s.length
s 和 t 由任意有效的 ASCII 字符组成
*/
package main

import (
	"fmt"
)

func isIsomorphic(s string, t string) bool {
	ls, lt := len(s), len(t)
	if ls != lt {
		return false
	}

	// mpS2T 保存 s 中字符到 t 中字符的映射
	mpS2T := make(map[byte]byte)

	// referenceCheck 存储已经映射过的字符，防止 s 中不同的字符映射到同一个字符上
	referenceCheck := make(map[byte]byte)
	for i := 0; i < ls; i++ {
		if _, ok := mpS2T[s[i]]; !ok {
			if _, ok := referenceCheck[t[i]]; ok && referenceCheck[t[i]] != s[i] {
				return false
			}
			mpS2T[s[i]] = t[i]
			referenceCheck[t[i]] = s[i]
		} else {
			if mpS2T[s[i]] != t[i] {
				return false
			}
		}

	}
	return true
}

// 简化写法。不使用 map 取值的第二个返回判断是否存在该键，而是直接与零值比较. 注意，这种写法可能偏慢
func isIsomorphicBetter(s string, t string) bool {
	ls, lt := len(s), len(t)
	if ls != lt {
		return false
	}

	// mpS2T 保存 s 中字符到 t 中字符的映射
	mpS2T := make(map[byte]byte)

	// mpT2S 存储 t 中字符到 s 中字符的映射，防止 s 中不同的字符映射到同一个字符上
	mpT2S := make(map[byte]byte)
	for i := 0; i < ls; i++ {
		if (mpS2T[s[i]] > 0 && mpS2T[s[i]] != t[i]) || (mpT2S[t[i]] > 0 && mpT2S[t[i]] != s[i]) {
			return false
		}
		mpS2T[s[i]] = t[i]
		mpT2S[t[i]] = s[i]

	}
	return true
}

// Test Case1: "egg", "add"	Output: true
// Test Case2: "foo", "bar"	Output: false
// Test Case3: "paper", "title"	Output: true
// Test Case4: "badc", "baba"	Output: false
// Test Case5: "bbbaaaba", "aaabbbba"	Output: false
func main() {
	var s, t string
	fmt.Scan(&s, &t)
	fmt.Println(isIsomorphic(s, t))
}
