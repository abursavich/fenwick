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
type Tree []int

// NewTree returns a new Tree with the inclusive range [min, max].
func NewTree(min, max int) Tree {
	t := make([]int, max-min+2)
	t[0] = min
	return t
}

// Add adds delta to index i.
func (t Tree) Add(i int, delta int) {
	if i = i - t[0] + 1; i < 1 {
		return
	}
	for k := len(t); i < k; i += lsb(i) {
		t[i] += delta
	}
}

// Sum returns the sum of values at indices inclusively up to i.
func (t Tree) Sum(i int) int {
	if i = i - t[0] + 1; i >= len(t) {
		i = len(t) - 1
	}
	n := 0
	for ; i > 0; i = i ^ lsb(i) {
		n += t[i]
	}
	return n
}

// Range returns the sum of values in the inclusive range [i, k].
func (t Tree) Range(i, k int) int {
	return t.Sum(k) - t.Sum(i-1)
}

// Value returns the value at index i.
func (t Tree) Value(i int) int {
	return t.Range(i, i)
}

// Min returns the minimum index in the range.
func (t Tree) Min() int {
	return t[0]
}

// Max returns the maximum index in the range.
func (t Tree) Max() int {
	return t[0] + len(t) - 2
}

// lsb returns the lowest set bit.
func lsb(i int) int {
	return i & -i
}
