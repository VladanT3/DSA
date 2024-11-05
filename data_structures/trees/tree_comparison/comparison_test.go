package tree_comparison

import "testing"

func TestTreeComparison(t *testing.T) {
	root1 := Node{
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
	root1.Left = &node1
	root1.Right = &node2
	node1.Left = &node3
	node1.Right = &node4
	node2.Left = &node5
	node2.Right = &node6

	root2 := Node{
		Value: 1,
	}
	node7 := Node{
		Value: 2,
	}
	node8 := Node{
		Value: 3,
	}
	node9 := Node{
		Value: 4,
	}
	node10 := Node{
		Value: 5,
	}
	node11 := Node{
		Value: 6,
	}
	node12 := Node{
		Value: 7,
	}
	root2.Left = &node7
	root2.Right = &node8
	node7.Left = &node9
	node7.Right = &node10
	node8.Left = &node11
	node8.Right = &node12

	res := Tree_Comparison(&root1, &root2)
	if !res {
		t.Errorf("\nexpected: %v\ngot: %v", true, res)
	}

	node7.Value = 13
	res = Tree_Comparison(&root1, &root2)
	if res {
		t.Errorf("\nexpected: %v\ngot: %v", false, res)
	}

	node7.Value = 2
	node1.Left = nil
	node4.Left = &node3
	res = Tree_Comparison(&root1, &root2)
	if res {
		t.Errorf("\nexpected: %v\ngot: %v", false, res)
	}
}
