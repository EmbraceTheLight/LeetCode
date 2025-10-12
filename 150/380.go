// 380. O(1) 时间插入、删除和获取随机元素
/*
实现RandomizedSet 类：
RandomizedSet() 初始化 RandomizedSet 对象
bool insert(int val) 当元素 val 不存在时，向集合中插入该项，并返回 true ；否则，返回 false 。
bool remove(int val) 当元素 val 存在时，从集合中移除该项，并返回 true ；否则，返回 false 。
int getRandom() 随机返回现有集合中的一项（测试用例保证调用此方法时集合中至少存在一个元素）。每个元素应该有 相同的概率 被返回。
你必须实现类的所有函数，并满足每个函数的 平均 时间复杂度为 O(1) 。

示例：
输入
["RandomizedSet", "insert", "remove", "insert", "getRandom", "remove", "insert", "getRandom"]
[[], [1], [2], [2], [], [1], [2], []]
输出
[null, true, false, true, 2, true, false, 2]

解释
RandomizedSet randomizedSet = new RandomizedSet();
randomizedSet.insert(1); // 向集合中插入 1 。返回 true 表示 1 被成功地插入。
randomizedSet.remove(2); // 返回 false ，表示集合中不存在 2 。
randomizedSet.insert(2); // 向集合中插入 2 。返回 true 。集合现在包含 [1,2] 。
randomizedSet.getRandom(); // getRandom 应随机返回 1 或 2 。
randomizedSet.remove(1); // 从集合中移除 1 ，返回 true 。集合现在包含 [2] 。
randomizedSet.insert(2); // 2 已在集合中，所以返回 false 。
randomizedSet.getRandom(); // 由于 2 是集合中唯一的数字，getRandom 总是返回 2 。

提示：
-2^31 <= val <= 2^31 - 1
最多调用 insert、remove 和 getRandom 函数 2 * 105 次
在调用 getRandom 方法时，数据结构中 至少存在一个 元素。
*/
package main

import (
	"fmt"
	"math/rand"
)

type RandomizedSet struct {
	Set  map[int]int
	Vals []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		Set:  make(map[int]int), // k: val, v: index in Vals
		Vals: make([]int, 0),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.Set[val]; ok {
		return false
	}
	this.Set[val] = len(this.Vals)
	this.Vals = append(this.Vals, val)
	return true
}

// ! 难点&技巧: 删除元素时，需要将待删除元素 rem 与数组的最后一个元素做交换，
// ! 这样就不需要一个一个地调整 this.Set 存储的位于数组中 rem 后面的值的索引，从而将时间复杂度降到O(1)。
func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.Set[val]; !ok {
		return false
	}

	/* 交换待删除元素与最后一个元素 */
	idx := this.Set[val]

	// 取得数组最后一个元素的值和下标
	lastIdx := len(this.Vals) - 1
	lastVal := this.Vals[lastIdx]

	// 更新 this.Set 中原最后一个元素的索引值
	this.Set[lastVal] = idx

	// 更新 this.Vals 待删除元素位置的值为最后一个元素的值
	this.Vals[idx] = lastVal

	// 删除待删除元素
	delete(this.Set, val)
	this.Vals = this.Vals[:lastIdx]
	return true
}

func (this *RandomizedSet) GetRandom() int {
	randIdx := rand.Intn(len(this.Vals))
	return this.Vals[randIdx]
}
func main() {
	// Test case 1
	{
		randomizedSet := Constructor()
		fmt.Println(randomizedSet.Insert(1))   // 向集合中插入 1 。返回 true 表示 1 被成功地插入。
		fmt.Println(randomizedSet.Remove(2))   // 返回 false ，表示集合中不存在 2 。
		fmt.Println(randomizedSet.Insert(2))   // 向集合中插入 2 。返回 true 。集合现在包含 [1,2] 。
		fmt.Println(randomizedSet.GetRandom()) // getRandom 应随机返回 1 或 2 。
		fmt.Println(randomizedSet.Remove(1))   // 从集合中移除 1 ，返回 true 。集合现在包含 [2] 。
		fmt.Println(randomizedSet.Insert(2))   // 2 已在集合中，所以返回 false 。
		fmt.Println(randomizedSet.GetRandom()) // 由于 2 是集合中唯一的数字，getRandom 总是返回 2 。}
	}
	fmt.Println()
	fmt.Println("======================================================")
	fmt.Println()
	// Test case 2
	{
		randomizedSet := Constructor()
		fmt.Println(randomizedSet.Insert(0))
		fmt.Println(randomizedSet.Insert(1))
		fmt.Println(randomizedSet.Remove(0))
		fmt.Println(randomizedSet.Insert(2))
		fmt.Println(randomizedSet.Remove(1))
		fmt.Println(randomizedSet.GetRandom())
	}
}
