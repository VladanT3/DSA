package QuickSort

import (
	"testing"

	SortingAlgorithms "github.com/VladanT3/DSA/sorting_algorithms"
)

func TestQuickSort(t *testing.T) {
	var test_case []int = []int{7, 6, 5, 4, 3, 2, 1}
	quickSort(test_case, 0, len(test_case)-1)
	if !SortingAlgorithms.EqualArrays(test_case, []int{1, 2, 3, 4, 5, 6, 7}) {
		t.Errorf("\nexpected: %d\ngot: %d", []int{1, 2, 3, 4, 5, 6, 7}, test_case)
	}
	test_case = []int{8, 7, 6, 5, 4, 3, 2, 1}
	quickSort(test_case, 0, len(test_case)-1)
	if !SortingAlgorithms.EqualArrays(test_case, []int{1, 2, 3, 4, 5, 6, 7, 8}) {
		t.Errorf("\nexpected: %d\ngot: %d", []int{1, 2, 3, 4, 5, 6, 7, 8}, test_case)
	}
	test_case = []int{7, 6, 2, 4, 6, 2, 1}
	quickSort(test_case, 0, len(test_case)-1)
	if !SortingAlgorithms.EqualArrays(test_case, []int{1, 2, 2, 4, 6, 6, 7}) {
		t.Errorf("\nexpected: %d\ngot: %d", []int{1, 2, 2, 4, 6, 6, 7}, test_case)
	}
	test_case = []int{7, 6, 2, 8, 4, 6, 2, 1}
	quickSort(test_case, 0, len(test_case)-1)
	if !SortingAlgorithms.EqualArrays(test_case, []int{1, 2, 2, 4, 6, 6, 7, 8}) {
		t.Errorf("\nexpected: %d\ngot: %d", []int{1, 2, 2, 4, 6, 6, 7, 8}, test_case)
	}
}
