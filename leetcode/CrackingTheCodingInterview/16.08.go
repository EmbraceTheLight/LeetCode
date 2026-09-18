// 16.08. 整数的英语表示
/*
给定一个整数，打印该整数的英文描述。
注意：本题与 273 题相同：https://leetcode.cn/problems/integer-to-english-words/
*/
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}
	var sb strings.Builder

	numMap := initNumMap()
	bitOfNum := getBitOfNum(num)
	numStr := strconv.Itoa(num)
	segmentSum := 0

	bit, bitIn3Num := bitOfNum, bitOfNum%3
	if bit >= 3 && bitIn3Num == 0 {
		bitIn3Num = 3
	}
	for i := 0; i < len(numStr); {
		var num1, num2 int
		num1 = int(numStr[i] - '0')
		if bit > 1 {
			num2 = int(numStr[i+1] - '0')
		}
		segmentSum += num1
		str, skip := getNumStr(num1, num2, bit, bitIn3Num, segmentSum, numMap)
		sb.WriteString(str)
		if skip == true {
			i += 2
			bitIn3Num -= 2
			bit -= 2
			if bitIn3Num == 0 {
				bitIn3Num = 3
				segmentSum = 0
			}
		} else {
			i += 1
			bitIn3Num -= 1
			bit -= 1
			if bitIn3Num == 0 {
				bitIn3Num = 3
				segmentSum = 0
			}
		}

	}
	return strings.TrimSpace(sb.String())
}

func getBitOfNum(num int) int {
	count := 0
	for num != 0 {
		count++
		num = num / 10
	}
	return count
}

func initNumMap() map[int]string {
	numMap := make(map[int]string)
	numMap[1] = "One"
	numMap[2] = "Two"
	numMap[3] = "Three"
	numMap[4] = "Four"
	numMap[5] = "Five"
	numMap[6] = "Six"
	numMap[7] = "Seven"
	numMap[8] = "Eight"
	numMap[9] = "Nine"
	numMap[10] = "Ten"
	numMap[11] = "Eleven"
	numMap[12] = "Twelve"
	numMap[13] = "Thirteen"
	numMap[14] = "Fourteen"
	numMap[15] = "Fifteen"
	numMap[16] = "Sixteen"
	numMap[17] = "Seventeen"
	numMap[18] = "Eighteen"
	numMap[19] = "Nineteen"
	numMap[20] = "Twenty"
	numMap[30] = "Thirty"
	numMap[40] = "Forty"
	numMap[50] = "Fifty"
	numMap[60] = "Sixty"
	numMap[70] = "Seventy"
	numMap[80] = "Eighty"
	numMap[90] = "Ninety"
	return numMap
}

// getNumStr 获取数字 num 的英语表示.
// num: 待表示为英文的数字
// nextNum: 下一位数字, 当 num 为 1 时使用
// bit: 当前数字 num 所处位数
// bitIn3Num: 在 3个一组的数字中的位数, 如 612,345, 数字 1 就位于 612 的第二位, 故 bit = 5, bitIn3Num = 2
// segmentSum: 当前 3个一组的数字的和, 如 1,000,000 其 000 的 segmentSum 就为 0, 此时不添加 "Thousand" 等后缀
// numToStrMap: 数字转英文标识映射表
func getNumStr(num int, nextNum, bit int, bitIn3Num int, segmentSum int, numToStrMap map[int]string) (str string, skipNextNum bool) {
	if segmentSum == 0 {
		return "", false
	}
	var sb strings.Builder
	base := numToStrMap[num]
	switch bitIn3Num {
	case 1:
		if base != "" {
			sb.WriteString(" " + base)
		}
		switch (bit - 1) / 3 {
		case 1:
			sb.WriteString(" Thousand")
		case 2:
			sb.WriteString(" Million")
		case 3:
			sb.WriteString(" Billion")
		case 4: // Trillion
			sb.WriteString(" Trillion")
		default:

		}

	case 2:
		if num == 1 {
			base = numToStrMap[num*10+nextNum]
			skipNextNum = true
		} else {
			base = numToStrMap[num*10]
		}
		if base == "" {
			return "", false
		}
		sb.WriteString(" " + base)
		if skipNextNum == true {
			switch (bit - 2) / 3 {
			case 1:
				sb.WriteString(" Thousand")
			case 2:
				sb.WriteString(" Million")
			case 3:
				sb.WriteString(" Billion")
			case 4: // Trillion
				sb.WriteString(" Trillion")
			default:

			}
		}
	case 3:
		if base == "" {
			return "", false
		}
		sb.WriteString(" " + base)
		sb.WriteString(" Hundred")
	}
	return sb.String(), skipNextNum
}

// 示例 1：
// 输入：123
// 输出："One Hundred Twenty Three"
//
// 示例 2：
// 输入：12345
// 输出："Twelve Thousand Three Hundred Forty Five"
//
// 示例 3：
// 输入：1234567
// 输出："One Million Two Hundred Thirty Four Thousand Five Hundred Sixty Seven"
//
// 示例 4：
// 输入：1234567891
// 输出："One Billion Two Hundred Thirty Four Million Five Hundred Sixty Seven Thousand Eight Hundred Ninety One"
func main() {
	var num int
	fmt.Scan(&num)
	fmt.Println(numberToWords(num))
}
