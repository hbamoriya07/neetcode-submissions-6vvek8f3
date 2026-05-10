// 1. Define the underlying type for the Min-Heap
type IntHeap []int

// 2. Implement sort.Interface methods (Len, Less, Swap)
func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] } // Min-Heap condition
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// 3. Implement heap.Interface methods (Push, Pop)
// Note: These use pointer receivers because they modify the slice's length
func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]     // Grab the last element
	*h = old[0 : n-1] // Shrink the slice
	return x
}

// 4. Define the KthLargest struct
type KthLargest struct {
	minHeap *IntHeap
	k       int
}

// Constructor initializes the object with the integer k and the initial stream
func Constructor(k int, nums []int) KthLargest {
	h := &IntHeap{}
	heap.Init(h) // Initialize the empty heap

	kl := KthLargest{
		minHeap: h,
		k:       k,
	}

	// Process initial numbers through the Add method to maintain size constraint
	for _, num := range nums {
		kl.Add(num)
	}

	return kl
}

// Add appends a value to the stream and returns the kth largest element
func (this *KthLargest) Add(val int) int {
	heap.Push(this.minHeap, val)

	// If the heap grows larger than k, evict the smallest element
	if this.minHeap.Len() > this.k {
		heap.Pop(this.minHeap)
	}

	// The root of the min-heap is the kth largest element
	return (*this.minHeap)[0]
}