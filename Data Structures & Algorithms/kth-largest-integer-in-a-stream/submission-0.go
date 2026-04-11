
type Heap struct {
	data []int
}

func (h *Heap) Insert(val int) {
	h.data = append(h.data, val)
	i := len(h.data) - 1

	for i > 0 {
		p := (i - 1) / 2
		if h.data[p] <= h.data[i] {
			break
		}
		h.data[p], h.data[i] = h.data[i], h.data[p]
		i = p
	}
}

func (h *Heap) ExtractMin() int {
	min := h.data[0]
	last := h.data[len(h.data)-1]

	h.data[0] = last
	h.data = h.data[:len(h.data)-1]

	i := 0
	n := len(h.data)

	for {
		l := 2*i + 1
		r := 2*i + 2
		smallest := i

		if l < n && h.data[l] < h.data[smallest] {
			smallest = l
		}
		if r < n && h.data[r] < h.data[smallest] {
			smallest = r
		}

		if smallest == i {
			break
		}

		h.data[i], h.data[smallest] = h.data[smallest], h.data[i]
		i = smallest
	}

	return min
}

func (h *Heap) Peek() int {
	return h.data[0]
}
type KthLargest struct {
	k    int
	heap *Heap
}

func Constructor(k int, nums []int) KthLargest {
	h := &Heap{}

	for _, n := range nums {
		h.Insert(n)
		if len(h.data) > k {
			h.ExtractMin()
		}
	}

	return KthLargest{k: k, heap: h}
}

func (kl *KthLargest) Add(val int) int {
	kl.heap.Insert(val)

	if len(kl.heap.data) > kl.k {
		kl.heap.ExtractMin()
	}

	return kl.heap.Peek()
}
