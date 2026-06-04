type Point struct {
	x, y int
	dist int
}

type MaxHeap []Point

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	return h[i].dist > h[j].dist
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(Point))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)

	item := old[n-1]
	*h = old[:n-1]

	return item
}

func kClosest(points [][]int, k int) [][]int {
	h := &MaxHeap{}
	heap.Init(h)

	for _, p := range points {
		dist := p[0]*p[0] + p[1]*p[1]

		heap.Push(h, Point{
			x:    p[0],
			y:    p[1],
			dist: dist,
		})

		if h.Len() > k {
			heap.Pop(h)
		}
	}

	res := make([][]int, 0, k)

	for h.Len() > 0 {
		p := heap.Pop(h).(Point)
		res = append(res, []int{p.x, p.y})
	}

	return res
}