/*
14.最长公共前缀
编写一个函数来查找字符串数组中的最长公共前缀。
如果不存在公共前缀，返回空字符串 ""。

示例 1：
输入：strs = ["flower","flow","flight"]
输出："fl"

示例 2：
输入：strs = ["dog","racecar","car"]
输出：""
解释：输入不存在公共前缀。

1 <= strs.length <= 200
0 <= strs[i].length <= 200
*/
package main

import (
	"fmt"
	"math"
)

func longestCommonPrefix(strs []string) string {
	l := len(strs)
	var ans string

	var minLen = math.MaxInt32
	for i := 0; i < l; i++ {
		minLen = min(minLen, len(strs[i]))
	}

Outer:
	for i := 0; i < minLen; i++ {
		cmp := strs[0][i]
		for m := 0; m < l; m++ {
			if strs[m][i] != cmp {
				break Outer
			}
		}
		ans += string(cmp)
	}
	return ans
}
func createStringArr() []string {
	fmt.Println("Input  string,\"-1\" to quit ")
	var str string
	ret := make([]string, 0)
	fmt.Scanf("%s\r\n", &str)
	for str != "-1" {
		ret = append(ret, str)
		fmt.Scanf("%s\n", &str)
	}
	return ret
}
func main() {
	a := createStringArr()
	fmt.Println(longestCommonPrefix(a))
}
