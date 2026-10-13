// 17.08. 马戏团人塔
/*
有个马戏团正在设计叠罗汉的表演节目，一个人要站在另一人的肩膀上。出于实际和美观的考虑，在上面的人要比下面的人矮一点且轻一点。
已知马戏团每个人的身高和体重，请编写代码计算叠罗汉最多能叠几个人。

提示：
height.length == weight.length <= 10000
*/
package main

import (
	"fmt"
	. "lc/pkg"
	"slices"
	"sort"
)

func bestSeqAtIndex(height []int, weight []int) int {
	n := len(height)
	var person [][]int = make([][]int, n) // person[i][0]: 第 i 个人的身高 	person[i][1]: 第 i 个人的体重
	for i := 0; i < n; i++ {
		person[i] = []int{height[i], weight[i]}
	}
	sort.Slice(person, func(i, j int) bool {
		// 身高做升序排序; 相同身高, 体重做降序排序, 保证第一次遇到某个身高的人时, 其体重是最重的
		return person[i][0] < person[j][0] || (person[i][0] == person[j][0] && person[i][1] > person[j][1])
	})

	// weightDp[i]: 截至第 i 个人, 体重的的递增子序列, 值为体重
	weightDp := make([]int, 0)

	for i := 0; i < n; i++ {
		if len(weightDp) == 0 || weightDp[len(weightDp)-1] < person[i][1] {
			weightDp = append(weightDp, person[i][1])
			continue
		} else if weightDp[len(weightDp)-1] == person[i][1] {
			continue
		} else {
			idx, _ := slices.BinarySearch(weightDp, person[i][1])
			weightDp[idx] = person[i][1]
		}
	}

	return len(weightDp)
}

// 示例：
// 输入：height = [65,70,56,75,60,68] weight = [100,150,90,190,95,110]
// 输出：6
// 解释：从上往下数，叠罗汉最多能叠 6 层：(56,90), (60,95), (65,100), (68,110), (70,150), (75,190)
func main() {
	fmt.Println("Input height:")
	height := CreateSlice[int]()
	fmt.Println("Input weight:")
	weight := CreateSlice[int]()
	fmt.Println(bestSeqAtIndex(height, weight))
}
