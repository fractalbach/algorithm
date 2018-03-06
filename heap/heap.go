// package heap is for practice with heaps and priority queues
//
// There is already a golang package for heaps:
// https://golang.org/src/container/heap/heap.go
//
// Some ideas & code are taken from there,
// but mostly this is a reimplementation.
//
package heap

import (
	"sort"
)

// any type that implements the heap interface can
// be used in this program.
type heapinterface interface {
	sort.Interface
	Push(interface{})
	Pop() interface{}
}

type heap []int

func parent(i int) int {
	return (i - 1) >> 1
}

func left(i int) int {
	return i<<1 + 1
}

func right(i int) int {
	return i<<1 + 2
}

// IsMaxHeap checks to see if the Max-Heap Property holds.
// Every node's value must be greater or equal to its parent.
// the Root node will be the greatest number.
func IsMaxHeap(h heap) bool {
	for i := 1; i < len(h); i++ {
		if !(h[parent(i)] >= h[i]) {
			return false
		}
	}
	return true
}

// IsMinHeap is similar to Max Heap, except that each
// node's parent must be equal or less than the roots.
// The Root node will be the smallest number.
func IsMinHeap(h heap) bool {
	for i := 1; i < len(h); i++ {
		if !(h[parent(i)] <= h[i]) {
			return false
		}
	}
	return true
}

// Height returns an positive integer value;
// If the heap were imagined as a tree.  It is the Height
// of that tree (also called depth?)
func Height(h heap) int {
	x := len(h)
	height := 0

	for x > 0 {
		x = x >> 1
		height++
	}
	return height
}

// MaxHeapify is recursive
func MaxHeapifyRecursive(h heap, i int) heap {
	l := left(i)
	r := right(i)
	largest := 0

	// Check the left side.
	if l <= len(h) && h[l] > h[i] {
		largest = l
	} else {
		largest = i
	}

	// Check the right side.
	if r <= len(h) && h[r] > h[largest] {
		largest = r
	}

	if largest != i {
		h[i], h[largest] = h[largest], h[i]
		h = MaxHeapify(h, largest)
	}

	return h
}

func MaxHeapify(h heap, i int) heap {

	for {
		l := left(i)
		r := right(i)
		largest := i

		// Check the left side.
		if l <= len(h) && h[l] > h[i] {
			largest = l
		}

		// Check the right side.
		if r <= len(h) && h[r] > h[largest] {
			largest = r
		}

		// Check if we need another iteration.
		if largest != i {
			h[i], h[largest] = h[largest], h[i]
			i = largest
			continue
		}
		break
	}
	return h
}

func MinHeapify(h heap, i int) heap {
	for {
		l := left(i)
		r := right(i)
		smallest := i

		// Check the left side.
		if l <= len(h) && h[l] < h[i] {
			smallest = l
		}

		// Check the right side.
		if r <= len(h) && h[r] > h[smallest] {
			smallest = r
		}

		// Check if we need another iteration.
		if smallest != i {
			h[i], h[smallest] = h[smallest], h[i]
			i = smallest
			continue
		}
		break
	}
	return h
}
