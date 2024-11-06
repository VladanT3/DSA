package binary_search_tree

import "testing"

func TestBST(t *testing.T) {
	root := Node{
		Value: 10,
	}
	node1 := Node{
		Value: 7,
	}
	node2 := Node{
		Value: 19,
	}
	node3 := Node{
		Value: 5,
	}
	node4 := Node{
		Value: 8,
	}
	node5 := Node{
		Value: 23,
	}
	node6 := Node{
		Value: 20,
	}
	node7 := Node{
		Value: 27,
	}
	root.Left = &node1
	root.Right = &node2
	node1.Left = &node3
	node1.Right = &node4
	node2.Right = &node5
	node5.Left = &node6
	node5.Right = &node7

	t.Run("Find", func(t *testing.T) {
		if !Find(&root, 20) {
			t.Errorf("\nexpected: %v\ngot: %v", true, false)
		}
	})
	t.Run("Find", func(t *testing.T) {
		if Find(&root, 3) {
			t.Errorf("\nexpected: %v\ngot: %v", false, true)
		}
	})
}
