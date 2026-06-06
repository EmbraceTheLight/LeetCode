/*
739. 每日温度
给定一个整数数组 temperatures ，表示每天的温度，
返回一个数组 answer ，其中 answer[i] 是指对于第 i 天，下一个更高温度出现在几天后。如果气温在这之后都不会升高，请在该位置用 0 来代替。

示例 1:
输入: temperatures = [73,74,75,71,69,72,76,73]
输出: [1,1,4,2,1,1,0,0]

示例 2:
输入: temperatures = [30,40,50,60]
输出: [1,1,1,0]

示例 3:
输入: temperatures = [30,60,90]
输出: [1,1,0]

1 <= temperatures.length <= 10^5
30 <= temperatures[i] <= 100
*/
package main

import (
	"fmt"
)

//	func dailyTemperatures(temperatures []int) []int {
//		var ans = make([]int, len(temperatures))
//
//		var stk []temperature
//		for i, v := range temperatures {
//			if len(stk) == 0 {
//				stk = append(stk, temperature{
//					t:   v,
//					idx: i,
//				})
//				continue
//			}
//			cmp := stk[len(stk)-1]
//			if cmp.t >= v { //温度不比前一天高，入栈
//				stk = append(stk, temperature{
//					t:   v,
//					idx: i,
//				})
//			} else {
//				l := len(stk)
//				var ti = l - 1
//				for ; ti >= 0; ti-- {
//					if stk[ti].t < v {
//						ans[stk[ti].idx] += l - ti
//					} else {
//						for tj := ti; tj >= 0; tj-- { //更新比新插入的值更高的值的计数
//							ans[stk[tj].idx] += l - ti - 1
//						}
//						break
//					}
//				}
//
//				stk = stk[:ti+1] //弹出比新插入的值更低的值
//
//				stk = append(stk, temperature{
//					t:   v,
//					idx: i,
//				}) //不要忘记将更高的温度入栈
//			}
//
//		}
//		var a = len(stk)
//		for i := 0; i < a; i++ {
//			ans[stk[i].idx] = 0 //将剩余温度置0，因为没有温度比它们高了
//		}
//		return ans
//	}
func dailyTemperatures(temperatures []int) []int {
	var ans = make([]int, len(temperatures))

	var stk []int
	for i, v := range temperatures {
		if len(stk) == 0 {
			stk = append(stk, i)
			continue
		}

		cmp := stk[len(stk)-1]
		for len(stk) > 0 && temperatures[cmp] < v {
			ans[cmp] = i - cmp
			stk = stk[:len(stk)-1] //弹出元素
			if len(stk) > 0 {
				cmp = stk[len(stk)-1]
			}
		}
		stk = append(stk, i)

	}
	return ans
}
func main() {
	fmt.Println("temperatures = ")
	var tmp int
	var temperatures []int
	fmt.Scan(&tmp)
	for tmp != -1 {
		temperatures = append(temperatures, tmp)
		fmt.Scan(&tmp)
	}
	fmt.Println(dailyTemperatures(temperatures))
}
