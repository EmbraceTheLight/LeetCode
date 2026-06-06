/*
146. LRU 缓存
请你设计并实现一个满足  LRU (最近最少使用) 缓存 约束的数据结构。
实现 LRUCache146 类：
LRUCache146(int capacity) 以 正整数 作为容量 capacity 初始化 LRU 缓存
int Get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
void Put(int key, int value) 如果关键字 key 已经存在，则变更其数据值 value ；如果不存在，则向缓存中插入该组 key-value 。如果插入操作导致关键字数量超过 capacity ，则应该 逐出 最久未使用的关键字。
函数 Get 和 Put 必须以 O(1) 的平均时间复杂度运行。
示例：
输入
["LRUCache146", "Put", "Put", "Get", "Put", "Get", "Put", "Get", "Get", "Get"]
[[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
输出
[null, null, null, 1, null, -1, null, -1, 3, 4]

解释
LRUCache146 lRUCache = new LRUCache146(2);
lRUCache.Put(1, 1); // 缓存是 {1=1}
lRUCache.Put(2, 2); // 缓存是 {1=1, 2=2}
lRUCache.Get(1);    // 返回 1
lRUCache.Put(3, 3); // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
lRUCache.Get(2);    // 返回 -1 (未找到)
lRUCache.Put(4, 4); // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
lRUCache.Get(1);    // 返回 -1 (未找到)
lRUCache.Get(3);    // 返回 3
lRUCache.Get(4);    // 返回 4

提示：

1 <= capacity <= 3000
0 <= key <= 10000
0 <= value <= 10^5
最多调用 2 * 10^5 次 Get 和 Put
*/
package main

import "fmt"

type node struct {
	Key  int
	Val  int
	Pre  *node
	Next *node
}
type LRUCache146 struct {
	mp         map[int]*node
	head, tail *node //伪首部、伪尾部
	Len        int
	Cap        int
}

func Constructor146(capacity int) LRUCache146 {
	var lc = new(LRUCache146)
	lc.mp = make(map[int]*node)
	lc.Len = 0
	lc.Cap = capacity

	lc.head = new(node)
	lc.tail = new(node)

	lc.tail.Pre = lc.head
	lc.head.Next = lc.tail

	return *lc
}

func (lc *LRUCache146) addToHead(n *node) {
	n.Pre = lc.head
	n.Next = lc.head.Next
	lc.head.Next.Pre = n
	lc.head.Next = n
	lc.mp[n.Key] = n
	lc.Len++
}
func (lc *LRUCache146) moveToHead(n *node) {
	lc.removeNode(n)
	lc.addToHead(n)
}
func (lc *LRUCache146) removeNode(n *node) {
	n.Pre.Next = n.Next
	n.Next.Pre = n.Pre
	lc.Len--
}
func (lc *LRUCache146) removeTail() {
	del := lc.tail.Pre
	del.Pre.Next = lc.tail
	delete(lc.mp, del.Key)
	lc.tail.Pre = del.Pre
	lc.Len--
}

func (lc *LRUCache146) Get(key int) int {
	ret, ok := lc.mp[key]
	if !ok {
		return -1
	}

	lc.moveToHead(ret)
	return ret.Val
}

func (lc *LRUCache146) Put(key int, value int) {
	ret, ok := lc.mp[key]
	if ok {
		lc.moveToHead(ret)
		ret.Val = value
	} else {
		ret = new(node)
		ret.Key = key
		ret.Val = value
		if lc.Len == lc.Cap {
			lc.removeTail()
		}
		lc.addToHead(ret)
	}
}

/**
 * Your LRUCache146 object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
func main() {
	var lRUCache = Constructor146(2)
	lRUCache.Put(1, 1)           // 缓存是 {1=1}
	lRUCache.Put(2, 2)           // 缓存是 {1=1, 2=2}
	fmt.Println(lRUCache.Get(1)) // 返回 1
	lRUCache.Put(3, 3)           // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
	fmt.Println(lRUCache.Get(2)) // 返回 -1 (未找到)
	lRUCache.Put(4, 4)           // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
	fmt.Println(lRUCache.Get(1)) // 返回 -1 (未找到)
	fmt.Println(lRUCache.Get(3)) // 返回 3
	fmt.Println(lRUCache.Get(4)) // 返回 4
}
