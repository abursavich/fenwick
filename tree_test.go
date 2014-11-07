// Copyright 2014 Andrew Bursavich. All rights reserved.
// Use of this source code is governed by The MIT License
// which can be found in the LICENSE file.

package fenwick_test

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
	"testing/quick"

	"github.com/abursavich/fenwick"
)

func TestBasic(tt *testing.T) {
	min, max := 0, 32
	t := fenwick.NewTree(min, max)
	for i := min; i <= max; i++ {
		t.Add(i, 1)
	}
	for i := min; i <= max; i++ {
		if got, exp := t.Prefix(i), i-min+1; got != exp {
			tt.Errorf("t.Sum(%d); got: %d; expected: %d", i, got, exp)
		}
		if got, exp := t.Value(i), 1; got != exp {
			tt.Errorf("t.Value(%d); got: %d; expected: %d", i, got, exp)
		}
		for k := i; k <= max; k++ {
			if got, exp := t.Range(i, k), k-i+1; got != exp {
				tt.Errorf("t.Sum(%d, %d); got: %d; expected: %d", i, k, got, exp)
			}
		}
	}
	if got := t.Min(); got != min {
		tt.Errorf("t.Min(); got: %d; expected: %d", got, min)
	}
	if got := t.Max(); got != max {
		tt.Errorf("t.Max(); got: %d; expected: %d", got, max)
	}
}

// TestPrimitiveFuzz tests Add and Prefix which are the
// operations upon which all others are contructed.
func TestPrimitiveFuzz(tt *testing.T) {
	tree := func(d []int) []int {
		t := fenwick.NewTree(0, len(d)-1)
		for i, v := range d {
			t.Add(i, v)
		}
		for i := range d {
			d[i] = t.Prefix(i)
		}
		return d
	}
	slow := func(d []int) []int {
		for i := 0; i < len(d)-2; i++ {
			d[i+1] += d[i]
		}
		return d
	}
	if err := quick.CheckEqual(tree, slow, nil); err != nil {
		tt.Error(err)
	}
}

func TestRangeFuzz(tt *testing.T) {
	minMax := func(d []int, k, j int) (min int, max int) {
		n := len(d)
		if n == 0 {
			return 0, 0
		}
		if k = k % n; k < 0 {
			k = -k
		}
		if j = j % n; j < 0 {
			j = -j
		}
		if j < k {
			return j, k
		}
		return k, j
	}
	tree := func(d []int, k, j int) int {
		t := fenwick.NewTree(0, len(d)-1)
		for i, v := range d {
			t.Add(i, v)
		}
		return t.Range(minMax(d, k, j))
	}
	slow := func(d []int, k, j int) int {
		if len(d) == 0 {
			return 0
		}
		n := 0
		for i, max := minMax(d, k, j); i <= max; i++ {
			n += d[i]
		}
		return n
	}
	if err := quick.CheckEqual(tree, slow, nil); err != nil {
		tt.Error(err)
	}
}

func TestValueFuzz(tt *testing.T) {
	fn := func(d []int) bool {
		t := fenwick.NewTree(0, len(d)-1)
		for i, v := range d {
			t.Add(i, v)
		}
		for i, v := range d {
			if v != t.Value(i) {
				return false
			}
		}
		return true
	}
	if err := quick.Check(fn, nil); err != nil {
		tt.Error(err)
	}
}

func TestSetFuzz(tt *testing.T) {
	fn := func(d []int) bool {
		t := fenwick.NewTree(0, len(d)-1)
		for i, v := range d {
			t.Add(i, rand.Int())
			t.Set(i, v)
		}
		for i, v := range d {
			if v != t.Value(i) {
				return false
			}
		}
		return true
	}
	if err := quick.Check(fn, nil); err != nil {
		tt.Error(err)
	}
}

// An example solution to the HackerRank Triplets challenge.
func ExampleTree() {
	// https://www.hackerrank.com/challenges/triplets
	//
	// There is an integer array d which does not contain
	// more than two elements of the same value. How many
	// distinct ascending triplets
	// (d[i] < d[j] < d[k], i < j < k) are present?
	n := readInt()
	d := make([]int, n)
	for i := range d {
		d[i] = readInt()
	}
	// Compress the range of the table while preserving
	// ordering by mapping all values into the range [0,k)
	// where k is the number of unique values in d.
	p := make(map[int][]int, n/2)
	u := make([]int, 0, n/2)
	for i, v := range d { // O(n)
		pv, ok := p[v]
		if !ok {
			u = append(u, v)
		}
		p[v] = append(pv, i)
	}
	sort.Ints(u)          // O(k*log(k))
	for i, v := range u { // O(n)
		for _, j := range p[v] {
			d[j] = i
		}
	}
	// Make a single pass over the values from left to right.
	// Use one Tree to track which values have been seen.
	// Use a second Tree to track the number of valid pairs by
	// querying the first Tree for seen values less than the
	// current value. Use a third Tree to track the number of
	// valid triplets by querying the second Tree for pairs of
	// values less than the current value.
	k := len(u) - 1
	seen := fenwick.NewTree(0, k)
	pair := fenwick.NewTree(0, k)
	trip := fenwick.NewTree(0, k)
	for _, v := range d { // O(n*log(k))
		seen.Set(v, 1)
		pair.Set(v, seen.Prefix(v-1))
		trip.Set(v, pair.Prefix(v-1))
	}
	// Print the sum of all triplet counts.
	fmt.Println(trip.Prefix(k)) // O(log(k))
}

func readInt() int { return 0 }
