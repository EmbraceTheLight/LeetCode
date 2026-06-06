package main

import "fmt"

func dfs17R(step int, digits string, letterMap map[byte][]byte, tmp []byte, ans *[]string) {
	if step == len(digits) {
		*ans = append(*ans, string(tmp))
		return
	}
	for _, letter := range letterMap[digits[step]] {
		tmp = append(tmp, letter)
		dfs17R(step+1, digits, letterMap, tmp, ans)
		tmp = tmp[:len(tmp)-1]
	}
}
func letterCombinationsR(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	var ans = make([]string, 0)
	
	var letterMapping = make(map[byte][]byte)
	letterMapping['2'] = []byte{'a', 'b', 'c'}
	letterMapping['3'] = []byte{'d', 'e', 'f'}
	letterMapping['4'] = []byte{'g', 'h', 'i'}
	letterMapping['5'] = []byte{'j', 'k', 'l'}
	letterMapping['6'] = []byte{'m', 'n', 'o'}
	letterMapping['7'] = []byte{'p', 'q', 'r', 's'}
	letterMapping['8'] = []byte{'t', 'u', 'v'}
	letterMapping['9'] = []byte{'w', 'x', 'y', 'z'}
	dfs17R(0, digits, letterMapping, []byte{}, &ans)
	return ans
}
func main() {
	var digits string
	fmt.Scan(&digits)
	fmt.Println(letterCombinationsR(digits))

}
