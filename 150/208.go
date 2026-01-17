// 208. 实现 Trie (前缀树)
/*
Trie（发音类似 "try"）或者说 前缀树 是一种树形数据结构，用于高效地存储和检索字符串数据集中的键。这一数据结构有相当多的应用情景，例如自动补全和拼写检查。

请你实现 Trie 类：

Trie() 初始化前缀树对象。
void insert(String word) 向前缀树中插入字符串 word 。
boolean search(String word) 如果字符串 word 在前缀树中，返回 true（即，在检索之前已经插入）；否则，返回 false 。
boolean startsWith(String prefix) 如果之前已经插入的字符串 word 的前缀之一为 prefix ，返回 true ；否则，返回 false 。


示例：

输入
["Trie", "insert", "search", "search", "startsWith", "insert", "search"]
[[], ["apple"], ["apple"], ["app"], ["app"], ["app"], ["app"]]
输出
[null, null, true, false, true, null, true]

解释
Trie trie = new Trie();
trie.insert("apple");
trie.search("apple");   // 返回 True
trie.search("app");     // 返回 False
trie.startsWith("app"); // 返回 True
trie.insert("app");
trie.search("app");     // 返回 True


提示：

1 <= word.length, prefix.length <= 2000
word 和 prefix 仅由小写英文字母组成
insert、search 和 startsWith 调用次数 总计 不超过 3 * 104 次
*/
package main

import "fmt"

type Trie struct {
	root *TrieNode
}

type TrieNode struct {
	isLeaf   bool          // isLeaf 表示 Trie 中是否有以当前节点表示的字母结尾的单词
	children [26]*TrieNode // 对应字母的 node 为 nil, 表示该字母在当前路径中没有存储
}

func Constructor() Trie {
	return Trie{root: &TrieNode{}}
}

func (t *Trie) Insert(word string) {
	curNode := t.root
	for i := 0; i < len(word); i++ {
		if curNode.children[word[i]-'a'] == nil {
			curNode.children[word[i]-'a'] = &TrieNode{}
		}
		curNode = curNode.children[word[i]-'a']
	}
	curNode.isLeaf = true
}

// searchPrefix 查找是否存在 str 为前缀的路径.
// 该方法的作用主要是用来供 Search 和 StartsWith 方法复用的
func (t *Trie) searchPrefix(str string) *TrieNode {
	curNode := t.root
	for i := 0; i < len(str); i++ {
		if curNode.children[str[i]-'a'] == nil {
			return nil
		}
		curNode = curNode.children[str[i]-'a']
	}
	return curNode
}

func (t *Trie) Search(word string) bool {
	node := t.searchPrefix(word)
	return node != nil && node.isLeaf
}

func (t *Trie) StartsWith(prefix string) bool {
	return t.searchPrefix(prefix) != nil
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
func main() {
	// Test Case1:
	{
		obj := Constructor()
		obj.Insert("apple")
		fmt.Println(obj.Search("apple"))   // 返回 True
		fmt.Println(obj.Search("app"))     // 返回 False
		fmt.Println(obj.StartsWith("app")) // 返回 True
		obj.Insert("app")
		fmt.Println(obj.Search("app")) // 返回 True
	}
}
