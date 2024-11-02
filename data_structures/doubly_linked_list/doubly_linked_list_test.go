package DoublyLinkedList

import (
	"reflect"
	"testing"
)

func TestDLL(t *testing.T) {
	var head Node = Node{
		Value: 1,
	}
	var node Node = Node{
		Value: 2,
		Prev:  &head,
	}
	var tail Node = Node{
		Value: 3,
		Prev:  &node,
	}
	head.Next = &node
	node.Next = &tail
	var test_dll DLL = DLL{
		Head:   &head,
		Tail:   &tail,
		Length: 3,
	}

	t.Run("GetHead", func(t *testing.T) {
		item := test_dll.GetHead()
		if item != 1 {
			t.Errorf("\nexpected: %d\ngot: %d", 1, item)
		}
	})
	t.Run("GetTail", func(t *testing.T) {
		item := test_dll.GetTail()
		if item != 3 {
			t.Errorf("\nexpected: %d\ngot: %d", 3, item)
		}
	})
	t.Run("GetAt", func(t *testing.T) {
		item := test_dll.GetAt(1)
		if item != 2 {
			t.Errorf("\nexpected: %d\ngot: %d", 2, item)
		}
	})
	t.Run("GetAt0", func(t *testing.T) {
		item := test_dll.GetAt(0)
		if item != 1 {
			t.Errorf("\nexpected: %d\ngot: %d", 1, item)
		}
	})
	t.Run("AddToStart", func(t *testing.T) {
		test_dll.AddToStart(0)
		item := test_dll.GetHead()
		if item != 0 {
			t.Errorf("\nexpected: %d\ngot: %d", 0, item)
		}
	})
	t.Run("AddToEnd", func(t *testing.T) {
		test_dll.AddToEnd(4)
		item := test_dll.GetTail()
		if item != 4 {
			t.Errorf("\nexpected: %d\ngot: %d", 4, item)
		}
	})
	t.Run("AddAt", func(t *testing.T) {
		test_dll.AddAt(5, 1)
		item := test_dll.GetAt(1)
		if item != 5 {
			t.Errorf("\nexpected: %d\ngot: %d", 5, item)
		}
	})
	t.Run("DeleteHead", func(t *testing.T) {
		test_dll.DeleteHead()
		item := test_dll.GetHead()
		if item != 5 {
			t.Errorf("\nexpected: %d\ngot: %d", 5, item)
		}
	})
	t.Run("DeleteTail", func(t *testing.T) {
		test_dll.DeleteTail()
		item := test_dll.GetTail()
		if item != 3 {
			t.Errorf("\nexpected: %d\ngot: %d", 3, item)
		}
	})
	t.Run("GetAll", func(t *testing.T) {
		items := test_dll.GetAll()
		if reflect.DeepEqual(items, []any{5, 1, 2, 3}) {
			t.Errorf("\nexpected: %v\ngot: %v", []any{5, 1, 2, 3}, items)
		}
	})
	t.Run("DeleteAt", func(t *testing.T) {
		test_dll.DeleteAt(2)
		item := test_dll.GetAt(2)
		if item != 3 {
			t.Errorf("\nexpected: %d\ngot: %d", 3, item)
		}
	})
}
