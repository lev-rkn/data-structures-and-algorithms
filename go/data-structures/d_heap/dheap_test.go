package dheap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDHeap(t *testing.T) {
	type elem struct {
		value    string
		priority int
	}
	const branchingFactor = 3
	compareFunc := func(higher, lower elem) bool {
		return higher.priority > lower.priority
	}
	q, _ := NewDHeap(compareFunc, branchingFactor, 0)

	elems := []elem{
		0: {value: "first", priority: 5},
		1: {value: "second", priority: 2},
		2: {value: "third", priority: 7},
		3: {value: "fourth", priority: 8},
		4: {value: "fifth", priority: 1},
	}

	q.Insert(elems[0])
	q.Insert(elems[1])
	q.Insert(elems[2])
	q.Insert(elems[3])
	q.Insert(elems[4])

	top, _ := q.Top()
	assert.Equal(t, elems[3], top)

	peek, _ := q.Peek()
	assert.Equal(t, elems[2], peek)

	assert.Equal(t, 4, q.Size())

	updated := elem{value: "fifth1", priority: 100}
	q.Update(elems[4], updated)
	peek, _ = q.Peek()
	assert.Equal(t, updated, peek)

	q.Remove(updated)

	top, _ = q.Top()
	assert.Equal(t, elems[2], top)

	peek, _ = q.Peek()
	assert.Equal(t, elems[0], peek)
}
