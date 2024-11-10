package adjmatrix

import (
	"reflect"
	"testing"
)

func TestAdjMatrixBFS(t *testing.T) {
	var matrix WeightedAdjacencyMatrix = [][]int{[]int{0, 0, 5, 0}, []int{5, 0, 0, 0}, []int{0, 0, 0, 5}, []int{5, 10, 0, 0}}
	res := AdjMatrixBFS(matrix, 0, 1)
	if !reflect.DeepEqual(res, []int{0, 2, 3, 1}) {
		t.Errorf("\nexpected: %v\ngot: %v", []int{0, 2, 3, 1}, res)
	}
	res = AdjMatrixBFS(matrix, 1, 3)
	if reflect.DeepEqual(res, []int{}) {
		t.Errorf("\nexpected: %v\ngot: %v", []int{}, res)
	}
}
