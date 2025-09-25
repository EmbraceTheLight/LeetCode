/*
354. 俄罗斯套娃信封问题  TODO: [hard]
给你一个二维整数数组 envelopes ，其中 envelopes[i] = [wi, hi] ，表示第 i 个信封的宽度和高度。
当另一个信封的宽度和高度都比这个信封大的时候，这个信封就可以放进另一个信封里，如同俄罗斯套娃一样。
请计算 最多能有多少个 信封能组成一组“俄罗斯套娃”信封（即可以把一个信封放到另一个信封里面）。
注意：不允许旋转信封。

示例 1：
输入：envelopes = [[5,4],[6,4],[6,7],[2,3]]
输出：3
解释：最多信封的个数为 3, 组合为: [2,3] => [5,4] => [6,7]。

示例 2：
输入：envelopes = [[1,1],[1,1],[1,1]]
输出：1

提示：
1 <= envelopes.length <= 10^5
envelopes[i].length == 2
1 <= wi, hi <= 10^5
*/
package main

import (
	"fmt"
	"sort"
	"lc/pkg"
)

// TODO：对于二维的信封，可以将其一维（比如宽度w）按照升序排序固定
// TODO：因为题目要求是严格递增的，所以针对每个宽度w，选择一个信封，将第二个维度h来降序排序，
// TODO：那么相同宽度的信封，其高度是依次递减的，不可能组成长度超过1的严格递增序列，保证了针对每个宽度信封只能选择一个
// TODO：然后就可以忽略w维度，求出h维度的最长严格递增子序列
func maxEnvelopes(envelopes [][]int) int {
	//TODO：宽度作为第一关键字升序，高度作为第二关键字降序
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] != envelopes[j][0] {
			return envelopes[i][0] < envelopes[j][0]
		} else {
			return envelopes[i][1] > envelopes[j][1]
		}
	})
	tmp := make([]int, 0) //tmp[i]表示高度h可以组成的长度为i的最长严格递增子序列的末尾元素最小值
	for _, e := range envelopes {
		h := e[1]
		if ind := sort.SearchInts(tmp, h); ind < len(tmp) {
			tmp[ind] = h
		} else {
			tmp = append(tmp, h)
		}
	}
	return len(tmp)
}
func main() {
	env := pkg.CreateIntSlice2[int]()
	fmt.Println(maxEnvelopes(env))
}
