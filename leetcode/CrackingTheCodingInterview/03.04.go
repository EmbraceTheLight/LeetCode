// 03.04. 化栈为队
/*

 */
package main

import "fmt"

type stack struct {
	nums []int
}

func (this *stack) push(v int) {
	this.nums = append(this.nums, v)
}
func (this *stack) pop() int {
	v := this.nums[len(this.nums)-1]
	this.nums = this.nums[:len(this.nums)-1]
	return v
}
func (this *stack) empty() bool {
	return len(this.nums) == 0
}
func (this *stack) len() int {
	return len(this.nums)
}
func (this *stack) peek() int {
	return this.nums[len(this.nums)-1]
}

type MyQueue struct {
	stkForPush *stack
	stkForPop  *stack
}

func Constructor0304() MyQueue {
	return MyQueue{
		stkForPush: &stack{nums: make([]int, 0)},
		stkForPop:  &stack{nums: make([]int, 0)},
	}
}

func (this *MyQueue) Push(x int) {
	this.stkForPush.push(x)
}

func (this *MyQueue) Pop() int {
	this.preparePop()
	return this.stkForPop.pop()
}

func (this *MyQueue) Peek() int {
	this.preparePop()
	return this.stkForPop.peek()
}

func (this *MyQueue) Empty() bool {
	return this.stkForPush.empty() && this.stkForPop.empty()
}

func (this *MyQueue) preparePop() {
	if this.stkForPop.empty() {
		for this.stkForPush.empty() != true {
			this.stkForPop.push(this.stkForPush.pop())
		}
	}
}

// leetcode submit region end(Prohibit modification and deletion)
func main() {
	queue := Constructor0304()

	queue.Push(1)
	queue.Push(2)
	fmt.Println(queue.Peek()) // 返回 1
	queue.Push(2)
	fmt.Println(queue.Peek())  // 返回 1
	fmt.Println(queue.Pop())   // 返回 1
	fmt.Println(queue.Empty()) // 返回 false
}
