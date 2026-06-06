package main

import "fmt"

// 双指针。关键在于如何移动左指针：
// ! 移动左指针直到子串中没有重复元素即可
func lengthOfLongestSubstringR(s string) int {
	if len(s) == 0 || len(s) == 1 {
		return len(s)
	}

	var ans = 0
	var left, right int
	byteMapping := make(map[byte]int)

	for right < len(s) {
		byteMapping[s[right]]++

		if byteMapping[s[right]] > 1 {
			ans = max(ans, right-left)

			// byteMapping[s[right]] > 1 说明当前 s[left,right] 子串中存在重复字符 s[right]
			for byteMapping[s[right]] > 1 {
				byteMapping[s[left]]--
				left++
			}
		}
		right++
	}

	// 这一步是为了防止 字符串中没有重复元素，导致ans = max(ans, right-left)没有执行到
	ans = max(ans, right-left)
	return ans
}

func main() {
	var s string
	fmt.Scan(&s)
	fmt.Println(lengthOfLongestSubstringR(s))
}
