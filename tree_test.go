// Copyright 2014 Andrew Bursavich. All rights reserved.
// Use of this source code is governed by The MIT License
// which can be found in the LICENSE file.

package fenwick

import (
	"math/rand"
	"testing"
	"testing/quick"
)

func TestBasic(tt *testing.T) {
	min, max := 0, 32
	t := NewTree(min, max)
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
		t := NewTree(0, len(d)-1)
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
		t := NewTree(0, len(d)-1)
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
		t := NewTree(0, len(d)-1)
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
		t := NewTree(0, len(d)-1)
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
