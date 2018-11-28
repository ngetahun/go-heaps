// Weak heap implementation
// reference: https://en.wikipedia.org/wiki/Weak_heap
// Implementation adopted from https://www.sanfoundry.com/cpp-program-implement-weak-heap/

package weak_heap

import (
	heap "github.com/theodesp/go-heaps"
)

// Weak heap data structure
type WeakHeap struct {
	root  *node
	nodes []*node
	// height int
}

type node struct {
	item     heap.Item
	children []*node
	parent   *node
	flag     uint8
}

func toggleFlag() {}

func New() *WeakHeap { return new(WeakHeap).init() }

func (w *WeakHeap) init() *WeakHeap {
	w.root = &node{}
	w.height = 0
	return w
}

// func construct() *WeakHeap {
// 	return nil
// }

// Insert to weak heap structure
func (w *WeakHeap) Insert(item heap.Item) heap.Item {
	newnode := &node{item: item}
	// w.root.children.append(newnode)
	return newnode
}

// Delete the minimum node from weak heap structure
func (w *WeakHeap) DeleteMin() heap.Item {}

// Find the minimum node in the weak heap structure
func (w *WeakHeap) FindMin() heap.Item {
	// get a sorted heap
	var temp *WeakHeap = sort(w, DESC)
	return temp.root.item
}

// clears weak heap
func (w *WeakHeap) Clear() {
	w.root = &node{}
}

func (w *WeakHeap) Size() int {
	return len(w.root.children)
}

// check if weak heap is empty
// returns a boolean value
func (w *WeakHeap) IsEmpty() bool {
	return w.root.item == nil
}

func (w *WeakHeap) merge() {}

// sorts the weak heap in an order given
func sort(w *WeakHeap, order int) *WeakHeap {
	return nil
}

// Insert item to the WeakHeap structure

func (w *WeakHeap) join(i int, j int) {}

func (w *WeakHeap) siftdown() {}

func (w *WeakHeap) siftup() {}
