package main

import (
	"fmt"
	"math"
)

func minWindowR(s string, t string) string {
	var targetLength = len(t) // 要匹配的字符数量
	var cmp = math.MaxInt
	var start, end int                  // 覆盖子串的左下标、右下标，可以看作是左右指针
	var ans string                      // 答案
	targetMapping := make(map[byte]int) // 目标串字符映射表。key为目标串中的字符，value为该字符出现的次数

	// 初始化目标串字符映射表
	for i := 0; i < targetLength; i++ {
		targetMapping[t[i]]++
	}

	// 移动子串的左下标，直到其字符 s[start] 第一次匹配到目标串中的一个字符
	for start < len(s) {
		if _, ok := targetMapping[s[start]]; !ok {
			start++
		} else {
			break
		}
	}

	end = start // 将左下标赋给右下标，开始遍历主串
	for ; end < len(s); end++ {

		// 找到了一个目标串中的字符
		if _, ok := targetMapping[s[end]]; ok {
			if targetMapping[s[end]] > 0 {
				targetLength--          // 将仍需匹配的字符串减一
				targetMapping[s[end]]-- // 将子串字符映射表中将该字符次数减一
				if targetLength == 0 {  // 当要匹配的字符数量为 0 时，表示 s[start,end] 串是一个覆盖子串。此时，就可以移动左指针，直到不满足覆盖字串的要求

					// 查看当前子串长度是否比上一个子串长度小，如果是，则设置 ans 为当前覆盖子串
					if cmp > end-start+1 {
						cmp = end - start + 1
						ans = s[start : end+1]
					}

					// 移动左指针，直至 s[start,end] 串不再是覆盖字串
					for start < end {

						// 当前左指针指向的字符是目标字符之一
						if _, ok := targetMapping[s[start]]; ok {
							targetMapping[s[start]]++
							start++ // 移动左指针

							// 注意这里下标是 start-1，因为刚刚自增了start
							// 如果 targetMapping[s[start - 1]] > 0，说明此时 s[start,end] 串不再是覆盖字串，targetLength++
							if targetMapping[s[start-1]] > 0 {
								targetLength++
								// 注意此时 end-start+2 要多加 1，因为此时 s[start - 1, end] 才是覆盖子串，而不是 s[start,end]
								if cmp > end-start+2 {
									cmp = end - start + 2
									ans = s[start-1 : end+1]
								}

								// 停止移动左指针，再次移动右指针，直到 s[start,end] 串再次成为覆盖子串
								break
							}

							// 当前左指针指向的字符不是目标字符，直接向右移动
						} else {
							start++
						}

					}
				}

				//  目标字符 s[end] 已经找完了，这里直接将 targetMapping[s[end]] 减一，即为负数，
				//  这表示当前 s[start,end] 串中有冗余的字符 s[end]
			} else {
				targetMapping[s[end]]--
			}
		}
	}

	return ans
}
func main() {
	var s, t string
	fmt.Scan(&s, &t)
	fmt.Println(minWindowR(s, t))
}
