package heap

import "testing"

func TestMinHeap(t *testing.T) {
	heap := MinHeap{}
	t.Run("InitLength", func(t *testing.T) {
		if heap.Length != 0 {
			t.Errorf("\nexpected: %d\ngot: %d", 0, heap.Length)
		}
	})
	t.Run("LengthAfterInsert", func(t *testing.T) {
		heap.Insert(5)
		heap.Insert(3)
		heap.Insert(69)
		heap.Insert(420)
		heap.Insert(4)
		heap.Insert(1)
		heap.Insert(8)
		heap.Insert(7)
		if heap.Length != 8 {
			t.Errorf("\nexpected: %d\ngot: %d", 8, heap.Length)
		}
	})
	t.Run("Delete1", func(t *testing.T) {
		out := heap.Delete()
		if out != 1 {
			t.Errorf("\nexpected: %d\ngot: %d", 1, out)
		}
	})
	t.Run("Delete3", func(t *testing.T) {
		out := heap.Delete()
		if out != 3 {
			t.Errorf("\nexpected: %d\ngot: %d", 3, out)
		}
	})
	t.Run("Delete4", func(t *testing.T) {
		out := heap.Delete()
		if out != 4 {
			t.Errorf("\nexpected: %d\ngot: %d", 4, out)
		}
	})
	t.Run("Delete5", func(t *testing.T) {
		out := heap.Delete()
		if out != 5 {
			t.Errorf("\nexpected: %d\ngot: %d", 5, out)
		}
	})
	t.Run("LengthAfterDelete", func(t *testing.T) {
		if heap.Length != 4 {
			t.Errorf("\nexpected: %d\ngot: %d", 4, heap.Length)
		}
	})
	t.Run("Delete7", func(t *testing.T) {
		out := heap.Delete()
		if out != 7 {
			t.Errorf("\nexpected: %d\ngot: %d", 7, out)
		}
	})
	t.Run("Delete8", func(t *testing.T) {
		out := heap.Delete()
		if out != 8 {
			t.Errorf("\nexpected: %d\ngot: %d", 8, out)
		}
	})
	t.Run("Delete69", func(t *testing.T) {
		out := heap.Delete()
		if out != 69 {
			t.Errorf("\nexpected: %d\ngot: %d", 69, out)
		}
	})
	t.Run("Delete420", func(t *testing.T) {
		out := heap.Delete()
		if out != 420 {
			t.Errorf("\nexpected: %d\ngot: %d", 420, out)
		}
	})
	t.Run("LengthAfterDelete2", func(t *testing.T) {
		if heap.Length != 0 {
			t.Errorf("\nexpected: %d\ngot: %d", 0, heap.Length)
		}
	})
}
