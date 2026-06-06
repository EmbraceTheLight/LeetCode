/*
402. 移掉 K 位数字
给你一个以字符串表示的非负整数 num 和一个整数 k ，移除这个数中的 k 位数字，使得剩下的数字最小。请你以字符串形式返回这个最小的数字。

示例 1 ：
输入：num = "1432219", k = 3
输出："1219"
解释：移除掉三个数字 4, 3, 和 2 形成一个新的最小的数字 1219 。

示例 2 ：
输入：num = "10200", k = 1
输出："200"
解释：移掉首位的 1 剩下的数字为 200. 注意输出不能有任何前导零。

示例 3 ：
输入：num = "10", k = 2
输出："0"
解释：从原数字移除所有的数字，剩余为空就是 0 。

1 <= k <= num.length <= 10^5
num 仅由若干位数字（0 - 9）组成
除了 0 本身之外，num 不含任何前导零
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func removeKdigits(num string, k int) string {
	if len(num) == k {
		return "0"
	}
	var ans string
	var cnt int //统计是否移除了k位数字
	var stack = make([]byte, 0)
	stack = append(stack, 0)
	l := len(num)
	i := 0
	for ; i < l; i++ {
		lastNum := stack[len(stack)-1]
		if num[i] >= lastNum {
			stack = append(stack, num[i])
		} else {
			for lastNum > num[i] && cnt < k {
				cnt++
				stack = stack[:len(stack)-1]
				lastNum = stack[len(stack)-1]
			}
			if cnt < k { //还未移除够数字
				stack = append(stack, num[i])
			} else {
				break
			}
		}
	}

	for cnt < k { //数字还未移除完，从栈顶开始移除
		stack = stack[:len(stack)-1]
		cnt++
	}
	builder := strings.Builder{}

	builder.Write(stack[1:])
	builder.Write([]byte(num[i:]))
	ans = builder.String()

	var idx int //记录从哪个位置开始没有前导0
	for idx = range ans {
		if ans[idx] != '0' {
			break
		}
	}
	return ans[idx:]
}
func main() {
	var k int
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Input k:")
	fmt.Scanf("%d \n", &k)

	fmt.Printf("Input s:")
	result, _, _ := reader.ReadLine()
	fmt.Println(removeKdigits(string(result), k))
}
