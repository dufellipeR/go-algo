package chapter_15

type Heap struct {
	data []int
}

func (h Heap) RootNode() any {
	return h.data[0]
}

func (h Heap) LastNode() any {
	return h.data[len(h.data)-1]
}

func (h Heap) LeftChildIndex(index int) int {
	return (index * 2) + 1
}

func (h Heap) RightChildIndex(index int) int {
	return (index * 2) + 2
}

func (h Heap) ParentIndex(index int) int {
	return (index - 1) / 2
}

func (h Heap) Insert(value int) {
	h.data = append(h.data, value)

	newNodeIndex := len(h.data) - 1

	for newNodeIndex > 0 && h.data[newNodeIndex] > h.data[h.ParentIndex(newNodeIndex)] {
		h.data[h.ParentIndex(newNodeIndex)], h.data[newNodeIndex] = h.data[newNodeIndex],
			h.data[h.ParentIndex(newNodeIndex)]

		newNodeIndex = h.ParentIndex(newNodeIndex)
	}
}
