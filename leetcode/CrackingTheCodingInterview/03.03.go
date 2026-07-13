// 03.03. 堆盘子
/*
堆盘子。设想有一堆盘子，堆太高可能会倒下来。因此，在现实生活中，盘子堆到一定高度时，我们就会另外堆一堆盘子。
请实现数据结构SetOfStacks，模拟这种行为。SetOfStacks应该由多个栈组成，并且在前一个栈填满时新建一个栈。
此外，SetOfStacks.push()和SetOfStacks.pop()应该与普通栈的操作方法相同（也就是说，pop()返回的值，应该跟只有一个栈时的情况一样）。
进阶：实现一个popAt(int index)方法，根据指定的子栈，执行pop操作。
当某个栈为空时，应当删除该栈。当栈中没有元素或不存在该栈时，pop，popAt 应返回 -1.
*/
package main

import "fmt"

type StackOfPlates struct {
	stack      [][]int
	capacity   int
	lenOfStack int
}

func Constructor0303(cap int) StackOfPlates {
	return StackOfPlates{
		stack:      [][]int{},
		capacity:   cap,
		lenOfStack: 0,
	}
}

func (this *StackOfPlates) Push(val int) {
	if this.capacity <= 0 {
		return
	}
	if this.lenOfStack == 0 {
		this.stack = append(this.stack, make([]int, 0, this.capacity))
		this.lenOfStack++
	} else if len(this.stack[this.lenOfStack-1]) == this.capacity {
		this.stack = append(this.stack, make([]int, 0, this.capacity))
		this.lenOfStack++
	}
	this.stack[this.lenOfStack-1] = append(this.stack[this.lenOfStack-1], val)
}

func (this *StackOfPlates) Pop() int {
	if this.lenOfStack == 0 {
		return -1
	}
	var ret int
	this.stack[this.lenOfStack-1], ret = this.pop(this.stack[this.lenOfStack-1])
	this.adjustStack(this.lenOfStack - 1)
	return ret
}

func (this *StackOfPlates) PopAt(index int) int {
	if this.lenOfStack == 0 || index >= this.lenOfStack {
		return -1
	}
	var ret int
	this.stack[index], ret = this.pop(this.stack[index])
	this.adjustStack(index)
	return ret
}

func (this *StackOfPlates) pop(stack []int) ([]int, int) {
	val := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	return stack, val
}

// adjustStack 检查 this.stack[index] 是否是空的, 如果是空的, 则将其从 this.stack 中剔除
func (this *StackOfPlates) adjustStack(index int) {
	if this.lenOfStack == 0 || index >= this.lenOfStack {
		return
	}
	if len(this.stack[index]) == 0 {
		this.stack = append(this.stack[:index], this.stack[index+1:]...)
		this.lenOfStack--
	}
}

// 示例 1：
//
//	输入：
//
// ["StackOfPlates", "push", "push", "popAt", "pop", "pop"]
// [[1], [1], [2], [1], [], []]
//
//	输出：
//
// [null, null, null, 2, 1, -1]
//
// 示例 2：
//
//	输入：
//
// ["StackOfPlates", "push", "push", "push", "popAt", "popAt", "popAt"]
// [[2], [1], [2], [3], [0], [0], [0]]
//
//	输出：
//
// [null, null, null, null, 2, 1, 3]
func main() {
	{
		obj := Constructor0303(0)
		obj.Push(1)
		obj.Push(2)
		fmt.Println(obj.PopAt(1))
		fmt.Println(obj.Pop())
		fmt.Println(obj.Pop())
	}

	fmt.Println()

	{
		obj := Constructor0303(2)
		obj.Push(1)
		obj.Push(2)
		obj.Push(3)
		fmt.Println(obj.PopAt(0))
		fmt.Println(obj.PopAt(0))
		fmt.Println(obj.PopAt(0))
	}

	fmt.Println()

	{
		obj := Constructor0303(6)
		fmt.Println(obj.Pop())
		fmt.Println(obj.Pop())
		fmt.Println(obj.PopAt(1))
		fmt.Println(obj.PopAt(3))
		fmt.Println(obj.Pop())
		obj.Push(40)
		obj.Push(10)
		obj.Push(44)
		obj.Push(44)
		fmt.Println(obj.Pop())
		obj.Push(24)
		obj.Push(42)
		fmt.Println(obj.PopAt(4))
		fmt.Println(obj.Pop())
		fmt.Println(obj.PopAt(0))
		obj.Push(42)
		fmt.Println(obj.PopAt(5))
	}
}
