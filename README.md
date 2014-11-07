fenwick [![GoDoc](https://godoc.org/github.com/abursavich/fenwick?status.svg)](https://godoc.org/github.com/abursavich/fenwick) [![Build Status](https://travis-ci.org/abursavich/fenwick.svg?branch=master)](https://travis-ci.org/abursavich/fenwick)
====

Package fenwick provides an implementation of a Fenwick Tree or Binary Indexed Tree which provides efficient manipulation and calculation of prefix sums on a table of values.

Given a table of size n, the tree requires O(n) space and operations take O(log(n)) time.

**Examples**:

```Go
tree := fenwick.NewTree(0, 10)
tree.Add(0, 3)
tree.Add(1, -2)
tree.Add(5, 10)
tree.Add(6, 42)
tree.Add(6, -42)
tree.Add(7, 3)
tree.Add(7, 4)
tree.Add(10, 42)
tree.Set(10, 100)
fmt.Println(tree.Value(5))     // 10
fmt.Println(tree.Value(6))     // 0
fmt.Println(tree.Value(7))     // 7
fmt.Println(tree.Value(10))    // 100
fmt.Println(tree.Prefix(0))    // 3
fmt.Println(tree.Prefix(1))    // 1
fmt.Println(tree.Prefix(6))    // 11
fmt.Println(tree.Range(1, 5))  // 8
fmt.Println(tree.Range(7, 10)) // 107
```

[HackerRank Triplets Challenge](https://www.hackerrank.com/challenges/triplets)

There is an integer array d which does not contain more than two elements of the same value. How many distinct ascending triples (d[i] < d[j] < d[k], i < j < k) are present?
```Go
func Triplets(d []int) int {
	// Compress the range of the table while preserving
	// ordering by mapping all values into the range [0,k)
	// where k is the number of unique values in d.
	n := len(d)
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
	// Use a second Tree to track how many seen values make a
	// valid pair with the current value. Use a third Tree to
	// determine how many pairs make a valid triplet with the
	// current value.
	k := len(u) - 1
	seen := fenwick.NewTree(0, k)
	pair := fenwick.NewTree(0, k)
	trip := fenwick.NewTree(0, k)
	for _, v := range d { // O(n*log(k))
		seen.Set(v, 1)
		pair.Set(v, seen.Prefix(v-1))
		trip.Set(v, pair.Prefix(v-1))
	}
	// Return the sum of all triplet counts.
	return trip.Prefix(k) // O(log(k))
}
```
