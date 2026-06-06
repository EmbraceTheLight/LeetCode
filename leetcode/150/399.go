// 399. 除法求值
/*
给你一个变量对数组 equations 和一个实数值数组 values 作为已知条件，其中 equations[i] = [Ai, Bi] 和 values[i] 共同表示等式 Ai / Bi = values[i] 。每个 Ai 或 Bi 是一个表示单个变量的字符串。
另有一些以数组 queries 表示的问题，其中 queries[j] = [Cj, Dj] 表示第 j 个问题，请你根据已知条件找出 Cj / Dj = ? 的结果作为答案。
返回 所有问题的答案 。如果存在某个无法确定的答案，则用 -1.0 替代这个答案。如果问题中出现了给定的已知条件中没有出现的字符串，也需要用 -1.0 替代这个答案。
注意：输入总是有效的。你可以假设除法运算中不会出现除数为 0 的情况，且不存在任何矛盾的结果。
注意：未在等式列表中出现的变量是未定义的，因此无法确定它们的答案。

示例 1：
输入：equations = [["a","b"],["b","c"]], values = [2.0,3.0], queries = [["a","c"],["b","a"],["a","e"],["a","a"],["x","x"]]
输出：[6.00000,0.50000,-1.00000,1.00000,-1.00000]
解释：
条件：a / b = 2.0, b / c = 3.0
问题：a / c = ?, b / a = ?, a / e = ?, a / a = ?, x / x = ?
结果：[6.0, 0.5, -1.0, 1.0, -1.0 ]
注意：x 是未定义的 => -1.0

示例 2：
输入：equations = [["a","b"],["b","c"],["bc","cd"]], values = [1.5,2.5,5.0], queries = [["a","c"],["c","b"],["bc","cd"],["cd","bc"]]
输出：[3.75000,0.40000,5.00000,0.20000]

示例 3：
输入：equations = [["a","b"]], values = [0.5], queries = [["a","b"],["b","a"],["a","c"],["x","y"]]
输出：[0.50000,2.00000,-1.00000,-1.00000]
提示：

1 <= equations.length <= 20
equations[i].length == 2
1 <= Ai.length, Bi.length <= 5
values.length == equations.length
0.0 < values[i] <= 20.0
1 <= queries.length <= 20
queries[i].length == 2
1 <= Cj.length, Dj.length <= 5
Ai, Bi, Cj, Dj 由小写英文字母与数字组成
*/
package main

import (
	"fmt"
	"lc/pkg"
)

// 暴力解法
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	equationMap := make(map[string]map[string]float64)
	for i := 0; i < len(equations); i++ {
		e := equations[i]
		eq1, eq2 := e[0], e[1]
		if _, ok := equationMap[eq1]; !ok {
			equationMap[eq1] = make(map[string]float64)
		}
		if _, ok := equationMap[eq2]; !ok {
			equationMap[eq2] = make(map[string]float64)
		}
		equationMap[eq1][eq1] = 1.0
		equationMap[eq1][eq2] = values[i]
		equationMap[eq2][eq1] = 1.0 / values[i]
		equationMap[eq2][eq2] = 1.0
	}
	buildEquationGraph(equationMap)
	var ans []float64
	for i := 0; i < len(queries); i++ {
		q := queries[i]
		eq1, eq2 := q[0], q[1]
		if value, ok := equationMap[eq1][eq2]; ok {
			ans = append(ans, value)
		} else {
			ans = append(ans, -1)
		}
	}
	return ans
}

// buildEquationGraph 构造方程图, 例如, 现在有 a->b, b->c, 则自动推导出 a->c 并更新 map
func buildEquationGraph(equationMap map[string]map[string]float64) {
	for { // 外层循环, 直到所有变量都有了关系, 包括间接关系, 每循环一轮, 就有一层更深的关系, 如 a->b, b->c, c->d, 则第一轮循环计算出 a->c, 第二轮循环则计算出 a->d
		finish := true
		for first, v := range equationMap { // 例如 a / b, k 为 a, v 为 map[a], 遍历 v, 得到 b 到其他变量的关系, 并加入队列, 最终算出 a 到其他变量的关系
			for second, v2 := range v {
				k2Map := equationMap[second]
				for third, v3 := range k2Map {
					if _, ok := equationMap[first][third]; !ok {
						finish = false
						equationMap[first][third] = v2 * v3
						equationMap[third][first] = 1 / (v2 * v3)
					}
				}
			}
		}
		if finish {
			break
		}
	}
}

