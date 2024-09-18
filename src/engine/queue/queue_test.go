package queue

import (
	"testing"
)

func TestQueue_Push(t *testing.T) {
	q := &Queue{}

	q.Push(1)
	q.Push(2)
	q.Push(3)

	if q.Len() != 3 {
		t.Errorf("Expected length 3, got %d", q.Len())
	}

	if q.head.item != 1 {
		t.Errorf("Expected head to be 1, got %v", q.head.item)
	}

	if q.tail.item != 3 {
		t.Errorf("Expected tail to be 3, got %v", q.tail.item)
	}
}

func TestQueue_Pop(t *testing.T) {
	q := &Queue{}

	q.Push(1)
	q.Push(2)
	q.Push(3)

	item := q.Pop()
	if item != 1 {
		t.Errorf("Expected to pop 1, got %v", item)
	}

	if q.Len() != 2 {
		t.Errorf("Expected length 2, got %d", q.Len())
	}

	item = q.Pop()
	if item != 2 {
		t.Errorf("Expected to pop 2, got %v", item)
	}

	item = q.Pop()
	if item != 3 {
		t.Errorf("Expected to pop 3, got %v", item)
	}

	if q.Len() != 0 {
		t.Errorf("Expected length 0, got %d", q.Len())
	}

	if q.Pop() != nil {
		t.Error("Expected to pop nil from an empty queue")
	}

	if q.Len() != 0 {
		t.Errorf("Expected length 0, got %d", q.Len())
	}
}

func TestQueue_PushPopLen(t *testing.T) {
	q := &Queue{}

	// Initial length should be 0
	if q.Len() != 0 {
		t.Errorf("Expected initial length 0, got %d", q.Len())
	}

	// Push items and check length
	q.Push(1)
	q.Push(2)
	q.Push(3)

	if q.Len() != 3 {
		t.Errorf("Expected length 3 after pushing 3 items, got %d", q.Len())
	}

	// Pop an item and check its value
	item := q.Pop()
	if item != 1 {
		t.Errorf("Expected to pop 1, got %v", item)
	}

	// Length should decrease
	if q.Len() != 2 {
		t.Errorf("Expected length 2 after popping 1 item, got %d", q.Len())
	}

	// Push another item and check length
	q.Push(4)

	if q.Len() != 3 {
		t.Errorf("Expected length 3 after pushing another item, got %d", q.Len())
	}

	// Pop remaining items and check values
	item = q.Pop()
	if item != 2 {
		t.Errorf("Expected to pop 2, got %v", item)
	}

	item = q.Pop()
	if item != 3 {
		t.Errorf("Expected to pop 3, got %v", item)
	}

	item = q.Pop()
	if item != 4 {
		t.Errorf("Expected to pop 4, got %v", item)
	}

	// Length should be back to 0
	if q.Len() != 0 {
		t.Errorf("Expected length 0 after popping all items, got %d", q.Len())
	}

	// Ensure popping from an empty queue returns nil
	if q.Pop() != nil {
		t.Error("Expected nil when popping from an empty queue")
	}

	// Length should still be 0
	if q.Len() != 0 {
		t.Errorf("Expected length 0 after popping all items, got %d", q.Len())
	}
}
