package lrucache

import "testing"

func TestLRUCache(t *testing.T) {
	test_lru := NewLRU(3)

	t.Run("GetFromEmpty", func(t *testing.T) {
		res := test_lru.get("foo")
		if res != nil {
			t.Errorf("\nexpected: %v\ngot: %v", nil, res)
		}
	})
	t.Run("AddNewToEmpty", func(t *testing.T) {
		test_lru.update("foo", 69)
		res := test_lru.get("foo")
		if res != 69 {
			t.Errorf("\nexpected: %v\ngot: %v", 69, res)
		}
	})
	t.Run("AddNewToLenOf1", func(t *testing.T) {
		test_lru.update("bar", 420)
		res := test_lru.get("bar")
		if res != 420 {
			t.Errorf("\nexpected: %v\ngot: %v", 420, res)
		}
	})
	t.Run("AddNewToLenOf2", func(t *testing.T) {
		test_lru.update("baz", 1337)
		res := test_lru.get("baz")
		if res != 1337 {
			t.Errorf("\nexpected: %v\ngot: %v", 1337, res)
		}
	})
	t.Run("AddNewToLenOf3", func(t *testing.T) {
		test_lru.update("ball", 69420)
		res := test_lru.get("ball")
		if res != 69420 {
			t.Errorf("\nexpected: %v\ngot: %v", 69420, res)
		}
	})
	t.Run("GetCurrHead", func(t *testing.T) {
		res := test_lru.get("ball")
		if res != 69420 {
			t.Errorf("\nexpected: %v\ngot: %v", 69420, res)
		}
	})
	t.Run("GetNonExistant", func(t *testing.T) {
		res := test_lru.get("foo")
		if res != nil {
			t.Errorf("\nexpected: %v\ngot: %v", nil, res)
		}
	})
	t.Run("GetCurrTail", func(t *testing.T) {
		res := test_lru.get("bar")
		if res != 420 {
			t.Errorf("\nexpected: %v\ngot: %v", 420, res)
		}
	})
	t.Run("AddFooAgain", func(t *testing.T) {
		test_lru.update("foo", 69)
		res := test_lru.get("foo")
		if res != 69 {
			t.Errorf("\nexpected: %v\ngot: %v", 69, res)
		}
	})
	t.Run("GetInTheMiddle", func(t *testing.T) {
		res := test_lru.get("bar")
		if res != 420 {
			t.Errorf("\nexpected: %v\ngot: %v", 420, res)
		}
	})
	t.Run("GetInTheMiddle2", func(t *testing.T) {
		res := test_lru.get("foo")
		if res != 69 {
			t.Errorf("\nexpected: %v\ngot: %v", 69, res)
		}
	})
	t.Run("GetNonExistant2", func(t *testing.T) {
		res := test_lru.get("baz")
		if res != nil {
			t.Errorf("\nexpected: %v\ngot: %v", nil, res)
		}
	})
}
