// 面试题 01.03. URL化
/*
URL化。编写一种方法，将字符串中的空格全部替换为%20。
假定该字符串尾部有足够的空间存放新增字符，并且知道字符串的“真实”长度。（注：用Java实现的话，请使用字符数组实现，以便直接在数组上操作。）

提示：
字符串长度在 [0, 500000] 范围内。
*/
package main

import (
	"fmt"
	. "lc/pkg"
)

func replaceSpaces(S string, length int) string {
	cntSpace := 0
	for i := 0; i < length; i++ {
		if S[i] == ' ' {
			cntSpace++
		}
	}
	b := make([]byte, length+2*cntSpace)
	for i := 0; i < length; i++ {
		b[i] = S[i]
	}
	idx1 := length + 2*cntSpace - 1
	idx2 := length - 1
	for ; idx2 >= 0; idx2-- {
		if b[idx2] == ' ' {
			b[idx1] = '0'
			b[idx1-1] = '2'
			b[idx1-2] = '%'
			idx1 -= 3
		} else {
			b[idx1] = b[idx2]
			idx1--
		}
	}
	return string(b)
}

// 示例 1：
// 输入："Mr John Smith    ", 13
// 输出："Mr%20John%20Smith"
//
// 示例 2：
// 输入："               ", 5
// 输出："%20%20%20%20%20"
func main() {
	var length int
	fmt.Println("Input str length:")
	fmt.Scan(&length)
	fmt.Println("Input string:")
	str := CreateString()
	fmt.Println(replaceSpaces(str, length))
}
