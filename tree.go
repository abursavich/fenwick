// Copyright 2014 Andrew Bursavich. All rights reserved.
// Use of this source code is governed by The MIT License
// which can be found in the LICENSE file.

// Package fenwick provides an implementation of a Fenwick Tree
// or Binary Indexed Tree.
package fenwick

// Tree represents a Fenwick Tree or Binary Indexed Tree
// which provides efficient manipulation and calculation
// of prefix sums of a table of values.
//
// Where n is the size of the table, it requires O(n) space
// and operations take O(log n) time.
type Tree struct {
	n []int
	m int
}

// NewTree returns a new Tree with the inclusive range [min, max].
func NewTree(min, max int) *Tree {
	return &Tree{
		n: make([]int, max-min+1),
		m: min,
	}
}

// Add adds delta to index i.
func (t *Tree) Add(i, delta int) {
	// convert the index to be one-based for lookup math.
	i = i - t.m + 1
	// increment every node on the path from the target to the root where
	// the index of the node is less than or equal to the target node.
	// in other words, increment the given node and every parent on its path
	// to the root where the child is on the left of the parent.
	for k := len(t.n); 0 < i && i <= k; i += lsb(i) {
		t.n[i-1] += delta // correct for the difference in indices
	}
}

// Prefix returns the sum of values in the inclusive range [min, i].
func (t *Tree) Prefix(i int) int {
	// convert the index to be one-based for lookup math.
	if i = i - t.m + 1; i > len(t.n) {
		i = len(t.n) // truncate to maximum range.
	}
	// accumulate every node on the path from the target to the root where
	// the index of the node is greater than or equal to the target node.
	// in other words, accumulate the given node and every parent on its path
	// to the root where the child is on the right of the parent.
	n := 0
	for ; i > 0; i = i ^ lsb(i) {
		n += t.n[i-1] // correct for the difference in indices
	}
	return n
}

// Range returns the sum of values in the inclusive range [i, k].
func (t *Tree) Range(i, k int) int {
	return t.Prefix(k) - t.Prefix(i-1)
}

// Value returns the value at index i.
func (t *Tree) Value(i int) int {
	return t.Range(i, i)
}

// Set sets the value v at index i.
func (t *Tree) Set(i, v int) {
	if vi := t.Value(i); vi != v {
		t.Add(i, v-vi)
	}
}

// Min returns the minimum index in the range.
func (t *Tree) Min() int {
	return t.m
}

// Max returns the maximum index in the range.
func (t *Tree) Max() int {
	return t.m + len(t.n) - 1
}

// lsb returns the lowest set bit.
func lsb(i int) int {
	return i & -i
}
