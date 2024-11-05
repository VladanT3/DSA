package breadth_first_search

type Node struct {
	Value any
	Left  *Node
	Right *Node
}

func Breadth_First_Search(root *Node) []any {
	var items []any
	var node_queue Node_Queue

	node_queue.Enqueue(root)
	traverse(root, &items, &node_queue)

	return items
}

func traverse(curr *Node, items *[]any, node_queue *Node_Queue) {
	if curr.Left != nil {
		node_queue.Enqueue(curr.Left)
	}
	if curr.Right != nil {
		node_queue.Enqueue(curr.Right)
	}

	*items = append(*items, node_queue.Dequeue().Value)
	if !node_queue.IsEmpty() {
		traverse(node_queue.Peek(), items, node_queue)
	} else {
		return
	}
}

type Node_Queue struct {
	elements []*Node
}

func (q *Node_Queue) Enqueue(node *Node) {
	q.elements = append(q.elements, node)
}

func (q *Node_Queue) Dequeue() *Node {
	node := q.elements[0]
	q.elements = q.elements[1:len(q.elements)]
	return node
}

func (q *Node_Queue) Peek() *Node {
	return q.elements[0]
}

func (q *Node_Queue) IsEmpty() bool {
	if len(q.elements) == 0 || q.elements == nil {
		return true
	}

	return false
}
