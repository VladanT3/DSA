package heap

type MinHeap struct {
	Length int
	data   []int
}

func (heap *MinHeap) Insert(value int) {
	heap.data = append(heap.data, value)
	heap.Length++
	heap.heapifyUp(heap.Length - 1)
}

func (heap *MinHeap) Delete() int {
	if heap.Length == 0 {
		return -1
	}

	out := heap.data[0]

	if heap.Length == 1 {
		heap.Length = 0
		heap.data = []int{}
		return out
	}
	heap.data[0] = heap.data[heap.Length-1]
	heap.Length--
	heap.heapifyDown(0)

	return out
}

func (heap *MinHeap) heapifyUp(new_elem_idx int) {
	if new_elem_idx == 0 {
		return
	}

	new_element := heap.data[new_elem_idx]
	parent_idx := getParent(new_elem_idx)

	if new_element < heap.data[parent_idx] {
		heap.data[new_elem_idx] = heap.data[parent_idx]
		heap.data[parent_idx] = new_element
		heap.heapifyUp(parent_idx)
	}
}

func (heap *MinHeap) heapifyDown(new_elem_idx int) {
	left_child_idx := getLeftChild(new_elem_idx)
	right_child_idx := getRightChild(new_elem_idx)

	if new_elem_idx >= heap.Length || left_child_idx >= heap.Length {
		return
	}

	new_element := heap.data[new_elem_idx]
	left_child := heap.data[left_child_idx]
	right_child := heap.data[right_child_idx]

	if left_child < right_child && new_element > left_child {
		heap.data[left_child_idx] = new_element
		heap.data[new_elem_idx] = left_child
		heap.heapifyDown(left_child_idx)
	} else if right_child < left_child && new_element > right_child {
		heap.data[right_child_idx] = new_element
		heap.data[new_elem_idx] = right_child
		heap.heapifyDown(right_child_idx)
	} else {
		return
	}
}

func getParent(i int) int {
	return (i - 1) / 2
}

func getLeftChild(i int) int {
	return (i * 2) + 1
}

func getRightChild(i int) int {
	return (i * 2) + 2
}
