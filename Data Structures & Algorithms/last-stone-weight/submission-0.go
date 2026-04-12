
// MaxHeap implementation
type MaxHeap []int

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	// Reverse for max-heap
	return h[i] > h[j]
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	val := old[n-1]
	*h = old[:n-1]
	return val
}

// Solution function
func lastStoneWeight(stones []int) int {
	h := &MaxHeap{}
	heap.Init(h)

	// Push all stones
	for _, s := range stones {
		heap.Push(h, s)
	}

	// Smash stones
	for h.Len() > 1 {
		y := heap.Pop(h).(int) // largest
		x := heap.Pop(h).(int) // second largest

		if y != x {
			heap.Push(h, y-x)
		}
	}

	if h.Len() == 0 {
		return 0
	}
	return heap.Pop(h).(int)
}