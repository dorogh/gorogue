package coll

// Prioritizable describes any value that exposes a floating-point priority used
// to order it inside a priority heap.
type Prioritizable interface {
	Priority() float64
}

// PriorityQueue maintains a min-heap of Prioritizable items so the smallest
// priority can be popped quickly.
type PriorityQueue[T Prioritizable] struct {
	items []T
}

// NewPriorityQueue returns an empty heap, optionally seeded with the provided
// items.
func NewPriorityQueue[T Prioritizable](items ...T) *PriorityQueue[T] {
	q := &PriorityQueue[T]{}
	if len(items) > 0 {
		q.items = append(q.items, items...)
		q.heapify()
	}
	return q
}

// Len returns the number of items currently stored.
func (q *PriorityQueue[T]) Len() int {
	return len(q.items)
}

// Push adds an item to the heap.
func (q *PriorityQueue[T]) Push(item T) {
	q.items = append(q.items, item)
	q.heapifyUp(len(q.items) - 1)
}

// Pop removes and returns the item with the lowest priority.
func (q *PriorityQueue[T]) Pop() (T, bool) {
	if len(q.items) == 0 {
		var zero T
		return zero, false
	}

	top := q.items[0]
	last := len(q.items) - 1
	q.items[0] = q.items[last]
	q.items = q.items[:last]
	if len(q.items) > 0 {
		q.heapifyDown(0)
	}

	return top, true
}

// Peek returns the item with the lowest priority without removing it.
func (q *PriorityQueue[T]) Peek() (T, bool) {
	if len(q.items) == 0 {
		var zero T
		return zero, false
	}
	return q.items[0], true
}

// Clear drops all items from the heap.
func (q *PriorityQueue[T]) Clear() {
	q.items = q.items[:0]
}

func (q *PriorityQueue[T]) heapify() {
	for i := len(q.items)/2 - 1; i >= 0; i-- {
		q.heapifyDown(i)
	}
}

func (q *PriorityQueue[T]) heapifyUp(idx int) {
	for idx > 0 {
		parent := (idx - 1) / 2
		if !q.less(idx, parent) {
			break
		}
		q.swap(idx, parent)
		idx = parent
	}
}

func (q *PriorityQueue[T]) heapifyDown(idx int) {
	n := len(q.items)
	for {
		left := 2*idx + 1
		if left >= n {
			break
		}
		smallest := left
		right := left + 1
		if right < n && q.less(right, left) {
			smallest = right
		}
		if !q.less(smallest, idx) {
			break
		}
		q.swap(idx, smallest)
		idx = smallest
	}
}

func (q *PriorityQueue[T]) less(i, j int) bool {
	return q.items[i].Priority() < q.items[j].Priority()
}

func (q *PriorityQueue[T]) swap(i, j int) {
	q.items[i], q.items[j] = q.items[j], q.items[i]
}
