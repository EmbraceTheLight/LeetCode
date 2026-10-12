// 17.06. 2出现的次数
/*
编写一个方法，计算从 0 到 n (含 n) 中数字 2 出现的次数。

提示：
n <= 10^9
*/
package main

import (
	"fmt"
)

// numberOf2sInRange 返回闭区间 [0, n] 中数字 2 出现的总次数。
// 算法分别统计个位、十位、百位等位置上数字 2 的出现次数；
// 每一位的数字都以 0～9 为一个循环，因此只需计算完整周期和最后一个不完整周期的贡献。
func numberOf2sInRange(n int) int {
	var total int
	// placeValue 表示当前统计位置的位权：1 为个位，10 为十位，100 为百位，以此类推。
	for placeValue := 1; placeValue <= n; placeValue = placeValue * 10 {
		// cycleLength 是当前位从 0 变化到 9 所需经历的完整周期长度。
		// 例如统计十位时 placeValue=10、cycleLength=100；每 100 个数，十位完成一次 0～9 的循环。
		cycleLength := placeValue * 10

		// fullCycles 表示 [0, n] 已经完整经历了多少个 cycleLength。
		// 每个完整周期中，当前位为 2 的数字恰好有 placeValue 个，
		// 例如十位在 0～99 中为 2 的数字是 20～29，共 10 个。
		fullCycles := n / cycleLength

		// cycleOffset 表示 n 在最后一个不完整周期中的偏移量，而不是 n 的当前位数字。
		// 它等价于“当前位数字 * placeValue + 当前位右侧的数字”。
		// 例如 n=325、统计十位时，cycleOffset=325%100=25，即已经走到该周期的第 25 个位置。
		cycleOffset := n % cycleLength

		// 在每个周期中，当前位为 2 的区间是 [2*placeValue, 3*placeValue-1]。
		// partialCount 计算最后一个不完整周期覆盖了该区间多少个数；+1 是因为区间包含起点。
		// 例如 n=325、placeValue=10 时，覆盖 320～325，因此 partialCount=25-20+1=6。
		partialCount := cycleOffset - 2*placeValue + 1

		// cycleOffset 尚未到达当前位为 2 的区间时，额外贡献为 0。
		if partialCount < 0 {
			partialCount = 0
		}
		// 当前位为 2 的完整区间只有 placeValue 个数。
		// 当 cycleOffset 已越过该区间（也就是当前位数字大于 2）时，partialCount 会大于 placeValue，
		// 但本周期最多只能贡献 placeValue 次，所以需要将其封顶。
		if partialCount > placeValue {
			partialCount = placeValue
		}

		// fullCycles*placeValue 是所有完整周期的贡献，partialCount 是最后一个不完整周期的贡献。
		total += fullCycles*placeValue + partialCount
	}
	return total
}

// 示例：
// 输入：25
// 输出：9
// 解释：(2, 12, 20, 21, 22, 23, 24, 25)(注意 22 应该算作两次)
func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(numberOf2sInRange(n))
}
