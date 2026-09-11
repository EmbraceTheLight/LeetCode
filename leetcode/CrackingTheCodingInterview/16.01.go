// 16.01. 交换数字
/*
编写一个函数，不用临时变量，直接交换numbers = [a, b]中a与b的值。

提示：
numbers.length == 2
-2147483647 <= numbers[i] <= 2147483647
*/
package main

import (
	"fmt"
	"lc/pkg"
)

func swapNumbers(numbers []int) []int {
	numbers[1], numbers[0] = numbers[1]^(numbers[0]^numbers[1]), numbers[0]^(numbers[0]^numbers[1])
	return numbers
}

// 示例：
// 输入: numbers = [1,2]
// 输出: [2,1]
func main() {
	fmt.Println(swapNumbers(pkg.CreateSlice[int]()))
}
