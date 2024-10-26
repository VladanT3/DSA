package InsertionSort

import (
	"testing"

	SortingAlgorithms "github.com/VladanT3/DSA/sorting_algorithms"
)

func TestInsertionSort(t *testing.T) {
	res := insertionSort([]int{7, 6, 5, 4, 3, 2, 1})
	if !SortingAlgorithms.EqualArrays(res, []int{1, 2, 3, 4, 5, 6, 7}) {
		t.Errorf("\nexpected: %d\ngot: %d", []int{1, 2, 3, 4, 5, 6, 7}, res)
	}
	res = insertionSort([]int{8, 7, 6, 5, 4, 3, 2, 1})
	if !SortingAlgorithms.EqualArrays(res, []int{1, 2, 3, 4, 5, 6, 7, 8}) {
		t.Errorf("\nexpected: %d\ngot: %d", []int{1, 2, 3, 4, 5, 6, 7, 8}, res)
	}
	res = insertionSort([]int{7, 6, 2, 4, 6, 2, 1})
	if !SortingAlgorithms.EqualArrays(res, []int{1, 2, 2, 4, 6, 6, 7}) {
		t.Errorf("\nexpected: %d\ngot: %d", []int{1, 2, 2, 4, 6, 6, 7}, res)
	}
	res = insertionSort([]int{7, 6, 2, 8, 4, 6, 2, 1})
	if !SortingAlgorithms.EqualArrays(res, []int{1, 2, 2, 4, 6, 6, 7, 8}) {
		t.Errorf("\nexpected: %d\ngot: %d", []int{1, 2, 2, 4, 6, 6, 7, 8}, res)
	}
}
