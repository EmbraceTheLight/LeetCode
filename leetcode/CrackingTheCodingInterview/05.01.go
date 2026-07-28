// 05.01. 插入
/*
编写一种方法，使 M 对应的二进制数字插入 N 对应的二进制数字的第 i ~ j 位区域，不足之处用 0 补齐。
题目保证从 i 位到 j 位足以容纳 M， 例如： M = 10011，则 i～j 区域至少可容纳 5 位。
*/
package main

import (
	"fmt"
)

func insertBits(N int, M int, i int, j int) int {
	idxN := i
	printBits(N)
	printBits(M)
	var tmp = 0b11111111111111111111111111111111
	for M != 0 {
		if M&1 == 0 {
			b := tmp &^ (1 << idxN) // 将 tmp (11111111111111111111111111111111) 的第 idxN 位置 0
			N = N & b               // 将 N 的第 idxN 位置 0
		} else {
			N = N | (1 << idxN) // 将 N 的第 idxN 位置 1
		}
		printBits(N)
		idxN++
		M = M >> 1
	}
	for idxN <= j {
		b := tmp &^ (1 << idxN)
		N = N & b
		idxN++
	}
	return N
}
func printBits(n int) {
	bits := make([]int, 0)
	for n != 0 {
		bits = append(bits, n&1)
		n = n >> 1
	}
	for i := len(bits) - 1; i >= 0; i-- {
		fmt.Print(bits[i])
	}
	fmt.Println()
}

// 示例 1：
//
//	输入：N = 1024(10000000000), M = 19(10011), i = 2, j = 6
//	输出：N = 1100(10001001100)
//
// 示例 2：
//
//	输入：N = 0, M = 31(11111), i = 0, j = 4
//	输出：N = 31(11111)
func main() {
	var n, m, i, j int
	fmt.Println("Input n,m,i,j:")
	fmt.Scan(&n, &m, &i, &j)
	fmt.Println(insertBits(n, m, i, j))
}
