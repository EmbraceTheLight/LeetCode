// 10.10. 数字流的秩
/*
假设你正在读取一串整数。每隔一段时间，你希望能找出数字 x 的秩(小于或等于 x 的值的个数)。
请实现数据结构和算法来支持这些操作，也就是说：
实现 track(int x) 方法，每读入一个数字都会调用该方法；
实现 getRankOfNumber(int x) 方法，返回小于或等于 x 的值的个数。

提示
x <= 50000
track 和 getRankOfNumber 方法的调用次数均不超过 2000 次
*/
package main

import "fmt"

type StreamRank struct {
	arr []int
}

func Constructor1010() StreamRank {
	return StreamRank{
		arr: make([]int, 0, 20),
	}
}

func (this *StreamRank) Track(x int) {
	if len(this.arr) == 0 {
		this.arr = append(this.arr, x)
		return
	}
	idx := this.getInsertIdx(x)
	this.arr = append(this.arr[:idx], append([]int{x}, this.arr[idx:]...)...)
}

func (this *StreamRank) GetRankOfNumber(x int) int {
	if len(this.arr) == 0 {
		return 0
	}
	return this.getInsertIdx(x)
}

func (this *StreamRank) getInsertIdx(x int) int {
	left, right := 0, len(this.arr)-1
	for left <= right {
		mid := (left + right) / 2
		if this.arr[mid] < x {
			left = mid + 1
		} else if this.arr[mid] > x {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

/**
 * Your StreamRank object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Track(x);
 * param_2 := obj.GetRankOfNumber(x);
 */
// 示例：
// 输入：
// ["StreamRank", "getRankOfNumber", "track", "getRankOfNumber"]
// [[], [1], [0], [0]]
// 输出：
// [null,0,null,1]
func main() {
	// case 1:
	fmt.Println("case 1:")
	{
		streamRank := Constructor1010()
		fmt.Println(streamRank.GetRankOfNumber(1))
		streamRank.Track(0)
		fmt.Println(streamRank.GetRankOfNumber(0))
	}
	fmt.Println()
	fmt.Println("case 2:")
	// case 2:
	{
		streamRank := Constructor1010()
		streamRank.Track(1)
		fmt.Println(streamRank.GetRankOfNumber(3)) // 1
		streamRank.Track(3)
		streamRank.Track(3)
		fmt.Println(streamRank.GetRankOfNumber(4)) // 3
		streamRank.Track(5)
		fmt.Println(streamRank.GetRankOfNumber(3))  // 3
		fmt.Println(streamRank.GetRankOfNumber(10)) // 4
	}
}
