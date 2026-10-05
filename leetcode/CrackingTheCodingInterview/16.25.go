// 16.25. LRU 缓存
/*
设计和构建一个“最近最少使用”缓存，该缓存会删除最近最少使用的项目。
缓存应该从键映射到值(允许你插入和检索特定键对应的值)，并在初始化时指定最大容量。当缓存被填满时，它应该删除最近最少使用的项目。
它应该支持以下操作： 获取数据 get 和 写入数据 put 。
获取数据 get(key) - 如果密钥 (key) 存在于缓存中，则获取密钥的值（总是正数），否则返回 -1。
写入数据 put(key, value) - 如果密钥不存在，则写入其数据值。当缓存容量达到上限时，它应该在写入新数据之前删除最近最少使用的数据值，从而为新的数据值留出空间。
*/
package main

import "fmt"

type list1625 struct {
	head *node1625
	tail *node1625
}

type node1625 struct {
	key      int
	value    int
	previous *node1625
	next     *node1625
}
type LRUCache1625 struct {
	cap  int
	mp   map[int]*node1625
	list *list1625
}

func Constructor1625(capacity int) LRUCache1625 {
	ret := LRUCache1625{
		cap: capacity,
		mp:  make(map[int]*node1625),
		list: &list1625{
			head: &node1625{},
			tail: &node1625{},
		},
	}
	ret.list.head.next = ret.list.tail
	ret.list.tail.previous = ret.list.head
	return ret
}

func (this *LRUCache1625) Get(key int) int {
	if _, ok := this.mp[key]; !ok {
		return -1
	}
	node := this.mp[key]
	this.moveToHead(node)
	return node.value
}

func (this *LRUCache1625) Put(key int, value int) {
	if val := this.Get(key); val != -1 {
		this.mp[key].value = value
		return
	}
	node := &node1625{key: key, value: value}
	this.mp[key] = node
	this.moveToHead(node)
	if len(this.mp) > this.cap {
		this.deleteTail()
	}
}
func (this *LRUCache1625) moveToHead(node *node1625) {
	previous, next := node.previous, node.next
	if previous != nil { // node 不是新插入的节点
		previous.next = next
		next.previous = previous
	}
	previous = this.list.head
	next = this.list.head.next
	this.list.head.next.previous = node
	this.list.head.next = node
	node.previous = previous
	node.next = next
}
func (this *LRUCache1625) deleteTail() {
	deletedNode := this.list.tail.previous
	previous := this.list.tail.previous.previous
	previous.next = this.list.tail
	this.list.tail.previous = previous

	deletedNode.next = nil
	deletedNode.previous = nil
	delete(this.mp, deletedNode.key)
}

// 示例：
// LRUCache cache = new LRUCache(2); // 缓存容量
// cache.put(1, 1);
// cache.put(2, 2);
// cache.get(1);       // 返回  1
// cache.put(3, 3);    // 该操作会使得密钥 2 作废
// cache.get(2);       // 返回 -1 (未找到)
// cache.put(4, 4);    // 该操作会使得密钥 1 作废
// cache.get(1);       // 返回 -1 (未找到)
// cache.get(3);       // 返回  3
// cache.get(4);       // 返回  4
func main() {
	cache := Constructor1625(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Println(cache.Get(1)) // 返回  1
	cache.Put(3, 3)           // 该操作会使得密钥 2 作废
	fmt.Println(cache.Get(2)) // 返回 -1 (未找到)
	cache.Put(4, 4)           // 该操作会使得密钥 1 作废
	fmt.Println(cache.Get(1)) // 返回 -1 (未找到)
	fmt.Println(cache.Get(3)) // 返回  3
	fmt.Println(cache.Get(4)) // 返回  4
}