func calcEquation2(equations [][]string, values []float64, queries [][]string) []float64 {
	n := len(equations)
	uf := UnionFind{}

	// 预留两倍空间, 因为有可能所有变量名都不同
	uf.init(2 * n)

	// 预处理, 构建变量名 -> id 的映射
	// 这样做是为了并查集操作的方便, 底层 parent 可以使用整型数组
	mp := make(map[string]int)
	id := 0
	for i := 0; i < n; i++ {
		eq1, eq2 := equations[i][0], equations[i][1]
		if _, ok := mp[eq1]; !ok {
			mp[eq1] = id
			id++
		}
		if _, ok := mp[eq2]; !ok {
			mp[eq2] = id
			id++
		}
		uf.union(mp[eq1], mp[eq2], values[i])
	}

	var ans []float64
	for i := 0; i < len(queries); i++ {
		eq1, eq2 := queries[i][0], queries[i][1]
		_, ok1 := mp[eq1]
		_, ok2 := mp[eq2]
		if !ok1 || !ok2 {
			ans = append(ans, -1)
		} else {
			ans = append(ans, uf.getValue(mp[eq1], mp[eq2]))
		}
	}
	return ans
}

// UnionFind 并查集, 压缩路径
type UnionFind struct {
	// parent[i] 表示第 i 个节点的父节点的 id, 初始时 parent[i] = i, 表示 i 自己是自己的父节点
	parent []int

	// weight[i] 表示第 i 个节点到其父节点的权重, 初始时 weight[i] = 1.0, 表示 i 到其根节点的权重为 1.0
	weight []float64
}

// 查找 id 为 x 的节点的根节点. 查找时路径压缩
func (uf *UnionFind) find(cur int) int {
	parent := uf.parent[cur]
	if parent != cur {
		uf.parent[cur] = uf.find(parent)
		uf.weight[cur] = uf.weight[cur] * uf.weight[parent]
	}
	return uf.parent[cur]
}

// numerator, denominator 分别为分子、分母的 id, value 为分子/分母的值
func (uf *UnionFind) union(numerator, denominator int, value float64) {
	rootNumerator := uf.find(numerator)
	rootDenominator := uf.find(denominator)
	// 二者已经在同一集合中, 不需要合并
	if rootNumerator == rootDenominator {
		return
	}
	uf.parent[rootNumerator] = rootDenominator

	// 合并时, 计算分子的根节点到分母的根节点的权重
	// 已知: 当前分子到其根节点的权重 uf.weight[numerator]; 当前分母到其根节点的权重 uf.weight[denominator]; 当前分子到的值 value. 求: 分子根节点到分母的根节点的权重 x
	// 分子到分母根节点有两条路: 1. 先到分子根节点, 再到分母根节点; 2. 先到分母根节点, 再到分子根节点, 两条路最终权重应该相等, 因此, 有下面的公式:
	// uf.weight[numerator] * x = uf.weight[denominator] * value, 即 x = uf.weight[denominator] * value / uf.weight[numerator], 最终得到分子根节点到分母根节点的权重 x, 完成合并
	uf.weight[rootNumerator] = uf.weight[denominator] * value / uf.weight[numerator]
}

func (uf *UnionFind) init(size int) {
	uf.parent = make([]int, size)
	uf.weight = make([]float64, size)
	for i := 0; i < size; i++ {
		uf.parent[i] = i
		uf.weight[i] = 1.0
	}
}

// 获得 分子/分母 的值. 如果分子/分母不在同一个集合中, 表明他们之间没有关系, 则返回 -1
func (uf *UnionFind) getValue(numeratorId, denominatorId int) float64 {
	rootNumerator := uf.find(numeratorId)
	rootDenominator := uf.find(denominatorId)
	if rootNumerator == rootDenominator {
		return uf.weight[numeratorId] / uf.weight[denominatorId]
	} else {
		return -1
	}
}

// Test Case1:	equations = [["a","b"],["b","c"]], values = [2.0,3.0], queries = [["a","c"],["b","a"],["a","e"],["a","a"],["x","x"]]		Output: [6.00000,0.50000,-1.00000,1.00000,-1.00000]
// Test Case2:	equations = [["a","b"],["b","c"],["bc","cd"]], values = [1.5,2.5,5.0], queries = [["a","c"],["c","b"],["bc","cd"],["cd","bc"]]		Output: [3.75000,0.40000,5.00000,0.20000]
// Test Case3:	equations = [["a","b"]], values = [0.5], queries = [["a","b"],["b","a"],["a","c"],["x","y"]]		Output: [0.50000,2.00000,-1.00000,-1.00000]
func main() {
	fmt.Println("Input equations")
	equations := pkg.CreateSlice2D[string]()

	fmt.Println("Input values")
	values := pkg.CreateSlice[float64]()

	fmt.Println("Input queries")
	queries := pkg.CreateSlice2D[string]()

	fmt.Println(calcEquation2(equations, values, queries))
}
