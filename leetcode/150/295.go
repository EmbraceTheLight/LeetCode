// 295. 数据流的中位数
/*
中位数是有序整数列表中的中间值。如果列表的大小是偶数，则没有中间值，中位数是两个中间值的平均值。

例如 arr = [2,3,4] 的中位数是 3 。
例如 arr = [2,3] 的中位数是 (2 + 3) / 2 = 2.5 。
实现 MedianFinder 类:

MedianFinder() 初始化 MedianFinder 对象。
void addNum(int num) 将数据流中的整数 num 添加到数据结构中。
double findMedian() 返回到目前为止所有元素的中位数。与实际答案相差 10-5 以内的答案将被接受。

提示:
-10^5 <= num <= 10^5
在调用 findMedian 之前，数据结构中至少有一个元素
最多 5 * 10^4 次调用 addNum 和 findMedian
*/
package main

import (
	"container/heap"
	"fmt"
)

// * 思路: 维护两个堆, 一个是小根堆 pqMin, 存放大于中位数的数, 另一个是大根堆 pqMax, 存放小于等于中位数的数
// * 插入时, 若还没有插入任何元素, 将其插入 pqMax 中, 若插入的数字小于当前中位数, 将其插入 pqMax. 否则插入 pqMin
// * 注意: 插入后需要调整堆, 如果 paMax 长度大于 pqMin 的长度 + 1, 则将 pqMax 堆顶元素移入 pqMin 中;
// * 若 paMin 长度大于 pqMax 的长度, 则将 pqMin 堆顶元素移入 pqMax 中
// * 这是为了保证每次取中位数时, pqMax 的长度等于 pqMin 的长度 + 1(奇数个元素) 或相等(偶数个元素)
// * 取中位数时, 若总元素个数为奇数, 则返回 pqMax 的堆顶元素; 若总元素个数为偶数, 则返回 pqMax 和 pqMin 堆顶元素的平均值
type MedianFinder struct {
	pqMin *priorityQueue295Min // 小根堆 存放大于中位数的数
	pqMax *priorityQueue295Max // 大根堆 存放小于等于中位数的数
}

func Constructor() MedianFinder {
	return MedianFinder{
		pqMax: &priorityQueue295Max{},
		pqMin: &priorityQueue295Min{},
	}
}

func (this *MedianFinder) AddNum(num int) {
	if this.pqMax.Len() == 0 && this.pqMin.Len() == 0 {
		heap.Push(this.pqMax, num)
		return
	}
	mid := this.FindMedian()
	if float64(num) <= mid {
		heap.Push(this.pqMax, num)
		if this.pqMax.Len() > this.pqMin.Len()+1 {
			heap.Push(this.pqMin, heap.Pop(this.pqMax))
		}
	} else {
		heap.Push(this.pqMin, num)
		if this.pqMax.Len() < this.pqMin.Len() {
			heap.Push(this.pqMax, heap.Pop(this.pqMin))
		}
	}
}

func (this *MedianFinder) FindMedian() float64 {
	total := this.pqMin.Len() + this.pqMax.Len()
	if total%2 == 1 {
		return float64((*this.pqMax)[0])
	} else {
		return float64((*this.pqMax)[0]+(*this.pqMin)[0]) / 2
	}
}

type priorityQueue295Min []int
type priorityQueue295Max []int

func (p *priorityQueue295Min) Len() int {
	return len(*p)
}

func (p *priorityQueue295Min) Less(i, j int) bool {
	return (*p)[i] < (*p)[j]
}

func (p *priorityQueue295Min) Swap(i, j int) {
	(*p)[i], (*p)[j] = (*p)[j], (*p)[i]
}

func (p *priorityQueue295Min) Push(x any) {
	*p = append(*p, x.(int))
}

func (p *priorityQueue295Min) Pop() any {
	ret := (*p)[p.Len()-1]
	*p = (*p)[:p.Len()-1]
	return ret
}

func (p *priorityQueue295Max) Len() int {
	return len(*p)
}

func (p *priorityQueue295Max) Less(i, j int) bool {
	return (*p)[i] > (*p)[j]
}

func (p *priorityQueue295Max) Swap(i, j int) {
	(*p)[i], (*p)[j] = (*p)[j], (*p)[i]
}

func (p *priorityQueue295Max) Push(x any) {
	*p = append(*p, x.(int))
}

func (p *priorityQueue295Max) Pop() any {
	ret := (*p)[p.Len()-1]
	*p = (*p)[:p.Len()-1]
	return ret
}

// ! 二分插入, 超时, 主要在于找到位置后移动后面的元素时间复杂度为 O(n)
//type MedianFinder struct {
//	list []int
//}
//
//func Constructor() MedianFinder {
//	return MedianFinder{
//		list: make([]int, 0),
//	}
//}
//
//func (this *MedianFinder) AddNum(num int) {
//	idx := sort.SearchInts(this.list, num)
//	this.list = append(this.list[:idx], append([]int{num}, this.list[idx:]...)...)
//}
//
//func (this *MedianFinder) FindMedian() float64 {
//	n := len(this.list)
//	if n%2 == 1 {
//		return float64(this.list[n/2])
//	} else {
//		return (float64(this.list[n/2-1]) + float64(this.list[n/2])) / 2
//	}
//}

// 示例 1：
// 输入
// ["MedianFinder", "addNum", "addNum", "findMedian", "addNum", "findMedian"]
// [[], [1], [2], [], [3], []]
// 输出
// [null, null, null, 1.5, null, 2.0]
// 解释
// MedianFinder medianFinder = new MedianFinder();
// medianFinder.addNum(1);    // arr = [1]
// medianFinder.addNum(2);    // arr = [1, 2]
// medianFinder.findMedian(); // 返回 1.5 ((1 + 2) / 2)
// medianFinder.addNum(3);    // arr[1, 2, 3]
// medianFinder.findMedian(); // return 2.0
func main() {
	// Test Case1
	{
		medianFinder := Constructor()
		medianFinder.AddNum(1)
		medianFinder.AddNum(2)
		fmt.Println(medianFinder.FindMedian()) // 1.5
		medianFinder.AddNum(3)
		fmt.Println(medianFinder.FindMedian()) // 2.0
	}
}
