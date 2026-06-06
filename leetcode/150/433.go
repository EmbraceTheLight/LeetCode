// 433. 最小基因变化
/*
基因序列可以表示为一条由 8 个字符组成的字符串，其中每个字符都是 'A'、'C'、'G' 和 'T' 之一。
假设我们需要调查从基因序列 start 变为 end 所发生的基因变化。一次基因变化就意味着这个基因序列中的一个字符发生了变化。
例如，"AACCGGTT" --> "AACCGGTA" 就是一次基因变化。
另有一个基因库 bank 记录了所有有效的基因变化，只有基因库中的基因才是有效的基因序列。（变化后的基因必须位于基因库 bank 中）
给你两个基因序列 start 和 end ，以及一个基因库 bank ，请你找出并返回能够使 start 变化为 end 所需的最少变化次数。如果无法完成此基因变化，返回 -1 。
注意：起始基因序列 start 默认是有效的，但是它并不一定会出现在基因库中。

示例 1：
输入：start = "AACCGGTT", end = "AACCGGTA", bank = ["AACCGGTA"]
输出：1

示例 2：
输入：start = "AACCGGTT", end = "AAACGGTA", bank = ["AACCGGTA","AACCGCTA","AAACGGTA"]
输出：2

示例 3：
输入：start = "AAAAACCC", end = "AACCCCCC", bank = ["AAAACCCC","AAACCCCC","AACCCCCC"]
输出：3

提示：
start.length == 8
end.length == 8
0 <= bank.length <= 10
bank[i].length == 8
start、end 和 bank[i] 仅由字符 ['A', 'C', 'G', 'T'] 组成
*/
package main

import (
	"fmt"
	"lc/pkg"
)

// convertInfo 记录当前基因序列以及到达当前序列的步数
type convertInfo struct {
	cur  string
	step int
}

func minMutation(startGene string, endGene string, bank []string) int {
	bankMap := make(map[string]bool)
	visit := make(map[string]bool)
	for _, gene := range bank {
		bankMap[gene] = true
	}
	genes := []string{"A", "C", "G", "T"}
	queue := []convertInfo{{cur: startGene, step: 0}}
	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]
		visit[top.cur] = true
		tmp := top.cur
		for i := 0; i < 8; i++ {
			for _, gene := range genes {
				tmp = top.cur[:i] + gene + top.cur[i+1:]
				if _, ok := bankMap[tmp]; !ok || visit[tmp] {
					continue
				}
				if tmp == endGene {
					return top.step + 1
				}
				queue = append(queue, convertInfo{cur: tmp, step: top.step + 1})
			}
		}
	}
	return -1
}

// Test Case1: start = "AACCGGTT", end = "AACCGGTA", bank = ["AACCGGTA"]	Output: 1
// Test Case2: start = "AACCGGTT", end = "AAACGGTA", bank = ["AACCGGTA","AACCGCTA","AAACGGTA"]	Output: 2
// Test Case3: start = "AAAAACCC", end = "AACCCCCC", bank = ["AAAACCCC","AAACCCCC","AACCCCCC"]	Output: 3
func main() {
	var start, end string
	fmt.Println("input start and end:")
	fmt.Scan(&start, &end)
	fmt.Println("input bank:")
	bank := pkg.CreateSlice[string]()
	fmt.Println(minMutation(start, end, bank))
}
