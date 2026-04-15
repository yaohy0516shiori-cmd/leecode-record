package main

import "container/heap"

// --------------------
// min heap
// --------------------
type MinHeap []int

func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(int))
}
func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// --------------------
// max heap
// --------------------
type MaxHeap []int

func (h MaxHeap) Len() int { return len(h) }
func (h MaxHeap) Less(i, j int) bool {
	return h[i] > h[j] // 反过来，变成最大堆
}
func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(int))
}
func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// --------------------
// MedianFinder
// --------------------
type MedianFinder struct {
	small *MaxHeap // 左半边，最大堆
	large *MinHeap // 右半边，最小堆
}

func Constructor() MedianFinder {
	small := &MaxHeap{}
	large := &MinHeap{}
	heap.Init(small)
	heap.Init(large)
	return MedianFinder{
		small: small,
		large: large,
	}
}

func (this *MedianFinder) AddNum(num int) {
	// 先放到 small
	heap.Push(this.small, num)

	// 保证 small 最大值 <= large 最小值
	if this.large.Len() > 0 && (*(this.small))[0] > (*(this.large))[0] {
		x := heap.Pop(this.small).(int)
		heap.Push(this.large, x)
	}

	// 平衡数量：small >= large，且最多多 1 个
	if this.small.Len() > this.large.Len()+1 {
		x := heap.Pop(this.small).(int)
		heap.Push(this.large, x)
	} else if this.large.Len() > this.small.Len() {
		x := heap.Pop(this.large).(int)
		heap.Push(this.small, x)
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.small.Len() > this.large.Len() {
		return float64((*(this.small))[0])
	}
	return float64((*(this.small))[0]+(*(this.large))[0]) / 2.0
}