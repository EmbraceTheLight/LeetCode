package main

import (
	"fmt"
)

func longestPalindromeR(s string) string {
	ans := string(s[0])
	var length int
	var cmp = 1
	for i := 0; i < len(s); i++ {
		length = 0

		cur, left, right := i, i, i

		for right < len(s) && s[right] == s[cur] {
			length++
			right++
		}

		for left >= 0 && s[cur] == s[left] {
			length++
			left--
		}
		length--

		for left >= 0 && right < len(s) {
			if s[left] == s[right] {
				length += 2
			} else {
				break
			}
			left--
			right++
		}

		if length > cmp {
			cmp = length
			start, end := left+1, right-1

			ans = s[start : end+1]
		}
	}
	return ans
}
func main() {
	var s string
	fmt.Scan(&s)
	fmt.Println(longestPalindromeR(s))
}
