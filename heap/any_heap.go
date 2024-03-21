type Heap struct {
	Data     []any
	lessFunc func(a, b any) bool
}

func (h Heap) Len() int           { return len(h.Data) }
func (h Heap) Less(i, j int) bool { return h.lessFunc(h.Data[i], h.Data[j]) }
func (h Heap) Swap(i, j int)      { h.Data[i], h.Data[j] = h.Data[j], h.Data[i] }

func (h *Heap) Push(x any) {
	h.Data = append(h.Data, x)
}

func (h *Heap) Pop() any {
  if h.IsEmpty() {
		panic("Error: Cannot pop, heap is empty!")
	}
	x := h.Data[(*h).Len()-1]
	h.Data = h.Data[:h.Len()-1]
	return x
}

func (h *Heap) Peek() any {
	if h.IsEmpty() {
		panic("Cannot peak, heap is empty!")
	}
	return h.Data[0]
}

func (h *Heap) IsEmpty() bool {
	return len(h.Data) == 0
}
