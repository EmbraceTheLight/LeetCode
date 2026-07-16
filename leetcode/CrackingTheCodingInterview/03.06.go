// 面试题 03.06. 动物收容所
/*
动物收容所。有家动物收容所只收容狗与猫，且严格遵守“先进先出”的原则。
在收养该收容所的动物时，收养人只能收养所有动物中“最老”（由其进入收容所的时间长短而定）的动物，或者可以挑选猫或狗（同时必须收养此类动物中“最老”的）。
换言之，收养人不能自由挑选想收养的对象。请创建适用于这个系统的数据结构，实现各种操作方法，比如enqueue、dequeueAny、dequeueDog和dequeueCat。允许使用Java内置的LinkedList数据结构。

enqueue方法有一个animal参数，animal[0]代表动物编号，animal[1]代表动物种类，其中 0 代表猫，1 代表狗。
dequeue*方法返回一个列表[动物编号, 动物种类]，若没有可以收养的动物，则返回[-1,-1]。

说明:

收纳所的最大容量为20000
*/
package main

import "fmt"

type animalObj struct {
	id      int
	typ     int
	version int // version 越大, animal 加入的时间越晚
}
type AnimalShelf struct {
	dogQueue []*animalObj
	catQueue []*animalObj
	version  int // version 越大, animal 加入的时间越晚
}

func Constructor0307() AnimalShelf {
	return AnimalShelf{
		dogQueue: make([]*animalObj, 0),
		catQueue: make([]*animalObj, 0),
		version:  1,
	}
}

func (this *AnimalShelf) Enqueue(animal []int) {
	a := &animalObj{
		id:      animal[0],
		typ:     animal[1],
		version: this.version,
	}
	this.version++
	if a.typ == 0 {
		this.catQueue = append(this.catQueue, a)
	} else {
		this.dogQueue = append(this.dogQueue, a)
	}
}

func (this *AnimalShelf) DequeueAny() []int {
	if len(this.catQueue) == 0 && len(this.dogQueue) == 0 {
		return []int{-1, -1}
	}
	if len(this.catQueue) == 0 {
		return this.DequeueDog()
	}
	if len(this.dogQueue) == 0 {
		return this.DequeueCat()
	}
	cat := this.catQueue[0]
	dog := this.dogQueue[0]
	if cat.version < dog.version {
		return this.DequeueCat()
	}
	return this.DequeueDog()
}

func (this *AnimalShelf) DequeueDog() []int {
	if len(this.dogQueue) == 0 {
		return []int{-1, -1}
	}
	ret := this.dogQueue[0]
	this.dogQueue = this.dogQueue[1:]
	return []int{ret.id, ret.typ}
}

func (this *AnimalShelf) DequeueCat() []int {
	if len(this.catQueue) == 0 {
		return []int{-1, -1}
	}
	ret := this.catQueue[0]
	this.catQueue = this.catQueue[1:]
	return []int{ret.id, ret.typ}
}

// 示例 1：
//
//	输入：
//
// ["AnimalShelf", "enqueue", "enqueue", "dequeueCat", "dequeueDog", "dequeueAny"]
// [[], [[0, 0]], [[1, 0]], [], [], []]
//
//	输出：
//
// [null,null,null,[0,0],[-1,-1],[1,0]]
// 示例 2：
//
//	输入：
//
// ["AnimalShelf", "enqueue", "enqueue", "enqueue", "dequeueDog", "dequeueCat", "dequeueAny"]
// [[], [[0, 0]], [[1, 0]], [[2, 1]], [], [], []]
//
//	输出：
//
// [null,null,null,null,[2,1],[0,0],[1,0]]
func main() {
	{
		queue := Constructor0307()
		queue.Enqueue([]int{0, 0})
		queue.Enqueue([]int{1, 0})
		fmt.Println(queue.DequeueCat()) // [0,0]
		fmt.Println(queue.DequeueDog()) // [-1,-1]
		fmt.Println(queue.DequeueAny()) // [1,0]
	}

	{
		queue := Constructor0307()
		queue.Enqueue([]int{0, 0})
		queue.Enqueue([]int{1, 0})
		queue.Enqueue([]int{2, 1})
		fmt.Println(queue.DequeueDog()) // [2,1]
		fmt.Println(queue.DequeueCat()) // [0,0]
		fmt.Println(queue.DequeueAny()) // [1,0]
	}
}
