package dheap

import "errors"

type DHeap[T comparable] struct {
	compare         func(a, b T) bool
	elems           []T
	elemToIndex     map[T]int
	branchingFactor int
}

func NewDHeap[T comparable](
	compareFunc func(a, b T) bool,
	branchingFactor int,
	capacity int,
) (*DHeap[T], error) {
	if branchingFactor < 2 || capacity < 0 {
		return nil, errors.New("invalid parameters")
	}

	return &DHeap[T]{
		compare:         compareFunc,
		elems:           make([]T, 0, capacity),
		elemToIndex:     make(map[T]int, capacity),
		branchingFactor: branchingFactor,
	}, nil
}

func (h *DHeap[T]) Peek() (T, error) {
	if len(h.elems) == 0 {
		var zero T
		return zero, errors.New("heap is empty")
	}
	return h.elems[0], nil
}

func (h *DHeap[T]) Size() int {
	return len(h.elems)
}

func (h *DHeap[T]) Insert(elem T) error {
	if _, ok := h.elemToIndex[elem]; ok {
		return errors.New("duplicate element not allowed")
	}

	h.elems = append(h.elems, elem)
	index := len(h.elems) - 1
	h.elemToIndex[elem] = index
	h.bubbleUp(index)

	return nil
}

func (h *DHeap[T]) Top() (T, error) {
	topElem, err := h.Peek()
	if err != nil {
		return topElem, err
	}

	h.elems[0] = h.popLastElem()

	h.pushDown(0)
	return topElem, nil
}

func (h *DHeap[T]) Update(elem T, updated T) error {
	index, ok := h.elemToIndex[elem]
	if !ok {
		return errors.New("element not found")
	}

	delete(h.elemToIndex, elem)
	h.updateElemAtIndex(index, updated)
	h.compareAndMove(elem, updated, index)

	return nil
}

func (h *DHeap[T]) Remove(elem T) error {
	index, ok := h.elemToIndex[elem]
	if !ok {
		return errors.New("element not found")
	}

	delete(h.elemToIndex, elem)
	h.elems[index] = h.popLastElem()
	h.compareAndMove(elem, h.elems[index], index)

	return nil
}

func (h *DHeap[T]) compareAndMove(oldElem, newElem T, index int )  {
	if oldElem == newElem {
		return
	}
	
	if h.compare(oldElem, h.elems[index]) {
		h.pushDown(index)
	} else {
		h.bubbleUp(index)
	}
}

func (h *DHeap[T]) bubbleUp(index int) {
	h.move(index, h.getParentIndex, func(next, current int)bool {
		return h.compare(h.elems[current], h.elems[next])
	})
}

func (h *DHeap[T]) pushDown(index int) {
	h.move(index, h.getBestChild, func(next, current int) bool {
		return h.compare(h.elems[next], h.elems[current])
	})
}

func (h *DHeap[T]) move(
	index int,
	direction func(int) int,
	condition func(int, int) bool) {
	current := h.elems[index]
	for {
		nextIndex := direction(index)
		if nextIndex < 0 || nextIndex >= len(h.elems) || !condition(nextIndex, index) {
			break
		}
		h.updateElemAtIndex(index, h.elems[nextIndex])
		index = nextIndex
	}
	h.updateElemAtIndex(index, current)
}

func (h *DHeap[T]) getParentIndex(index int) int {
	return (index - 1) / h.branchingFactor
}

func (h *DHeap[T]) getBestChild(index int) (childIndex int) {
	firstChildIndex := h.branchingFactor*index + 1
	bestChildIndex := firstChildIndex
	for i := firstChildIndex+1; i < len(h.elems) && i <= firstChildIndex+h.branchingFactor; i++ {
		if h.compare(h.elems[i], h.elems[bestChildIndex]) {
			bestChildIndex = i
		}
	}
	return bestChildIndex
}

func (h *DHeap[T]) updateElemAtIndex(index int, elem T) {
	h.elems[index] = elem
	h.elemToIndex[elem] = index
}

func (h *DHeap[T]) popLastElem() T {
	n := len(h.elems)
	lastElem := h.elems[n-1]
	h.elems = h.elems[:n-1]
	return lastElem
}
