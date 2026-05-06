// 93. 复原 IP 地址
/*
有效 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。
例如："0.1.2.201" 和 "192.168.1.1" 是 有效 IP 地址，但是 "0.011.255.245"、"192.168.1.312" 和 "192.168@1.1" 是 无效 IP 地址。
给定一个只包含数字的字符串 s ，用以表示一个 IP 地址，返回所有可能的有效 IP 地址，
这些地址可以通过在 s 中插入 '.' 来形成。你 不能 重新排序或删除 s 中的任何数字。你可以按 任何 顺序返回答案。

提示：
1 <= s.length <= 20
s 仅由数字组成
*/
package main

import (
	"fmt"
	"strconv"
)

func restoreIpAddresses(s string) []string {
	var ans []string
	if len(s) < 4 {
		return []string{}
	}
	dfs93(0, 0, s, "", &ans)
	return ans
}
func dfs93(step int, lastIdx int, s string, cur string, ans *[]string) {
	ls := len(s)
	if step == 4 {
		// 正好分配完所有数字
		if ls == lastIdx {
			*ans = append(*ans, cur)
		}
		return
	}

	// isValid 检查字符串合法:
	// - 是否包含前导 0
	// - 网段是否大于 255
	isValid := func(s string) bool {
		// 存在前导 0
		if s[0] == '0' && len(s) > 1 {
			return false
		}
		num, err := strconv.Atoi(s)
		if err != nil {
			return false
		}
		return num < 256
	}

	for i := lastIdx; i < lastIdx+3 && i < ls; i++ {
		subnet := s[lastIdx : i+1]
		if isValid(subnet) == false {
			return
		}
		newCur := subnet
		if cur != "" {
			newCur = cur + "." + newCur
		}
		dfs93(step+1, i+1, s, newCur, ans)
	}
}

// 示例 1：
// 输入：s = "25525511135"
// 输出：["255.255.11.135","255.255.111.35"]
//
// 示例 2：
// 输入：s = "0000"
// 输出：["0.0.0.0"]
//
// 示例 3：
// 输入：s = "101023"
// 输出：["1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3","101.0.2.3"]
func main() {
	var ipStr string
	fmt.Println(restoreIpAddresses(ipStr))
}
