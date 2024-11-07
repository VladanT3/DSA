package binary_search_tree

import "testing"

func TestBST(t *testing.T) {
	root := Node{
		Value:       10,
		Nodes_Left:  3,
		Nodes_Right: 4,
	}
	node1 := Node{
		Value:       7,
		Nodes_Left:  1,
		Nodes_Right: 1,
	}
	node2 := Node{
		Value:       19,
		Nodes_Right: 3,
	}
	node3 := Node{
		Value: 5,
	}
	node4 := Node{
		Value: 8,
	}
	node5 := Node{
		Value:       23,
		Nodes_Left:  1,
		Nodes_Right: 1,
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

	t.Run("Insert", func(t *testing.T) {
		Insert(&root, 6)
		if !Find(&root, 6) {
			t.Errorf("\nexpected: %v\ngot: %v", true, false)
		}
	})
	t.Run("Insert", func(t *testing.T) {
		Insert(&root, 18)
		Insert(&root, 16)
		if !Find(&root, 16) {
			t.Errorf("\nexpected: %v\ngot: %v", true, false)
		}
	})

	t.Run("Delete", func(t *testing.T) {
		Delete(&root, 5)
		Delete(&root, 19)
		if Find(&root, 5) {
			t.Errorf("\nexpected: %v\ngot: %v", false, true)
		}
		if Find(&root, 19) {
			t.Errorf("\nexpected: %v\ngot: %v", false, true)
		}
		if !Find(&root, 6) {
			t.Errorf("\nexpected: %v\ngot: %v", true, false)
		}
		if !Find(&root, 20) {
			t.Errorf("\nexpected: %v\ngot: %v", true, false)
		}
	})
}
