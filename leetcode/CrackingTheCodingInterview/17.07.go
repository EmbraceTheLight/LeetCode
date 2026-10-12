// 17.07. 婴儿名字
/*
每年，政府都会公布一万个最常见的婴儿名字和它们出现的频率，也就是同名婴儿的数量。
有些名字有多种拼法，例如，John 和 Jon 本质上是相同的名字，但被当成了两个名字公布出来。
给定两个列表，一个是名字及对应的频率，另一个是本质相同的名字对。
设计一个算法打印出每个真实名字的实际频率。注意，如果 John 和 Jon 是相同的，并且 Jon 和 Johnny 相同，则 John 与 Johnny 也相同，即它们有传递和对称性。

在结果列表中，选择 字典序最小 的名字作为真实名字。

提示：
names.length <= 100000
*/
package main

import (
	"fmt"
	. "lc/pkg"
	"strconv"
	"strings"
)

type unionMap1707 struct {
	um map[string]string
}

func trulyMostPopular(names []string, synonyms []string) []string {
	nameMap := make(map[string]string) // key: 名称, 包含同义词  value: 真实名称
	nameCount := make(map[string]int)  // key: 真实名称         value: 频率

	set := make(map[string]struct{})
	for _, name := range synonyms {
		tmp := getNameFromSynonyms(name)
		set[tmp[0]] = struct{}{}
		set[tmp[1]] = struct{}{}
	}
	um := unionMap1707{
		um: make(map[string]string),
	}
	um.init(set)

	for _, name := range synonyms {
		tmp := getNameFromSynonyms(name)
		name1, name2 := tmp[0], tmp[1]
		um.join(name1, name2)
	}

	for n := range set {
		nameMap[n] = um.find(n)
	}

	for _, str := range names {
		// 有些名称可能不在并查集中. 将其作为根为自身的普通节点
		name, count := getNameAndCount(str)
		if _, ok := nameMap[name]; !ok {
			nameMap[name] = name
		}
		nameCount[nameMap[name]] += count
	}

	var ans []string
	for name, count := range nameCount {
		ans = append(ans, formatNameCount(name, count))
	}
	return ans
}

func (um *unionMap1707) init(nameSet map[string]struct{}) {
	for n := range nameSet {
		um.um[n] = n
	}
}
func (um *unionMap1707) find(name string) string {
	if um.um[name] != name {
		um.um[name] = um.find(um.um[name])
	}
	return um.um[name]
}

func (um *unionMap1707) join(name1, name2 string) {
	root1 := um.find(name1)
	root2 := um.find(name2)
	if root1 != root2 {
		if root1 < root2 {
			um.um[root2] = root1
		} else {
			um.um[root1] = root2
		}
	}
}
func getNameFromSynonyms(synonymsStr string) []string {
	str := strings.Split(synonymsStr, ",")
	return []string{strings.TrimPrefix(str[0], "("), strings.TrimSuffix(str[1], ")")}
}

func getNameAndCount(str string) (string, int) {
	tmp := strings.Split(str, "(")
	num, _ := strconv.Atoi(strings.TrimSuffix(tmp[1], ")"))
	return tmp[0], num
}

func formatNameCount(name string, count int) string {
	return name + "(" + strconv.Itoa(count) + ")"
}

// 示例：
// 输入：names = ["John(15)","Jon(12)","Chris(13)","Kris(4)","Christopher(19)"], synonyms = ["(Jon,John)","(John,Johnny)","(Chris,Kris)","(Chris,Christopher)"]
// 输出：["John(27)","Chris(36)"]
func main() {
	fmt.Println("Input names:")
	names := CreateSlice[string]()
	fmt.Println("Input synonyms:")
	synonyms := CreateSlice[string]()
	fmt.Println(trulyMostPopular(names, synonyms))
}
