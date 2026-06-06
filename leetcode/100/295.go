// 295. 数据流的中位数
/*
中位数是有序整数列表中的中间值。如果列表的大小是偶数，则没有中间值，中位数是两个中间值的平均值。

例如 arr = [2,3,4] 的中位数是 3。
例如 arr = [2,3] 的中位数是 (2 + 3) / 2 = 2.5 。
实现 MedianFinder 类:
MedianFinder() 初始化 MedianFinder 对象。
void addNum(int num) 将数据流中的整数 num 添加到数据结构中。
double findMedian() 返回到目前为止所有元素的中位数。与实际答案相差 10^-5 以内的答案将被接受。


// 提示:
-10^5 <= num <= 10^5
在调用 findMedian 之前，数据结构中至少有一个元素
最多 5 * 10^4 次调用 addNum 和 findMedian
*/
package main

import (
	"container/heap"
	"fmt"
)

type IHeap interface {
	Len() int
	Push(x any)
	Pop() any
	Less(i, j int) bool
	Swap(i, j int)
}
type bigRoot295 []int

func (br *bigRoot295) Top() int {
	return (*br)[0]
}
func (br *bigRoot295) Len() int {
	return len(*br)
}

func (br *bigRoot295) Push(x any) {
	*br = append(*br, x.(int))
}

func (br *bigRoot295) Pop() any {
	ret := (*br)[len(*br)-1]
	*br = (*br)[:len(*br)-1]
	return ret
}

func (br *bigRoot295) Less(i, j int) bool {
	return (*br)[i] > (*br)[j]
}

func (br *bigRoot295) Swap(i, j int) {
	(*br)[i], (*br)[j] = (*br)[j], (*br)[i]
}

type smallRoot295 []int

func (sr *smallRoot295) Top() int {
	return (*sr)[0]
}
func (sr *smallRoot295) Len() int {
	return len(*sr)
}

func (sr *smallRoot295) Push(x any) {
	*sr = append(*sr, x.(int))
}

func (sr *smallRoot295) Pop() any {
	ret := (*sr)[len(*sr)-1]
	*sr = (*sr)[:len(*sr)-1]
	return ret
}

func (sr *smallRoot295) Less(i, j int) bool {
	return (*sr)[i] < (*sr)[j]
}

func (sr *smallRoot295) Swap(i, j int) {
	(*sr)[i], (*sr)[j] = (*sr)[j], (*sr)[i]
}

type MedianFinder struct {
	br *bigRoot295
	sr *smallRoot295
}

func Constructor295() MedianFinder {
	br, sr := &bigRoot295{}, &smallRoot295{}
	heap.Init(br)
	heap.Init(sr)
	return MedianFinder{br: br, sr: sr}
}

func (this *MedianFinder) AddNum(num int) {
	len1, len2 := this.sr.Len(), this.br.Len()
	length := len1 + len2
	if length == 0 {
		heap.Push(this.br, num)
		return
	}
	mid := this.FindMedian()
	if float64(num) > mid {
		heap.Push(this.sr, num)
		if this.br.Len() < this.sr.Len() {
			x := heap.Pop(this.sr).(int)
			heap.Push(this.br, x)
		}
	} else {
		heap.Push(this.br, num)
		if this.br.Len() > this.sr.Len()+1 {
			x := heap.Pop(this.br).(int)
			heap.Push(this.sr, x)
		}
	}

}

func (this *MedianFinder) FindMedian() float64 {
	length := this.br.Len() + this.sr.Len()
	if length%2 == 0 {
		return float64(this.br.Top()+this.sr.Top()) / 2
	}
	return float64(this.br.Top())
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
// 示例 1：
// 输入
// ["MedianFinder", "addNum", "addNum", "findMedian", "addNum", "findMedian"]
// [[], [1], [2], [], [3], []]
// 输出
// [null, null, null, 1.5, null, 2.0]
//
// 解释
// MedianFinder medianFinder = new MedianFinder();
// medianFinder.addNum(1);    // arr = [1]
// medianFinder.addNum(2);    // arr = [1, 2]
// medianFinder.findMedian(); // 返回 1.5 ((1 + 2) / 2)
// medianFinder.addNum(3);    // arr[1, 2, 3]
// medianFinder.findMedian(); // return 2.0
func main() {
	{
		medianFinder := Constructor295()
		medianFinder.AddNum(1)                 // arr = [1]
		medianFinder.AddNum(2)                 // arr = [1, 2]
		fmt.Println(medianFinder.FindMedian()) // 返回 1.5 ((1 + 2) / 2)
		medianFinder.AddNum(3)                 // arr[1, 2, 3]
		fmt.Println(medianFinder.FindMedian()) // return 2.0
	}
	{
		medianFinder := Constructor295()
		medianFinder.AddNum(0)                 // arr = [0]
		medianFinder.AddNum(0)                 // arr = [0, 0]
		fmt.Println(medianFinder.FindMedian()) // 返回 0.0 ((0 + 0) / 2)
		medianFinder.AddNum(3)                 // arr[0, 0, 3]
		fmt.Println(medianFinder.FindMedian()) // return 0.0
	}
}
