package tree_comparison

type Node struct {
	Value any
	Left  *Node
	Right *Node
}

func Tree_Comparison(curr1 *Node, curr2 *Node) bool {
	if curr1 == nil && curr2 == nil {
		return true
	} else if curr1 == nil || curr2 == nil {
		return false
	}

	if curr1.Value == curr2.Value {
		return Tree_Comparison(curr1.Left, curr2.Left) && Tree_Comparison(curr1.Right, curr2.Right)
	} else {
		return false
	}
}
