package DoublyLinkedList

type Node struct {
	Value any
	Next  *Node
	Prev  *Node
}

type DLL struct {
	Head   *Node
	Tail   *Node
	Length int
}

func (dll *DLL) GetHead() any {
	return dll.Head.Value
}

func (dll *DLL) GetTail() any {
	return dll.Tail.Value
}

func (dll *DLL) GetAt(pos int) any {
	if pos > dll.Length-1 {
		return nil
	} else if pos < 0 {
		return nil
	} else if pos == 0 {
		return dll.GetHead()
	} else if pos == dll.Length-1 {
		return dll.GetTail()
	}

	var curr *Node = dll.Head
	for i := 0; i < pos; i++ {
		curr = curr.Next
	}

	return curr.Value
}

func (dll *DLL) AddToStart(item any) {
	var new_node Node = Node{
		Value: item,
		Next:  dll.Head,
	}
	dll.Head.Prev = &new_node
	dll.Head = &new_node
	dll.Length++
}

func (dll *DLL) AddToEnd(item any) {
	var new_node Node = Node{
		Value: item,
		Prev:  dll.Tail,
	}
	dll.Tail.Next = &new_node
	dll.Tail = &new_node
	dll.Length++
}

func (dll *DLL) AddAt(item any, pos int) {
	if pos == 0 {
		dll.AddToStart(item)
		return
	} else if pos == dll.Length-1 {
		dll.AddToEnd(item)
		return
	} else if pos > dll.Length-1 || pos < 0 {
		return
	}

	var new_node Node = Node{
		Value: item,
	}

	if pos < (dll.Length-1)/2 {
		var curr *Node = dll.Head
		for i := 0; i < pos; i++ {
			curr = curr.Next
		}
		new_node.Next = curr
		new_node.Prev = curr.Prev
		curr.Prev.Next = &new_node
		curr.Prev = &new_node
	} else {
		var curr *Node = dll.Tail
		for i := dll.Length - 1; i > pos; i-- {
			curr = curr.Prev
		}
		new_node.Next = curr
		new_node.Prev = curr.Prev
		curr.Prev.Next = &new_node
		curr.Prev = &new_node
	}

	dll.Length++
}

func (dll *DLL) DeleteHead() {
	dll.Head = dll.Head.Next
	dll.Head.Prev.Next = nil
	dll.Head.Prev = nil
	dll.Length--
}

func (dll *DLL) DeleteTail() {
	dll.Tail = dll.Tail.Prev
	dll.Tail.Next.Prev = nil
	dll.Tail.Next = nil
	dll.Length--
}

func (dll *DLL) DeleteAt(pos int) {
	if pos == 0 {
		dll.DeleteHead()
		return
	} else if pos == dll.Length-1 {
		dll.DeleteTail()
		return
	} else if pos > dll.Length-1 {
		return
	}

	if pos < (dll.Length-1)/2 {
		var curr *Node = dll.Head
		for i := 0; i < pos; i++ {
			curr = curr.Next
		}
		curr.Prev.Next = curr.Next
		curr.Next.Prev = curr.Prev
		curr.Next = nil
		curr.Prev = nil
	} else {
		var curr *Node = dll.Tail
		for i := dll.Length - 1; i > pos; i-- {
			curr = curr.Prev
		}
		curr.Next.Prev = curr.Prev
		curr.Prev.Next = curr.Next
		curr.Next = nil
		curr.Prev = nil
	}

	dll.Length--
}

func (dll *DLL) GetAll() []any {
	var items []any
	var curr *Node = dll.Head
	for i := 0; i < dll.Length-1; i++ {
		items = append(items, curr.Value)
		curr = curr.Next
	}

	return items
}
