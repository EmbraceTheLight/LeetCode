// 03.02. 栈的最小值
/*
请设计一个栈，除了常规栈支持的pop与push函数以外，还支持min函数，该函数返回栈元素中的最小值。
执行push、pop和min操作的时间复杂度必须为O(1)。

示例：
MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin();   --> 返回 -3.
minStack.pop();
minStack.top();      --> 返回 0.
minStack.getMin();   --> 返回 -2.
*/
package main

import "fmt"

type MinStack struct {
	stack []int
	mins  []int
}

/** initialize your data structure here. */
func Constructor0302() MinStack {
	return MinStack{
		stack: make([]int, 0),
		mins:  make([]int, 0),
	}
}

func (this *MinStack) len() int {
	return len(this.stack)
}
func (this *MinStack) Push(x int) {
	if this.len() == 0 {
		this.stack = append(this.stack, x)
		this.mins = append(this.mins, x)
		return
	}
	if x < this.mins[this.len()-1] {
		this.mins = append(this.mins, x)
	} else {
		this.mins = append(this.mins, this.mins[this.len()-1])
	}
	this.stack = append(this.stack, x)
}

func (this *MinStack) Pop() {
	this.mins = this.mins[:this.len()-1]
	this.stack = this.stack[:this.len()-1]
}

func (this *MinStack) Top() int {
	return this.stack[this.len()-1]
}

func (this *MinStack) GetMin() int {
	return this.mins[this.len()-1]
}

func main() {
	obj := Constructor0302()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)
	fmt.Println(obj.GetMin())
	obj.Pop()
	fmt.Println(obj.Top())
	fmt.Println(obj.GetMin())
}
