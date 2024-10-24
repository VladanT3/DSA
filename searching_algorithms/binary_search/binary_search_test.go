package BinarySearch

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	res := binarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 1)
	if res != 0 {
		t.Errorf("expected: 0\ngot: %d", res)
	}
	res = binarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 3)
	if res != 2 {
		t.Errorf("expected: 2\ngot: %d", res)
	}
	res = binarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 5)
	if res != 4 {
		t.Errorf("expected: 4\ngot: %d", res)
	}
	res = binarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 6)
	if res != 5 {
		t.Errorf("expected: 5\ngot: %d", res)
	}
	res = binarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 8)
	if res != 7 {
		t.Errorf("expected: 7\ngot: %d", res)
	}
	res = binarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 1)
	if res != 0 {
		t.Errorf("expected: 0\ngot: %d", res)
	}
	res = binarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 3)
	if res != 2 {
		t.Errorf("expected: 2\ngot: %d", res)
	}
	res = binarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5)
	if res != 4 {
		t.Errorf("expected: 4\ngot: %d", res)
	}
	res = binarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 6)
	if res != 5 {
		t.Errorf("expected: 5\ngot: %d", res)
	}
	res = binarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 7)
	if res != 6 {
		t.Errorf("expected: 6\ngot: %d", res)
	}
	res = binarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 9)
	if res != 8 {
		t.Errorf("expected: 8\ngot: %d", res)
	}
}
