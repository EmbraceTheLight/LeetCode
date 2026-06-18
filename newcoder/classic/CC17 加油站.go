// CC17 加油站
/*
描述
环形路上有n个加油站，第i个加油站的汽油量是gas[i].
你有一辆车，车的油箱可以无限装汽油。从加油站i走到下一个加油站（i+1）花费的油量是cost[i]，你从一个加油站出发，刚开始的时候油箱里面没有汽油。
求从哪个加油站出发可以在环形路上走一圈。返回加油站的下标，如果没有答案的话返回-1。
注意：答案保证唯一。
*/
package main

import (
	"fmt"
	. "lc/pkg"
)

func canCompleteCircuit(gas []int, cost []int) int {
	// write code here
	ans := -1
	n := len(gas)
	start := 0
	gasLeft := 0
	minVal := 0
	for i := 0; i < n; i++ {
		gasLeft = gasLeft + gas[i] - cost[i]
		if gasLeft < minVal {
			minVal = gasLeft
			start = (i + 1) % n
		}
	}

	cnt := 0
	gasLeft = 0
	ans = start
	for i := start; cnt < n; cnt, i = cnt+1, (i+1)%n {
		gasLeft = gasLeft + gas[i] - cost[i]
		if gasLeft < 0 {
			ans = -1
			break
		}
	}
	return ans
}

// 示例1
// 输入: [2,3,1],[3,1,2]
// 返回值: 1
func main() {
	fmt.Println("Input gas:")
	gas := CreateSlice[int]()
	fmt.Println("Input cost:")
	cost := CreateSlice[int]()
	fmt.Println(canCompleteCircuit(gas, cost))
}
