// 68. 文本左右对齐
/*
给定一个单词数组 words 和一个长度 maxWidth ，重新排版单词，使其成为每行恰好有 maxWidth 个字符，且左右两端对齐的文本。
你应该使用 “贪心算法” 来放置给定的单词；也就是说，尽可能多地往每行中放置单词。必要时可用空格 ' ' 填充，使得每行恰好有 maxWidth 个字符。
要求尽可能均匀分配单词间的空格数量。如果某一行单词间的空格不能均匀分配，则左侧放置的空格数要多于右侧的空格数。
文本的最后一行应为左对齐，且单词之间不插入额外的空格。

注意:
单词是指由非空格字符组成的字符序列。
每个单词的长度大于 0，小于等于 maxWidth。
输入单词数组 words 至少包含一个单词。


示例 1:
输入: words = ["This", "is", "an", "example", "of", "text", "justification."], maxWidth = 16
输出:
[
   "This    is    an",
   "example  of text",
   "justification.  "
]

示例 2:
输入:words = ["What","must","be","acknowledgment","shall","be"], maxWidth = 16
输出:
[
  "What   must   be",
  "acknowledgment  ",
  "shall be        "
]
解释: 注意最后一行的格式应为 "shall be    " 而不是 "shall     be",
     因为最后一行应为左对齐，而不是左右两端对齐。
     第二行同样为左对齐，这是因为这行只包含一个单词。

示例 3:
输入:words = ["Science","is","what","we","understand","well","enough","to","explain","to","a","computer.","Art","is","everything","else","we","do"]，maxWidth = 20
输出:
[
  "Science  is  what we",
  "understand      well",
  "enough to explain to",
  "a  computer.  Art is",
  "everything  else  we",
  "do                  "
]


提示:
1 <= words.length <= 300
1 <= words[i].length <= 20
words[i] 由小写英文字母和符号组成
1 <= maxWidth <= 100
words[i].length <= maxWidth
*/
package main

import (
	"bytes"
	"fmt"
	"lc/pkg"
	"strings"
)

// * 自己推导出来的，一遍过，值得鼓励！
func fullJustify(words []string, maxWidth int) []string {
	n := len(words)
	ans := make([]string, 0)

	lineStrs := make([]string, 0)
	strsSumLength := 0
	for i := 0; i < n; i++ {
		strsSumLength += len(words[i])
		lineStrs = append(lineStrs, words[i])

		// 每个单词之间至少有一个空格，算上空格，若整行的长度大于 maxWidth, 减去最后一个单词，拼接当前行，之后开始下一行
		if strsSumLength+(len(lineStrs)-1) > maxWidth {
			strsSumLength -= len(words[i])
			lineStrs = lineStrs[:len(lineStrs)-1]

			// 拼接单词
			var sb strings.Builder
			whiteSpaceLen := maxWidth - strsSumLength // 除去字符，还剩下的空格数
			if len(lineStrs) == 1 {
				sb.WriteString(lineStrs[0])
				sb.WriteString(strings.Repeat(" ", maxWidth-len(lineStrs[0])))
			} else {
				wsBetweenWord := whiteSpaceLen / (len(lineStrs) - 1)  // 每个单词之间的空格数量
				remainingSpace := whiteSpaceLen % (len(lineStrs) - 1) // 用于处理平均分配之后还剩下的空格
				for j := 0; j < len(lineStrs); j++ {
					sb.WriteString(lineStrs[j])
					if j < len(lineStrs)-1 {
						sb.WriteString(strings.Repeat(" ", wsBetweenWord))
					}
					if remainingSpace > 0 {
						sb.WriteString(" ")
						remainingSpace--
					}
				}
			}

			ans = append(ans, sb.String())

			// 重置
			i-- // 回退到上一个单词
			strsSumLength = 0
			lineStrs = make([]string, 0)
		}
	}

	// 最后一行算上空格，长度肯定小于等于 maxWidth，需要单独处理
	var sb strings.Builder
	for i := 0; i < len(lineStrs); i++ {
		sb.WriteString(lineStrs[i])
		if i < len(lineStrs)-1 {
			sb.WriteString(" ")
		}
	}
	sb.WriteString(strings.Repeat(" ", maxWidth-len(sb.String())))
	ans = append(ans, sb.String())
	return ans
}

// Test Case 1: maxLength: 16	words: ["This", "is", "an", "example", "of", "text", "justification."]
// Output: ["This    is    an", "example  of text", "justification.  "]

// Test Case 2: maxLength: 16	words: ["What","must","be","acknowledgment","shall","be"]
// Output: ["What   must   be", "acknowledgment  ", "shall be        "]

// Test Case 3: maxLength: 20	words: ["Science","is","what","we","understand","well","enough","to","explain","to","a","computer.","Art","is","everything","else","we","do"]
// Output: ["Science  is  what we", "understand      well", "enough to explain to", "a  computer.  Art is", "everything  else  we", "do                  "]
func main() {
	fmt.Println("Input the maxWidth:")
	var maxWidth int
	fmt.Scan(&maxWidth)
	fmt.Println("Input the words:")
	words := pkg.CreateSlice[string]()
	ret := fullJustify(words, maxWidth)
	str := ""
	var buf bytes.Buffer
	for _, s := range ret {
		for i := 0; i < len(s); i++ {
			if s[i] == ' ' {
				buf.WriteByte('-')
				continue
			}
			buf.WriteByte(s[i])
		}
		buf.WriteString(",\n")
	}
	str = buf.String()
	fmt.Println(str)
}
