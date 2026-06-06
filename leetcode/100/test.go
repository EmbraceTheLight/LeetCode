package main

import "fmt"

func main() {
	//head := pkg.CreateTree()
	//pkg.PrintTreeByLevel(head)
	var dp = make([][]int, 2*3, 2*5)
	dp[0][0] = 1
	fmt.Printf("%#v", dp)
}
