package queue

type Queue struct {
	elements []any
}

func (q *Queue) Enqueue(item any) {
	q.elements = append(q.elements, item)
}

func (q *Queue) Dequeue() any {
	item := q.elements[0]
	q.elements = q.elements[1:len(q.elements)]
	return item
}

func (q *Queue) Peek() any {
	return q.elements[0]
}

func (q *Queue) IsEmpty() bool {
	if len(q.elements) == 0 || q.elements == nil {
		return true
	}

	return false
}
