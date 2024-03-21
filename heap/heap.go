/*
  An implementation of minHeap taken from https://pkg.go.dev/container/heap with some added functions.
  Change line 8 to ahcieve maxHeap
*/
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] } // Flip this to get maxHeap
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *IntHeap) Peek() any {
    return (*h)[0]
}

func (h *IntHeap) IsEmpty() bool {
    return len(*h) == 0
}
