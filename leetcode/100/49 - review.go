package main

import (
	"fmt"
	"sort"
)

func groupAnagramsR(strs []string) [][]string {
	var ans [][]string
	words := make(map[string][]string)
	for _, str := range strs {
		strSlice := []byte(str)
		sort.Slice(strSlice, func(i, j int) bool {
			return strSlice[i] < strSlice[j]
		})
		words[string(strSlice)] = append(words[string(strSlice)], str)
	}

	for _, v := range words {
		ans = append(ans, v)
	}
	return ans
}
func main() {
	var strs []string
	var tmp string
	fmt.Scan(&tmp)
	for tmp != "-1" {
		strs = append(strs, tmp)
		fmt.Scan(&tmp)
	}
	fmt.Println(groupAnagramsR(strs))
}
