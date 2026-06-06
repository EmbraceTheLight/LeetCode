// 134. 加油站
// ! 不必强求一定要理解这道题，因为设计到一些数学证明。
// https://leetcode-cn.com/problems/gas-station/
package main

import "fmt"

/*
在一条环路上有 n 个加油站，其中第 i 个加油站有汽油 gas[i] 升。
你有一辆油箱容量无限的的汽车，从第 i 个加油站开往第 i+1 个加油站需要消耗汽油 cost[i] 升。你从其中的一个加油站出发，开始时油箱为空。
给定两个整数数组 gas 和 cost ，如果你可以按顺序绕环路行驶一周，则返回出发时加油站的编号，否则返回 -1 。如果存在解，则 保证 它是 唯一 的。

示例 1:
输入: gas = [1,2,3,4,5], cost = [3,4,5,1,2]
输出: 3
解释:
从 3 号加油站(索引为 3 处)出发，可获得 4 升汽油。此时油箱有 = 0 + 4 = 4 升汽油
开往 4 号加油站，此时油箱有 4 - 1 + 5 = 8 升汽油
开往 0 号加油站，此时油箱有 8 - 2 + 1 = 7 升汽油
开往 1 号加油站，此时油箱有 7 - 3 + 2 = 6 升汽油
开往 2 号加油站，此时油箱有 6 - 4 + 3 = 5 升汽油
开往 3 号加油站，你需要消耗 5 升汽油，正好足够你返回到 3 号加油站。
因此，3 可为起始索引。

示例 2:
输入: gas = [2,3,4], cost = [3,4,3]
输出: -1
解释:
你不能从 0 号或 1 号加油站出发，因为没有足够的汽油可以让你行驶到下一个加油站。
我们从 2 号加油站出发，可以获得 4 升汽油。 此时油箱有 = 0 + 4 = 4 升汽油
开往 0 号加油站，此时油箱有 4 - 3 + 2 = 3 升汽油
开往 1 号加油站，此时油箱有 3 - 3 + 3 = 3 升汽油
你无法返回 2 号加油站，因为返程需要消耗 4 升汽油，但是你的油箱只有 3 升汽油。
因此，无论怎样，你都不可能绕环路行驶一周。

提示:
n == gas.length == cost.length
1 <= n <= 10^5
0 <= gas[i], cost[i] <= 10^4
输入保证答案唯一。
*/

// 朴素做法，一个一个站试，O(n^2)，会超时
func canCompleteCircuitTimeOut(gas []int, cost []int) int {
	var ans int = -1
	n := len(gas)
	remainings := make([]int, n) // remainings[i] 表示从 第 i 个加油站出发，到第 i+1 个加油站还剩多少汽油
	s := 0
	for i := 0; i < n; i++ {
		remainings[i] = gas[i] - cost[i]
		s += remainings[i]
	}
	if s < 0 {
		return -1
	}

	for i := 0; i < n; i++ {
		if remainings[i] < 0 {
			continue
		}

		remaining := 0
		count := 0
		for ; count < n; count++ {
			nextStation := (count + i) % n
			remaining += remainings[nextStation]
			if remaining < 0 {
				break
			}
		}
		if count == n {
			ans = i
			break
		}
	}
	return ans
}

// O(n) 解法
func canCompleteCircuit(gas []int, cost []int) int {
	var ans int = -1
	n := len(gas)
	remainings := make([]int, n) // remainings[i] 表示从 第 i 个加油站出发，到第 i+1 个加油站还剩多少汽油
	s := 0
	for i := 0; i < n; i++ {
		remainings[i] = gas[i] - cost[i]
		s += remainings[i]
	}

	// 若汽油不足，则无法完成一周，否则一定有解。
	if s < 0 {
		return -1
	}

	// 计算前缀和, 计算完毕后 remainings[i] 表示从 第 0 个加油站出发，到第 i 个加油站还剩多少汽油
	// remainings[i] - remainings[i-1]如果大于 0， 就说明从第 i - 1 个加油站出发，到第 i 个加油站还能多剩余一些油
	for i := 1; i < n; i++ {
		remainings[i] += remainings[i-1]
	}

	// 假设先从第 0 个加油站出发，寻找一个加油站，这个站是从加油站 0 出发后能到达的消耗最多的站点
	// 比如 remainings = [-2, -4, -6, -3, 0]，到达第二个站点时亏空油最多。
	// 因为题目保证答案唯一，因此就从第三个站点开始，在到达第二站之前有足够的余量到达第二站，就可以保证绕环路一周。
	minRemIdx := 0
	for i := 0; i < n; i++ {
		if remainings[minRemIdx] > remainings[i] {
			minRemIdx = i
		}
	}

	// 因为是环路，因此需要对结果取模
	ans = (minRemIdx + 1) % n
	return ans
}
func main() {
	// Test Case 1, Expected Output: 3
	gas := []int{1, 2, 3, 4, 5}
	cost := []int{3, 4, 5, 1, 2}
	fmt.Println(canCompleteCircuit(gas, cost))

	// Test Case 2, Expected Output: -1
	gas = []int{2, 3, 4}
	cost = []int{3, 4, 3}
	fmt.Println(canCompleteCircuit(gas, cost))
}
