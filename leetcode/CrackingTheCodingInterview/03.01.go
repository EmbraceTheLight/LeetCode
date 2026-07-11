// 03.01. 三合一
/*
三合一。描述如何只用一个数组来实现三个栈。
你应该实现push(stackNum, value)、pop(stackNum)、isEmpty(stackNum)、peek(stackNum)方法。stackNum表示栈下标，value表示压入的值。
构造函数会传入一个stackSize参数，代表每个栈的大小。

提示：
0 <= stackNum <= 2
*/
package main

import "fmt"

type TripleInOne struct {
	stack []int
	sLen  [3]int
	size  int
}

func Constructor(stackSize int) TripleInOne {
	return TripleInOne{
		stack: make([]int, stackSize*3),
		size:  stackSize,
	}
}

func (this *TripleInOne) Push(stackNum int, value int) {
	sLen := this.sLen[stackNum]
	if sLen >= this.size {
		return
	}
	targetIdx := stackNum + sLen*3
	this.stack[targetIdx] = value
	this.sLen[stackNum]++
}

func (this *TripleInOne) Pop(stackNum int) int {
	return this.top(stackNum, true)
}

func (this *TripleInOne) Peek(stackNum int) int {
	return this.top(stackNum, false)
}

func (this *TripleInOne) top(stackNum int, pop bool) int {
	sLen := this.sLen[stackNum]
	if sLen == 0 {
		return -1
	}
	targetIdx := stackNum + (sLen-1)*3
	val := this.stack[targetIdx]
	if pop {
		this.sLen[stackNum]--
	}
	return val
}
func (this *TripleInOne) IsEmpty(stackNum int) bool {
	return this.sLen[stackNum] == 0
}

/**
 * Your TripleInOne object will be instantiated and called as such:
 * obj := Constructor(stackSize);
 * obj.Push(stackNum,value);
 * param_2 := obj.Pop(stackNum);
 * param_3 := obj.Peek(stackNum);
 * param_4 := obj.IsEmpty(stackNum);
 */
// 示例 1：
// 输入：
// ["TripleInOne", "push", "push", "pop", "pop", "pop", "isEmpty"]
// [[1], [0, 1], [0, 2], [0], [0], [0], [0]]
// 输出：
// [null, null, null, 1, -1, -1, true]
// 说明：当栈为空时`pop, peek`返回-1，当栈满时`push`不压入元素。
//
// 示例 2：
// 输入：
// ["TripleInOne", "push", "push", "push", "pop", "pop", "pop", "peek"]
// [[2], [0, 1], [0, 2], [0, 3], [0], [0], [0], [0]]
// 输出：
// [null, null, null, null, 2, 1, -1, -1]
func main() {
	{
		obj := Constructor(1)
		obj.Push(0, 1)
		obj.Push(0, 2)
		fmt.Println(obj.Pop(0))
		fmt.Println(obj.Pop(0))
		fmt.Println(obj.Pop(0))
		fmt.Println(obj.IsEmpty(0))
	}
	fmt.Println()

	{
		obj := Constructor(2)
		obj.Push(0, 1)
		obj.Push(0, 2)
		obj.Push(0, 3)
		fmt.Println(obj.Pop(0))
		fmt.Println(obj.Pop(0))
		fmt.Println(obj.Pop(0))
		fmt.Println(obj.Peek(0))
	}
}
