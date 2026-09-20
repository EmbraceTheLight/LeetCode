// 16.10 生存人数
/*
给定 N 个人的出生年份和死亡年份，第 i 个人的出生年份为 birth[i]，死亡年份为 death[i]，实现一个方法以计算生存人数最多的年份。
你可以假设所有人都出生于 1900 年至 2000 年（含 1900 和 2000 ）之间。
如果一个人在某一年的任意时期处于生存状态，那么他应该被纳入那一年的统计中。
例如，生于 1908 年、死于 1909 年的人应当被列入 1908 年和 1909 年的计数。
如果有多个年份生存人数相同且均为最大值，输出其中最小的年份。

提示：
0 < birth.length == death.length <= 10000
birth[i] <= death[i]
*/
package main

import (
	"fmt"
	. "lc/pkg"
)

func maxAliveYear(birth []int, death []int) int {
	birthSlice := [101]int{}
	deathSlice := [101]int{}
	for i := 0; i < len(birth); i++ {
		birthSlice[birth[i]-1900]++
	}
	for i := 0; i < len(death); i++ {
		if death[i] < 2000 {
			deathSlice[death[i]-1900+1]++ // 死亡年份归于下一年
		}
	}
	birthPrefixSum := make([]int, len(birthSlice))
	deathPrefixSum := make([]int, len(deathSlice))
	birthPrefixSum[0] = birthSlice[0]
	deathPrefixSum[0] = deathSlice[0]
	for i := 1; i < len(birthSlice); i++ {
		birthPrefixSum[i] = birthPrefixSum[i-1] + birthSlice[i]
	}
	for i := 1; i < len(deathSlice); i++ {
		deathPrefixSum[i] = deathPrefixSum[i-1] + deathSlice[i]
	}
	var ans int
	var cntOfSurvival int
	for i := 0; i < len(birthSlice); i++ {
		cnt := birthPrefixSum[i] - deathPrefixSum[i]
		if cnt > cntOfSurvival {
			cntOfSurvival = cnt
			ans = i + 1900
		}
	}
	return ans
}

// 示例：
// 输入：
// birth = [1900, 1901, 1950]
// death = [1948, 1951, 2000]
// 输出： 1901
func main() {
	birth := CreateSlice[int]()
	death := CreateSlice[int]()
	for i := 0; i < len(birth); i++ {
		if birth[i] == death[i] {
			fmt.Printf("%d\t%d\t%d\n", birth[i], birth[i], death[i])
		}
	}
	fmt.Println(maxAliveYear(birth, death))
}
