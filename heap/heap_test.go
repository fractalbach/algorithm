package heap

import (
	"fmt"
	"testing"
)

var (
	print = fmt.Println
	asdf  = heap{5, 4, 3, 7, 9}

	h1 = heap{16, 14, 10, 8, 7, 9, 3, 2, 4, 1}
	h2 = heap{10, 4, 5, 1, 2}

	h11            = heap{23, 17, 14, 6, 13, 10, 1, 5, 7, 12}
	hSorted1       = heap{1, 2, 3, 4, 5, 6, 7}
	hSorted2       = heap{1, 3, 8, 12, 20, 22, 23, 40, 912}
	hAlmostSorted1 = heap{7, 1, 2, 3, 4, 5, 6}
)

func TestBasic(t *testing.T) {
	keyExamples := []struct {
		i, parent, left, right int
	}{
		{1, 0, 3, 4},
		{2, 0, 5, 6},
		{3, 1, 7, 8},
		{4, 1, 9, 10},
		{5, 2, 11, 12},
		{6, 2, 13, 14},
	}
	for _, ex := range keyExamples {

		// example input
		i := ex.i

		//Test parent() Function
		p := parent(i)
		if p != ex.parent {
			t.Errorf("Bad Parent(%v); got: (%v), want: (%v). ", i, p, ex.parent)
		}

		// Test left() function
		l := left(i)
		if l != ex.left {
			t.Errorf("Bad Left(%v); got: (%v), want: (%v). ", i, l, ex.left)
		}

		// Test right() function
		r := right(i)
		if r != ex.right {
			t.Errorf("Bad Right(%v); got: (%v), want: (%v). ", i, r, ex.right)
		}
	}
}

/*
func TestPrint(t *testing.T) {
	print(
		h,
		parent(1),
		parent(2),
	)
}
*/
func TestPropertyMax(t *testing.T) {
	cases := []struct {
		i    heap
		want bool
	}{
		{h1, true},
		{h2, true},
		{heap{10}, true},
		{heap{1, 2}, false},
	}
	for _, h := range cases {
		got := IsMaxHeap(h.i)
		if h.want != got {
			t.Errorf("want: %v, got: %v", h.want, got)
		}
	}
}

func TestPropertyMin(t *testing.T) {
	cases := []struct {
		i    heap
		want bool
	}{
		{h1, false},
		{h2, false},
		{heap{10}, true},
		{heap{1, 2}, true},
	}
	for _, h := range cases {
		got := IsMinHeap(h.i)
		if h.want != got {
			t.Errorf("want: %v, got: %v", h.want, got)
		}
	}
	print("      First:", IsMinHeap(hSorted1))
	print("      Second:", IsMinHeap(hSorted2))
}

func TestHeight(t *testing.T) {
	cases := []struct {
		i    heap
		want int
	}{
		{h1, 4},
		{h2, 3},
		{heap{}, 0},
		{heap{1}, 1},
		{heap{1, 2}, 2},
		{heap{1, 2, 2}, 2},
		{heap{1, 2, 2, 3}, 3},
		{heap{1, 2, 2, 3, 3}, 3},
		{heap{1, 2, 2, 3, 3, 3}, 3},
		{heap{1, 2, 2, 3, 3, 3, 3}, 3},
		{heap{1, 2, 2, 3, 3, 3, 3, 4}, 4},
		{heap{1, 2, 2, 3, 3, 3, 3, 4, 4}, 4},
		{heap{1, 2, 2, 3, 3, 3, 3, 4, 4, 4}, 4},
	}
	for _, h := range cases {
		got := Height(h.i)
		if h.want != got {
			t.Errorf("want: %v, got: %v", h.want, got)
		}
	}
}

func Test_Max_Heapify(t *testing.T) {
	h := heap{16, 4, 10, 14, 7, 9, 3, 2, 8, 1}
	print("   before:", h)
	h = MaxHeapify(h, 1)
	print("   after:", h)

	// Should now be a max heap:
	if !IsMaxHeap(h) {
		t.Errorf("%v is not a max heap", h)
	} else {
		print("   Max-Heapify was successful!")
	}
}

func Test_Min_Heapify(t *testing.T) {
	h := hAlmostSorted1
	print("   before:", h)
	h = BuildMinHeap(h)
	print("   after:", h)

	// Confirm that MinHeapify worked.
	if !IsMinHeap(h) {
		t.Errorf("%v is not a min heap", h)
	} else {
		print("   Min-Heapify was successful!")
	}
}
