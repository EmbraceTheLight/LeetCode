// 面试题 03.05. 栈排序
/*
栈排序。 编写程序，对栈进行排序使最小元素位于栈顶。
最多只能使用一个其他的临时栈存放数据，但不得将元素复制到别的数据结构（如数组）中。
该栈支持如下操作：push、pop、peek 和 isEmpty。当栈为空时，peek 返回 -1。

提示：
栈中的元素数目在[0, 5000]范围内。
*/
package main

import "fmt"

type stack0305 struct {
	nums []int
}

func (this *stack0305) push(v int) {
	this.nums = append(this.nums, v)
}
func (this *stack0305) pop() int {
	if this.empty() == true {
		return -1
	}
	v := this.nums[this.len()-1]
	this.nums = this.nums[:this.len()-1]
	return v
}

func (this *stack0305) empty() bool {
	return this.len() == 0
}
func (this *stack0305) len() int {
	return len(this.nums)
}
func (this *stack0305) peek() int {
	if this.empty() == true {
		return -1
	}
	return this.nums[len(this.nums)-1]
}

type SortedStack struct {
	stk *stack0305
	tmp *stack0305
}

func Constructor0305() SortedStack {
	return SortedStack{
		stk: &stack0305{},
		tmp: &stack0305{},
	}
}

func (this *SortedStack) Push(val int) {
	if this.stk.empty() == true {
		this.stk.push(val)
		return
	}
	for this.stk.empty() == false && this.stk.peek() < val {
		this.tmp.push(this.stk.pop())
	}
	this.stk.push(val)
	for this.tmp.empty() == false {
		this.stk.push(this.tmp.pop())
	}
}

func (this *SortedStack) Pop() {
	_ = this.stk.pop()
}

func (this *SortedStack) Peek() int {
	return this.stk.peek()
}

func (this *SortedStack) IsEmpty() bool {
	return this.stk.empty()
}

// 示例 1：
//
//	输入：
//
// ["SortedStack", "push", "push", "peek", "pop", "peek"]
// [[], [1], [2], [], [], []]
//
//	输出：
//
// [null,null,null,1,null,2]
//
// 示例 2：
//
//	输入：
//
// ["SortedStack", "pop", "pop", "push", "pop", "isEmpty"]
// [[], [], [], [1], [], []]
//
//	输出：
//
// [null,null,null,null,null,true]
func main() {
	sortedStack := Constructor0305()
	sortedStack.Push(1)
	sortedStack.Push(2)
	fmt.Println(sortedStack.Peek())
	sortedStack.Pop()
	fmt.Println(sortedStack.Peek())
}
