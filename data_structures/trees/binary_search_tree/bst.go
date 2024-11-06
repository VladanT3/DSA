package binary_search_tree

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func Find(node *Node, value int) bool {
	if node == nil {
		return false
	}

	if node.Value == value {
		return true
	} else if node.Value < value {
		return Find(node.Right, value)
	} else {
		return Find(node.Left, value)
	}
}
