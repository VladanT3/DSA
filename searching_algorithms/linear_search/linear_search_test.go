package LinearSearch

import (
	"testing"
)

func TestLinearSearch(t *testing.T) {
	res := linearSearch([]int{3, 6, 8, 2, 6, 3, 5, 0, 2, 7}, 7)
	if res != 9 {
		t.Errorf("expected: 9\ngot: %d", res)
	}
	res = linearSearch([]int{3, 6, 8, 2, 6, 3, 5, 0, 2, 7}, 10)
	if res != -1 {
		t.Errorf("expected: -1\ngot: %d", res)
	}
}
