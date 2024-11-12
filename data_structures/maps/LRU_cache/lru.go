package lrucache

type Node struct {
	value any
	next  *Node
	prev  *Node
}

type LRU struct {
	lookup         map[any]*Node
	reverse_lookup map[*Node]any
	length         int
	capacity       int
	head           *Node
	tail           *Node
}

func NewLRU(cap int) *LRU {
	lru := new(LRU)
	lru.length = 0
	lru.capacity = cap
	lru.head, lru.tail = nil, nil
	lru.lookup = make(map[any]*Node, cap)
	lru.reverse_lookup = make(map[*Node]any, cap)

	return lru
}

func (lru *LRU) update(key any, value any) {
	node, ok := lru.lookup[key]
	if !ok {
		if lru.length == lru.capacity {
			tail := lru.tail
			tail.prev.next = nil
			lru.tail = tail.prev
			tail.prev = nil
			delete(lru.lookup, lru.reverse_lookup[tail])
			delete(lru.reverse_lookup, tail)
			lru.length--
		}
		new_node := Node{
			value: value,
			next:  lru.head,
		}
		if lru.head != nil {
			lru.head.prev = &new_node
		}
		if lru.tail == nil {
			lru.tail = &new_node
		}
		lru.head = &new_node
		lru.length++
		lru.lookup[key] = &new_node
		lru.reverse_lookup[&new_node] = key
	} else {
		node.value = value
		if node == lru.head {
			return
		} else if node == lru.tail {
			node.prev.next = nil
			lru.tail = node.prev
		} else {
			node.next.prev = node.prev
			node.prev.next = node.next
		}

		node.prev = nil
		node.next = lru.head
		lru.head = node
	}
}

func (lru *LRU) get(key any) any {
	node, ok := lru.lookup[key]
	if !ok {
		return nil
	}

	if node == lru.head {
		return node.value
	} else if node == lru.tail {
		node.prev.next = nil
		lru.tail = node.prev
	} else {
		node.next.prev = node.prev
		node.prev.next = node.next
	}

	node.prev = nil
	node.next = lru.head
	lru.head = node
	node.next.prev = node

	return node.value
}
