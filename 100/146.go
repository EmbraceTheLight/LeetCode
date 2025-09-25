/*
146.LRU 缓存
请你设计并实现一个满足  LRU (最近最少使用) 缓存 约束的数据结构。
实现 LRUCache146 类：
LRUCache146(int capacity) 以 正整数 作为容量 capacity 初始化 LRU 缓存
int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
void put(int key, int value) 如果关键字 key 已经存在，则变更其数据值 value ；如果不存在，则向缓存中插入该组 key-value 。
如果插入操作导致关键字数量超过 capacity ，则应该 逐出 最久未使用的关键字。
函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。
*/
package main

import "fmt"

type LRUNode struct {
	Key   int
	Value int
	Next  *LRUNode
	Pre   *LRUNode
}
type LRUCache struct {
	Head *LRUNode
	Tail *LRUNode
	Mp   map[int]*LRUNode
	Len  int
	Size int
}

// UpdateLRUCache 更新链表
func (L *LRUCache) UpdateLRUCache(ret *LRUNode) {
	if ret != L.Head {
		pre := ret.Pre
		pre.Next = ret.Next
		ret.Next.Pre = pre
		ret.Next = L.Head
		if ret == L.Tail { //处理查询的是尾部元素的特殊情况
			L.Tail = pre
		}
		L.Head.Pre = ret
		L.Head = ret
		L.Tail.Next = L.Head //重要！更新L.Tail的next指针
	}
}
func Constructor(capacity int) LRUCache {
	return LRUCache{
		Head: nil,
		Tail: nil,
		Mp:   make(map[int]*LRUNode),
		Len:  0,
		Size: capacity,
	}
}

func (L *LRUCache) Get(key int) int {
	ret, ok := L.Mp[key]
	if !ok {
		return -1
	} else {
		L.UpdateLRUCache(ret)
		return ret.Value
	}
}

func (L *LRUCache) Put(key int, value int) {
	if L.Head == nil { //目前LRUCache为空
		head := new(LRUNode)
		head.Key = key
		head.Value = value
		head.Next = head
		head.Pre = head
		L.Head = head
		L.Tail = head
		L.Mp[key] = head
		L.Len = 1
	} else {
		//更新已存在的值
		if v, ok := L.Mp[key]; ok {
			v.Value = value
			//更新其位置
			L.UpdateLRUCache(v)
			return
		}
		if L.Len == L.Size { //队列已满，需要逐出一个元素
			k := L.Tail.Key
			newTail := L.Tail.Pre
			newTail.Next = L.Head
			L.Head.Pre = newTail
			L.Tail = newTail
			delete(L.Mp, k)
			L.Len--
		}

		//正常插入
		newHead := new(LRUNode)
		newHead.Key = key
		newHead.Value = value
		newHead.Next = L.Head
		newHead.Pre = L.Tail
		L.Head.Pre = newHead
		L.Tail.Next = newHead
		L.Head = newHead
		L.Mp[key] = newHead
		L.Len++
	}
}

/**
 * Your LRUCache146 object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
func main() {
	obj := Constructor(7)
	//obj.Put(1, 1)
	//obj.Put(2, 2)
	//fmt.Println(obj.Get(1))
	//obj.Put(3, 3)
	//fmt.Println(obj.Get(2))
	//obj.Put(4, 4)
	//fmt.Println(obj.Get(1))
	//fmt.Println(obj.Get(3))
	//fmt.Println(obj.Get(4))
	obj.Put(6, 11)
	obj.Put(9, 10)
	obj.Put(2, 19)
	fmt.Println(obj.Get(2))
	obj.Put(5, 25)
	obj.Put(9, 22)
	obj.Put(5, 5)
	obj.Put(1, 30)
	obj.Put(9, 12)
	fmt.Println(obj.Get(5))
	fmt.Println(obj.Get(9))
	obj.Put(4, 30)
	obj.Put(9, 3)
	obj.Put(6, 14)
	obj.Put(3, 1)
	obj.Put(10, 11)
	obj.Put(2, 14)
	fmt.Println(obj.Get(5))
	fmt.Println(obj.Get(4))
	obj.Put(11, 4)
	obj.Put(12, 24)
	fmt.Println(obj.Get(6))

}
