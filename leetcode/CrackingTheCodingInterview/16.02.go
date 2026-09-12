// 16.02. 单词频率
/*
设计一个方法，找出任意指定单词在一本书中的出现频率。
你的实现应该支持如下操作：
WordsFrequency(book)构造函数，参数为字符串数组构成的一本书
get(word)查询指定单词在书中出现的频率

提示：
book[i]中只包含小写字母
1 <= book.length <= 100000
1 <= book[i].length <= 10
get函数的调用次数不会超过100000
*/
package main

import "fmt"

type WordsFrequency struct {
	root *treeNode
}
type treeNode struct {
	count int
	node  [26]*treeNode
}

func Constructor1602(book []string) WordsFrequency {
	root := &treeNode{}
	for _, v := range book {
		tmp := root
		for i := range v {
			if tmp.node[v[i]-'a'] == nil {
				tmp.node[v[i]-'a'] = &treeNode{}
			}
			tmp = tmp.node[v[i]-'a']
		}
		tmp.count++
	}
	return WordsFrequency{root: root}
}

func (this *WordsFrequency) Get(word string) int {
	tmp := this.root
	for i := range word {
		if tmp.node[word[i]-'a'] == nil {
			return 0
		}
		tmp = tmp.node[word[i]-'a']
	}
	return tmp.count
}

// 示例：
// WordsFrequency wordsFrequency = new WordsFrequency({"i", "have", "an", "apple", "he", "have", "a", "pen"});
// wordsFrequency.get("you"); //返回0，"you"没有出现过
// wordsFrequency.get("have"); //返回2，"have"出现2次
// wordsFrequency.get("an"); //返回1
// wordsFrequency.get("apple"); //返回1
// wordsFrequency.get("pen"); //返回1
func main() {
	{
		wordsFrequency := Constructor1602([]string{"i", "have", "an", "apple", "he", "have", "a", "pen"})

		fmt.Println(wordsFrequency.Get("you"))   //返回0，"you"没有出现过
		fmt.Println(wordsFrequency.Get("have"))  //返回2，"have"出现2次
		fmt.Println(wordsFrequency.Get("an"))    //返回1
		fmt.Println(wordsFrequency.Get("apple")) //返回1
		fmt.Println(wordsFrequency.Get("pen"))   //返回1
	}
}
