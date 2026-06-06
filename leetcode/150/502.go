// 502.IPO
/*
假设 力扣（LeetCode）即将开始 IPO 。为了以更高的价格将股票卖给风险投资公司，力扣 希望在 IPO 之前开展一些项目以增加其资本。
由于资源有限，它只能在 IPO 之前完成最多 k 个不同的项目。帮助 力扣 设计完成最多 k 个不同项目后得到最大总资本的方式。
给你 n 个项目。对于每个项目 i ，它都有一个纯利润 profits[i] ，和启动该项目需要的最小资本 capital[i] 。
最初，你的资本为 w。当你完成一个项目时，你将获得纯利润，且利润将被添加到你的总资本中。
总而言之，从给定项目中选择 最多 k 个不同项目的列表，以最大化最终资本，并输出最终可获得的最多资本。
答案保证在 32 位有符号整数范围内。



提示：
1 <= k <= 10^5
0 <= w <= 10^9
n == profits.length
n == capital.length
1 <= n <= 10^5
0 <= profits[i] <= 10^4
0 <= capital[i] <= 10^9
*/
package main

import (
	"container/heap"
	"fmt"
	"lc/pkg"
	"sort"
)

type profitCapital struct {
	profit  int
	capital int
}

type pc []*profitCapital
type priorityQueue struct {
	pq pc
}

func (pq *priorityQueue) Pop() any {
	ret := pq.pq[pq.Len()-1]
	pq.pq = pq.pq[:pq.Len()-1]
	return ret
}

func (pq *priorityQueue) Push(x any) {
	pq.pq = append(pq.pq, x.(*profitCapital))
}
func (pq *priorityQueue) Len() int {
	return len(pq.pq)
}
func (pq *priorityQueue) Swap(i, j int) {
	pq.pq[i], pq.pq[j] = pq.pq[j], pq.pq[i]
}
func (pq *priorityQueue) Less(i, j int) bool {
	return pq.pq[i].profit > pq.pq[j].profit
}

func (q *pc) Push(x *profitCapital) {
	*q = append(*q, x)
}

func (q *pc) Len() int {
	return len(*q)
}
func (q *pc) Pop() *profitCapital {
	ret := (*q)[q.Len()-1]
	*q = (*q)[:q.Len()-1]
	return ret
}

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	pq := &priorityQueue{pq: make([]*profitCapital, 0)}
	heap.Init(pq)

	tmppq1 := make(pc, 0)
	for i := 0; i < len(profits); i++ {
		heap.Push(pq, &profitCapital{
			profit:  profits[i],
			capital: capital[i],
		})
	}

	var project *profitCapital
	for i := 0; i < k; i++ {
		project = heap.Pop(pq).(*profitCapital)
		for pq.Len() > 0 && project.capital > w {
			tmppq1.Push(project)
			project = heap.Pop(pq).(*profitCapital)
		}
		if project.capital > w {
			break
		}
		w += project.profit
		for tmppq1.Len() > 0 {
			heap.Push(pq, tmppq1.Pop())
		}
	}
	return w
}

// * 思路: 使用大根堆储存利润, 每次遍历成本小于当前资本的项目, 将其利润存入大根堆, 取出利润最大的项目, 更新资本
func findMaximizedCapitalBetter(k int, w int, profits []int, capital []int) int {
	bigRootHeap := &hp{sli: make([]int, 0)}
	heap.Init(bigRootHeap)
	type pc struct {
		profit  int
		capital int
	}
	pcSli := make([]*pc, 0)
	for i := 0; i < len(profits); i++ {
		pcSli = append(pcSli, &pc{
			profit:  profits[i],
			capital: capital[i],
		})
	}
	sort.Slice(pcSli, func(i, j int) bool {
		return pcSli[i].capital < pcSli[j].capital
	})

	for i := 0; k > 0; k-- {
		for i < len(profits) && pcSli[i].capital <= w {
			heap.Push(bigRootHeap, pcSli[i].profit)
			i++
		}
		if bigRootHeap.Len() == 0 {
			break
		}
		w += heap.Pop(bigRootHeap).(int)
	}
	return w
}

type hp struct {
	sli []int
}

func (hp *hp) Len() int {
	return len(hp.sli)
}
func (hp *hp) Less(i, j int) bool {
	return hp.sli[i] > hp.sli[j]
}
func (hp *hp) Swap(i, j int) {
	hp.sli[i], hp.sli[j] = hp.sli[j], hp.sli[i]
}

func (hp *hp) Push(x any) {
	hp.sli = append(hp.sli, x.(int))
}
func (hp *hp) Pop() any {
	ret := hp.sli[hp.Len()-1]
	hp.sli = hp.sli[:hp.Len()-1]
	return ret
}

// 示例 1：
// 输入：k = 2, w = 0, profits = [1,2,3], capital = [0,1,1]
// 输出：4
// 解释：
// 由于你的初始资本为 0，你仅可以从 0 号项目开始。
// 在完成后，你将获得 1 的利润，你的总资本将变为 1。
// 此时你可以选择开始 1 号或 2 号项目。
// 由于你最多可以选择两个项目，所以你需要完成 2 号项目以获得最大的资本。
// 因此，输出最后最大化的资本，为 0 + 1 + 3 = 4。
//
// 示例 2：
// 输入：k = 3, w = 0, profits = [1,2,3], capital = [0,1,2]
// 输出：6
func main() {
	fmt.Println("Input k,w:")
	var k, w int
	fmt.Scan(&k, &w)
	fmt.Println("Input profits:")
	profits := pkg.CreateSlice[int]()
	fmt.Println("Input capitals:")
	capitals := pkg.CreateSlice[int]()
	fmt.Println(findMaximizedCapital(k, w, profits, capitals))
	fmt.Println(findMaximizedCapitalBetter(k, w, profits, capitals))
}
