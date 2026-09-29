// 16.18. 模式匹配
/*
你有两个字符串，即pattern和value。 pattern字符串由字母"a"和"b"组成，用于描述字符串中的模式。
例如，字符串"catcatgocatgo"匹配模式"aabab"（其中"cat"是"a"，"go"是"b"），
该字符串也匹配像"a"、"ab"和"b"这样的模式。
但需注意"a"和"b"不能同时表示相同的字符串。编写一个方法判断value字符串是否匹配pattern字符串。

提示：
0 <= len(pattern) <= 1000
0 <= len(value) <= 1000
你可以假设pattern只包含字母"a"和"b"，value仅包含小写字母。
*/
package main

import (
	"fmt"
	"strings"
)

func patternMatching(pattern string, value string) bool {
	if pattern == "" {
		return false
	}
	if value == "" {
		if strings.Contains(pattern, "a") && strings.Contains(pattern, "b") {
			return false
		}
		return true
	}

	// 统一化处理
	// 若第一个是 b, 翻转 pattern 中的所有 a 和 b 的值, 即 a -> b, b -> a
	// 这样 pattern 中的第一个字符一定是 a. 只要 pattern 形式不变即可
	p := []byte(pattern)
	if pattern[0] == 'b' {
		for i := range p {
			if p[i] == 'a' {
				p[i] = 'b'
			} else {
				p[i] = 'a'
			}
		}
		pattern = string(p)
	}
	cntA, cntB := 0, 0

	secondCharCount, tmp := 0, false
	for i := range pattern {
		if pattern[i] != 'a' && !tmp {
			secondCharCount = i
			tmp = true
		}
		if pattern[i] == 'a' {
			cntA++
		} else {
			cntB++
		}
	}

	var dfs1618 func(pattern, value string, a, b string, step, curCntA, curCntB int) bool
	dfs1618 = func(pattern, value string, a, b string, step, curCntA, curCntB int) bool {
		if a == "" && b == "" {
			return false
		}
		idx := curCntA*len(a) + curCntB*len(b)
		if step == len(pattern) {
			return idx == len(value)
		}
		if step > 0 {
			if (pattern[step-1] == 'a' && idx > len(value)) || (pattern[step-1] == 'b' && idx > len(value)) {
				return false
			}
		}

		if pattern[step] == 'a' && a != value[idx:idx+len(a)] {
			return false
		} else if pattern[step] == 'b' && b != value[idx:idx+len(b)] {
			return false
		}
		if pattern[step] == 'a' {
			return dfs1618(pattern, value, a, b, step+1, curCntA+1, curCntB)
		} else {
			return dfs1618(pattern, value, a, b, step+1, curCntA, curCntB+1)
		}
	}

	// cntA * len(a) + cntB * len(B) == len(value)
	// 只要确定 a 的长度, b 的长度就确定了: len(b) = (len(value) - cntA * len(a)) / cntB
	for i := 0; i <= len(value); i++ {
		a := value[0:i]
		b := ""

		if len(value)-cntA*len(a) < 0 {
			break
		}
		if cntB != 0 && (len(value)-cntA*len(a))%cntB != 0 {
			continue
		}

		if cntB == 0 {
			b = ""
		} else {
			lenB := (len(value) - cntA*len(a)) / cntB
			b = value[secondCharCount*len(a) : secondCharCount*len(a)+lenB]
		}
		if a == b {
			continue
		}
		if dfs1618(pattern, value, a, b, 0, 0, 0) == true {
			return true
		}
	}
	return false
}

// 示例 1：
// 输入： pattern = "abba", value = "dogcatcatdog"
// 输出： true
//
// 示例 2：
// 输入： pattern = "abba", value = "dogcatcatfish"
// 输出： false
//
// 示例 3：
// 输入： pattern = "aaaa", value = "dogcatcatdog"
// 输出： false
//
// 示例 4：
// 输入： pattern = "abba", value = "dogdogdogdog"
// 输出： true
// 解释： "a"="dogdog",b=""，反之也符合规则
func main() {
	var pattern, value string
	fmt.Println("Input pattern:")
	fmt.Scan(&pattern)
	fmt.Println("Input value:")
	fmt.Scan(&value)
	fmt.Println(patternMatching(pattern, value))
}
