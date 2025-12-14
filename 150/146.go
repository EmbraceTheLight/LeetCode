// 146. LRU 缓存
/*
请你设计并实现一个满足  LRU (最近最少使用) 缓存 约束的数据结构。
实现 LRUCache 类：
LRUCache(int Capacity) 以 正整数 作为容量 Capacity 初始化 LRU 缓存
int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
void put(int key, int value) 如果关键字 key 已经存在，则变更其数据值 value ；如果不存在，则向缓存中插入该组 key-value 。如果插入操作导致关键字数量超过 Capacity ，则应该 逐出 最久未使用的关键字。
函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。

示例：
输入
["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
[[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
输出
[null, null, null, 1, null, -1, null, -1, 3, 4]

解释
LRUCache lRUCache = new LRUCache(2);
lRUCache.put(1, 1); // 缓存是 {1=1}
lRUCache.put(2, 2); // 缓存是 {1=1, 2=2}
lRUCache.get(1);    // 返回 1
lRUCache.put(3, 3); // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
lRUCache.get(2);    // 返回 -1 (未找到)
lRUCache.put(4, 4); // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
lRUCache.get(1);    // 返回 -1 (未找到)
lRUCache.get(3);    // 返回 3
lRUCache.get(4);    // 返回 4


提示：
1 <= Capacity <= 3000
0 <= key <= 10000
0 <= value <= 10^5
最多调用 2 * 10^5 次 get 和 put
*/
package main

import "fmt"

type LRUCache struct {
	Capacity int
	Head     *Node146
	Tail     *Node146
	Mp       map[int]*Node146
}

type Node146 struct {
	Key  int
	Val  int
	Prev *Node146
	Next *Node146
}

func Constructor146(capacity int) *LRUCache {
	cache := &LRUCache{
		Capacity: capacity,
		Mp:       make(map[int]*Node146),
	}
	head := &Node146{}
	tail := &Node146{}
	head.Next = tail
	head.Prev = tail
	tail.Next = head
	tail.Prev = head
	cache.Head = head
	cache.Tail = tail
	return cache
}

func (lc *LRUCache) Get(key int) int {
	node, ok := lc.Mp[key]
	if !ok {
		return -1
	}
	lc.moveToHead(node)
	return node.Val
}

func (lc *LRUCache) Put(key int, value int) {
	node, ok := lc.Mp[key]
	if ok {
		node.Val = value
		delete(lc.Mp, key)
		lc.moveToHead(node)
		lc.Mp[key] = node
		return
	}

	newNode := &Node146{Key: key, Val: value}
	if len(lc.Mp) >= lc.Capacity {
		lc.RemoveLastNode()
	}
	lc.insert(newNode)
}

func (lc *LRUCache) moveToHead(node *Node146) {
	prev, next := node.Prev, node.Next

	prev.Next = next
	next.Prev = prev

	node.Next = lc.Head.Next
	node.Prev = lc.Head
	lc.Head.Next.Prev = node
	lc.Head.Next = node
}

// insert 将新节点插入到头节点之后
func (lc *LRUCache) insert(node *Node146) {
	node.Next = lc.Head.Next
	node.Prev = lc.Head
	lc.Head.Next.Prev = node
	lc.Head.Next = node
	lc.Mp[node.Key] = node
}

func (lc *LRUCache) RemoveLastNode() {
	lastNode := lc.Tail.Prev
	delete(lc.Mp, lastNode.Key)
	lastNode.Prev.Next = lc.Tail
	lc.Tail.Prev = lastNode.Prev
}
func main() {
	{
		lRUCache := Constructor146(2)
		lRUCache.Put(1, 1)           // 缓存是 {1=1}
		lRUCache.Put(2, 2)           // 缓存是 {1=1, 2=2}
		fmt.Println(lRUCache.Get(1)) // 返回 1
		lRUCache.Put(3, 3)           // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
		fmt.Println(lRUCache.Get(2)) // 返回 -1 (未找到)
		lRUCache.Put(4, 4)           // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
		fmt.Println(lRUCache.Get(1)) // 返回 -1 (未找到)
		fmt.Println(lRUCache.Get(3)) // 返回 3
		fmt.Println(lRUCache.Get(4)) // 返回 4
		fmt.Println()
	}

	{
		lRUCache := Constructor146(1)
		lRUCache.Put(2, 1)           // 缓存是 {2=1}
		fmt.Println(lRUCache.Get(2)) // 返回 1
		lRUCache.Put(3, 2)           // 缓存是 {3=2}
		fmt.Println(lRUCache.Get(2)) // 返回 -1 (未找到)
		fmt.Println(lRUCache.Get(3)) // 返回 2
	}

}
