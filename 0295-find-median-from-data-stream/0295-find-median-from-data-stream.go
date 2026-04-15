package main

// --------------------
// 手写最小堆
// --------------------
type MinHeap struct {
	data []int
}

func (h *MinHeap) Len() int {
	return len(h.data)
}

func (h *MinHeap) Top() int {
	return h.data[0]
}

func (h *MinHeap) Push(val int) {
	h.data = append(h.data, val)
	h.siftUp(len(h.data) - 1)
}

func (h *MinHeap) Pop() int {
	n := len(h.data)
	if n == 1 {
		x := h.data[0]
		h.data = h.data[:0]
		return x
	}

	top := h.data[0]
	h.data[0] = h.data[n-1]
	h.data = h.data[:n-1]
	h.siftDown(0)
	return top
}

func (h *MinHeap) siftUp(i int) {
	for i > 0 {
		parent := (i - 1) / 2
		if h.data[parent] <= h.data[i] {
			break
		}
		h.data[parent], h.data[i] = h.data[i], h.data[parent]
		i = parent
	}
}

func (h *MinHeap) siftDown(i int) {
	n := len(h.data)
	for {
		left := 2*i + 1
		right := 2*i + 2
		smallest := i

		if left < n && h.data[left] < h.data[smallest] {
			smallest = left
		}
		if right < n && h.data[right] < h.data[smallest] {
			smallest = right
		}
		if smallest == i {
			break
		}
		h.data[i], h.data[smallest] = h.data[smallest], h.data[i]
		i = smallest
	}
}

// --------------------
// 手写最大堆
// --------------------
type MaxHeap struct {
	data []int
}

func (h *MaxHeap) Len() int {
	return len(h.data)
}

func (h *MaxHeap) Top() int {
	return h.data[0]
}

func (h *MaxHeap) Push(val int) {
	h.data = append(h.data, val)
	h.siftUp(len(h.data) - 1)
}

func (h *MaxHeap) Pop() int {
	n := len(h.data)
	if n == 1 {
		x := h.data[0]
		h.data = h.data[:0]
		return x
	}

	top := h.data[0]
	h.data[0] = h.data[n-1]
	h.data = h.data[:n-1]
	h.siftDown(0)
	return top
}

func (h *MaxHeap) siftUp(i int) {
	for i > 0 {
		parent := (i - 1) / 2
		if h.data[parent] >= h.data[i] {
			break
		}
		h.data[parent], h.data[i] = h.data[i], h.data[parent]
		i = parent
	}
}

func (h *MaxHeap) siftDown(i int) {
	n := len(h.data)
	for {
		left := 2*i + 1
		right := 2*i + 2
		largest := i

		if left < n && h.data[left] > h.data[largest] {
			largest = left
		}
		if right < n && h.data[right] > h.data[largest] {
			largest = right
		}
		if largest == i {
			break
		}
		h.data[i], h.data[largest] = h.data[largest], h.data[i]
		i = largest
	}
}

// --------------------
// MedianFinder
// --------------------
type MedianFinder struct {
	small MaxHeap // 左边最大堆
	large MinHeap // 右边最小堆
}

func Constructor() MedianFinder {
	return MedianFinder{
		small: MaxHeap{data: []int{}},
		large: MinHeap{data: []int{}},
	}
}

func (this *MedianFinder) AddNum(num int) {
	// 先放进左边最大堆
	this.small.Push(num)

	// 保证 small.Top() <= large.Top()
	if this.large.Len() > 0 && this.small.Top() > this.large.Top() {
		x := this.small.Pop()
		this.large.Push(x)
	}

	// 保证数量平衡
	if this.small.Len() > this.large.Len()+1 {
		x := this.small.Pop()
		this.large.Push(x)
	} else if this.large.Len() > this.small.Len() {
		x := this.large.Pop()
		this.small.Push(x)
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.small.Len() > this.large.Len() {
		return float64(this.small.Top())
	}
	return float64(this.small.Top()+this.large.Top()) / 2.0
}