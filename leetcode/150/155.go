// 155. 最小栈
/*
设计一个支持 push，pop，top 操作，并能在常数时间内检索到最小元素的栈。

实现 MinStack 类:
MinStack() 初始化堆栈对象。
void push(int val) 将元素val推入堆栈。
void pop() 删除堆栈顶部的元素。
int top() 获取堆栈顶部的元素。
int getMin() 获取堆栈中的最小元素。


示例 1:
输入：
["MinStack","push","push","push","getMin","pop","top","getMin"]
[[],[-2],[0],[-3],[],[],[],[]]

输出：
[null,null,null,null,-3,null,0,-2]

解释：
MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin();   --> 返回 -3.
minStack.pop();
minStack.top();      --> 返回 0.
minStack.getMin();   --> 返回 -2.


提示：

-2^31 <= val <= 2^31 - 1
pop、top 和 getMin 操作总是在 非空栈 上调用
push, pop, top, and getMin最多被调用 3 * 104 次
*/
package main

import (
	"container/heap"
	"fmt"
)

// ! 下次使用辅助栈实现。辅助栈中存放当前栈中的最小值，每次push时，将当前值与辅助栈顶部的最小值比较，
// ! 若当前值小于辅助栈顶部的最小值，则将当前值压入辅助栈, 否则, 将最小值入栈
// ! pop时，若当前值等于辅助栈顶部的最小值，则pop掉辅助栈顶部的最小值。
// ! top时，返回辅助栈顶部的最小值。getMin时，返回辅助栈顶部的最小值。
type minHeap []int

func (h *minHeap) Len() int {
	return len(*h)
}

func (h *minHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *minHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *minHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *minHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type MinStack struct {
	len     int
	minHeap minHeap
	stack   []int
}

func Constructor155() *MinStack {
	ms := &MinStack{
		len:     0,
		minHeap: minHeap{},
		stack:   []int{},
	}
	heap.Init(&ms.minHeap)
	return ms
}

func (ms *MinStack) Push(val int) {
	ms.stack = append(ms.stack, val)
	ms.len++
	heap.Push(&ms.minHeap, val)
}

func (ms *MinStack) Pop() {
	val := ms.stack[ms.len-1]
	ms.stack = ms.stack[:ms.len-1]
	idx := 0
	for i, v := range ms.minHeap {
		if v == val {
			idx = i
			break
		}
	}
	heap.Remove(&ms.minHeap, idx)
	ms.len--
}

func (ms *MinStack) Top() int {
	return ms.stack[ms.len-1]
}

func (ms *MinStack) GetMin() int {
	return ms.minHeap[0]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
func main() {
	// Test case 1
	{
		minStack := Constructor155()
		minStack.Push(-2)
		minStack.Push(0)
		minStack.Push(-3)
		fmt.Println(minStack.GetMin())
		minStack.Pop()
		fmt.Println(minStack.Top())
		fmt.Println(minStack.GetMin())
	}

	fmt.Println()
	// Test case 2
	{
		minStack := Constructor155()
		minStack.Push(-2)
		minStack.Push(0)
		minStack.Push(-1)
		fmt.Println(minStack.GetMin())
		fmt.Println(minStack.Top())
		minStack.Pop()
		fmt.Println(minStack.GetMin())
	}
}
