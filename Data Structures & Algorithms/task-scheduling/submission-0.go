type maxHeap []int

func(h maxHeap) Len() int {
	return len(h)
}

func(h maxHeap) Less(i, j int) bool { return h[i] > h[j] }

func (h *maxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1] // Get the last element
	*h = old[:n-1] // Shrink the slice
	return x       // Return the element itself, not the whole slice
}

func(h maxHeap) Swap(i, j int) {h[i], h[j] = h[j], h[i]}

func(h *maxHeap) Push(x interface{}) {*h = append(*h, x.(int))}

type CDtasks struct {
	Freq int
	CoolDown int
}

func leastInterval(tasks []byte, n int) int {
	freq := make([]int, 26)

	for _, i := range tasks {
		freq[i - 'A']++
	}

	h := &maxHeap{}
	heap.Init(h)

	for _, f := range freq{
		if f > 0 {
			heap.Push(h, f)
		}
	}

	queue := []CDtasks{}
	time := 0 

 	for h.Len() > 0 || len(queue) > 0 {
		  time++ 
		  if h.Len() > 0 {
			currFreq := heap.Pop(h).(int)
			currFreq--
			if currFreq > 0 {
				queue = append(queue, CDtasks{
					Freq: currFreq,
					CoolDown: time + n,
				})
			}
		}

		if len(queue) > 0 && queue[0].CoolDown == time {
			heap.Push(h, queue[0].Freq)
			queue = queue[1:]
		}
	}

	return time
}
