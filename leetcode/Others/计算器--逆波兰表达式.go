package main

import (
	"bufio"
	"fmt"
	"os"
)

var mpOp = make(map[byte]int)

func init() {
	mpOp['('] = 5
	mpOp[')'] = 0
	mpOp['+'] = 1
	mpOp['-'] = 1
	mpOp['*'] = 2
	mpOp['/'] = 2
}

func calculate(s string) int {
	var idx = 0
	var stackNums = make([]int, 0)
	var stackOpts = make([]byte, 0)
	var ans int

	//预处理，向数字栈中压入0,防止第一个为符号‘-’的情况
	stackNums = append(stackNums, 0)
	for ; idx < len(s); idx++ {
		if s[idx] == ' ' {
			continue
		} else if s[idx] <= '9' && s[idx] >= '0' { //读到数字
			var num int
			for idx < len(s) && (s[idx] <= '9' && s[idx] >= '0') {
				num = num*10 + int(s[idx]-'0')
				idx++
			}
			idx-- //索引减一，否则会跳过下一个字符
			//将数字压入数字栈中
			stackNums = append(stackNums, num)
		} else { //读到符号
			if s[idx] == '(' { //匹配的是左括号，直接入栈
				stackOpts = append(stackOpts, s[idx])
				//清除左括号后面的空格
				for ; idx < len(s); idx++ {
					if s[idx+1] != ' ' {
						break
					}
				}
				if idx < len(s)-1 && s[idx+1] == '-' {
					stackNums = append(stackNums, 0) //若左括号后边一个字符即为负号，向数字栈中添加0
				}
			} else if s[idx] == ')' { //匹配到右括号，则直到匹配到左括号才结束本次运算，并将结果放入数字栈中，即将括号内容展开并计算
				helperCalc(&stackNums, &stackOpts, ')') //计算值直到遇到左括号或符号栈为空或运算符优先级高于上一个运算符
			} else { //匹配到的是运算符，则比较当前运算符与栈顶运算符的优先级，若优先级低于或等于栈顶运算符，则将栈中运算符弹出直到清空栈或遇到优先级更低的运算符,否则将运算符压入栈中
				var cmp byte
				if len(stackOpts) > 0 {
					cmp = stackOpts[len(stackOpts)-1]
				}
				if mpOp[s[idx]] <= mpOp[cmp] { //当前符号优先级<=运算符栈栈顶运算符优先级
					helperCalc(&stackNums, &stackOpts, s[idx]) //计算值直到遇到左括号或符号栈为空或运算符优先级高于上一个运算符
				}
				stackOpts = append(stackOpts, s[idx]) //将该运算符压栈
			}
		}
	}

	//计算将括号完全展开后的式子
	for len(stackOpts) > 0 {
		op := stackOpts[len(stackOpts)-1]
		num1 := stackNums[len(stackNums)-2]
		num2 := stackNums[len(stackNums)-1]
		res := calcNums(num1, num2, op)
		stackOpts = stackOpts[:len(stackOpts)-1]
		stackNums = stackNums[:len(stackNums)-2]
		stackNums = append(stackNums, res)
	}
	ans = stackNums[len(stackNums)-1] //运算全部完毕，取数字栈最后一个元素了，即为答案
	return ans
}
func helperCalc(stackNums *[]int, stackOpts *[]byte, curOp byte) (result int) {
	var pre byte //当前栈顶运算符
	pre = (*stackOpts)[len(*stackOpts)-1]

	for mpOp[curOp] <= mpOp[pre] { //匹配到左括号或符号栈为空或当前运算符优先级大于栈顶运算符为止
		if pre == '(' { //遇到左括号
			if curOp == ')' { //若当前运算符为右括号，清除它并退出循环
				*stackOpts = (*stackOpts)[:len(*stackOpts)-1]
			}
			break
		}
		if len(*stackOpts) == 0 {
			break
		}
		*stackOpts = (*stackOpts)[:len(*stackOpts)-1]
		num1 := (*stackNums)[len(*stackNums)-2] //注意操作数顺序
		num2 := (*stackNums)[len(*stackNums)-1]
		*stackNums = (*stackNums)[:len(*stackNums)-2]
		tmp := calcNums(num1, num2, pre)
		result += tmp
		*stackNums = append(*stackNums, tmp)
		if len(*stackOpts) > 0 {
			pre = (*stackOpts)[len(*stackOpts)-1]
		}
	}
	return
}

func calcNums(num1, num2 int, op byte) (ret int) {
	switch op {
	case '+':
		ret = num1 + num2
	case '-':
		ret = num1 - num2
	case '*':
		ret = num1 * num2
	case '/':
		ret = num1 / num2
	}
	return
}
func main() {
	fmt.Printf("Input s:")
	reader := bufio.NewReader(os.Stdin)
	result, _, _ := reader.ReadLine()
	fmt.Println(calculate(string(result)))
}
