package breadth_first_search

import (
	"reflect"
	"testing"
)

func TestBFS(t *testing.T) {
	root := Node{
		Value: 1,
	}
	node1 := Node{
		Value: 2,
	}
	node2 := Node{
		Value: 3,
	}
	node3 := Node{
		Value: 4,
	}
	node4 := Node{
		Value: 5,
	}
	node5 := Node{
		Value: 6,
	}
	node6 := Node{
		Value: 7,
	}
	root.Left = &node1
	root.Right = &node2
	node1.Left = &node3
	node1.Right = &node4
	node2.Left = &node5
	node2.Right = &node6

	res := Breadth_First_Search(&root)
	if !reflect.DeepEqual(res, []any{1, 2, 3, 4, 5, 6, 7}) {
		t.Errorf("\nexpected: %v\ngot: %v", []any{1, 2, 3, 4, 5, 6, 7}, res)
	}
}
