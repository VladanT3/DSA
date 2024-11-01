package stack

import (
	"reflect"
	"testing"
)

func TestStack(t *testing.T) {
	var test_stack Stack = Stack{
		elements: []any{1, 2, 3, 4, 5},
	}

	t.Run("Push", func(t *testing.T) {
		test_stack.Push(6)
		if !reflect.DeepEqual(test_stack.elements, []any{1, 2, 3, 4, 5, 6}) {
			t.Errorf("\nexpected: %v\ngot: %v", []any{1, 2, 3, 4, 5, 6}, test_stack.elements)
		}
	})
	t.Run("Pop", func(t *testing.T) {
		test_item := test_stack.Pop()
		if test_item != 6 {
			t.Errorf("\nexpected: %d\ngot: %d", 6, test_item)
		}
	})
	t.Run("Peek", func(t *testing.T) {
		peeked_item := test_stack.Peek()
		if peeked_item != 5 {
			t.Errorf("\nexpected: %d\ngot: %d", 5, peeked_item)
		}
	})
	t.Run("IsEmpty", func(t *testing.T) {
		if test_stack.IsEmpty() {
			t.Errorf("\nexpected: %v\ngot: %v", false, test_stack.IsEmpty())
		}
	})
}
