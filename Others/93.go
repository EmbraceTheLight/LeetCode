/*
93. 复原 IP 地址
有效 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。

例如："0.1.2.201" 和 "192.168.1.1" 是 有效 IP 地址，但是 "0.011.255.245"、"192.168.1.312" 和 "192.168@1.1" 是 无效 IP 地址。
给定一个只包含数字的字符串 s ，用以表示一个 IP 地址，返回所有可能的有效 IP 地址，这些地址可以通过在 s 中插入 '.' 来形成。你 不能 重新排序或删除 s 中的任何数字。你可以按 任何 顺序返回答案。

示例 1：
输入：s = "25525511135"
输出：["255.255.11.135","255.255.111.35"]

示例 2：
输入：s = "0000"
输出：["0.0.0.0"]

示例 3：
输入：s = "101023"
输出：["1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3","101.0.2.3"]

1 <= s.length <= 20
s 仅由数字组成
*/
package main

import (
	"fmt"
	"strconv"
)

func dfs93(s string, step int, fields string, ans *[]string) {
	l := len(s)
	if step == 0 && s[0] == '0' {
		return
	}
	if l > (4-step+1)*3 {
		return
	}

	var cnt int
	var num int
	if step == 4 {
		for i := 0; i < len(s); i++ {
			num = num*10 + int(s[i]-'0')
			cnt++
			if cnt == 1 && num == 0 && len(s) != 1 { //有前导0且不止包含0
				return
			}
		}
		if num <= 255 {
			fields = fields + strconv.Itoa(num)
			*ans = append(*ans, fields)
		}

	} else {
		for i := 0; i < 3; i++ {
			cnt++
			num = num*10 + int(s[i]-'0')
			if len(s[i+1:]) < 4-step || num > 255 {
				break
			}
			str := strconv.Itoa(num) + "."
			dfs93(s[i+1:], step+1, fields+str, ans)
			if cnt == 1 && num == 0 { //字段有前导0
				break
			}
		}
	}

}

func restoreIpAddresses(s string) []string {
	var ans = make([]string, 0)
	dfs93(s, 1, "", &ans)
	return ans
}
func main() {
	var str string
	fmt.Scanf("%s\n", &str)
	fmt.Println(restoreIpAddresses(str))
}
