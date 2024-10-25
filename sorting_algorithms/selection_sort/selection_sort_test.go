package SelectionSort

import (
	"testing"

	SortingAlgorithms "github.com/VladanT3/DSA/sorting_algorithms"
)

func TestSelectionSort(t *testing.T) {
	res := selectionSort([]int{7, 6, 5, 4, 3, 2, 1})
	if !SortingAlgorithms.EqualArrays(res, []int{1, 2, 3, 4, 5, 6, 7}) {
		t.Errorf("\nexpected: %d\ngot: %d", []int{1, 2, 3, 4, 5, 6, 7}, res)
	}
	res = selectionSort([]int{7, 5, 6, 3, 4, 1, 2})
	if !SortingAlgorithms.EqualArrays(res, []int{1, 2, 3, 4, 5, 6, 7}) {
		t.Errorf("\nexpected: %d\ngot: %d", []int{1, 2, 3, 4, 5, 6, 7}, res)
	}
	res = selectionSort([]int{7, 5, 6, 3, 4, 1, 2, 8})
	if !SortingAlgorithms.EqualArrays(res, []int{1, 2, 3, 4, 5, 6, 7, 8}) {
		t.Errorf("\nexpected: %d\ngot: %d", []int{1, 2, 3, 4, 5, 6, 7, 8}, res)
	}
	res = selectionSort([]int{7, 3, 6, 3, 2, 1, 2, 8})
	if !SortingAlgorithms.EqualArrays(res, []int{1, 2, 2, 3, 3, 6, 7, 8}) {
		t.Errorf("\nexpected: %d\ngot: %d", []int{1, 2, 2, 3, 3, 6, 7, 8}, res)
	}
}
