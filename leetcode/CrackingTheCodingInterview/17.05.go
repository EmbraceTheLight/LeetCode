// 17.05 字母与数字
/*
给定一个放有字母和数字的数组，找到最长的子数组，且包含的字母和数字的个数相同。
返回该子数组，若存在多个最长子数组，返回左端点下标值最小的子数组。若不存在这样的数组，返回一个空数组。

提示：
array.length <= 100000
*/
package main

import (
	"fmt"
	. "lc/pkg"
)

func findLongestSubarray(array []string) []string {
	/*
	* 思路：前缀和求出前 i 个字符中有多少个数字/字母 numCount、lettCount.
	* 若 numCount[i] == letterCount[i], 则说明 array[0, i] 中的数字和字母个数相等, 为候选答案
	* 对于 i, 假设数字 - 字母个数结果为 numCount[i] - lettCount[i]，
	* 如果对于 j > i, 存在 numCount[j] - lettCount[j] == numCount[i] - lettCount[i], 则说明 array[i+1, j] 中的数字和字母个数相等, 即为候选答案
	 */
	isDigit := func(b byte) bool {
		return b >= '0' && b <= '9'
	}
	start, end := -1, -1
	mp := make(map[int]int, len(array))    // key: array[0, i] 中数字的个数减去字母的个数. value: 第一次出现 key 时的下标
	numCount := make([]int, len(array))    // array[0, i] 中数字的个数
	letterCount := make([]int, len(array)) // array[0, i] 中字母的个数
	if isDigit(array[0][0]) {
		numCount[0] = 1
	} else {
		letterCount[0] = 1
	}
	mp[numCount[0]-letterCount[0]] = 0
	for i := 1; i < len(array); i++ {
		numCount[i] = numCount[i-1]
		letterCount[i] = letterCount[i-1]
		if isDigit(array[i][0]) == true {
			numCount[i] = numCount[i-1] + 1
		} else {
			letterCount[i] = letterCount[i-1] + 1
		}
		if numCount[i] == letterCount[i] {
			if i+1 > end-start+1 {
				start, end = 0, i+1
			}
		}
		if idx, ok := mp[numCount[i]-letterCount[i]]; ok {
			if i-idx+1 > end-start+1 {
				start, end = idx+1, i+1
			}
		} else {
			mp[numCount[i]-letterCount[i]] = i
		}
	}
	if start == -1 {
		return []string{}
	}
	return array[start:end]
}

// 示例 1：
// 输入：["A","1","B","C","D","2","3","4","E","5","F","G","6","7","H","I","J","K","L","M"]
// 输出：["A","1","B","C","D","2","3","4","E","5","F","G","6","7"]
//
// 示例 2：
// 输入：["A","A"]
// 输出：[]
func main() {
	array := CreateSlice[string]()
	fmt.Println(findLongestSubarray(array))
}
