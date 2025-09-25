/*
468. 验证IP地址
给定一个字符串 queryIP。如果是有效的 IPv4 地址，返回 "IPv4" ；如果是有效的 IPv6 地址，返回 "IPv6" ；如果不是上述类型的 IP 地址，返回 "Neither" 。
有效的IPv4地址 是 “x1.x2.x3.x4” 形式的IP地址。 其中 0 <= xi <= 255 且 xi 不能包含 前导零。
例如: “192.168.1.1” 、 “192.168.1.0” 为有效IPv4地址， “192.168.01.1” 为无效IPv4地址; “192.168.1.00” 、 “192.168@1.1” 为无效IPv4地址。

一个有效的IPv6地址 是一个格式为“x1:x2:x3:x4:x5:x6:x7:x8” 的IP地址，其中:
1 <= xi.length <= 4
xi 是一个 十六进制字符串 ，可以包含数字、小写英文字母( 'a' 到 'f' )和大写英文字母( 'A' 到 'F' )。
在 xi 中允许前导零。
例如 "2001:0db8:85a3:0000:0000:8a2e:0370:7334" 和 "2001:db8:85a3:0:0:8A2E:0370:7334" 是有效的 IPv6 地址，
而 "2001:0db8:85a3::8A2E:037j:7334" 和 "02001:0db8:85a3:0000:0000:8a2e:0370:7334" 是无效的 IPv6 地址。
示例 1：
输入：queryIP = "172.16.254.1"
输出："IPv4"
解释：有效的 IPv4 地址，返回 "IPv4"

示例 2：
输入：queryIP = "2001:0db8:85a3:0:0:8A2E:0370:7334"
输出："IPv6"
解释：有效的 IPv6 地址，返回 "IPv6"

示例 3：
输入：queryIP = "256.256.256.256"
输出："Neither"
解释：既不是 IPv4 地址，又不是 IPv6 地址
*/
package main

import (
	"fmt"
)

func ToUpper(char byte) byte {
	if char <= 'f' && char >= 'a' {
		return char - 32
	}
	return char
}
func isAtoF(char byte) bool {
	a := ToUpper(char)
	return a >= 'A' && a <= 'F'
}
func isNum(char byte) bool {
	return char >= '0' && char <= '9'
}

func isIPv4(s string) bool {
	if isNum(s[0]) == false {
		return false
	}
	l := len(s)
	var cntField int
	for i := 0; i < l; {
		var cnt0 bool //是否有前导0
		var cnt int   //统计每段数字个数
		var num int
		for i < l && isNum(s[i]) == true {
			if s[i] == '0' && num == 0 {
				cnt0 = true
			}
			cnt++
			num = num*10 + int(s[i]-'0')
			i++
		}
		if num > 255 || (cnt > 1 && cnt0 == true) || (i < l && s[i] != '.') {
			return false
		}
		cntField++
		if cntField == 4 && len(s[i:]) > 0 {
			return false
		}
		i++
	}
	if cntField == 4 {
		return true
	} else {
		return false
	}
}

func isIPv6(s string) bool {
	l := len(s)
	if l < 3 {
		return false
	}
	if !isNum(s[0]) && !isAtoF(ToUpper(s[0])) { //首字符不是有效字符
		return false
	}
	var cntField = 0
	for i := 0; cntField < 8 && i < l; {
		var cnt int
		for i < l && (isNum(s[i]) || isAtoF(ToUpper(s[i]))) {
			i++
			cnt++
		}
		if cnt > 4 || cnt == 0 { //长度不符合要求或有不符合要求的字符
			return false
		}
		cntField++ //字段个数+1
		if cntField < 8 {
			if s[i] == ':' {
				i++
			} else {
				return false
			}
		}

		if cntField == 8 && len(s[i:]) > 0 { //后面还有字符
			return false
		}
		//if i < l && s[i+1] == ':' {
		//	if double == true { //已经使用过一次"::"了，不允许再次使用
		//		return false
		//	}
		//	double = true
		//}
	}
	if cntField == 8 {
		return true
	} else {
		return false
	}
}
func validIPAddress(queryIP string) string {
	l := len(queryIP)
	var chkV4orV6 int
	for i := 0; i < l; i++ {
		if queryIP[i] == '.' {
			chkV4orV6 = 1
			break
		}
		if queryIP[i] == ':' {
			chkV4orV6 = 2
			break
		}
	}
	if chkV4orV6 == 1 {
		if isIPv4(queryIP) == true {
			return "IPv4"
		} else {
			return "Neither"
		}
	} else if chkV4orV6 == 2 {
		if isIPv6(queryIP) == true {
			return "IPv6"
		} else {
			return "Neither"
		}
	}
	return "Neither"
}
func main() {
	var str string
	fmt.Scanf("%s\n", &str)
	fmt.Println(validIPAddress(str))
}
