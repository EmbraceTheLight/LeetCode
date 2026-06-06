package main

import "fmt"

func min2(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func mergeAlternately(word1 string, word2 string) string {
	var ret []uint8
	m := min2(len(word1), len(word2))
	var i int = 0
	for ; i < m; i++ {
		ret = append(ret, word1[i])
		ret = append(ret, word2[i])
	}
	if m == len(word1) {
		n := len(word2)
		for ; i < n; i++ {
			ret = append(ret, word2[i])
		}
	} else {
		n := len(word1)
		for ; i < n; i++ {
			ret = append(ret, word1[i])
		}
	}
	return string(ret)
}
func main() {
	var word1 string
	var word2 string
	fmt.Scan(&word1)
	fmt.Scan(&word2)
	fmt.Println(mergeAlternately(word1, word2))
}
