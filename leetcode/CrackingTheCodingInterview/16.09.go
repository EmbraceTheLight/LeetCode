// 16.09. 运算
/*
请实现整数数字的乘法、减法和除法运算，运算结果均为整数数字，程序中只允许使用加法运算符和逻辑运算符，允许程序中出现正负常数，不允许使用位运算。
你的实现应该支持如下操作：
Operations() 构造函数
minus(a, b) 减法，返回a - b
multiply(a, b) 乘法，返回a * b
divide(a, b) 除法，返回a / b

提示：
你可以假设函数输入一定是有效的，例如不会出现除法分母为0的情况
单个用例的函数调用次数不会超过1000次
*/
package main

import (
	"fmt"
)

type Operations struct {
	minusOne int // 代替减一操作：i + minusOne 等价于 i - 1，避免直接使用减法运算符
}

func Constructor1609() Operations {
	return Operations{
		minusOne: -1,
	}
}

func (this *Operations) Minus(a int, b int) int {
	// a-b 等价于 a+(-b)，因此减法的关键是先得到 b 的相反数。
	return a + this.negate(b)
}
func (this *Operations) negate(num int) int {
	if num == 0 {
		return 0
	}
	var ans int // 累加与 num 符号相反的增量，最终得到 num 的相反数
	step := 1   // 每次尝试的增量；方向与 num 相反，并按 1、2、4、8……倍增
	if num > 0 {
		step = -1
	}
	steps := make([]int, 0) // 保存所有未越过 0 的倍增量，例如 13 对应 -1、-2、-4、-8
	steps = append(steps, step)
	// 先生成可使用的倍增量，下一步会越过 0 时停止。
	for {
		step = step + step
		if (num < 0 && num+step > 0) || (num > 0 && num+step < 0) {
			break
		}
		steps = append(steps, step)
	}
	// 从最大的增量开始选择，使 num 尽快靠近 0；选中的增量之和就是 -num。
	for i := len(steps) + this.minusOne; i >= 0 && num != 0; i = i + this.minusOne {
		if (num > 0 && num+steps[i] < 0) || (num < 0 && num+steps[i] > 0) {
			continue
		}
		ans += steps[i]
		num += steps[i]
	}
	return ans
}

func (this *Operations) half(num int) int {
	if num == 0 {
		return 0
	}
	var ans int // num 每减少 2*step，ans 就增加 step，因此最终 ans 等于 num/2
	step := 1   // 与 num 同号，表示当前可以加入答案的倍增量
	if num < 0 {
		step = -1
	}
	steps := make([]int, 0) // 保存 1、2、4、8……或对应的负数，用于从大到小计算一半
	steps = append(steps, step)
	// 若答案增加 tmpStep，原数就需要扣除 2*tmpStep；提前生成所有不会越过 0 的 tmpStep。
	for {
		tmpStep := step + step
		if (num < 0 && this.Minus(num, tmpStep+tmpStep) > 0) || (num > 0 && this.Minus(num, tmpStep+tmpStep) < 0) {
			break
		}
		step = tmpStep
		steps = append(steps, step)

	}
	// 贪心选取最大的 step：能从 num 中扣除 2*step，就把 step 计入 ans。
	for i := len(steps) + this.minusOne; i >= 0; {
		if (num < 0 && this.Minus(num, steps[i]+steps[i]) > 0) || (num > 0 && this.Minus(num, steps[i]+steps[i]) < 0) {
			i = i + this.minusOne
			continue
		}
		num = this.Minus(num, steps[i]+steps[i])
		ans += steps[i]
		if num == 0 {
			break
		}
	}
	return ans
}
func (this *Operations) Multiply(a int, b int) int {
	if a > b {
		return this.multiply(a, b)
	}
	return this.multiply(b, a)
}

func (this *Operations) multiply(big, small int) int {
	var flag bool // flag 决定最后是否再进行一次运算. 与 base 配合 示例: big = 3, small = 2, 则 ans = 3 + 3 = 6, flag = false; big = 7, small = 6, 则 ans = (14 + 14) + 14(base) = 42, flag = true
	var base int  // 当前 ans 的值. 用于决定最后是否需要再加一次, 与 flag 配合
	if big == 0 || small == 0 {
		return 0
	}
	if big < 0 {
		// 两个操作数都为负数时，同时取反，并交换位置以继续保持 big、small 的含义。
		tmpBig, tmpSmall := this.negate(big), this.negate(small)
		small, big = tmpBig, tmpSmall
	}
	var needNegate bool // 两个操作数异号时，先按正数相乘，最后再将结果取反
	if big > 0 && small < 0 {
		needNegate = true
		small = this.negate(small)
		if big < small {
			big, small = small, big
		}
	}
	var ans int = big
	for small > 1 {
		tmpSmall := this.half(small)
		if tmpSmall+tmpSmall != small {
			base += ans
			flag = true
		}
		ans += ans
		small = tmpSmall
	}
	if flag == true {
		ans = ans + base
	}
	if needNegate {
		return this.negate(ans)
	}
	return ans
}

func (this *Operations) Divide(a int, b int) int {
	var ans int   // 记录每次成功扣除的 b 的倍数之和，即最终的商
	var flag bool // true 表示 a、b 同号，最终商为非负数；false 表示最终商为负数
	if (a < 0 && b < 0) || (a > 0 && b > 0) {
		flag = true
	} else {
		flag = false
	}

	// 先统一转换为正数处理，最后再根据 flag 恢复商的符号。
	if a < 0 {
		a = this.negate(a)
	}
	if b < 0 {
		b = this.negate(b)
	}
	// bSlice 与 multiples 一一对应：bSlice[i] 等于原始除数的 multiples[i] 倍。
	// 例如除数为 3 时，两者分别为 [3,6,12,24] 和 [1,2,4,8]。
	bSlice := make([]int, 0)
	bSlice = append(bSlice, b)

	multiples := make([]int, 0) // 保存 bSlice 中每个数对应的是原始除数的多少倍
	multiple := 1
	multiples = append(multiples, multiple)
	// 除数和对应倍数同步翻倍，直到下一份倍增后的除数大于被除数 a。
	for {
		b += b
		multiple += multiple
		if this.Minus(a, b) < 0 {
			break
		}
		multiples = append(multiples, multiple)
		bSlice = append(bSlice, b)
	}
	// 从最大的倍增除数开始尝试扣除；每成功扣除一次，就累加它对应的倍数。
	for i := len(bSlice) + this.minusOne; i >= 0; i = i + this.minusOne {
		res := this.Minus(a, bSlice[i])
		if res < 0 {
			continue
		}
		a = res
		ans += multiples[i]
	}
	if flag == false {
		return this.negate(ans)
	}
	return ans
}

/**
 * Your Operations object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Minus(a,b);
 * param_2 := obj.Multiply(a,b);
 * param_3 := obj.Divide(a,b);
 */
// 示例：
// Operations operations = new Operations();
// operations.minus(1, 2); //返回-1
// operations.multiply(3, 4); //返回12
// operations.divide(5, -2); //返回-2
func main() {
	obj := Constructor1609()
	fmt.Println(obj.Divide(-2147483648, 1))
}
