package BubbleSort

import (
	"testing"

	SortingAlgorithms "github.com/VladanT3/DSA/sorting_algorithms"
)

func TestBubbleSort(t *testing.T) {
	res := bubbleSort([]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1})
	if !SortingAlgorithms.EqualArrays(res, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}) {
		t.Errorf("expected: %d\ngot: %d", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, res)
	}
	res = bubbleSort([]int{10, 8, 9, 3, 4, 2, 1, 5, 7, 6})
	if !SortingAlgorithms.EqualArrays(res, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}) {
		t.Errorf("expected: %d\ngot: %d", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, res)
	}
	res = bubbleSort([]int{1, 3, 2, 5, 4, 7, 6, 8, 9})
	if !SortingAlgorithms.EqualArrays(res, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}) {
		t.Errorf("expected: %d\ngot: %d", []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, res)
	}
	res = bubbleSort([]int{1, 3, 6, 5, 6, 7, 9, 8, 9})
	if !SortingAlgorithms.EqualArrays(res, []int{1, 3, 5, 6, 6, 7, 8, 9, 9}) {
		t.Errorf("expected: %d\ngot: %d", []int{1, 3, 5, 6, 6, 7, 8, 9, 9}, res)
	}
	res = bubbleSort([]int{9, 8, 9, 7, 6, 5, 6, 3, 1})
	if !SortingAlgorithms.EqualArrays(res, []int{1, 3, 5, 6, 6, 7, 8, 9, 9}) {
		t.Errorf("expected: %d\ngot: %d", []int{1, 3, 5, 6, 6, 7, 8, 9, 9}, res)
	}
}
