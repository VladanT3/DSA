package TreeTraversal

type Node struct {
	Value any
	Left  *Node
	Right *Node
}

func PreOrder(root *Node) []any {
	var items []any
	items = append(items, root.Value)
	traversePreOrder(&items, root.Left)
	traversePreOrder(&items, root.Right)
	return items
}

func traversePreOrder(items *[]any, curr *Node) {
	*items = append(*items, curr.Value)
	if curr.Left != nil {
		traversePreOrder(items, curr.Left)
	}
	if curr.Right != nil {
		traversePreOrder(items, curr.Right)
	}

	return
}

func InOrder(root *Node) []any {
	var items []any
	traverseInOrder(&items, root.Left)
	items = append(items, root.Value)
	traverseInOrder(&items, root.Right)
	return items
}

func traverseInOrder(items *[]any, curr *Node) {
	if curr.Left != nil {
		traverseInOrder(items, curr.Left)
	}
	*items = append(*items, curr.Value)
	if curr.Right != nil {
		traverseInOrder(items, curr.Right)
	}
}

func PostOrder(root *Node) []any {
	var items []any
	traversePostOrder(&items, root.Left)
	traversePostOrder(&items, root.Right)
	items = append(items, root.Value)
	return items
}

func traversePostOrder(items *[]any, curr *Node) {
	if curr.Left != nil {
		traversePostOrder(items, curr.Left)
	}
	if curr.Right != nil {
		traversePostOrder(items, curr.Right)
	}
	*items = append(*items, curr.Value)
}
