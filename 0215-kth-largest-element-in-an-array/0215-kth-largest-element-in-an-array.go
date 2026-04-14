func findKthLargest(nums []int, k int) int {
    heap:=[]int{}
    for _, i := range nums{
        if len(heap)<k{
            heap=append(heap,i)
            downwards(heap, len(heap)-1)
			continue
        }
		if i > heap[0] {
			heap[0] = i
			up(heap, 0, len(heap))
		}
    }
    return heap[0]
}

func downwards(heap []int, i int){
    for i>0{
        parent:=(i-1)/2 // heap special nature
        if heap[parent]<=heap[i]{
            break
        }
        heap[parent],heap[i]=heap[i],heap[parent]
        i=parent // find parent's parent
    }
}

func up(heap []int, i int, n int){
    for {
        left:=2*i+1
        right:=2*i+2
        parent:=i
        if left<n && heap[left]<heap[parent]{
            parent=left
        }
        if right<n && heap[right]<heap[parent]{
            parent=right
        }
        if parent==i{
            break
        }
        heap[i],heap[parent]=heap[parent],heap[i]
        i=parent
    }
}