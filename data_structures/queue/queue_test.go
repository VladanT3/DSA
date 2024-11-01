package queue

import (
	"reflect"
	"testing"
)

func TestQueue(t *testing.T) {
	var test_queue Queue = Queue{
		elements: []any{1, 2, 3, 4, 5},
	}
	t.Run("Enqueue", func(t *testing.T) {
		test_queue.Enqueue(6)
		if !reflect.DeepEqual(test_queue.elements, []any{1, 2, 3, 4, 5, 6}) {
			t.Errorf("\nexpected: %v\ngot: %v", []any{1, 2, 3, 4, 5, 6}, test_queue.elements)
		}
	})
	t.Run("Dequeue", func(t *testing.T) {
		test_item := test_queue.Dequeue()
		if test_item != 1 {
			t.Errorf("\nexpected: %d\ngot: %d", 1, test_item)
		}
	})
	t.Run("Peek", func(t *testing.T) {
		peeked_item := test_queue.Peek()
		if peeked_item != 2 {
			t.Errorf("\nexpected: %d\ngot: %d", 2, peeked_item)
		}
	})
	t.Run("IsEmptyNoElems", func(t *testing.T) {
		if test_queue.IsEmpty() {
			t.Errorf("\nexpected: %v\ngot: %v", false, test_queue.IsEmpty())
		}
	})
	t.Run("IsEmptyNil", func(t *testing.T) {
		test_queue.elements = nil
		if !test_queue.IsEmpty() {
			t.Errorf("\nexpected: %v\ngot: %v", true, test_queue.IsEmpty())
		}
	})
}
