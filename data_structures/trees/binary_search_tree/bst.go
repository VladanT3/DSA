package binary_search_tree

type Node struct {
	Value       int
	Left        *Node
	Right       *Node
	Nodes_Left  int
	Nodes_Right int
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

func Insert(curr *Node, new_value int) {
	if new_value <= curr.Value {
		if curr.Left != nil {
			curr.Nodes_Left++
			Insert(curr.Left, new_value)
		} else {
			curr.Nodes_Left++
			curr.Left = &Node{
				Value: new_value,
			}
		}
	} else {
		if curr.Right != nil {
			curr.Nodes_Right++
			Insert(curr.Right, new_value)
		} else {
			curr.Nodes_Right++
			curr.Right = &Node{
				Value: new_value,
			}
		}
	}
}

func Delete(curr *Node, value_to_remove int) {
	if !Find(curr, value_to_remove) {
		return
	}
	deleteHelper(curr, curr, value_to_remove)
}

func deleteHelper(curr *Node, parent *Node, value_to_remove int) {
	if value_to_remove < curr.Value {
		curr.Nodes_Left--
		deleteHelper(curr.Left, curr, value_to_remove)
	} else if value_to_remove > curr.Value {
		curr.Nodes_Right--
		deleteHelper(curr.Right, curr, value_to_remove)
	} else {
		if curr.Left == nil && curr.Right == nil { //if no children
			if parent.Left.Value == curr.Value { //if curr is on parent's left
				parent.Left = nil
			} else if parent.Right.Value == curr.Value { //if curr is on parent's right
				parent.Right = nil
			}
		} else if curr.Left == nil || curr.Right == nil { //if only one child
			if curr.Left != nil { //if child is on left
				if parent.Left.Value == curr.Value { //if curr is on parent's left
					parent.Left = curr.Left
				} else if parent.Right.Value == curr.Value { //if curr is on parent's right
					parent.Right = curr.Left
				}
				curr.Left = nil
			} else if curr.Right != nil { //if child is on right
				if parent.Left.Value == curr.Value { //if curr is on parent's left
					parent.Left = curr.Right
				} else if parent.Right.Value == curr.Value { //if curr is on parent's right
					parent.Right = curr.Right
				}
				curr.Right = nil
			}
		} else { //if 2 or more children
			if curr.Nodes_Left >= curr.Nodes_Right {
				child := findLargest(curr.Left, curr)
				if curr.Left != nil {
					child.Left = curr.Left
				}
				if curr.Right != nil {
					child.Right = curr.Right
				}
				child.Nodes_Left = curr.Nodes_Left - 1
				child.Nodes_Right = curr.Nodes_Right
				curr.Left = nil
				curr.Right = nil
				if parent.Left.Value == curr.Value { //if curr is on parent's left
					parent.Left = child
				} else if parent.Right.Value == curr.Value { //if curr is on parent's right
					parent.Right = child
				}
			} else {
				child := findSmallest(curr.Right, curr)
				if curr.Left != nil {
					child.Left = curr.Left
				}
				if curr.Right != nil {
					child.Right = curr.Right
				}
				child.Nodes_Left = curr.Nodes_Left
				child.Nodes_Right = curr.Nodes_Right - 1
				curr.Left = nil
				curr.Right = nil
				if parent.Left.Value == curr.Value { //if curr is on parent's left
					parent.Left = child
				} else if parent.Right.Value == curr.Value { //if curr is on parent's right
					parent.Right = child
				}
			}
		}
	}
}

func findSmallest(curr *Node, parent *Node) *Node {
	if curr.Left == nil {
		parent.Left = nil
		return curr
	}
	return findSmallest(curr.Left, curr)
}

func findLargest(curr *Node, parent *Node) *Node {
	if curr.Right == nil {
		parent.Right = nil
		return curr
	}
	return findLargest(curr.Right, curr)
}
