package main

import (
	"fmt"
	"lc/pkg"
)

func trapR2(height []int) int {
	var ans int
	var left, right int
	var stk []int
	for left = 0; left < len(height) && height[left] == 0; left++ {
	}

	stk = append(stk, left)
	right = left + 1
	for ; right < len(height); right++ {
		if height[right] <= height[stk[len(stk)-1]] {
			stk = append(stk, right)
		} else {
			for len(stk) > 1 && height[right] > height[stk[len(stk)-1]] {
				bottom := stk[len(stk)-1]
				l := stk[len(stk)-2]
				width := right - l - 1
				h := min(height[l], height[right]) - height[bottom]
				ans += width * h
				stk = stk[:len(stk)-1]
			}
			if len(stk) == 1 && height[right] > height[stk[len(stk)-1]] {
				stk = stk[:len(stk)-1]
			}
			stk = append(stk, right)
		}
	}

	return ans
}
func main() {
	heights := pkg.CreateIntSlice[int]()
	fmt.Println(trapR2(heights))
}
