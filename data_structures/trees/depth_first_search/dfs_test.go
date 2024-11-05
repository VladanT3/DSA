package Depth_First_Search

import (
	"reflect"
	"testing"
)

func TestDFS(t *testing.T) {
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
	root.Right = &node4
	node1.Left = &node2
	node1.Right = &node3
	node4.Left = &node5
	node4.Right = &node6
	t.Run("PreOrder", func(t *testing.T) {
		res := PreOrder(&root)
		if !reflect.DeepEqual(res, []any{1, 2, 3, 4, 5, 6, 7}) {
			t.Errorf("\nexpected: %v\ngot: %v", []any{1, 2, 3, 4, 5, 6, 7}, res)
		}
	})

	root.Value = 4
	node1.Value = 2
	node2.Value = 1
	node3.Value = 3
	node4.Value = 6
	node5.Value = 5
	node6.Value = 7
	t.Run("InOrder", func(t *testing.T) {
		res := InOrder(&root)
		if !reflect.DeepEqual(res, []any{1, 2, 3, 4, 5, 6, 7}) {
			t.Errorf("\nexpected: %v\ngot: %v", []any{1, 2, 3, 4, 5, 6, 7}, res)
		}
	})

	root.Value = 7
	node1.Value = 3
	node2.Value = 1
	node3.Value = 2
	node4.Value = 6
	node5.Value = 4
	node6.Value = 5
	t.Run("PostOrder", func(t *testing.T) {
		res := PostOrder(&root)
		if !reflect.DeepEqual(res, []any{1, 2, 3, 4, 5, 6, 7}) {
			t.Errorf("\nexpected: %v\ngot: %v", []any{1, 2, 3, 4, 5, 6, 7}, res)
		}
	})
}
