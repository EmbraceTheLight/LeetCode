// 17.01. 不用加号的加法
/*
设计一个函数把两个数字相加。不得使用 + 或者其他算术运算符。

提示：
a, b 均可能是负数或 0
结果不会溢出 32 位整数
*/
package main

import "fmt"

func add(a int, b int) int {
	var ans int32
	carry := 0
	for i := 0; i < 32; i++ {
		if (a>>i)&1 == 1 && ((b>>i)&1 == 1 || carry == 1) {
			if (b>>i)&1 == 1 && carry == 1 {
				ans = ans | (1 << i)
			}
			carry = 1
		} else if (b>>i)&1 == 1 && ((a>>i)&1 == 1 || carry == 1) {
			if (a>>i)&1 == 1 && carry == 1 {
				ans = ans | (1 << i)
			}
			carry = 1
		} else {
			if (a>>i)&1 == 1 || (b>>i)&1 == 1 || carry == 1 {
				ans = ans | (1 << i)
			}
			carry = 0
		}
	}
	return int(ans)
}

// 示例：
// 输入：a = 1, b = 1
// 输出：2
func main() {
	var a, b int
	fmt.Println("Input a,b:")
	fmt.Scan(&a, &b)
	fmt.Println(add(a, b))
}
